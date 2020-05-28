# MetadataProxyConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Secret** | **string** | Secret word or phrase to access metadata server. | [default to null]
**ServerCertificates** | **[]string** | Valid certificates should be configured. The validity of certificates is not checked. Certificates are managed through /infra/certificates API on Policy. | [optional] [default to null]
**EdgeClusterPath** | **string** | Edge clusters configured on MP are auto-discovered by Policy and create corresponding read-only intent objects. | [default to null]
**PreferredEdgePaths** | **[]string** | Edge nodes should be members of edge cluster configured in edge_cluster_path. | [optional] [default to null]
**CryptoProtocols** | **[]string** | The cryptographic protocols listed here are supported by the metadata proxy. TLSv1.1 and TLSv1.2 are supported by default | [optional] [default to null]
**ServerAddress** | **string** | This field is a URL. Example formats - http://1.2.3.4:3888/path, http://text-md-proxy:5001/. Port number should be between 3000-9000. | [default to null]
**EnableStandbyRelocation** | **bool** | Only auto-placed metadata proxies are considered for relocation. Must be FALSE, when the preferred_edge_paths property is configured. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

