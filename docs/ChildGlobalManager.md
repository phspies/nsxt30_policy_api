# ChildGlobalManager

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**MarkedForDelete** | **bool** | If this field is set to true, delete operation is triggered on the intent tree. This resource along with its all children in intent tree will be deleted. This is a cascade delete and should only be used if intent object along with its all children are to be deleted. This does not support deletion of single non-leaf node within the tree and should be used carefully.  | [optional] [default to false]
**GlobalManager** | [***GlobalManager**](GlobalManager.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

