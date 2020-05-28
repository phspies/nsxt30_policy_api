# SegmentSubnet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpConfig** | [***SegmentDhcpConfig**](SegmentDhcpConfig.md) |  | [optional] [default to null]
**GatewayAddress** | **string** | Gateway IP address in CIDR format for both IPv4 and IPv6.  | [optional] [default to null]
**DhcpRanges** | **[]string** | DHCP address ranges are used for dynamic IP allocation. Supports address range and CIDR formats. First valid host address from the first value is assigned to DHCP server IP address. Existing values cannot be deleted or modified, but additional DHCP ranges can be added.  | [optional] [default to null]
**Network** | **string** | Network CIDR for this subnet calculated from gateway_addresses and prefix_len.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

