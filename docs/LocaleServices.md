# LocaleServices

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**PreferredEdgePaths** | **[]string** | Policy paths to edge nodes. Specified edge is used as preferred edge cluster member when failover mode is set to PREEMPTIVE, not applicable otherwise.  | [optional] [default to null]
**HaVipConfigs** | [**[]Tier0HaVipConfig**](Tier0HaVipConfig.md) | This configuration can be defined only for Active-Standby Tier0 gateway to provide redundancy. For mulitple external interfaces, multiple HA VIP configs must be defined and each config will pair exactly two external interfaces. The VIP will move and will always be owned by the Active node. When this property is configured, configuration of dynamic-routing is not allowed. | [optional] [default to null]
**EdgeClusterPath** | **string** | Policy path to edge cluster. Auto-assigned on Tier0 if associated enforcement-point has only one edge cluster.  | [optional] [default to null]
**RouteRedistributionConfig** | [***Tier0RouteRedistributionConfig**](Tier0RouteRedistributionConfig.md) |  | [optional] [default to null]
**RouteRedistributionTypes** | **[]string** | Enable redistribution of different types of routes on Tier-0. This property is only valid for locale-service under Tier-0. This property is deprecated, please use \&quot;route_redistribution_config\&quot; property to configure redistribution rules.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

