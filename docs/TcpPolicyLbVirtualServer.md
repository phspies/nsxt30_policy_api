# TcpPolicyLbVirtualServer

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

