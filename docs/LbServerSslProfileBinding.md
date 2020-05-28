# LbServerSslProfileBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAuthCaPaths** | **[]string** | If server auth type is REQUIRED, server certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified.  | [optional] [default to null]
**ClientCertificatePath** | **string** | To support client authentication (load balancer acting as a client authenticating to the backend server), client certificate can be specified in the server-side SSL profile binding  | [optional] [default to null]
**ServerAuth** | **string** | Server authentication mode. | [optional] [default to SERVER_AUTH.AUTO_APPLY]
**CertificateChainDepth** | **int64** | Authentication depth is used to set the verification depth in the server certificates chain.  | [optional] [default to 3]
**ServerAuthCrlPaths** | **[]string** | A Certificate Revocation List (CRL) can be specified in the server-side SSL profile binding to disallow compromised server certificates.  | [optional] [default to null]
**SslProfilePath** | **string** | Server SSL profile defines reusable, application-independent server side SSL properties.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

