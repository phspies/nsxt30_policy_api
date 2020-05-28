# DistributedFloodProtectionProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IcmpActiveFlowLimit** | **int64** | If this field is empty, firewall will not set a limit to active ICMP connections. | [optional] [default to null]
**TcpHalfOpenConnLimit** | **int64** | If this field is empty, firewall will not set a limit to half open TCP connections. | [optional] [default to null]
**UdpActiveFlowLimit** | **int64** | If this field is empty, firewall will not set a limit to active UDP connections. | [optional] [default to null]
**ResourceType** | **string** | GatewayFloodProtectionProfile is used for all Tier0 and Tier1 gateways. DistributedFloodProtectionProfile is used for all Transport Nodes.  | [default to null]
**OtherActiveConnLimit** | **int64** | If this field is empty, firewall will not set a limit to other active connections. besides UDP, ICMP and half open TCP connections. | [optional] [default to null]
**EnableRstSpoofing** | **bool** | If set to true, rst spoofing will be enabled. Flag is used only for distributed firewall profiles. | [optional] [default to false]
**EnableSyncache** | **bool** | If set to true, sync cache will be enabled. Flag is used only for distributed firewall profiles. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

