# LbEdgeNodeUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The property identifies the load balancer node usage type.  | [default to null]
**NodePath** | **string** | The property identifies the node path for load balancer node usage. For example, node_path&#x3D;/infra/sites/default/enforcement-points/default /edge-clusters/85175e0b-4d74-461d-83e1-f3b785adef9c/edge-nodes /86e077c0-449f-11e9-87c8-02004eb37029.  | [default to null]
**CurrentMediumLoadBalancerCount** | **int64** | The count of medium load balancer services configured on the node.  | [optional] [default to null]
**RemainingXlargeLoadBalancerCount** | **int64** | The remaining count of xlarge load balancer services which can be configured on the given edge node.  | [optional] [default to null]
**Severity** | **string** | The severity calculation is based on current credit usage percentage of load balancer for one node.  | [optional] [default to null]
**PoolMemberCapacity** | **int64** | Pool member capacity means maximum number of pool members which can be configured on the given edge node.  | [optional] [default to null]
**CurrentVirtualServerCount** | **int64** | The count of virtual servers configured on the node.  | [optional] [default to null]
**CurrentXlargeLoadBalancerCount** | **int64** | The count of xlarge load balancer services configured on the node.  | [optional] [default to null]
**CurrentPoolCount** | **int64** | The count of pools configured on the node.  | [optional] [default to null]
**RemainingSmallLoadBalancerCount** | **int64** | The remaining count of small load balancer services which can be configured on the given edge node.  | [optional] [default to null]
**CurrentPoolMemberCount** | **int64** | The count of pool members configured on the node.  | [optional] [default to null]
**LoadBalancerCreditCapacity** | **int64** | The load balancer credit capacity means the maximum credits which can be used for load balancer configuration for the given edge node.  | [optional] [default to null]
**EdgeClusterPath** | **string** | The path of edge cluster which contains the edge node.  | [optional] [default to null]
**CurrentLoadBalancerCredits** | **int64** | The current load balancer credits means the current credits used on the node. For example, configuring a medium load balancer on a node consumes 10 credits. If there are 2 medium instances configured on a node, the current credit number is 2 * 10 &#x3D; 20.  | [optional] [default to null]
**RemainingLargeLoadBalancerCount** | **int64** | The remaining count of large load balancer services which can be configured on the given edge node.  | [optional] [default to null]
**CurrentLargeLoadBalancerCount** | **int64** | The count of large load balancer services configured on the node.  | [optional] [default to null]
**FormFactor** | **string** | The form factor of the given edge node.  | [optional] [default to null]
**UsagePercentage** | **float64** | The usage percentage of the edge node for load balancer. The value is the larger value between load balancer credit usage percentage and pool member usage percentage for the edge node.  | [optional] [default to null]
**CurrentSmallLoadBalancerCount** | **int64** | The count of small load balancer services configured on the node.  | [optional] [default to null]
**RemainingMediumLoadBalancerCount** | **int64** | The remaining count of medium load balancer services which can be configured on the given edge node.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

