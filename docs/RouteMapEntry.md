# RouteMapEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action for the route map entry  | [default to null]
**CommunityListMatches** | [**[]CommunityMatchCriteria**](CommunityMatchCriteria.md) | Community list match criteria for route map. Properties community_list_matches and prefix_list_matches are mutually exclusive and cannot be used in the same route map entry.  | [optional] [default to null]
**Set** | [***RouteMapEntrySet**](RouteMapEntrySet.md) |  | [optional] [default to null]
**PrefixListMatches** | **[]string** | Prefix list match criteria for route map. Properties community_list_matches and prefix_list_matches are mutually exclusive and cannot be used in the same route map entry.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

