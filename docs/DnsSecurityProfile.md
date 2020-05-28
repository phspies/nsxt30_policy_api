# DnsSecurityProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Ttl** | **int64** | Time to live for DNS cache entry in seconds. Valid TTL values are between 3600 to 864000. However, this field accepts values between 0 through 864000. We define TTL type based on the value of TTL as follows: TTL 0 - cached entry never expires. TTL 1 to 3599 - invalid input and error is thrown TTL 3600 to 864000 - ttl is set to user input TTL field not set by user - TTL type is &#x27;AUTO&#x27; and ttl value is set from DNS response packet.  User defined TTL value is used only when it is betweeen 3600 to 864000.  | [optional] [default to 86400]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

