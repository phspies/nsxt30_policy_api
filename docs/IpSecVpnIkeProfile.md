# IpSecVpnIkeProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**DigestAlgorithms** | **[]string** | Algorithm to be used for message digest during Internet Key Exchange(IKE) negotiation. Default is SHA2_256. | [optional] [default to null]
**EncryptionAlgorithms** | **[]string** | Encryption algorithm is used during Internet Key Exchange(IKE) negotiation. Default is AES_128. | [optional] [default to null]
**DhGroups** | **[]string** | Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14. | [optional] [default to null]
**SaLifeTime** | **int64** | Life time for security association. Default is 86400 seconds (1 day). | [optional] [default to 86400]
**IkeVersion** | **string** | IKE protocol version to be used. IKE-Flex will initiate IKE-V2 and responds to both IKE-V1 and IKE-V2. | [optional] [default to IKE_VERSION.V2]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

