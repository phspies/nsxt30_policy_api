# IpAddressPoolBlockSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Specifies whether the IpAddressPoolSubnet is to be carved out of a IpAddressBlock or will be specified by the user | [default to null]
**IpBlockPath** | **string** | The path of the IpAddressBlock from which the subnet is to be created. | [default to null]
**StartIp** | **string** | For internal system use Only. Represents start ip address of the subnet from IP block. Subnet ip adddress will start from this ip address. | [optional] [default to null]
**AutoAssignGateway** | **bool** | If this property is set to true, the first IP in the range will be reserved for gateway. | [optional] [default to true]
**Size** | **int64** | The size parameter is required for subnet creation. It must be specified during creation but cannot be changed later. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

