# Tier1Interface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]InterfaceSubnet**](InterfaceSubnet.md) | Specify IP address and network prefix for interface.  | [default to null]
**UrpfMode** | **string** | Unicast Reverse Path Forwarding mode | [optional] [default to URPF_MODE.STRICT]
**Mtu** | **int32** | Maximum transmission unit (MTU) specifies the size of the largest packet that a network protocol can transmit.  | [optional] [default to null]
**Ipv6ProfilePaths** | **[]string** | Configrue IPv6 NDRA profile. Only one NDRA profile can be configured.  | [optional] [default to null]
**SegmentPath** | **string** | Policy path of Segment to which interface is connected to.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

