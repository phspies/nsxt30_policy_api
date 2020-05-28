# RouteAdvertisementRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to advertise filtered routes to the connected Tier0 gateway. PERMIT: Enables the advertisment DENY: Disables the advertisement  | [default to ACTION.PERMIT]
**Subnets** | **[]string** | Network CIDRs to be routed.  | [optional] [default to null]
**PrefixOperator** | **string** | Prefix operator to filter subnets. GE prefix operator filters all the routes with prefix length greater than or equal to the subnets configured. EQ prefix operator filter all the routes with prefix length equal to the subnets configured.  | [optional] [default to PREFIX_OPERATOR.GE]
**Name** | **string** | Display name should be unique.  | [default to null]
**RouteAdvertisementTypes** | **[]string** | Enable different types of route advertisements. When not specified, routes to IPSec VPN local-endpoint subnets (TIER1_IPSEC_LOCAL_ENDPOINT) are automatically advertised.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

