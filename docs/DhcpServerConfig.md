# DhcpServerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**ServerAddresses** | **[]string** | DHCP server address in CIDR format. Both IPv4 and IPv6 address families are supported. Prefix length should be less than or equal to 30 for IPv4 address family and less than or equal to 126 for IPv6. When not specified, IPv4 value is auto-assigned to 100.96.0.1/30. Ignored when this object is configured at a Segment.  | [optional] [default to null]
**PreferredEdgePaths** | **[]string** | Policy paths to edge nodes on which the DHCP servers run. The first edge node is assigned as active edge, and second one as stanby edge. If only one edge node is specified, the DHCP servers will run without HA support. When this property is not specified, edge nodes are auto-assigned during realization of the DHCP server.  | [optional] [default to null]
**ServerAddress** | **string** | DHCP server address in CIDR format. Prefix length should be less than or equal to 30. DHCP server is deployed as DHCP relay service. This property is deprecated, use server_addresses instead. Both properties cannot be specified together with different new values.  | [optional] [default to null]
**EdgeClusterPath** | **string** | Edge cluster path. Auto assigned if only one edge cluster is configured on enforcement-point. Modifying edge cluster will reallocate DHCP server to the new edge cluster. Please note that re-allocating edge-cluster will result in losing of all exisitng DHCP lease information. Change edge cluster only when losing DHCP leases is not a real problem, e.g. cross-site migration or failover and all client hosts will be reboot and get new IP addresses.  | [optional] [default to null]
**LeaseTime** | **int64** | IP address lease time in seconds.  | [optional] [default to 86400]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

