# Infra

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectivityStrategy** | **string** | The connectivity strategy is deprecated. Use default layer3 rule, /infra/domains/default/security-policies/default-layer3-security-policy/rules/default-layer3-rule. This field indicates the default connectivity policy for the infra or tenant space WHITELIST - Adds a default drop rule. Administrator can then use \&quot;allow\&quot; rules (aka whitelist) to allow traffic between groups BLACKLIST - Adds a default allow rule. Admin can then use \&quot;drop\&quot; rules (aka blacklist) to block traffic between groups WHITELIST_ENABLE_LOGGING - Whitelising with logging enabled BLACKLIST_ENABLE_LOGGING - Blacklisting with logging enabled NONE - No default rules are added.  | [optional] [default to null]
**Domains** | [**[]Domain**](Domain.md) | This field is used while creating or updating the infra space.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

