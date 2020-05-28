# FloodProtectionProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IcmpActiveFlowLimit** | **int64** | If this field is empty, firewall will not set a limit to active ICMP connections. | [optional] [default to null]
**TcpHalfOpenConnLimit** | **int64** | If this field is empty, firewall will not set a limit to half open TCP connections. | [optional] [default to null]
**UdpActiveFlowLimit** | **int64** | If this field is empty, firewall will not set a limit to active UDP connections. | [optional] [default to null]
**ResourceType** | **string** | GatewayFloodProtectionProfile is used for all Tier0 and Tier1 gateways. DistributedFloodProtectionProfile is used for all Transport Nodes.  | [default to null]
**OtherActiveConnLimit** | **int64** | If this field is empty, firewall will not set a limit to other active connections. besides UDP, ICMP and half open TCP connections. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

