# BgpRouteFiltering

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutRouteFilters** | **[]string** | Specify path of prefix-list or route map to filter routes for OUT direction. When not specified, a built-in prefix-list named &#x27;prefixlist-out-default&#x27; is automatically applied.  | [optional] [default to null]
**InRouteFilters** | **[]string** | Specify path of prefix-list or route map to filter routes for IN direction.  | [optional] [default to null]
**Enabled** | **bool** | Flag to enable address family. | [optional] [default to true]
**AddressFamily** | **string** | Address family type. If not configured, this property automatically derived for IPv4 &amp; IPv6 peer configuration. | [optional] [default to null]
**MaximumRoutes** | **int32** | Maximum number of routes for the address family.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

