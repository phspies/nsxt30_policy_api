# L3VpnContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IkeLogLevel** | **string** | Log level for internet key exchange (IKE).  | [optional] [default to IKE_LOG_LEVEL.INFO]
**Enabled** | **bool** | If true, enable L3Vpn Service for given tier-0. Enabling/disabling this service affects all L3Vpns under the given tier-0.  | [optional] [default to true]
**BypassRules** | [**[]L3VpnRule**](L3VpnRule.md) | Bypass L3Vpn rules that will be shared across L3Vpns. Only Bypass action is supported on these L3Vpn rules.  | [optional] [default to null]
**AvailableLocalAddresses** | [**[]PolicyIpAddressInfo**](PolicyIPAddressInfo.md) | Local gateway IPv4 addresses available for configuration of each L3Vpn.  | [optional] [default to null]
**Label** | **string** | Policy path referencing Label. A label is used as a mechanism to group route-based L3Vpns in order to apply edge firewall rules on members&#x27; VTIs.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

