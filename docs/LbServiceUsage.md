# LbServiceUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**PoolCapacity** | **int64** | Pool capacity means maximum number of pools which could be configured in the given load balancer service.  | [optional] [default to null]
**ServiceSize** | **string** | The size of load balancer service. | [optional] [default to null]
**Severity** | **string** | The severity calculation is based on the largest usage percentage from virtual servers, pools and pool members for one load balancer service.  | [optional] [default to null]
**PoolMemberCapacity** | **int64** | Pool member capacity means maximum number of pool members which could be configured in the given load balancer service.  | [optional] [default to null]
**CurrentVirtualServerCount** | **int64** | The current number of virtual servers which has been configured in the given load balancer service.  | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated. | [optional] [default to null]
**CurrentPoolCount** | **int64** | The current number of pools which has been configured in the given load balancer service.  | [optional] [default to null]
**CurrentPoolMemberCount** | **int64** | The current number of pool members which has been configured in the given load balancer service.  | [optional] [default to null]
**UsagePercentage** | **float64** | The usage percentage is the largest usage percentage from virtual servers, pools and pool members for the load balancer service. If the property relax_scale_validation is set as true for LBService, it is possible that the value is larger than 100.0. For example, if SMALL LBS is deployed on MEDIUM edge node and configured with MEDIUM LBS virtual server scale number, LBS usage percentage is shown larger than 100.0.  | [optional] [default to null]
**VirtualServerCapacity** | **int64** | Virtual server capacity means maximum number of virtual servers which could be configured in the given load balancer service.  | [optional] [default to null]
**ServicePath** | **string** | LBService object path. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

