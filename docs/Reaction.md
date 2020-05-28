# Reaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Events** | [**[]Event**](Event.md) | Events that provide contextual variables about what the reaction should react to. This field can be interpreted as the WHAT of the Reaction, or simply as \&quot;If This\&quot; Clause.  | [default to null]
**Actions** | [**[]Action**](Action.md) | Actions that need to be taken when the events occur. These actions must appear in the order that they need to be taken in. This field can be interpreted as the HOW of the Reaction, or simply as \&quot;Then That\&quot;.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

