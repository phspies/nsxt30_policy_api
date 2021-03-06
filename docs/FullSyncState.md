# FullSyncState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**LastUpateTime** | **int64** | Timestamp of last update, could be progress or success or error. | [optional] [default to null]
**StartTime** | **int64** | Timestamp of Full Sync start. | [optional] [default to null]
**Errors** | **[]string** | Errors occurred during full sync.  | [optional] [default to null]
**LastCompletedStage** | **string** | The current stage of full sync completion for ongoing sync. When Local Manager (LM) receives full sync data from AR, LM starts with workflow to prserve the state and restore the full sync from where it has left off in case of change of leadership of the service to different NSX node or LM is restarted. LM starts the full sync workflow with state INITIAL capturing the AR full sync id and data location details. The stage/state transition follows the order given below INITIAL - Full sync started PROCESSED_FULLSYNC_DATA - Compelted processing the full state data                           provided by AR PRCESSED_DELTAS - Completed processing pending delta changes provided                   by AR. DELETED_STALE_ENTITIES - Completed deletion of all global entities on                          LM that are not in GM anymore COMPLETED - Full sync handling is completed on LM ERROR - Full sync failed with errors on LM, in which case AR will         re-attempt full sync later point in time for the LM ABORTED - Indicates that the full sync cancelled as per user request  | [optional] [default to null]
**FullSyncId** | **string** | Full sync id generated by Async Replicator (AR) service.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

