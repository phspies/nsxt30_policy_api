# RouteBasedL3VpnSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | - A Policy Based L3Vpn is a configuration in which protect rules to match local and remote subnet needs to be defined. Tunnel is established for each pair of local and remote subnet defined in protect rules. - A Route Based L3Vpn is more flexible, more powerful and recommended over policy based. IP Tunnel subnet is created and all traffic routed through tunnel subnet (commonly known as VTI) is sent over tunnel. Routes can be learned through BGP. A route based L3Vpn is required when using redundant L3Vpn.  | [default to null]
**RoutingConfigPath** | **string** | This is a deprecated field. Any specified value is not saved and will be ignored.  | [optional] [default to null]
**TunnelSubnets** | [**[]TunnelSubnet**](TunnelSubnet.md) | Virtual tunnel interface (VTI) port IP subnets to be used to configure route-based L3Vpn session. A max of one tunnel subnet is allowed.  | [default to null]
**DefaultRuleLogging** | **bool** | Indicates if logging should be enabled for the default whitelisting rule for the VTI interface.  | [optional] [default to false]
**ForceWhitelisting** | **bool** | The default firewall rule Action is set to DROP if true otherwise set to ALLOW.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

