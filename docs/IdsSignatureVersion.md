# IdsSignatureVersion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Status** | **string** | This flag tells the status of the signatures under a version. OUTDATED: It means the signatures under this version are outdated and new version is available. LATEST: It means the signatures of this version are up to date.  | [optional] [default to null]
**ChangeLog** | **string** | Represents the version&#x27;s change log.  | [optional] [default to null]
**UpdateTime** | **int64** | Time when this version was downloaded and saved.  | [optional] [default to null]
**UserUploaded** | **bool** | Flag which tells whether te SIgnature version is uploaded by user or not.  | [optional] [default to null]
**State** | **string** | This flag tells which Version is currently active. ACTIVE: It means the signatures under this version is currently been used  under IDS Profiles. NOTACTIVE: It means signatures of this version are available but not  being used in IDS Profiles.  | [optional] [default to null]
**VersionId** | **string** | Represents the version id.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

