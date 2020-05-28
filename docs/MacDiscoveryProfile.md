# MacDiscoveryProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**MacLearningAgingTime** | **int32** | Indicates how long learned MAC address remain. | [optional] [default to 600]
**MacLearningEnabled** | **bool** | Allowing source MAC address learning | [default to null]
**MacLimitPolicy** | **string** | The policy after MAC Limit is exceeded | [optional] [default to MAC_LIMIT_POLICY.ALLOW]
**RemoteOverlayMacLimit** | **int32** | This property specifies the limit on the maximum number of MACs learned for a remote virtual machine&#x27;s MAC to VTEP binding per overlay logical switch.  | [optional] [default to 2048]
**MacLimit** | **int32** | The maximum number of MAC addresses that can be learned on this port | [optional] [default to 4096]
**MacChangeEnabled** | **bool** | Allowing source MAC address change | [optional] [default to false]
**UnknownUnicastFloodingEnabled** | **bool** | Allowing flooding for unlearned MAC for ingress traffic | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

