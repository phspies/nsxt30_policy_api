# IngressRateLimiter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to null]
**ResourceType** | **string** | Type rate limiter  | [default to null]
**PeakBandwidth** | **int32** | The peak bandwidth rate is used to support burst traffic. | [optional] [default to 0]
**AverageBandwidth** | **int32** | You can use the average bandwidth to reduce network congestion. | [optional] [default to 0]
**BurstSize** | **int32** | The burst duration is set in the burst size setting. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

