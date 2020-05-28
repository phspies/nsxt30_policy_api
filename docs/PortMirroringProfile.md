# PortMirroringProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Direction** | **string** | Port mirroring profile direction | [optional] [default to DIRECTION.BIDIRECTIONAL]
**ProfileType** | **string** | Allows user to select type of port mirroring session. | [optional] [default to PROFILE_TYPE.REMOTE_L3_SPAN]
**SnapLength** | **int32** | If this property is set, the packet will be truncated to the provided length. If this property is unset, entire packet will be mirrored.  | [optional] [default to null]
**EncapsulationType** | **string** | User can provide Mirror Destination type e.g GRE, ERSPAN_TWO or ERSPAN_THREE.If profile type is REMOTE_L3_SPAN, encapsulation type is used else ignored. | [optional] [default to ENCAPSULATION_TYPE.GRE]
**ErspanId** | **int32** | Used by physical switch for the mirror traffic forwarding. Must be provided and only effective when encapsulation type is ERSPAN type II or type III.  | [optional] [default to 0]
**GreKey** | **int32** | User-configurable 32-bit key only for GRE | [optional] [default to 0]
**DestinationGroup** | **string** | Data from source group will be copied to members of destination group. Only IPSET group and group with membership criteria VM is supported. IPSET group allows only three ip&#x27;s.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

