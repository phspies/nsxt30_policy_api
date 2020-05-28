# PolicyDnsForwarderZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**DnsDomainNames** | **[]string** | List of domain names on which conditional forwarding is based. This field is required if the DNS Zone is being used for a conditional forwarder. This field will also be used for conditional reverse lookup. Example 1, if for one of the zones, one of the entries in the fqdn is example.com, all the DNS requests under the domain example.com will be served by the corresponding upstream DNS server. Example 2, if for one of the zones, one of the entries in the fqdn list is \&quot;13.12.30.in-addr.arpa\&quot;, reverse lookup for 30.12.13.0/24 will go to the corresponding DNS server.  | [optional] [default to null]
**UpstreamServers** | **[]string** | Max of 3 DNS servers can be configured | [default to null]
**SourceIp** | **string** | The source IP used by the DNS Forwarder zone.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

