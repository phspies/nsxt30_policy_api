# SegmentAdvancedConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectivity** | **string** | Connectivity configuration to manually connect (ON) or disconnect (OFF) a logical entity from network topology. Only valid for Tier1 Segment.  | [optional] [default to CONNECTIVITY.ON]
**LocalEgressRoutingPolicies** | [**[]LocalEgressRoutingEntry**](LocalEgressRoutingEntry.md) | An ordered list of routing policies to forward traffic to the next hop.  | [optional] [default to null]
**Multicast** | **bool** | Enable multicast for a segment. Only applicable for segments connected to Tier0 gateway.  | [optional] [default to null]
**UplinkTeamingPolicyName** | **string** | The name of the switching uplink teaming policy for the Segment. This name corresponds to one of the switching uplink teaming policy names listed in TransportZone associated with the Segment. See transport_zone_path property above for more details. When this property is not specified, the segment will not have a teaming policy associated with it and the host switch&#x27;s default teaming policy will be used by MP. | [optional] [default to null]
**AddressPoolPaths** | **[]string** | Policy path to IP address pools.  | [optional] [default to null]
**Hybrid** | **bool** | When set to true, all the ports created on this segment will behave in a hybrid fashion. The hybrid port indicates to NSX that the VM intends to operate in underlay mode, but retains the ability to forward egress traffic to the NSX overlay network. This property is only applicable for segment created with transport zone type OVERLAY_STANDARD. This property cannot be modified after segment is created.  | [optional] [default to false]
**LocalEgress** | **bool** | This property is used to enable proximity routing with local egress. When set to true, logical router interface (downlink) connecting Segment to Tier0/Tier1 gateway is configured with prefix-length 32.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

