# IdentityGroupInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | **string** | Each LDAP object is uniquely identified by its distinguished name (DN). A DN is a sequence of relative distinguished names (RDN) connected by commas. e.g. CN&#x3D;Larry Cole,CN&#x3D;admin,DC&#x3D;corp,DC&#x3D;acme,DC&#x3D;com. A valid fully qualified distinguished name should be provided to include specific groups else the create / update realization of the Group containing an invalid/ partial DN will fail. This value is valid only if it matches to exactly 1 LDAP object on the LDAP server.  | [default to null]
**DomainBaseDistinguishedName** | **string** | This is the base distinguished name for the domain where this particular group resides. (e.g. dc&#x3D;example,dc&#x3D;com) Each active directory domain has a domain naming context (NC), which contains domain-specific data. The root of this naming context is represented by a domain&#x27;s distinguished name (DN) and is typically referred to as the NC head.  | [default to null]
**Sid** | **string** | A security identifier (SID) is a unique value of variable length used to identify a trustee. A SID consists of the following components: The revision level of the SID structure; A 48-bit identifier authority value that identifies the authority that issued the SID; A variable number of subauthority or relative identifier (RID) values that uniquely identify the trustee relative to the authority that issued the SID. This field is only populated for Microsoft Active Directory identity store.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

