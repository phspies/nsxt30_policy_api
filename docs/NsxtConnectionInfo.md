# NsxtConnectionInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforcementPointAddress** | **string** | Value of this property could be Hostname or IP. For instance: - On an NSX-T MP running on default port, the value could be \&quot;10.192.1.1\&quot; - On an NSX-T MP running on custom port, the value could be \&quot;192.168.1.1:32789\&quot; - On an NSX-T MP in VMC deployments, the value could be \&quot;192.168.1.1:5480/nsxapi\&quot;  | [default to null]
**ResourceType** | **string** | Resource Type of Enforcement Point Connection Info. | [default to null]
**Username** | **string** | Username. | [optional] [default to null]
**TransportZoneIds** | **[]string** | Transport Zone UUIDs on enforcement point. Transport zone information is required for creating logical L2, L3 constructs on enforcement point. Max 1 transport zone ID. This is a deprecated property. The transport zone id is now auto populated from enforcement point and its value can be read using APIs GET /infra/sites/site-id/enforcement-points/enforcementpoint-id/transport-zones and GET /infra/sites/site-id/enforcement-points/enforcementpoint-id/transport-zones/transport-zone-id. The value passed through this property will be ignored.  | [optional] [default to null]
**Password** | **string** | Password. | [optional] [default to null]
**EdgeClusterIds** | **[]string** | Edge Cluster UUIDs on enforcement point. Edge cluster information is required for creating logical L2, L3 constructs on enforcement point. Max 1 edge cluster ID. This is a deprecated property. The edge cluster id is now auto populated from enforcement point and its value can be read using APIs GET /infra/sites/site-id/enforcement-points/enforcementpoint-id/edge-clusters and GET /infra/sites/site-id/enforcement-points/enforcementpoint-1/edge-clusters/edge-cluster-id. The value passed through this property will be ignored.  | [optional] [default to null]
**Thumbprint** | **string** | Thumbprint of EnforcementPoint in the form of a SHA-256 hash represented in lower case HEX.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

