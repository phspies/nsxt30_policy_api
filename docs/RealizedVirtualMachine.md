# RealizedVirtualMachine

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealizationSpecificIdentifier** | **string** | Realization id of this object | [optional] [default to null]
**IntentReference** | **[]string** | Desire state paths of this object | [optional] [default to null]
**State** | **string** | Realization state of this object | [default to null]
**RealizationApi** | **string** | Realization API of this object on enforcement point | [optional] [default to null]
**Alarms** | [**[]PolicyAlarmResource**](PolicyAlarmResource.md) | Alarm info detail | [optional] [default to null]
**RuntimeError** | **string** | It define the root cause for runtime error.  | [optional] [default to null]
**RuntimeStatus** | **string** | Possible values could be UP, DOWN, UNKNOWN, DEGRADED This list is not exhaustive.  | [optional] [default to null]
**HostId** | **string** | Id of the host on which the vm exists. | [optional] [default to null]
**LocalIdOnHost** | **string** | Id of the vm unique within the host. | [optional] [default to null]
**PowerState** | **string** | Current power state of this virtual machine in the system. | [optional] [default to null]
**ComputeIds** | **[]string** | List of external compute ids of the virtual machine in the format &#x27;id-type-key:value&#x27; , list of external compute ids [&#x27;uuid:xxxx-xxxx-xxxx-xxxx&#x27;, &#x27;moIdOnHost:moref-11&#x27;, &#x27;instanceUuid:xxxx-xxxx-xxxx-xxxx&#x27;] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

