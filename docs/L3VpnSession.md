# L3VpnSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | - A Policy Based L3Vpn is a configuration in which protect rules to match local and remote subnet needs to be defined. Tunnel is established for each pair of local and remote subnet defined in protect rules. - A Route Based L3Vpn is more flexible, more powerful and recommended over policy based. IP Tunnel subnet is created and all traffic routed through tunnel subnet (commonly known as VTI) is sent over tunnel. Routes can be learned through BGP. A route based L3Vpn is required when using redundant L3Vpn.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

