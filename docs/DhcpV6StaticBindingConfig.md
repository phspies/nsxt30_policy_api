# DhcpV6StaticBindingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**SntpServers** | **[]string** | SNTP server IP addresses. | [optional] [default to null]
**PreferredTime** | **int64** | Preferred time, in seconds. If this value is not provided, the value of lease_time*0.8 will be used.  | [optional] [default to null]
**LeaseTime** | **int64** | Lease time, in seconds. | [optional] [default to 86400]
**MacAddress** | **string** | The MAC address of the client host. Either client-duid or mac-address, but not both.  | [default to null]
**IpAddresses** | **[]string** | When not specified, no ip address will be assigned to client host.  | [optional] [default to null]
**DnsNameservers** | **[]string** | When not specified, no DNS nameserver will be set to client host.  | [optional] [default to null]
**DomainNames** | **[]string** | When not specified, no domain name will be assigned to client host.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

