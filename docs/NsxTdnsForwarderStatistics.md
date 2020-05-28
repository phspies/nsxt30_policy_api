# NsxTdnsForwarderStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**EnforcementPointPath** | **string** | Policy path referencing the enforcement point from where the statistics are fetched.  | [optional] [default to null]
**QueriesForwarded** | **int64** | The total number of forwarded DNS queries | [optional] [default to null]
**CachedEntries** | **int64** | The total number of cached entries | [optional] [default to null]
**DefaultForwarderStatistics** | [***NsxTdnsForwarderZoneStatistics**](NsxTDNSForwarderZoneStatistics.md) |  | [optional] [default to null]
**QueriesAnsweredLocally** | **int64** | The total number of queries answered from local cache | [optional] [default to null]
**UsedCacheStatistics** | [**[]NsxTPerNodeUsedCacheStatistics**](NsxTPerNodeUsedCacheStatistics.md) | The statistics of used cache | [optional] [default to null]
**ConfiguredCacheSize** | **int64** | The configured cache size, in kb | [optional] [default to null]
**Timestamp** | **int64** | Time stamp of the current statistics, in ms | [optional] [default to null]
**ConditionalForwarderStatistics** | [**[]NsxTdnsForwarderZoneStatistics**](NsxTDNSForwarderZoneStatistics.md) | The statistics of conditional forwarder zones | [optional] [default to null]
**TotalQueries** | **int64** | The total number of received DNS queries | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

