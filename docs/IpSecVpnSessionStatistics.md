# IpSecVpnSessionStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**EnforcementPointPath** | **string** | Policy Path referencing the enforcement point wehere the statistics are fetched.  | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated.  | [optional] [default to null]
**IkeTrafficStatistics** | [***IpSecVpnIkeTrafficStatistics**](IPSecVpnIkeTrafficStatistics.md) |  | [optional] [default to null]
**IkeStatus** | [***IpSecVpnIkeSessionStatus**](IPSecVpnIkeSessionStatus.md) |  | [optional] [default to null]
**PolicyStatistics** | [**[]IpSecVpnPolicyTrafficStatistics**](IPSecVpnPolicyTrafficStatistics.md) | Gives aggregate traffic statistics across all ipsec tunnels and individual tunnel statistics.  | [optional] [default to null]
**AggregateTrafficCounters** | [***IpSecVpnTrafficCounters**](IPSecVpnTrafficCounters.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

