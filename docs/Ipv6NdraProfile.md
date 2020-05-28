# Ipv6NdraProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**RaMode** | **string** | RA Mode | [default to RA_MODE.SLAAC_DNS_THROUGH_RA]
**RaConfig** | [***RaConfig**](RAConfig.md) |  | [default to null]
**RetransmitInterval** | **int64** | The time, in milliseconds, between retransmitted neighbour solicitation messages. A value of 0 means unspecified.  | [optional] [default to 1000]
**DnsConfig** | [***RaDnsConfig**](RaDNSConfig.md) |  | [optional] [default to null]
**ReachableTimer** | **int64** | Neighbour reachable time duration in milliseconds. A value of 0 means unspecified.  | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

