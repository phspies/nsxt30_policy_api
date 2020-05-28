# PolicyDraft

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**LockComments** | **string** | Comments for a policy draft lock/unlock. | [optional] [default to null]
**Locked** | **bool** | Indicates whether a draft should be locked. If the draft is locked by an user, then no other user would be able to modify or publish this draft. Once the user releases the lock, other users can then modify or publish this draft.  | [optional] [default to false]
**UserArea** | [***Infra**](Infra.md) |  | [optional] [default to null]
**LockModifiedBy** | **string** | ID of the user who last modified the lock for a policy draft.  | [optional] [default to null]
**LockModifiedTime** | **int64** | Policy draft locked/unlocked time in epoch milliseconds. | [optional] [default to null]
**RefDraftPath** | **string** | When specified, a manual draft will be created w.r.t. the specified draft. If not specified, manual draft will be created w.r.t. the current published configuration. For an auto draft, this will always be null.  | [optional] [default to null]
**SystemArea** | [***Infra**](Infra.md) |  | [optional] [default to null]
**IsAutoDraft** | **bool** | Flag to indicate whether draft is auto created. True indicates that the draft is an auto draft. False indicates that the draft is a manual draft.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

