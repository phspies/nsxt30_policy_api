# LbNodeUsageSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarm** | [***PolicyRuntimeAlarm**](PolicyRuntimeAlarm.md) |  | [optional] [default to null]
**EnforcementPointPath** | **string** | Policy Path referencing the enforcement point where the info is fetched.  | [optional] [default to null]
**CurrentLoadBalancerCredits** | **int64** | Current load balancer credits in use for all nodes. For example, configuring a medium load balancer on a node consumes 10 credits. If there are 2 medium instances configured, the current load balancer credit number is 2 * 10 &#x3D; 20.  | [optional] [default to null]
**LoadBalancerCreditCapacity** | **int64** | The load balancer credit capacity means the maximum credits which can be used for load balancer service configuration for all nodes.  | [optional] [default to null]
**NodeUsages** | [**[]LbNodeUsage**](LBNodeUsage.md) | The property identifies all LB node usages. By default, it is not included in response. It exists when parameter ?include_usages&#x3D;true.  | [optional] [default to null]
**Severity** | **string** | The severity calculation is based on overall credit usage percentage of load balancer for all nodes.  | [optional] [default to null]
**PoolMemberCapacity** | **int64** | Pool member capacity means maximum number of pool members which can be configured on all nodes.  | [optional] [default to null]
**NodeCounts** | [**[]LbNodeCountPerSeverity**](LBNodeCountPerSeverity.md) | The property identifies array of node count for each severity.  | [optional] [default to null]
**CurrentPoolMemberCount** | **int64** | The overall count of pool members configured on all nodes.  | [optional] [default to null]
**UsagePercentage** | **float64** | The overall usage percentage of all nodes for load balancer. The value is the larger value between overall pool member usage percentage and overall load balancer credit usage percentage.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

