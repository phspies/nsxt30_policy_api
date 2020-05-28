# ApplicationTier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**IpRange** | **string** | This represents the subnet that is associated with tier. If this is specified, size property is ignored.  | [optional] [default to null]
**AccessType** | **string** | There are three kinds of Access Types supported for an Application. Public  - Tier is accessible from external networks and its IP is picked up from public IP           addresses from VHC configuration unless specified explicitly by user. Private - Tier is accessbile only within the application and its IP is picked up from           private IP addresses from VHC configuration unless specified explicitly by user. Shared  - Tier is accessible within the scope of VHC and its IP is picked up from private           IP addresses from VHC configuration unless specified explicitly by user.  | [optional] [default to ACCESS_TYPE.SHARED]
**Size** | **string** | ONE  - 1       XXS - 8 XS   - 16        S - 32 M    - 64        L - 128 XL   - 256     XXL - 512 XXXL - 1024  | [optional] [default to SIZE.XXS]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

