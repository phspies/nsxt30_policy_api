# PolicyServiceInstanceStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**ServiceInstanceId** | **string** | PolicyServiceInsatnce path | [optional] [default to null]
**InstanceRuntimeStatistics** | [**[]InstanceRuntimeStatistic**](InstanceRuntimeStatistic.md) | Statistics for the data NICs for all the runtimes associated with this service instance.  | [optional] [default to null]
**EnforcementPointPath** | **string** | Enforcement point path, forward slashes must be escaped using %2F.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

