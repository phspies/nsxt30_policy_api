# LbHttpRequestCookieCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inverse** | **bool** | A flag to indicate whether reverse the match result of this condition | [optional] [default to false]
**Type_** | **string** | Type of load balancer rule condition | [default to null]
**MatchType** | **string** | Match type of cookie value. | [optional] [default to MATCH_TYPE.REGEX]
**CookieName** | **string** | Cookie name. | [default to null]
**CookieValue** | **string** | Cookie value. | [default to null]
**CaseSensitive** | **bool** | If true, case is significant when comparing cookie value.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

