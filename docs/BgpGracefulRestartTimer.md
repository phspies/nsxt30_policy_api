# BgpGracefulRestartTimer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartTimer** | **int64** | Maximum time taken (in seconds) for a BGP session to be established after a restart. This can be used to speed up routing convergence by its peer in case the BGP speaker does not come back up after a restart. If the session is not re-established within this timer,  the receiving speaker will delete all the stale routes from that peer.  | [optional] [default to 180]
**StaleRouteTimer** | **int64** | Maximum time (in seconds) before stale routes are removed from the RIB (Routing Information Base) when BGP restarts.  | [optional] [default to 600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

