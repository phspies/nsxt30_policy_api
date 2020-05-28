# TraceflowConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Timeout** | **int64** | Maximum time in seconds the management plane will wait for observation result to be sent by opsAgent.  | [optional] [default to 10]
**IsTransient** | **bool** | This field indicates if intent is transient and will be cleaned up by the system if set to true | [optional] [default to true]
**Packet** | [***PacketData**](PacketData.md) |  | [default to null]
**SegmentPortPath** | **string** | Segment Port Path or UUID | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

