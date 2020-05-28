# GenericPolicyRealizedResource

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
**SitePath** | **string** | The site where this entity resides. | [optional] [default to null]
**IntentPaths** | **[]string** | Collection of intent paths | [optional] [default to null]
**EnforcementPointPath** | **string** | The path of the enforcement point. | [optional] [default to null]
**ExtendedAttributes** | [**[]AttributeVal**](AttributeVal.md) | Collection of type specific properties | [optional] [default to null]
**EntityType** | **string** | Type of realized entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

