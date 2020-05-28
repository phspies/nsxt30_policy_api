# L3VpnRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Action** | **string** | Action to exchange data with or without protection. PROTECT - Allows to exchange data with ipsec protection. Protect rules are defined per L3Vpn. BYPASS - Allows to exchange data without ipsec protection. Bypass rules are defined per L3VpnContext and affects all policy based L3Vpns. Bypass rules are prioritized over protect rules.  | [optional] [default to ACTION.PROTECT]
**Sources** | [**[]L3VpnSubnet**](L3VpnSubnet.md) | List of local subnets used in policy-based L3Vpn.  | [default to null]
**SequenceNumber** | **int32** | This field is used to resolve conflicts between multiple L3VpnRules associated with a single L3Vpn or L3VpnContext.  | [optional] [default to null]
**Destinations** | [**[]L3VpnSubnet**](L3VpnSubnet.md) | List of remote subnets used in policy-based L3Vpn.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

