# L7PolicyLbPersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**Persistence** | **string** | This field indicates the persistence method used for the PolicyLbVirtualServer. - COOKIE persistence allows related client connections, identified by the same cookie in HTTP requests [Refer to HTTP Cookie for details on HTTP cookies], to be redirected to the same server. Load balancer does not maintain any persistence table for cookie persistence. Instead, it encodes the necessary information in the HTTP cookie value sent to client and relies on the client to store it and send it back in subsequent related HTTP requests. Hence there is no limit on the number of cookie persistence entries that can be supported. - This object is not required and persistence is disabled by default  | [optional] [default to PERSISTENCE.COOKIE]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

