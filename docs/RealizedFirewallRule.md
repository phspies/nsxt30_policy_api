# RealizedFirewallRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealizationSpecificIdentifier** | **string** | Realization id of this object | [optional] [default to null]
**IntentReference** | **[]string** | Desire state paths of this object | [optional] [default to null]
**State** | **string** | Realization state of this object | [default to null]
**RealizationApi** | **string** | Realization API of this object on enforcement point | [optional] [default to null]
**Alarms** | [**[]PolicyAlarmResource**](PolicyAlarmResource.md) | Alarm info detail | [optional] [default to null]
**RuntimeError** | **string** | It define the root cause for runtime error.  | [optional] [default to null]
**RuntimeStatus** | **string** | Possible values could be UP, DOWN, UNKNOWN, DEGRADED This list is not exhaustive.  | [optional] [default to null]
**Disabled** | **bool** | Flag to disable rule. Disabled will only be persisted but never provisioned/realized. | [optional] [default to null]
**Sources** | [**[]ResourceReference**](ResourceReference.md) | List of sources. Null will be treated as any. | [optional] [default to null]
**Direction** | **string** | Rule direction in case of stateless firewall rules. This will only considered if section level parameter is set to stateless. Default to IN_OUT if not specified. | [optional] [default to DIRECTION.IN_OUT]
**Services** | [**[]ResourceReference**](ResourceReference.md) | List of the services. Null will be treated as any. | [optional] [default to null]
**Action** | **string** | Action enforced on the packets which matches the firewall rule. | [optional] [default to null]
**Destinations** | [**[]ResourceReference**](ResourceReference.md) | List of the destinations. Null will be treated as any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

