# SegmentSecurityProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**BpduFilterEnable** | **bool** | Indicates whether BPDU filter is enabled. BPDU filtering is enabled by default.  | [optional] [default to true]
**RaGuardEnabled** | **bool** | Enable or disable Router Advertisement Guard.  | [optional] [default to false]
**DhcpClientBlockV6Enabled** | **bool** | Filters DHCP server and/or client IPv6 traffic. DHCP server blocking is enabled and client blocking is disabled by default.  | [optional] [default to false]
**NonIpTrafficBlockEnabled** | **bool** | A flag to block all traffic except IP/(G)ARP/BPDU.  | [optional] [default to false]
**BpduFilterAllow** | **[]string** | Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering. List of allowed MACs - 01:80:c2:00:00:00, 01:80:c2:00:00:01, 01:80:c2:00:00:02, 01:80:c2:00:00:03,                        01:80:c2:00:00:04, 01:80:c2:00:00:05, 01:80:c2:00:00:06, 01:80:c2:00:00:07,                        01:80:c2:00:00:08, 01:80:c2:00:00:09, 01:80:c2:00:00:0a, 01:80:c2:00:00:0b,                        01:80:c2:00:00:0c, 01:80:c2:00:00:0d, 01:80:c2:00:00:0e, 01:80:c2:00:00:0f,                        00:e0:2b:00:00:00, 00:e0:2b:00:00:04, 00:e0:2b:00:00:06, 01:00:0c:00:00:00,                        01:00:0c:cc:cc:cc, 01:00:0c:cc:cc:cd, 01:00:0c:cd:cd:cd, 01:00:0c:cc:cc:c0,                        01:00:0c:cc:cc:c1, 01:00:0c:cc:cc:c2, 01:00:0c:cc:cc:c3, 01:00:0c:cc:cc:c4,                        01:00:0c:cc:cc:c5, 01:00:0c:cc:cc:c6, 01:00:0c:cc:cc:c7  | [optional] [default to null]
**DhcpServerBlockEnabled** | **bool** | Filters DHCP server and/or client traffic. DHCP server blocking is enabled and client blocking is disabled by default.  | [optional] [default to true]
**RateLimitsEnabled** | **bool** | Enable or disable Rate Limits  | [optional] [default to false]
**RateLimits** | [***TrafficRateLimits**](TrafficRateLimits.md) |  | [optional] [default to null]
**DhcpClientBlockEnabled** | **bool** | Filters DHCP server and/or client traffic. DHCP server blocking is enabled and client blocking is disabled by default.  | [optional] [default to false]
**DhcpServerBlockV6Enabled** | **bool** | Filters DHCP server and/or client IPv6 traffic. DHCP server blocking is enabled and client blocking is disabled by default.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

