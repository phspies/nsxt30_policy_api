# ConstraintTarget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Attribute name of the target entity. | [optional] [default to null]
**PathPrefix** | **string** | Path prefix of the entity to apply constraint. This is required to further disambiguiate if multiple policy entities share the same resource type. Example - Edge FW and DFW use the same resource type CommunicationMap, CommunicationEntry, Group, etc.  | [optional] [default to null]
**TargetResourceType** | **string** | Resource type of the target entity. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

