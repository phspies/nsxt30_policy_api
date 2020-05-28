# LbVirtualServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**LbPersistenceProfilePath** | **string** | Path to optional object that enables persistence on a virtual server allowing related client connections to be sent to the same backend server. Persistence is disabled by default.  | [optional] [default to null]
**AccessListControl** | [***LbAccessListControl**](LBAccessListControl.md) |  | [optional] [default to null]
**PoolPath** | **string** | The server pool(LBPool) contains backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application.  | [optional] [default to null]
**LogSignificantEventOnly** | **bool** | The property log_significant_event_only can take effect only when access_log_enabled is true. If log_significant_event_only is true, significant events are logged in access log. For L4 virtual server, significant event means unsuccessful(error or dropped) TCP/UDP connections. For L7 virtual server, significant event means unsuccessful connections or HTTP/HTTPS requests which have error response code(e.g. 4xx, 5xx).  | [optional] [default to false]
**Rules** | [**[]LbRule**](LBRule.md) | Load balancer rules allow customization of load balancing behavior using match/action rules. Currently, load balancer rules are supported for only layer 7 virtual servers with LBHttpProfile.  | [optional] [default to null]
**DefaultPoolMemberPorts** | **[]string** | Default pool member ports when member port is not defined.  | [optional] [default to null]
**ServerSslProfileBinding** | [***LbServerSslProfileBinding**](LBServerSslProfileBinding.md) |  | [optional] [default to null]
**ApplicationProfilePath** | **string** | The application profile defines the application protocol characteristics. It is used to influence how load balancing is performed. Currently, LBFastTCPProfile, LBFastUDPProfile and LBHttpProfile, etc are supported.  | [default to null]
**AccessLogEnabled** | **bool** | If access log is enabled, all HTTP requests sent to L7 virtual server are logged to the access log file. Both successful returns information responses(1xx), successful responses(2xx), redirection messages(3xx) and unsuccessful requests, backend server returns 4xx or 5xx, are logged to access log, if enabled. All L4 virtual server connections are also logged to the access log if enabled. The non-significant events such as successful requests are not logged if log_significant_event_only is set to true.  | [optional] [default to false]
**MaxConcurrentConnections** | **int64** | To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited.  | [optional] [default to null]
**MaxNewConnectionRate** | **int64** | To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited.  | [optional] [default to null]
**LbServicePath** | **string** | virtual servers can be associated to LBService(which is similar to physical/virtual load balancer), LB virtual servers, pools and other entities could be defined independently, the LBService identifier list here would be used to maintain the relationship of LBService and other LB entities.  | [optional] [default to null]
**ClientSslProfileBinding** | [***LbClientSslProfileBinding**](LBClientSslProfileBinding.md) |  | [optional] [default to null]
**SorryPoolPath** | **string** | When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool.  | [optional] [default to null]
**IpAddress** | **string** | Configures the IP address of the LBVirtualServer where it receives all client connections and distributes them among the backend servers.  | [default to null]
**Ports** | **[]string** | Ports contains a list of at least one port or port range such as \&quot;80\&quot;, \&quot;1234-1236\&quot;. Each port element in the list should be a single port or a single port range.  | [default to null]
**Enabled** | **bool** | Flag to enable the load balancer virtual server. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

