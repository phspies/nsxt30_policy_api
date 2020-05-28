# EnforcementPoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ConnectionInfo** | [***EnforcementPointConnectionInfo**](EnforcementPointConnectionInfo.md) |  | [default to null]
**Version** | **string** | Version of the Enforcement point. | [optional] [default to null]
**AutoEnforce** | **bool** | Auto enforce flag suggests whether the policy objects shall be automatically enforced on this enforcement point or not. When this flag is set to true, all policy objects will be automatically enforced on this enforcement point. If this flag is set to false, user shall rely on the usual means of realization, i.e., deployment maps.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

