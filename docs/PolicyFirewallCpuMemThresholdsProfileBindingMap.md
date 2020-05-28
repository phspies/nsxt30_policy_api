# PolicyFirewallCpuMemThresholdsProfileBindingMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilePath** | **string** | PolicyPath of associated Profile | [default to null]
**SequenceNumber** | **int64** | Sequence number is used to resolve conflicts when two profiles get applied to a single node. Lower value gets higher precedence. Two binding maps having the same profile path should have the same sequence number.  | [default to null]
**TransportNodes** | [**[]PolicyResourceReference**](PolicyResourceReference.md) | References of transport nodes on which the profile intended to be applied.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

