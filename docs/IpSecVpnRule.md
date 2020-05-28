# IpSecVpnRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Sources** | [**[]IpSecVpnSubnet**](IPSecVpnSubnet.md) | List of local subnets. Specifying no value is interpreted as 0.0.0.0/0.  | [optional] [default to null]
**Action** | **string** | PROTECT - Protect rules are defined per policy based IPSec VPN session. BYPASS - Bypass rules are defined per IPSec VPN service and affects all policy based IPSec VPN sessions. Bypass rules are prioritized over protect rules.  | [optional] [default to ACTION.PROTECT]
**Enabled** | **bool** | A flag to enable/disable the rule. | [optional] [default to true]
**Logged** | **bool** | A flag to enable/disable the logging for the rule. | [optional] [default to false]
**SequenceNumber** | **int32** | A sequence number is used to give a priority to an IPSecVpnRule. | [optional] [default to null]
**Destinations** | [**[]IpSecVpnSubnet**](IPSecVpnSubnet.md) | List of peer subnets. Specifying no value is interpreted as 0.0.0.0/0.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

