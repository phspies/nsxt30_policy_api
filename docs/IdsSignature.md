# IdsSignature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Cvssv3** | **string** | Signature cvssv3 score.  | [optional] [default to null]
**Cvssv2** | **string** | Signature cvssv2 score.  | [optional] [default to null]
**ClassType** | **string** | Class type of Signature.  | [optional] [default to null]
**SignatureRevision** | **string** | Represents revision of the Signature.  | [optional] [default to null]
**ProductAffected** | **string** | Product affected by this signature.  | [optional] [default to null]
**Flow** | **string** | Flow established from server, from client etc.  | [optional] [default to null]
**Severity** | **string** | Represents the severity of the Signature.  | [optional] [default to null]
**Urls** | **[]string** | List of mitre attack URLs pertaining to signature  | [optional] [default to null]
**AttackTarget** | **string** | Target of the signature.  | [optional] [default to null]
**SignatureId** | **string** | Represents the Signature&#x27;s id.  | [optional] [default to null]
**Cves** | **[]string** | CVE score  | [optional] [default to null]
**Categories** | **[]string** | Represents the internal categories a signature belongs to.  | [optional] [default to null]
**Name** | **string** | Signature name.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

