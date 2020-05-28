# LocalEgressRoutingEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NexthopAddress** | **string** | Next hop address for proximity routing.  | [default to null]
**PrefixListPaths** | **[]string** | The destination address of traffic matching a prefix-list is forwarded to the nexthop_address. Traffic matching a prefix list with Action DENY will be dropped. Individual prefix-lists specified could have different actions.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

