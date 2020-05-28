# ContainerProject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**NetworkStatus** | **string** | Network status of container project. | [optional] [default to null]
**ContainerClusterId** | **string** | Identifier of the container cluster to which this project/namespace belongs. | [optional] [default to null]
**OriginProperties** | [**[]KeyValuePair**](KeyValuePair.md) | Array of additional specific properties of container project in key-value format.  | [optional] [default to null]
**ExternalId** | **string** | External identifier of the container project. | [default to null]
**NetworkErrors** | [**[]NetworkError**](NetworkError.md) | List of network errors related to container project. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
