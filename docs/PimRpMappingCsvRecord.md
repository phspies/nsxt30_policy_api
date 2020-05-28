# PimRpMappingCsvRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRp** | **bool** | Value of this field will be true if this edge transport node acts as rendezvous point, otherwise false.  | [optional] [default to null]
**TransportNode** | **string** | Transport node uuid or policy path. | [optional] [default to null]
**Group** | **string** | Multicast group address. | [optional] [default to null]
**Source** | **string** | Source of learning RP information. Either Static RP configured or RP learned via BSR (Bootstrap Router).  | [optional] [default to null]
**OutgoingInterface** | **string** | Outgoing/Egress interface for multicast traffic. | [optional] [default to null]
**RpAddress** | **string** | RP (Randezvous Point) address. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

