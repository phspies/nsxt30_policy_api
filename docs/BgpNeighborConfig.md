# BgpNeighborConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**OutRouteFilters** | **[]string** | Specify path of prefix-list or route map to filter routes for OUT direction. When not specified, a built-in prefix-list named &#x27;prefixlist-out-default&#x27; is automatically applied. This property is deprecated, use route_filtering instead. Specifying different values for both properties will result in error.  | [optional] [default to null]
**GracefulRestartMode** | **string** | If mode is DISABLE, then graceful restart and helper modes are disabled. If mode is GR_AND_HELPER, then both graceful restart and helper modes are enabled. If mode is HELPER_ONLY, then helper mode is enabled. HELPER_ONLY mode is the ability for a BGP speaker to indicate its ability to preserve forwarding state during BGP restart. GRACEFUL_RESTART mode is the ability of a BGP speaker to advertise its restart to its peers.  | [optional] [default to null]
**Bfd** | [***BgpBfdConfig**](BgpBfdConfig.md) |  | [optional] [default to null]
**KeepAliveTime** | **int32** | Interval (in seconds) between keep alive messages sent to peer.  | [optional] [default to 60]
**MaximumHopLimit** | **int32** | Maximum number of hops allowed to reach BGP neighbor.  | [optional] [default to 1]
**HoldDownTime** | **int32** | Wait time in seconds before declaring peer dead.  | [optional] [default to 180]
**InRouteFilters** | **[]string** | Specify path of prefix-list or route map to filter routes for IN direction. This property is deprecated, use route_filtering instead. Specifying different values for both properties will result in error.  | [optional] [default to null]
**RemoteAsNum** | **string** | 4 Byte ASN of the neighbor in ASPLAIN Format | [default to null]
**RouteFiltering** | [**[]BgpRouteFiltering**](BgpRouteFiltering.md) | Enable address families and route filtering in each direction.  | [optional] [default to null]
**SourceAddresses** | **[]string** | Source addresses should belong to Tier0 external or loopback interface IP Addresses . BGP peering is formed from all these addresses. This property is mandatory when maximum_hop_limit is greater than 1.  | [optional] [default to null]
**Password** | **string** | Specify password for BGP neighbor authentication. Empty string (\&quot;\&quot;) clears existing password.  | [optional] [default to null]
**AllowAsIn** | **bool** | Flag to enable allowas_in option for BGP neighbor | [optional] [default to false]
**NeighborAddress** | **string** | Neighbor IP Address | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

