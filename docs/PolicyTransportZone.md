# PolicyTransportZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IsDefault** | **bool** | Flag to indicate if the transport zone is the default one. Only one transport zone can be the default one for a given transport zone type.  | [optional] [default to false]
**TzType** | **string** | Transport Zone Type.  | [optional] [default to null]
**NsxId** | **string** | UUID of transport zone on NSX-T enforcement point. | [optional] [default to null]
**UplinkTeamingPolicyNames** | **[]string** | The names of switching uplink teaming policies that all transport nodes in this transport zone support. Uplinkin teaming policies are only valid for VLAN backed transport zones. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

