# SiteInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SitePath** | **string** | For the local manager this needs to be set to &#x27;default&#x27;. This represents the path of the site which is managed by Global Manager.  | [default to null]
**TransportZonePaths** | **[]string** | The transport zone has to be set when creating VHC on Local manager. If not set for local manager, default transport zone will be used. For the Global Manager the transport zone path will be picked up from the site.  | [optional] [default to null]
**EdgeClusterPaths** | **[]string** | The edge cluster on which the networking elements for the VHC will be created.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

