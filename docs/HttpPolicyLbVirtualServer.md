# HttpPolicyLbVirtualServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLogEnabled** | **bool** | If access log is enabled, all HTTP requests sent to an L7 virtual server are logged to the access log file. Both successful requests (backend server returns 2xx) and unsuccessful requests (backend server returns 4xx or 5xx) are logged to access log, if enabled.  | [optional] [default to false]
**RouterPath** | **string** | Path to router type object that PolicyLbVirtualServer connects to. The only supported router object is Network.  | [default to null]
**LbPersistenceProfile** | **string** | Path to optional object that enables persistence on a virtual server allowing related client connections to be sent to the same backend server. Persistence is disabled by default.  | [optional] [default to null]
**TrafficSource** | **string** |  | [optional] [default to null]
**IpAddress** | **string** | Configures the IP address of the PolicyLbVirtualServer where it receives all client connections and distributes them among the backend servers.  | [default to null]
**Ports** | **[]string** | Ports contains a list of at least one port or port range such as \&quot;80\&quot;, \&quot;1234-1236\&quot;. Each port element in the list should be a single port or a single port range.  | [default to null]
**ResourceType** | **string** |  | [default to null]
**InsertClientIpHeader** | **bool** | Backend web servers typically log each request they handle along with the requesting client IP address. These logs are used for debugging, analytics and other such purposes. If the deployment topology requires enabling SNAT on the load balancer, then server will see the client as the SNAT IP which defeats the purpose of logging. To work around this issue, load balancer can be configured to insert XFF HTTP header with the original client IP address. Backend servers can then be configured to log the IP address in XFF header instead of the source IP address of the connection. If XFF header is not present in the incoming request, load balancer inserts a new XFF header with the client IP address.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

