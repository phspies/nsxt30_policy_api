# IpDiscoveryProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ArpNdBindingTimeout** | **int32** | This property controls the ARP and ND cache timeout period. It is recommended that this property be greater than the ARP/ND cache timeout on the VM.  | [optional] [default to 10]
**IpV6DiscoveryOptions** | [***IPv6DiscoveryOptions**](IPv6DiscoveryOptions.md) |  | [optional] [default to null]
**DuplicateIpDetection** | [***DuplicateIpDetectionOptions**](DuplicateIPDetectionOptions.md) |  | [optional] [default to null]
**TofuEnabled** | **bool** | Indicates whether \&quot;Trust on First Use(TOFU)\&quot; paradigm is enabled. | [optional] [default to true]
**IpV4DiscoveryOptions** | [***IPv4DiscoveryOptions**](IPv4DiscoveryOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

