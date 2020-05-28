# PolicyLbMonitorProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**MonitorPort** | **int32** | Typically, monitors perform healthchecks to Group members using the member IP address and pool_port. However, in some cases, customers prefer to run healthchecks against a different port than the pool member port which handles actual application traffic. In such cases, the port to run healthchecks against can be specified in the monitor_port value.  | [optional] [default to null]
**Timeout** | **int64** | Timeout specified in seconds.  After a healthcheck is initiated, if it does not complete within a certain period, then also the healthcheck is considered to be unsuccessful. Completing a healthcheck within timeout means establishing a connection (TCP or SSL), if applicable, sending the request and receiving the response, all within the configured timeout.  | [optional] [default to 15]
**FallCount** | **int64** | Only if a healthcheck fails consecutively for a specified number of times, given with fall_count, to a member will the member status be marked DOWN.  | [optional] [default to 3]
**Interval** | **int64** | Active healthchecks are initiated periodically, at a configurable interval (in seconds), to each member of the Group.  | [optional] [default to 5]
**RiseCount** | **int64** | Once a member is DOWN, a specified number of consecutive successful healthchecks specified by rise_count will bring the member back to UP state.  | [optional] [default to 3]
**ResourceType** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

