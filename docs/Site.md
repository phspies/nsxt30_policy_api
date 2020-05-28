# Site

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**FederationConfig** | [***GmFederationSiteConfig**](GmFederationSiteConfig.md) |  | [optional] [default to null]
**FailIfRttExceeded** | **bool** | Fail onboarding if maximum RTT exceeded.  | [optional] [default to true]
**SiteNumber** | **int64** | 12-bit system generated site number | [optional] [default to null]
**MaximumRtt** | **int64** | If provided and fail_if_rtt_exceeded is true, onboarding of the site will fail if measured RTT is greater than this value.  | [optional] [default to 250]
**FailIfRtepMisconfigured** | **bool** | Both the local site and the remote site must have edge clusters correctly configured and remote tunnel endpoint (RTEP) interfaces must be defined, or onboarding will fail.  | [optional] [default to true]
**SiteConnectionInfo** | [**[]SiteNodeConnectionInfo**](SiteNodeConnectionInfo.md) | To onboard a site, the connection information (username, password, and API thumbprint) for at least one NSX manager node in the remote site must be provided. Once the site has been successfully onboarded, the site_connection_info is discarded and authentication to the remote site occurs using an X.509 client certificate.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

