# PolicyTraceflowObservationDroppedLogical

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePathIndex** | **int64** | The index of service path that is a chain of services represents the point where the traceflow packet was dropped.  | [optional] [default to null]
**ComponentId** | **string** | The id of the component that dropped the traceflow packet. | [optional] [default to null]
**AclRulePath** | **string** | The path of the ACL rule that was applied to forward the traceflow packet | [optional] [default to null]
**NatRulePath** | **string** | The path of the NAT rule that was applied to forward the traceflow packet | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

