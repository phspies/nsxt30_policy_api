# Tier0

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**FederationConfig** | [***FederationGatewayConfig**](FederationGatewayConfig.md) |  | [optional] [default to null]
**DefaultRuleLogging** | **bool** | Indicates if logging should be enabled for the default whitelisting rule. This filed is deprecated and recommended to change Rule logging filed. Note that this filed is not synchornied with default logging field.  | [optional] [default to false]
**FailoverMode** | **string** | Determines the behavior when a Tier-0 instance in ACTIVE-STANDBY high-availability mode restarts after a failure. If set to PREEMPTIVE, the preferred node will take over, even if it causes another failure. If set to NON_PREEMPTIVE, then the instance that restarted will remain secondary. This property is not used when the ha_mode property is set to ACTIVE_ACTIVE. Only applicable when edge cluster is configured in Tier0 locale-service.  | [optional] [default to FAILOVER_MODE.NON_PREEMPTIVE]
**VrfConfig** | [***Tier0VrfConfig**](Tier0VrfConfig.md) |  | [optional] [default to null]
**DisableFirewall** | **bool** | Disable or enable gateway fiewall. | [optional] [default to false]
**InternalTransitSubnets** | **[]string** | Specify subnets that are used to assign addresses to logical links connecting service routers and distributed routers. Only IPv4 addresses are supported. When not specified, subnet 169.254.0.0/24 is assigned by default in ACTIVE_ACTIVE HA mode or 169.254.0.0/28 in ACTIVE_STANDBY mode.  | [optional] [default to null]
**TransitSubnets** | **[]string** | Specify transit subnets that are used to assign addresses to logical links connecting tier-0 and tier-1s. Both IPv4 and IPv6 addresses are supported. When not specified, subnet 100.64.0.0/16 is configured by default.  | [optional] [default to null]
**HaMode** | **string** | Specify high-availability mode for Tier-0. Default is ACTIVE_ACTIVE. When ha_mode is changed from ACTIVE_ACTIVE to ACTIVE_STANDBY, inter SR iBGP (in BGP) is disabled. Changing ha_mode from ACTIVE_STANDBY to ACTIVE_ACTIVE will enable inter SR iBGP (in BGP) and previously configured preferred edge nodes (in Tier0 locale-service) are removed.  | [optional] [default to HA_MODE.ACTIVE]
**ForceWhitelisting** | **bool** | This filed is deprecated and recommended to change Rule action filed. Note that this filed is not synchornied with default rule field.  | [optional] [default to false]
**IntersiteConfig** | [***IntersiteGatewayConfig**](IntersiteGatewayConfig.md) |  | [optional] [default to null]
**RdAdminField** | **string** | If you are using EVPN service, then route distinguisher administrator address should be defined if you need auto generation of route distinguisher on your VRF configuration.  | [optional] [default to null]
**Ipv6ProfilePaths** | **[]string** | IPv6 NDRA and DAD profiles configuration on Tier0. Either or both NDRA and/or DAD profiles can be configured.  | [optional] [default to null]
**DhcpConfigPaths** | **[]string** | DHCP configuration for Segments connected to Tier-0. DHCP service is configured in relay mode.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

