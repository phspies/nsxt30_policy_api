# PortAttachment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficTag** | **int64** | Not valid when type field is INDEPENDENT, mainly used to identify traffic from different ports in container use case.  | [optional] [default to null]
**AllocateAddresses** | **string** | Indicate how IP will be allocated for the port | [optional] [default to null]
**HyperbusMode** | **string** | Flag to indicate if hyperbus configuration is required. | [optional] [default to HYPERBUS_MODE.DISABLE]
**ContextType** | **string** | Set to PARENT when type field is CHILD. Read only field. | [optional] [default to null]
**ContextId** | **string** | If type is CHILD and the parent port is on the same segment as the child port, then this field should be VIF ID of the parent port. If type is CHILD and the parent port is on a different segment, then this field should be policy path of the parent port. If type is INDEPENDENT/STATIC, then this field should be transport node ID.  | [optional] [default to null]
**Type_** | **string** | Type of port attachment. STATIC is added to replace INDEPENDENT. INDEPENDENT type and PARENT type are deprecated. | [optional] [default to null]
**AppId** | **string** | ID used to identify/look up a child attachment behind a parent attachment  | [optional] [default to null]
**Id** | **string** | VIF UUID on NSX Manager. If the attachement type is PARENT, this property is required. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

