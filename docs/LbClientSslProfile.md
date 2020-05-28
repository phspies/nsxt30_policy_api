# LbClientSslProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionCacheEnabled** | **bool** | SSL session caching allows SSL client and server to reuse previously negotiated security parameters avoiding the expensive public key operation during handshake.  | [optional] [default to true]
**SessionCacheTimeout** | **int64** | Session cache timeout specifies how long the SSL session parameters are held on to and can be reused.  | [optional] [default to 300]
**CipherGroupLabel** | **string** | It is a label of cipher group which is mostly consumed by GUI.  | [optional] [default to null]
**IsFips** | **bool** | This flag is set to true when all the ciphers and protocols are FIPS compliant. It is set to false when one of the ciphers or protocols are not FIPS compliant..  | [optional] [default to null]
**IsSecure** | **bool** | This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure.  | [optional] [default to null]
**PreferServerCiphers** | **bool** | During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client&#x27;s preference.  | [optional] [default to true]
**Ciphers** | **[]string** | Supported SSL cipher list to client side. | [optional] [default to null]
**Protocols** | **[]string** | SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

