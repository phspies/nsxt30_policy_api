# Segment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Subnets** | [**[]SegmentSubnet**](SegmentSubnet.md) | Subnet configuration. Max 1 subnet | [optional] [default to null]
**ConnectivityPath** | **string** | Policy path to the connecting Tier-0 or Tier-1. Valid only for segments created under Infra.  | [optional] [default to null]
**ExtraConfigs** | [**[]SegmentExtraConfig**](SegmentExtraConfig.md) | This property could be used for vendor specific configuration in key value string pairs, the setting in extra_configs will be automatically inheritted by segment ports in the Segment.  | [optional] [default to null]
**AdvancedConfig** | [***SegmentAdvancedConfig**](SegmentAdvancedConfig.md) |  | [optional] [default to null]
**AddressBindings** | [**[]PortAddressBindingEntry**](PortAddressBindingEntry.md) | Static address binding used for the Segment. | [optional] [default to null]
**FederationConfig** | [***FederationConnectivityConfig**](FederationConnectivityConfig.md) |  | [optional] [default to null]
**MacPoolId** | **string** | Mac pool id that associated with a Segment. | [optional] [default to null]
**MetadataProxyPaths** | **[]string** | Policy path to metadata proxy configuration. Multiple distinct MD proxies can be configured. | [optional] [default to null]
**VlanIds** | **[]string** | VLAN ids for a VLAN backed Segment. Can be a VLAN id or a range of VLAN ids specified with &#x27;-&#x27; in between.  | [optional] [default to null]
**DomainName** | **string** | DNS domain name | [optional] [default to null]
**OverlayId** | **int32** | Used for overlay connectivity of segments. The overlay_id should be allocated from the pool as definied by enforcement-point. If not provided, it is auto-allocated from the default pool on the enforcement-point.  | [optional] [default to null]
**DhcpConfigPath** | **string** | Policy path to DHCP server or relay configuration to use for all IPv4 &amp; IPv6 subnets configured on this segment.  | [optional] [default to null]
**LsId** | **string** | This property is deprecated. The property will continue to work as expected for existing segments. The segments that are newly created with ls_id will be ignored. Sepcify pre-creted logical switch id for Segment.  | [optional] [default to null]
**L2Extension** | [***L2Extension**](L2Extension.md) |  | [optional] [default to null]
**AdminState** | **string** | Represents Desired state of the Segment | [optional] [default to ADMIN_STATE.UP]
**BridgeProfiles** | [**[]BridgeProfileConfig**](BridgeProfileConfig.md) | Multiple distinct L2 bridge profiles can be configured. | [optional] [default to null]
**Type_** | **string** | Segment type based on configuration.  | [optional] [default to null]
**TransportZonePath** | **string** | Policy path to the transport zone. Supported for VLAN backed segments as well as Overlay Segments. This field is required for VLAN backed Segments. Auto assigned if only one transport zone exists in the enforcement point. Default transport zone is auto assigned for overlay segments if none specified.  | [optional] [default to null]
**ReplicationMode** | **string** | If this field is not set for overlay segment, then the default of MTEP will be used.  | [optional] [default to REPLICATION_MODE.MTEP]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

