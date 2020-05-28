# L4PolicyLbPersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**Persistence** | **string** | This field indicates the persistence method used for the PolicyLbVirtualServer. - SOURCE_IP persistence ensures all connections from a client (identified by IP address) are sent to the same backend server for a specified period. - This object is not required and persistence is disabled by default  | [optional] [default to PERSISTENCE.IP]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

