# IpfixSwitchCollectionInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IpfixCollectorProfilePaths** | **[]string** | Policy path for IPFIX collector profiles. IPFIX data from these logical segments will be sent to all specified IPFIX collectors.  | [default to null]
**IdleTimeout** | **int32** | The time in seconds after a Flow is expired if no more packets matching this Flow are received by the cache.  | [default to 300]
**SourceLogicalSegmentPaths** | **[]string** | Policy path for source tier-1 segment. IPFIX data from these logical segments will be sent IPFIX collector.  | [default to null]
**MaxFlows** | **int64** | The maximum number of flow entries in each exporter flow cache.  | [optional] [default to 16384]
**ObservationDomainId** | **int32** | An identifier that is unique to the exporting process and used to meter the Flows.  | [optional] [default to 0]
**ActiveTimeout** | **int32** | The time in seconds after a flow is expired even if more packets matching this flow are received by the cache.  | [default to 300]
**PacketSampleProbability** | **float64** | The probability in percentage that a packet is sampled, in range 0-100. The probability is equal for every packet.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

