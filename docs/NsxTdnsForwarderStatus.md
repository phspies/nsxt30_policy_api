# NsxTdnsForwarderStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**EnforcementPointPath** | **string** | Policy path referencing the enforcement point from where the status is fetched.  | [optional] [default to null]
**Status** | **string** | UP means the DNS forwarder is working correctly on the active transport node and the stand-by transport node (if present). Failover will occur if either node goes down. DOWN means the DNS forwarder is down on both active transport node and standby node (if present). The DNS forwarder does not function in this situation. Error means there is some error on one or both transport node, or no status was reported from one or both transport nodes. The DNS forwarder may be working (or not working). NO_BACKUP means DNS forwarder is working in only one transport node, either because it is down on the standby node, or no standby is configured. An forwarder outage will occur if the active node goes down.  | [optional] [default to null]
**Timestamp** | **int64** | Time stamp of the current status, in ms | [optional] [default to null]
**ExtraMessage** | **string** | Extra message, if available | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

