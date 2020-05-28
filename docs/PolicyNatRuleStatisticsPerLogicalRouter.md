# PolicyNatRuleStatisticsPerLogicalRouter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated.  | [optional] [default to null]
**PerNodeStatistics** | [**[]PolicyNatRuleStatisticsPerTransportNode**](PolicyNatRuleStatisticsPerTransportNode.md) | Detailed Rule statistics per logical router.  | [optional] [default to null]
**Statistics** | [***PolicyNatRuleCounters**](PolicyNATRuleCounters.md) |  | [optional] [default to null]
**RouterPath** | **string** | Path of the router.  | [optional] [default to null]
**EnforcementPointPath** | **string** | Policy Path referencing the enforcement point from where the statistics are fetched.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

