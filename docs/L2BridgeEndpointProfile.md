# L2BridgeEndpointProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**FailoverMode** | **string** | Failover mode for the edge bridge cluster | [optional] [default to FAILOVER_MODE.PREEMPTIVE]
**HaMode** | **string** | High avaialability mode can be active-active or active-standby. High availability mode cannot be modified after realization. | [optional] [default to HA_MODE.STANDBY]
**EdgePaths** | **[]string** | List of policy paths to edge nodes. Edge allocation for L2 bridging. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

