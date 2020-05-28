# PolicyAlarmResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Absolute path of this object | [optional] [default to null]
**ParentPath** | **string** | Path of its parent | [optional] [default to null]
**UniqueId** | **string** | This is a UUID generated by the GM/LM to uniquely identify entites in a federated environment. For entities that are stretched across multiple sites, the same ID will be used on all the stretched sites.  | [optional] [default to null]
**RelativePath** | **string** | Path relative from its parent | [optional] [default to null]
**SourceReference** | **string** | path of the object on which alarm is created | [optional] [default to null]
**Message** | **string** | error message to describe the issue | [optional] [default to null]
**ErrorDetails** | [***PolicyApiError**](PolicyApiError.md) |  | [optional] [default to null]
**SourceSiteId** | **string** | This field will refer to the source site on which the alarm is generated. This field is populated by GM, when it receives corresponding notification from LM.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
