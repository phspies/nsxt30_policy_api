# EndpointRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ServiceProfiles** | **[]string** | The policy paths of service profiles are listed here. It pecifies what services are applied on the group. Currently only one is allowed.  | [default to null]
**Groups** | **[]string** | We need paths as duplicate names may exist for groups under different domains. In order to specify all groups, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values.  | [default to null]
**SequenceNumber** | **int32** | This field is used to resolve conflicts between multiple entries under EndpointPolicy. It will be system default value when not specified by user.  | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

