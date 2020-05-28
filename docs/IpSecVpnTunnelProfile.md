# IpSecVpnTunnelProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ExtendedAttributes** | [**[]AttributeVal**](AttributeVal.md) | Collection of type specific properties. As of now, to hold encapsulation mode and transform protocol.  | [optional] [default to null]
**DigestAlgorithms** | **[]string** | Algorithm to be used for message digest. Default digest algorithm is implicitly covered by default encryption algorithm \&quot;AES_GCM_128\&quot;. | [optional] [default to null]
**EncryptionAlgorithms** | **[]string** | Encryption algorithm to encrypt/decrypt the messages exchanged between IPSec VPN initiator and responder during tunnel negotiation. Default is AES_GCM_128. | [optional] [default to null]
**EnablePerfectForwardSecrecy** | **bool** | If true, perfect forward secrecy (PFS) is enabled. | [optional] [default to true]
**DhGroups** | **[]string** | Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14. | [optional] [default to null]
**DfPolicy** | **string** | Defragmentation policy helps to handle defragmentation bit present in the inner packet. COPY copies the defragmentation bit from the inner IP packet into the outer packet. CLEAR ignores the defragmentation bit present in the inner packet. | [optional] [default to DF_POLICY.COPY]
**SaLifeTime** | **int64** | SA life time specifies the expiry time of security association. Default is 3600 seconds. | [optional] [default to 3600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

