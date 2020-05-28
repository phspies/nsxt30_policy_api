# Tier0HaVipConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInterfacePaths** | **[]string** | Policy paths to Tier0 external interfaces which are to be paired to provide redundancy. Floating IP will be owned by one of these interfaces depending upon which edge node is Active. | [default to null]
**VipSubnets** | [**[]InterfaceSubnet**](InterfaceSubnet.md) | Array of IP address subnets which will be used as floating IP addresses. | [default to null]
**Enabled** | **bool** | Flag to enable this HA VIP config. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

