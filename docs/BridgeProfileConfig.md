# BridgeProfileConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkTeamingPolicyName** | **string** | The name of the switching uplink teaming policy for the bridge endpoint. This name corresponds to one fot he switching uplink teaming policy names listed in teh transport zone. When this property is not specified, the teaming policy is assigned by MP. | [optional] [default to null]
**VlanIds** | **[]string** | VLAN specification for bridge endpoint. Either VLAN ID or VLAN ranges can be specified. Not both. | [optional] [default to null]
**BridgeProfilePath** | **string** | Same bridge profile can be configured on different segments. Each bridge profile on a segment must unique. | [default to null]
**VlanTransportZonePath** | **string** | VLAN transport zone should belong to the enforcment-point as the transport zone specified in the segment. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

