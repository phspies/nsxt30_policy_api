# Vhc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**SiteInfos** | [**[]SiteInfo**](SiteInfo.md) | Information related to sites applicable for given VHC.  | [default to null]
**PublicIpAddresses** | **[]string** | This is set of IP addresses that will be used for Public Application tiers.  | [default to null]
**PrivateIpAddresses** | **[]string** | This is set of IP addresses that will be used for Shared and Private Application tiers.  | [default to null]
**Capabilities** | **[]string** | Type of Services to be made available for the applications defined under VHC.  | [optional] [default to null]
**Tier0s** | **[]string** | The tier 0 has to be pre-created before VHC is created. The tier 0 typically provides connectivity to external world. List of sites for VHC has to be subset of sites where the tier 0 spans.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

