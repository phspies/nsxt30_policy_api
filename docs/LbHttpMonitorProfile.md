# LbHttpMonitorProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorPort** | **int32** | Typically, monitors perform healthchecks to Group members using the member IP address and pool_port. However, in some cases, customers prefer to run healthchecks against a different port than the pool member port which handles actual application traffic. In such cases, the port to run healthchecks against can be specified in the monitor_port value. For ICMP monitor, monitor_port is not required.  | [optional] [default to null]
**FallCount** | **int64** | Only if a healthcheck fails consecutively for a specified number of times, given with fall_count, to a member will the member status be marked DOWN.  | [optional] [default to 3]
**Interval** | **int64** | Active healthchecks are initiated periodically, at a configurable interval (in seconds), to each member of the Group.  | [optional] [default to 5]
**RiseCount** | **int64** | Once a member is DOWN, a specified number of consecutive successful healthchecks specified by rise_count will bring the member back to UP state.  | [optional] [default to 3]
**Timeout** | **int64** | Timeout specified in seconds.  After a healthcheck is initiated, if it does not complete within a certain period, then also the healthcheck is considered to be unsuccessful. Completing a healthcheck within timeout means establishing a connection (TCP or SSL), if applicable, sending the request and receiving the response, all within the configured timeout.  | [optional] [default to 5]
**ResponseStatusCodes** | **[]int32** | The HTTP response status code should be a valid HTTP status code.  | [optional] [default to null]
**RequestMethod** | **string** | The health check method for HTTP monitor type. | [optional] [default to REQUEST_METHOD.GET]
**RequestBody** | **string** | String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST.  | [optional] [default to null]
**ResponseBody** | **string** | If HTTP response body match string (regular expressions not supported) is specified (using LBHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful.  | [optional] [default to null]
**RequestUrl** | **string** | For HTTP active healthchecks, the HTTP request url sent can be customized and can include query parameters.  | [optional] [default to /]
**RequestVersion** | **string** | HTTP request version. | [optional] [default to REQUEST_VERSION.11_]
**RequestHeaders** | [**[]LbHttpRequestHeader**](LbHttpRequestHeader.md) | Array of HTTP request headers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

