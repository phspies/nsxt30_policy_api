# Tier0VrfConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**RouteTargets** | [**[]VrfRouteTargets**](VrfRouteTargets.md) | Route targets. | [optional] [default to null]
**EvpnTransitVni** | **int32** | L3 VNI associated with the VRF for overlay traffic. VNI must be unique and belong to configured VNI pool.  | [optional] [default to null]
**Tier0Path** | **string** | Default tier0 path. Cannot be modified after realization.  | [default to null]
**RouteDistinguisher** | **string** | Route distinguisher. ASN:&lt;number&gt; or IPAddress:&lt;number&gt;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

