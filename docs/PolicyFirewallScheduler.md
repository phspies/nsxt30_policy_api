# PolicyFirewallScheduler

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**TimeInterval** | [**[]PolicyTimeIntervalValue**](PolicyTimeIntervalValue.md) | The recurring time interval in a day during which the schedule will be applicable. It should not be present when the recurring flag is false.  | [optional] [default to null]
**EndTime** | **string** | If recurring field is set false, then this field must be present. The schedule will be enforced till the end time of the specified end date. If recurring field is set true, then this field should not be present.  | [optional] [default to null]
**EndDate** | **string** | End date on which schedule to end. Example, 12/22/2019.  | [default to null]
**Timezone** | **string** | Host Timezone to be used to enforce firewall rules.  | [default to null]
**StartTime** | **string** | Time in 24 hour and minutes in multiple of 30. Example, 9:00. If recurring field is set false, then this field must be present. The schedule will start getting enforced from the start time of the specified start date. If recurring field is set true, then this field should not be present.  | [optional] [default to null]
**Recurring** | **bool** | Flag to indicate whether firewall schedule recurs or not. The default value is true and it should be set to false when the firewall schedule does not recur and is a one time time interval.  | [default to true]
**Days** | **[]string** | Days of week on which rules will be enforced. If property is omitted, then days of the week will not considered while calculating the firewall schedule. It should not be present when the recurring flag is false.  | [optional] [default to null]
**StartDate** | **string** | Start date on which schedule to start. Example, 02/22/2019.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

