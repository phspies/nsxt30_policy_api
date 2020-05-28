# BgpGracefulRestartConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | If mode is DISABLE, then graceful restart and helper modes are disabled. If mode is GR_AND_HELPER, then both graceful restart and helper modes are enabled. If mode is HELPER_ONLY, then helper mode is enabled. HELPER_ONLY mode is the ability for a BGP speaker to indicate its ability to preserve forwarding state during BGP restart. GRACEFUL_RESTART mode is the ability of a BGP speaker to advertise its restart to its peers.  | [optional] [default to MODE.HELPER_ONLY]
**Timer** | [***BgpGracefulRestartTimer**](BgpGracefulRestartTimer.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

