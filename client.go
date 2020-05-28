package nsxt30policy

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the NSX-T Manager API API v3.0.0.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg     *Configuration
	Context context.Context
	common  service

	// API Services

	FederationInfraApi *FederationInfraApiService

	FederationLocalApi *FederationLocalApiService

	PolicyApi *PolicyApiService

	PolicyAuthorizationApi *PolicyAuthorizationApiService

	PolicyCapacityApi *PolicyCapacityApiService

	PolicyInfraApi *PolicyInfraApiService

	PolicyInventoryApi *PolicyInventoryApiService

	PolicyNetworkingApi *PolicyNetworkingApiService

	PolicyOperationsApi *PolicyOperationsApiService

	PolicySecurityApi *PolicySecurityApiService

	PolicySystemApi *PolicySystemApiService

	SearchApi *SearchSearchApiService

	SystemAdministrationApi *SystemAdministrationApiService
}

type service struct {
	client *APIClient
}

func GetContext(cfg *Configuration) context.Context {
	if len(cfg.ClientAuthCertFile) == 0 && cfg.RemoteAuth == false {
		auth := BasicAuth{UserName: cfg.UserName,
			Password: cfg.Password}
		return context.WithValue(context.Background(), ContextBasicAuth, auth)
	}

	return context.Background()
}

//Perform 'session create' and save the xsrf & session id to the default headers
func GetDefaultHeaders(client *APIClient) error {

	XSRF_TOKEN := "X-XSRF-TOKEN"
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)
	queryParams := url.Values{}
	formParams := url.Values{}

	requestHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// For remote Auth (vIDM use case), construct the REMOTE auth header
	remoteAuthHeader := ""
	if client.cfg.RemoteAuth {
		auth := client.cfg.UserName + ":" + client.cfg.Password
		encoded := base64.StdEncoding.EncodeToString([]byte(auth))
		remoteAuthHeader = "Remote " + encoded
		requestHeaders["Authorization"] = remoteAuthHeader
	}

	path := strings.TrimSuffix(client.cfg.BasePath, "v1") + "session/create"
	// Call session create
	r, err := client.prepareRequest(
		client.Context,
		path,
		httpMethod,
		postBody,
		requestHeaders,
		queryParams,
		formParams,
		fileName,
		fileBytes)
	if err != nil {
		return fmt.Errorf("Failed to create session: %s.", err)
	}
	response, err := client.callAPI(r)
	if err != nil || response == nil {
		return fmt.Errorf("Failed to create session: %s.", err)
	}

	if client.cfg.DefaultHeader == nil {
		client.cfg.DefaultHeader = make(map[string]string)
	}

	// Go over the headers
	for k, v := range response.Header {
		if strings.ToLower("Set-Cookie") == strings.ToLower(k) {
			r, _ := regexp.Compile("JSESSIONID=.*?;")
			result := r.FindString(v[0])
			if result != "" {
				client.cfg.DefaultHeader["Cookie"] = result
			}
		}
		if strings.ToLower(XSRF_TOKEN) == strings.ToLower(k) {
			client.cfg.DefaultHeader[XSRF_TOKEN] = v[0]
		}
	}

	response.Body.Close()

	// For remote Auth (vIDM use case), construct the REMOTE auth header
	if client.cfg.RemoteAuth {
		client.cfg.DefaultHeader["Authorization"] = remoteAuthHeader
	}
	return nil
}

