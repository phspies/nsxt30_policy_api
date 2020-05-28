# L2VpnService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**EncapIpPool** | **[]string** | IP Pool to allocate local and peer endpoint IPs for L2VpnSession logical tap.  | [optional] [default to null]
**Mode** | **string** | Specify an L2VPN service mode as SERVER or CLIENT. | [optional] [default to MODE.SERVER]
**EnableHub** | **bool** | This property applies only in SERVER mode. If set to true, traffic from any client will be replicated to all other clients. If set to false, traffic received from clients is only replicated to the local VPN endpoint.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

