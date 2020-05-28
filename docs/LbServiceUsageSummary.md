# LbServiceUsageSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolUsagePercentage** | **float64** | Overall pool usage percentage for all load balancer services.  | [optional] [default to null]
**PoolCapacity** | **int64** | Pool capacity means maximum number of pools which can be configured for all load balancer services.  | [optional] [default to null]
**PoolMemberSeverity** | **string** | The severity calculation is based on the overall usage percentage of pool members for all load balancer services.  | [optional] [default to null]
**PoolMemberCapacity** | **int64** | Pool capacity means maximum number of pool members which can be configured for all load balancer services.  | [optional] [default to null]
**PoolMemberUsagePercentage** | **float64** | Overall pool member usage percentage for all load balancer services.  | [optional] [default to null]
**CurrentVirtualServerCount** | **int64** | The current count of virtual servers configured for all load balancer services.  | [optional] [default to null]
**CurrentPoolCount** | **int64** | The current count of pools configured for all load balancer services.  | [optional] [default to null]
**CurrentPoolMemberCount** | **int64** | The current count of pool members configured for all load balancer services.  | [optional] [default to null]
**PoolSeverity** | **string** | The severity calculation is based on the overall usage percentage of pools for all load balancer services.  | [optional] [default to null]
**ServiceUsages** | [**[]LbServiceUsage**](LBServiceUsage.md) | The property identifies all lb service usages. By default, it is not included in response. It exists when parameter ?include_usages&#x3D;true.  | [optional] [default to null]
**VirtualServerUsagePercentage** | **float64** | Overall virtual server usage percentage for all load balancer services.  | [optional] [default to null]
**VirtualServerSeverity** | **string** | The severity calculation is based on the overall usage percentage of virtual servers for all load balancer services.  | [optional] [default to null]
**ServiceCounts** | [**[]LbServiceCountPerSeverity**](LBServiceCountPerSeverity.md) | The service count for each load balancer usage severity.  | [optional] [default to null]
**VirtualServerCapacity** | **int64** | Virtual server capacity means maximum number of virtual servers which can be configured for all load balancer services.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

