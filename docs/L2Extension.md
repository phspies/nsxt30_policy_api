# L2Extension

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L2vpnPaths** | **[]string** | Policy paths corresponding to the associated L2 VPN sessions  | [optional] [default to null]
**LocalEgress** | [***LocalEgress**](LocalEgress.md) |  | [optional] [default to null]
**L2vpnPath** | **string** | This property has been deprecated. Please use the property l2vpn_paths for setting the paths of associated L2 VPN session. This property will continue to work as expected to provide backwards compatibility. However, when both l2vpn_path and l2vpn_paths properties are specified, only l2vpn_paths is used.  | [optional] [default to null]
**TunnelId** | **int32** | Tunnel ID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

