# LbService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**AccessLogEnabled** | **bool** | Flag to enable access log | [optional] [default to null]
**ConnectivityPath** | **string** | LBS could be instantiated (or created) on the Tier-1, etc. For now, only the Tier-1 object is supported.  | [optional] [default to null]
**ErrorLogLevel** | **string** | Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log.  | [optional] [default to ERROR_LOG_LEVEL.INFO]
**RelaxScaleValidation** | **bool** | If relax_scale_validation is true, the scale validations for virtual servers/pools/pool members/rules are relaxed for load balancer service. When load balancer service is deployed on edge nodes, the scale of virtual servers/pools/pool members for the load balancer service should not exceed the scale number of the largest load balancer size which could be configured on a certain edge form factor. For example, the largest load balancer size supported on a MEDIUM edge node is MEDIUM. So one SMALL load balancer deployed on MEDIUM edge nodes can support the scale number of MEDIUM load balancer. It is not recommended to enable active monitors if relax_scale_validation is true due to performance consideration. If relax_scale_validation is false, scale numbers should be validated for load balancer service.  | [optional] [default to false]
**Enabled** | **bool** | Flag to enable the load balancer service. | [optional] [default to true]
**Size** | **string** | Load balancer service size. | [optional] [default to SIZE.SMALL]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

