# HttpsPolicyLbVirtualServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsertClientIpHeader** | **bool** | Backend web servers typically log each request they handle along with the requesting client IP address. These logs are used for debugging, analytics and other such purposes. If the deployment topology requires enabling SNAT on the load balancer, then server will see the client as the SNAT IP which defeats the purpose of logging. To work around this issue, load balancer can be configured to insert XFF HTTP header with the original client IP address. Backend servers can then be configured to log the IP address in XFF header instead of the source IP address of the connection. If XFF header is not present in the incoming request, load balancer inserts a new XFF header with the client IP address.  | [optional] [default to false]
**ClientSslSettings** | **string** | Security settings representing various security settings when the VirtualServer acts as an SSL server - BASE_SECURE_111317 - MODERATE_SECURE_111317 - HIGH_SECURE_111317  | [optional] [default to CLIENT_SSL_SETTINGS.HIGH_SECURE_111317]
**ClientSslCertificateIds** | **[]string** | Client-side SSL profile binding allows multiple certificates, for different hostnames, to be bound to the same virtual server. The setting is used when load balancer acts as an SSL server and terminating the client SSL connection  | [optional] [default to null]
**DefaultClientSslCertificateId** | **string** | The setting is used when load balancer acts as an SSL server and terminating the client SSL connection.  A default certificate should be specified which will be used if the server does not host multiple hostnames on the same IP address or if the client does not support SNI extension.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

