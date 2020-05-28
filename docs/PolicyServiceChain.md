# PolicyServiceChain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ReversePathServiceProfiles** | **[]string** | Reverse path service profiles are applied to egress traffic and is optional. 2 different set of profiles can be defined for forward and reverse path. If not defined, the reverse of the forward path service profile is applied. | [optional] [default to null]
**PathSelectionPolicy** | **string** | Path selection policy can be - ANY - Service Insertion is free to redirect to any service path regardless of any load balancing considerations or flow pinning. LOCAL - Preference to be given to local service insances. REMOTE - Preference to be given to the SVM co-located on the same host. ROUND_ROBIN - All active service paths are hit with equal probability. | [optional] [default to PATH_SELECTION_POLICY.ANY]
**ServiceSegmentPath** | **[]string** | Path to service segment using which the traffic needs to be redirected. | [default to null]
**ForwardPathServiceProfiles** | **[]string** | Forward path service profiles are applied to ingress traffic. | [default to null]
**FailurePolicy** | **string** | Failure policy for the service defines the action to be taken i.e to allow or to block the traffic during failure scenarios. | [optional] [default to FAILURE_POLICY.ALLOW]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

