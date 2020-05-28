# BasePolicyServiceInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**PartnerServiceName** | **string** | Unique name of Partner Service in the Marketplace | [default to null]
**TransportType** | **string** | Transport to be used while deploying Service-VM. | [optional] [default to TRANSPORT_TYPE.L2_BRIDGE]
**DeploymentMode** | **string** | Deployment mode specifies how the partner appliance will be deployed i.e. in HA or standalone mode. | [optional] [default to DEPLOYMENT_MODE.ACTIVE_STANDBY]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

