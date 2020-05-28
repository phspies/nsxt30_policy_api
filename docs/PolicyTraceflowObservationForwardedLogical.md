# PolicyTraceflowObservationForwardedLogical

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePathIndex** | **int64** | The path index of the service insertion component | [optional] [default to null]
**ComponentId** | **string** | The id of the component that forwarded the traceflow packet. | [optional] [default to null]
**SpoofguardVlanId** | **int64** | This field specified the VLAN id a traceflow packet matched in the whitelist in spoofguard. | [optional] [default to null]
**ResendType** | **string** | ARP_UNKNOWN_FROM_CP - Unknown ARP query result emitted by control plane ND_NS_UNKNOWN_FROM_CP - Unknown neighbor solicitation query result emitted by control plane UNKNOWN - Unknown resend type | [optional] [default to null]
**LportName** | **string** | The name of the logical port through which the traceflow packet was forwarded. | [optional] [default to null]
**AclRuleId** | **int64** | The id of the acl rule that was applied to forward the traceflow packet | [optional] [default to null]
**ServiceIndex** | **int64** | The index of the service insertion component | [optional] [default to null]
**Vni** | **int32** | VNI for the logical network on which the traceflow packet was forwarded. | [optional] [default to null]
**DstComponentName** | **string** | The name of the destination component to which the traceflow packet was forwarded. | [optional] [default to null]
**NatRuleId** | **int64** | The ID of the NAT rule that was applied to forward the traceflow packet | [optional] [default to null]
**TranslatedSrcIp** | **string** | The translated source IP address of VPN/NAT | [optional] [default to null]
**TranslatedDstIp** | **string** | The translated destination IP address of VNP/NAT | [optional] [default to null]
**SpoofguardMac** | **string** | The source MAC address of form: \&quot;^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$\&quot;. For example: 00:00:00:00:00:00.  | [optional] [default to null]
**DstComponentType** | **string** | The type of the destination component to which the traceflow packet was forwarded. | [optional] [default to null]
**LportId** | **string** | The id of the logical port through which the traceflow packet was forwarded. | [optional] [default to null]
**DstComponentId** | **string** | The id of the destination component to which the traceflow packet was forwarded. | [optional] [default to null]
**SpoofguardIp** | **string** | This field specified the prefix IP address a traceflow packet matched in the whitelist in spoofguard. | [optional] [default to null]
**ServiceTtl** | **int64** | The ttl of the service insertion component | [optional] [default to null]
**SvcNhMac** | **string** | MAC address of nexthop for service insertion(SI) in service VM(SVM) where the traceflow packet was received.  | [optional] [default to null]
**AclRulePath** | **string** | The path of the ACL rule that was applied to forward the traceflow packet | [optional] [default to null]
**NatRulePath** | **string** | The path of the NAT rule that was applied to forward the traceflow packet | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

