# DnsSecurityProfileBindingMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilePath** | **string** | PolicyPath of associated Profile | [default to null]
**SequenceNumber** | **int64** | Sequence number used to resolve conflicts betweeen two profiles applied on the same group. Lower sequence number takes higher precedence. Two binding maps applied to the same profile must have the same sequence number. User defined sequence numbers range from 1 through 100,000. System defined sequence numbers range from 100,001 through 200,000.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

