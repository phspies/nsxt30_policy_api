# PolicyLbVirtualServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**AccessLogEnabled** | **bool** | If access log is enabled, all HTTP requests sent to an L7 virtual server are logged to the access log file. Both successful requests (backend server returns 2xx) and unsuccessful requests (backend server returns 4xx or 5xx) are logged to access log, if enabled.  | [optional] [default to false]
**RouterPath** | **string** | Path to router type object that PolicyLbVirtualServer connects to. The only supported router object is Network.  | [default to null]
**LbPersistenceProfile** | **string** | Path to optional object that enables persistence on a virtual server allowing related client connections to be sent to the same backend server. Persistence is disabled by default.  | [optional] [default to null]
**TrafficSource** | **string** |  | [optional] [default to null]
**IpAddress** | **string** | Configures the IP address of the PolicyLbVirtualServer where it receives all client connections and distributes them among the backend servers.  | [default to null]
**Ports** | **[]string** | Ports contains a list of at least one port or port range such as \&quot;80\&quot;, \&quot;1234-1236\&quot;. Each port element in the list should be a single port or a single port range.  | [default to null]
**ResourceType** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

