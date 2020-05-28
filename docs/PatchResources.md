# PatchResources

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Reaction Action resource type.  | [default to null]
**Injections** | [**[]Injection**](Injection.md) | Injections holding keys (variables) and their corresponding values. | [optional] [default to null]
**Body** | [***interface{}**](interface{}.md) | Patch body representing a Hierarchical Patch payload. The resources included in the body are patched replacing the injections&#x27; keys with their actual values.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

