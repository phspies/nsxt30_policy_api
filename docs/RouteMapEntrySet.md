# RouteMapEntrySet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferGlobalV6NextHop** | **bool** | For incoming and import route_maps on receiving both v6 global and v6 link-local address for the route, prefer to use the global address as the next hop. By default, it prefers the link-local next hop.  | [optional] [default to null]
**Med** | **int32** | Multi exit descriminator (MED) is a hint to BGP neighbors about the preferred path into an autonomous system (AS) that has multiple entry points. A lower MED value is preferred over a higher value.  | [optional] [default to null]
**LocalPreference** | **int64** | Local preference indicates the degree of preference for one BGP route over other BGP routes. The path with highest local preference is preferred.  | [optional] [default to 100]
**Weight** | **int32** | Weight is used to select a route when multiple routes are available to the same network. Route with the highest weight is preferred.  | [optional] [default to null]
**AsPathPrepend** | **string** | AS path prepend to influence route selection.  | [optional] [default to null]
**Community** | **string** | Set BGP regular or large community for matching routes. A maximum of one value for each community type separated by space. Well-known community name, community value in aa:nn (2byte:2byte) format for regular community and community value in aa:bb:nn (4byte:4byte:4byte) format for large community are supported.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

