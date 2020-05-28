# TcpMaximumSegmentSizeClamping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSegmentSize** | **int64** | MSS defines the maximum amount of data that a host is willing to accept in a single TCP segment. This field is set in TCP header during connection establishment. To avoid packet fragmentation, you can set this field depending on uplink MTU and VPN overhead. This is an optional field and in case it is left unconfigured, best possible MSS value will be calculated based on effective mtu of uplink interface. Supported MSS range is 216 to 8960.  | [optional] [default to null]
**Direction** | **string** | Specifies the traffic direction for which to apply MSS Clamping.  | [optional] [default to DIRECTION.NONE]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

