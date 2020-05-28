# IntersiteGatewayConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntersiteTransitSubnet** | **string** | IPv4 subnet for inter-site transit segment connecting service routers across sites for stretched gateway. For IPv6 link local subnet is auto configured.  | [optional] [default to 169.254.32.0/20]
**PrimarySitePath** | **string** | Primary egress site for gateway. T0/T1 gateway in Active/Standby mode supports stateful services on primary site. In this mode primary site must be set if gateway is stretched to more than one site. For T0 gateway in Active/Active primary site is optional field. If set then secondary site prefers routes learned from primary over locally learned routes. This field is not applicable for T1 gateway with no services.  | [optional] [default to null]
**LastAdminActiveEpoch** | **int64** | Epoch(in seconds) is auto updated based on system current timestamp when primary locale service is updated. It is used for resolving conflict during site failover. If system clock not in sync then User can optionally override this. New value must be higher than the current value.  | [optional] [default to null]
**FallbackSites** | **[]string** | Fallback site to be used as new primary site on current primary site failure. Disaster recovery must be initiated via API/UI. Fallback site configuration is supported only for T0 gateway. T1 gateway will follow T0 gateway&#x27;s primary site during disaster recovery.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

