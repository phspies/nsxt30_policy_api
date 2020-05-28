# GlobalManager

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**FailIfRttExceeded** | **bool** | Fail onboarding if maximum RTT exceeded.  | [optional] [default to true]
**FederationId** | **string** | Internally generated UUID to the federation of Global Manager.  | [optional] [default to null]
**Mode** | **string** | There can be at most one ACTIVE global manager and one STANDBY global manager. In order to add a STANDBY manager, there must be an ACTIVE manager defined.  | [default to null]
**ConnectionInfo** | [**[]SiteNodeConnectionInfo**](SiteNodeConnectionInfo.md) | To create a standby GM, the connection information (username, password, and API thumbprint) for at least one NSX manager node in the remote site must be provided. Once the GM has been successfully onboarded, the connection_info is discarded and authentication to the standby GM occurs using an X.509 client certificate.  | [optional] [default to null]
**MaximumRtt** | **int64** | If provided and fail_if_rtt_exceeded is true, onboarding of the site will fail if measured RTT is greater than this value.  | [optional] [default to 250]
**SiteId** | **string** | UUID of the site where Global manager is running. This is the Site Manager generated UUID for every NSX deployment.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

