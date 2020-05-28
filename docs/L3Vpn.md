# L3Vpn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**RemotePrivateAddress** | **string** | This field is used to resolve conflicts in case of a remote site being behind NAT as remote public ip address is not enough. If it is not the case the remote public address should be provided here. If not provided, the value of this field is set to remote_public_address.  | [optional] [default to null]
**TunnelDigestAlgorithms** | **[]string** | Algorithm to be used for message digest during tunnel establishment. Default algorithm is empty.  | [optional] [default to null]
**Passphrases** | **[]string** | List of IPSec pre-shared keys used for IPSec authentication. If not specified, the older passphrase values are retained if there are any.  | [optional] [default to null]
**EnablePerfectForwardSecrecy** | **bool** | If true, perfect forward secrecy (PFS) is enabled.  | [optional] [default to true]
**IkeDigestAlgorithms** | **[]string** | Algorithm to be used for message digest during Internet Key Exchange(IKE) negotiation. Default is SHA2_256.  | [optional] [default to null]
**IkeVersion** | **string** | IKE protocol version to be used. IKE-Flex will initiate IKE-V2 and responds to both IKE-V1 and IKE-V2.  | [optional] [default to IKE_VERSION.V2]
**IkeEncryptionAlgorithms** | **[]string** | Algorithm to be used during Internet Key Exchange(IKE) negotiation. Default is AES_128.  | [optional] [default to null]
**LocalAddress** | **string** | IPv4 address of local gateway | [default to null]
**L3vpnSession** | [***L3VpnSession**](L3VpnSession.md) |  | [default to null]
**DhGroups** | **[]string** | Diffie-Hellman group to be used if PFS is enabled. Default group is GROUP14.  | [optional] [default to null]
**TunnelEncryptionAlgorithms** | **[]string** | Encryption algorithm to encrypt/decrypt the messages exchanged between IPSec VPN initiator and responder during tunnel negotiation. Default is AES_GCM_128.  | [optional] [default to null]
**Enabled** | **bool** | Flag to enable L3Vpn. Default is enabled.  | [optional] [default to true]
**RemotePublicAddress** | **string** | Public IPv4 address of remote gateway | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

