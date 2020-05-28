# PolicyServiceProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**VendorTemplateKey** | **string** | The vendor template key property of actual vendor template. This should be used when multiple templates with same name exist. | [optional] [default to null]
**Attributes** | [**[]Attribute**](Attribute.md) | List of attributes specific to a partner for which the service is created. These attributes are passed on to the partner appliance and are opaque to NSX. If a vendor template exposes configurable parameters, then their values are specified here. | [optional] [default to null]
**RedirectionAction** | **string** | The redirection action represents if the packet is exclusively redirected to the service, or if a copy is forwarded to the service. Redirection action is not applicable to guest introspection service. | [optional] [default to null]
**VendorTemplateName** | **string** | Name of the vendor template for which this Service Profile is being created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

