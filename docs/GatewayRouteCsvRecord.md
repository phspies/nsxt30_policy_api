# GatewayRouteCsvRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LrComponentId** | **string** | Logical router component(Service Router/Distributed Router) id | [optional] [default to null]
**NextHop** | **string** | The IP of the next hop | [optional] [default to null]
**LrComponentType** | **string** | Logical router component(Service Router/Distributed Router) type | [optional] [default to null]
**Network** | **string** | CIDR network address | [default to null]
**RouteType** | **string** | Route type (USER, CONNECTED, NSX_INTERNAL,..) | [default to null]
**Interface_** | **string** | The policy path of the interface which is used as the next hop | [optional] [default to null]
**EdgePath** | **string** | Edge node policy path.  | [optional] [default to null]
**AdminDistance** | **int64** | The admin distance of the next hop | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

