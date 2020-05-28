# IpAddressPoolStaticSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Specifies whether the IpAddressPoolSubnet is to be carved out of a IpAddressBlock or will be specified by the user | [default to null]
**DnsNameservers** | **[]string** | The collection of upto 3 DNS servers for the subnet. | [optional] [default to null]
**Cidr** | **string** | Subnet representation is a network address and prefix length | [default to null]
**GatewayIp** | **string** | The default gateway address on a layer-3 router. | [optional] [default to null]
**AllocationRanges** | [**[]IpPoolRange**](IpPoolRange.md) | A collection of IPv4 or IPv6 IP Pool Ranges. | [default to null]
**DnsSuffix** | **string** | The DNS suffix for the DNS server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

