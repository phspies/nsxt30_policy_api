# IpSecVpnService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IkeLogLevel** | **string** | Log level for internet key exchange (IKE). | [optional] [default to IKE_LOG_LEVEL.INFO]
**BypassRules** | [**[]IpSecVpnRule**](IPSecVpnRule.md) | Bypass policy rules are configured using VPN service. Bypass rules always have higher priority over protect rules and they affect all policy based vpn sessions associated with the IPSec VPN service. Protect rules are defined per policy based vpn session.  | [optional] [default to null]
**HaSync** | **bool** | Enable/disable IPSec HA state sync. IPSec HA state sync can be disabled if in case there are performance issues w.r.t. the state sync messages. | [optional] [default to true]
**Enabled** | **bool** | If true, enable VPN services for given locale service. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

