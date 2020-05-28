# AbstractSpace

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ConnectivityStrategy** | **string** | The connectivity strategy is deprecated. Use default layer3 rule, /infra/domains/default/security-policies/default-layer3-security-policy/rules/default-layer3-rule. This field indicates the default connectivity policy for the infra or tenant space WHITELIST - Adds a default drop rule. Administrator can then use \&quot;allow\&quot; rules (aka whitelist) to allow traffic between groups BLACKLIST - Adds a default allow rule. Admin can then use \&quot;drop\&quot; rules (aka blacklist) to block traffic between groups WHITELIST_ENABLE_LOGGING - Whitelising with logging enabled BLACKLIST_ENABLE_LOGGING - Blacklisting with logging enabled NONE - No default rules are added.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

