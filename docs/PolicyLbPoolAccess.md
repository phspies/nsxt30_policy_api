# PolicyLbPoolAccess

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**PoolPort** | **int32** | Port for LoadBalancer to send connections to the PolicyLbPoolAccess&#x27;s Group. Pool_port could be optional, if it is not specified, LB will use PolicyLbVirtualServer port to connect to backend servers. If the PolicyLbMonitorProfile is configured in PolicyLbPoolAccess and active monitor IP protocol is TCP/UDP(which requires TCP or UDP port number), monitor_port should be specified if pool_port is unset.  | [optional] [default to null]
**IpPortList** | [**[]IpAddressPortPair**](IPAddressPortPair.md) | IP Port list for applications within the Group to allow for non-uniform port usage by applications  | [optional] [default to null]
**SourceNat** | **string** | Depending on the topology, Source NAT (SNAT) may be required to ensure traffic from the server destined to the client is received by the load balancer. SNAT can be enabled per pool. If SNAT is not enabled for a pool, then load balancer uses the client IP and port (spoofing) while establishing connections to the servers. This is referred to as no-SNAT or TRANSPARENT mode.  SNAT is enabled by default and will use the load balancer interface IP and an ephemeral port as the source IP and port of the server side connection.  | [optional] [default to SOURCE_NAT.ENABLED]
**Algorithm** | **string** | Load balanding algorithm controls how the incoming connections are distributed among the members. - ROUND_ROBIN - requests to the application servers are distributed in a round-robin fashion, - LEAST_CONNECTION - next request is assigned to the server with the least number of active connections  | [optional] [default to ALGORITHM.ROUND_ROBIN]
**LbMonitorProfile** | **string** | Path of the PolicyLbMonitorProfile to actively monitor the PolicyLbPoolAccess&#x27;s Group  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

