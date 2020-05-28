# BgpRoutingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**InterSrIbgp** | **bool** | Flag to enable inter SR IBGP configuration. When not specified, inter SR IBGP is automatically enabled if Tier-0 is created in ACTIVE_ACTIVE ha_mode.  | [optional] [default to null]
**LocalAsNum** | **string** | Specify BGP AS number for Tier-0 to advertize to BGP peers. AS number can be specified in ASPLAIN (e.g., \&quot;65546\&quot;) or ASDOT (e.g., \&quot;1.10\&quot;) format. Empty string disables BGP feature. It is required by normal tier0 but not required in vrf tier0.  | [optional] [default to null]
**GracefulRestart** | **bool** | Flag to enable graceful restart. This field is deprecated, please use graceful_restart_config parameter for graceful restart configuration. If both parameters are set and consistent with each other (i.e. graceful_restart&#x3D;false and graceful_restart_mode&#x3D;HELPER_ONLY OR graceful_restart&#x3D;true and graceful_restart_mode&#x3D;GR_AND_HELPER) then this is allowed, but if inconsistent with each other then this is not allowed and validation error will be thrown.  | [optional] [default to null]
**RouteAggregations** | [**[]RouteAggregationEntry**](RouteAggregationEntry.md) | List of routes to be aggregated.  | [optional] [default to null]
**Enabled** | **bool** | Flag to enable BGP configuration. Disabling will stop feature and BGP peering.  | [optional] [default to null]
**GracefulRestartConfig** | [***BgpGracefulRestartConfig**](BgpGracefulRestartConfig.md) |  | [optional] [default to null]
**MultipathRelax** | **bool** | Flag to enable BGP multipath relax option. | [optional] [default to null]
**Ecmp** | **bool** | Flag to enable ECMP.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

