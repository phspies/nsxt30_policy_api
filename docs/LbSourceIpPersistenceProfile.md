# LbSourceIpPersistenceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistenceShared** | **bool** | Persistence shared setting indicates that all LBVirtualServers that consume this LBPersistenceProfile should share the same persistence mechanism when enabled.  Meaning, persistence entries of a client accessing one virtual server will also affect the same client&#x27;s connections to a different virtual server. For example, say there are two virtual servers vip-ip1:80 and vip-ip1:8080 bound to the same Group g1 consisting of two servers (s11:80 and s12:80). By default, each virtual server will have its own persistence table or cookie. So, in the earlier example, there will be two tables (vip-ip1:80, p1) and (vip-ip1:8080, p1) or cookies. So, if a client connects to vip1:80 and later connects to vip1:8080, the second connection may be sent to a different server than the first.  When persistence_shared is enabled, then the second connection will always connect to the same server as the original connection. For COOKIE persistence type, the same cookie will be shared by multiple virtual servers. For SOURCE_IP persistence type, the persistence table will be shared across virtual servers. For GENERIC persistence type, the persistence table will be shared across virtual servers which consume the same persistence profile in LBRule actions.  | [optional] [default to false]
**ResourceType** | **string** | The resource_type property identifies persistence profile type.  | [default to null]
**Purge** | **string** | Persistence purge setting. | [optional] [default to PURGE.FULL]
**HaPersistenceMirroringEnabled** | **bool** | Persistence entries are not synchronized to the HA peer by default.  | [optional] [default to false]
**Timeout** | **int64** | When all connections complete (reference count reaches 0), persistence entry timer is started with the expiration time.  | [optional] [default to 300]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

