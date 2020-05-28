# IPv6DiscoveryOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NdSnoopingConfig** | [***NdSnoopingConfig**](NdSnoopingConfig.md) |  | [optional] [default to null]
**DhcpSnoopingV6Enabled** | **bool** | Enable this method will snoop the DHCPv6 message transaction which a VM makes with a DHCPv6 server. From the transaction, we learn the IPv6 addresses assigned by the DHCPv6 server to this VM along with its lease time.  | [optional] [default to false]
**VmtoolsV6Enabled** | **bool** | Enable this method will learn the IPv6 addresses which are configured on interfaces of a VM with the help of the VMTools software.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

