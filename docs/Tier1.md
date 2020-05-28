# Tier1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**FederationConfig** | [***FederationGatewayConfig**](FederationGatewayConfig.md) |  | [optional] [default to null]
**DefaultRuleLogging** | **bool** | Indicates if logging should be enabled for the default whitelisting rule. This filed is deprecated and recommended to change Rule logging filed. Note that this filed is not synchornied with default logging field.  | [optional] [default to false]
**RouteAdvertisementRules** | [**[]RouteAdvertisementRule**](RouteAdvertisementRule.md) | Route advertisement rules and filtering | [optional] [default to null]
**PoolAllocation** | **string** | Supports edge node allocation at different sizes for routing and load balancer service to meet performance and scalability requirements.   ROUTING: Allocate edge node to provide routing services.   LB_SMALL, LB_MEDIUM, LB_LARGE, LB_XLARGE: Specify size of load balancer service that will be configured on TIER1 gateway.  | [optional] [default to POOL_ALLOCATION.ROUTING]
**Tier0Path** | **string** | Specify Tier-1 connectivity to Tier-0 instance.  | [optional] [default to null]
**EnableStandbyRelocation** | **bool** | Flag to enable standby service router relocation. Standby relocation is not enabled until edge cluster is configured for Tier1.  | [optional] [default to false]
**DisableFirewall** | **bool** | Disable or enable gateway fiewall. | [optional] [default to false]
**FailoverMode** | **string** | Determines the behavior when a Tier-1 instance restarts after a failure. If set to PREEMPTIVE, the preferred node will take over, even if it causes another failure. If set to NON_PREEMPTIVE, then the instance that restarted will remain secondary. Only applicable when edge cluster is configured in Tier1 locale-service.  | [optional] [default to FAILOVER_MODE.NON_PREEMPTIVE]
**ForceWhitelisting** | **bool** | This filed is deprecated and recommended to change Rule action filed. Note that this filed is not synchornied with default rule field.  | [optional] [default to false]
**IntersiteConfig** | [***IntersiteGatewayConfig**](IntersiteGatewayConfig.md) |  | [optional] [default to null]
**DhcpConfigPaths** | **[]string** | DHCP configuration for Segments connected to Tier-1. DHCP service is enabled in relay mode.  | [optional] [default to null]
**Ipv6ProfilePaths** | **[]string** | Configuration IPv6 NDRA and DAD profiles. Either or both NDRA and/or DAD profiles can be configured.  | [optional] [default to null]
**QosProfile** | [***GatewayQosProfileConfig**](GatewayQosProfileConfig.md) |  | [optional] [default to null]
**Type_** | **string** | Tier1 connectivity type for reference. Property value is not validated with Tier1 configuration.   ROUTED: Tier1 is connected to Tier0 gateway and routing is enabled.   ISOLATED: Tier1 is not connected to any Tier0 gateway.   NATTED: Tier1 is in ROUTED type with NAT configured locally.  | [optional] [default to null]
**RouteAdvertisementTypes** | **[]string** | Enable different types of route advertisements. When not specified, routes to IPSec VPN local-endpoint subnets (TIER1_IPSEC_LOCAL_ENDPOINT) are automatically advertised.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

