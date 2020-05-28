# SegmentDhcpV6Config

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAddress** | **string** | IP address of the DHCP server in CIDR format. The server_address is mandatory in case this segment has provided a dhcp_config_path and it represents a DHCP server config. If this SegmentDhcpConfig is a SegmentDhcpV4Config, the address must be an IPv4 address. If this is a SegmentDhcpV6Config, the address must be an IPv6 address. This address must not overlap the ip-ranges of the subnet, or the gateway address of the subnet, or the DHCP static-binding addresses of this segment.  | [optional] [default to null]
**DnsServers** | **[]string** | IP address of DNS servers for subnet. DNS server IP address must belong to the same address family as segment gateway_address property.  | [optional] [default to null]
**LeaseTime** | **int64** | DHCP lease time in seconds. When specified, this property overwrites lease time configured DHCP server config.  | [optional] [default to 86400]
**ResourceType** | **string** |  | [default to null]
**ExcludedRanges** | **[]string** | Excluded addresses to define dynamic ip allocation ranges. | [optional] [default to null]
**SntpServers** | **[]string** | IPv6 address of SNTP servers for subnet.  | [optional] [default to null]
**PreferredTime** | **int64** | The length of time that a valid address is preferred. When the preferred lifetime expires, the address becomes deprecated.  | [optional] [default to null]
**DomainNames** | **[]string** | Domain names for subnet.  | [optional] [default to null]
**Options** | [***DhcpV6Options**](DhcpV6Options.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
