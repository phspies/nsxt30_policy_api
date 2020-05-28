# SecurityPolicyStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalSectionId** | **string** | Realized id of the section on NSX MP. Policy Manager can create more than one section per SecurityPolicy, in which case this identifier helps to distinguish between the multiple sections created.  | [optional] [default to null]
**ResultCount** | **int64** | Total count for rule statistics | [optional] [default to null]
**Results** | [**[]RuleStatistics**](RuleStatistics.md) | List of rule statistics. | [optional] [default to null]
**LrPath** | **string** | Path of the LR on which the section is applied in case of Gateway Firewall.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

