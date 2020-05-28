# DhcpV4StaticBindingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**GatewayAddress** | **string** | When not specified, gateway address is auto-assigned from segment configuration.  | [optional] [default to null]
**HostName** | **string** | Hostname to assign to the host.  | [optional] [default to null]
**MacAddress** | **string** | MAC address of the host.  | [default to null]
**LeaseTime** | **int64** | DHCP lease time in seconds.  | [optional] [default to 86400]
**IpAddress** | **string** | IP assigned to host. The IP address must belong to the subnet, if any, configured on Segment.  | [default to null]
**Options** | [***DhcpV4Options**](DhcpV4Options.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

