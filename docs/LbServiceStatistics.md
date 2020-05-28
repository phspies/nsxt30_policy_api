# LbServiceStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**Pools** | [**[]LbPoolStatistics**](LBPoolStatistics.md) | Statistics of load balancer pools | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated. | [optional] [default to null]
**VirtualServers** | [**[]LbVirtualServerStatistics**](LBVirtualServerStatistics.md) | Statistics of load balancer virtual servers. | [optional] [default to null]
**ServicePath** | **string** | load balancer service identifier. | [optional] [default to null]
**Statistics** | [***LbServiceStatisticsCounter**](LBServiceStatisticsCounter.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