func InitHttpClient(cfg *Configuration) error {

	tlsConfig := &tls.Config{
		//MinVersion:               tls.VersionTLS12,
		//PreferServerCipherSuites: true,
		InsecureSkipVerify: cfg.Insecure,
	}

	if len(cfg.ClientAuthCertFile) > 0 {
		cert, err := tls.LoadX509KeyPair(cfg.ClientAuthCertFile,
			cfg.ClientAuthKeyFile)
		if err != nil {
			return err
		}

		tlsConfig.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		}

	}

	if len(cfg.CAFile) > 0 {
		caCert, err := ioutil.ReadFile(cfg.CAFile)
		if err != nil {
			return err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig.RootCAs = caCertPool
	}

	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{Proxy: http.ProxyFromEnvironment,
		TLSClientConfig:     tlsConfig,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	cfg.HTTPClient = &http.Client{Transport: transport}
	return nil
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) (*APIClient, error) {

	if cfg.HTTPClient == nil {
		err := InitHttpClient(cfg)

		if cfg.HTTPClient == nil {
			return nil, err
		}
	}

	c := &APIClient{}
	c.cfg = cfg
	c.Context = GetContext(cfg)
	c.common.client = c
	err := GetDefaultHeaders(c)
	if err != nil {
		return nil, err
	}

	// API Services
	c.FederationInfraApi = (*FederationInfraApiService)(&c.common)
	c.FederationLocalApi = (*FederationLocalApiService)(&c.common)
	c.PolicyApi = (*PolicyApiService)(&c.common)
	c.PolicyAuthorizationApi = (*PolicyAuthorizationApiService)(&c.common)
	c.PolicyCapacityApi = (*PolicyCapacityApiService)(&c.common)
	c.PolicyInfraApi = (*PolicyInfraApiService)(&c.common)
	c.PolicyInventoryApi = (*PolicyInventoryApiService)(&c.common)
	c.PolicyNetworkingApi = (*PolicyNetworkingApiService)(&c.common)
	c.PolicyOperationsApi = (*PolicyOperationsApiService)(&c.common)
	c.PolicySecurityApi = (*PolicySecurityApiService)(&c.common)
	c.PolicySystemApi = (*PolicySystemApiService)(&c.common)
	c.SearchApi = (*SearchSearchApiService)(&c.common)
	c.SystemAdministrationApi = (*SystemAdministrationApiService)(&c.common)

	return c, nil
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

func (c *APIClient) shouldRetryOnStatus(code int) bool {
	for _, s := range c.cfg.RetriesConfiguration.RetryOnStatuses {
		if code == s {
			return true
		}
	}
	return false
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {

	// keep the initial request body string
	var requestBodyString string
	if request.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		requestBodyString = buf.String()
		request.Body = ioutil.NopCloser(strings.NewReader(requestBodyString))
	}

	// first run
	localVarHttpResponse, err := c.callAPIInternal(request)
	config := c.cfg.RetriesConfiguration
	maxRetries := int(math.Max(2, float64(config.MaxRetries)))
	// loop until not getting the retry-able error, or until max retries
	for n_try := 1; n_try < maxRetries; n_try++ {
		if localVarHttpResponse == nil || c.shouldRetryOnStatus(localVarHttpResponse.StatusCode) {
			status := 0
			if localVarHttpResponse != nil {
				status = localVarHttpResponse.StatusCode
				localVarHttpResponse.Body.Close()
			}
			log.Printf("[DEBUG] Retrying request %s %s for the %d time because of status %d", request.Method, request.URL, n_try, status)
			// sleep a random increasing time
			float_delay := float64(rand.Intn(config.RetryMinDelay * n_try))
			fixed_delay := time.Duration(math.Min(float64(config.RetryMaxDelay), float_delay))
			time.Sleep(fixed_delay * time.Millisecond)
			// reset Request.Body
			if request.Body != nil {
				request.Body = ioutil.NopCloser(strings.NewReader(requestBodyString))
			}
			// perform the request again
			localVarHttpResponse, err = c.callAPIInternal(request)
		} else {
			// Non retry-able response
			return localVarHttpResponse, err
		}
	}
	// max retries exceeded
	return localVarHttpResponse, err
}

func (c *APIClient) callAPIInternal(request *http.Request) (*http.Response, error) {
	localVarHttpResponse, err := c.cfg.HTTPClient.Do(request)

	if err == nil && localVarHttpResponse != nil && (localVarHttpResponse.StatusCode == 400 || localVarHttpResponse.StatusCode == 500) {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		err = reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form paramters and file if available.
	if len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	// Setup path and query paramters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.URL.Host = c.cfg.Host
	}

	if c.cfg.Scheme != "" {
		localVarRequest.URL.Scheme = c.cfg.Scheme
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	// Walk through any authentication.
	if ctx != nil {
		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// Added for supporting NSX 3.0 Container Inventory API
func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}