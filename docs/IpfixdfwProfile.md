# IpfixdfwProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IpfixDfwCollectorProfilePath** | **string** | Policy path for IPFIX collector profiles. IPFIX data from these logical segments will be sent to all specified IPFIX collectors.  | [default to null]
**Priority** | **int32** | This priority field is used to resolve conflicts in Segment Ports which are covered by more than one IPFIX profiles. The IPFIX exporter will send records to Collectors in highest priority profile (lowest number) only.  | [optional] [default to 0]
**ActiveFlowExportTimeout** | **int32** | For long standing active flows, IPFIX records will be sent per timeout period in minutes.  | [default to 1]
**ObservationDomainId** | **int32** | An identifier that is unique to the exporting process and used to meter the flows.  | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

