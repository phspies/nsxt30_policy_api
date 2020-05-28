# StaticRouteBfdPeer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**BfdProfilePath** | **string** | Bfd Profile is not supported for IPv6 networks. | [optional] [default to null]
**PeerAddress** | **string** | Only IPv4 addresses are supported. Only a single BFD config per peer address is allowed. | [default to null]
**SourceAddresses** | **[]string** | Array of Tier0 external interface IP addresses. BFD peering is established from all these source addresses to the neighbor specified in peer_address. Only IPv4 addresses are supported. | [optional] [default to null]
**Enabled** | **bool** | Flag to enable BFD peer. | [optional] [default to true]
**Scope** | **[]string** | Represents the array of policy paths of locale services where this BFD peer should get relalized on. The locale service service and this BFD peer must belong to the same router. Default scope is empty.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

