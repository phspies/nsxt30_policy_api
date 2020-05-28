# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddorRemoveGroupExternalIDMembers**](PolicyInventoryApi.md#AddorRemoveGroupExternalIDMembers) | **Post** /infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Add or Remove external id based members from/to a Group 
[**AddorRemoveGroupExternalIDMembers_0**](PolicyInventoryApi.md#AddorRemoveGroupExternalIDMembers_0) | **Post** /global-infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Add or Remove external id based members from/to a Group 
[**AddorRemoveGroupIPAddresses**](PolicyInventoryApi.md#AddorRemoveGroupIPAddresses) | **Post** /infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Add or Remove IP Addresses from/to a Group 
[**AddorRemoveGroupIPAddresses_0**](PolicyInventoryApi.md#AddorRemoveGroupIPAddresses_0) | **Post** /global-infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Add or Remove IP Addresses from/to a Group 
[**AddorRemoveGroupMACAddresses**](PolicyInventoryApi.md#AddorRemoveGroupMACAddresses) | **Post** /infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Add or Remove MAC Addresses from/to a Group 
[**AddorRemoveGroupMACAddresses_0**](PolicyInventoryApi.md#AddorRemoveGroupMACAddresses_0) | **Post** /global-infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Add or Remove MAC Addresses from/to a Group 
[**AddorRemoveGroupPathMembers**](PolicyInventoryApi.md#AddorRemoveGroupPathMembers) | **Post** /infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Add or Remove path based members from/to a Group 
[**AddorRemoveGroupPathMembers_0**](PolicyInventoryApi.md#AddorRemoveGroupPathMembers_0) | **Post** /global-infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Add or Remove path based members from/to a Group 
[**CreateOrReplaceTier0Group**](PolicyInventoryApi.md#CreateOrReplaceTier0Group) | **Put** /infra/tier-0s/{tier-0-id}/groups/{group-id} | Create or update Group under specified Tier-0
[**CreateOrReplaceTier0Group_0**](PolicyInventoryApi.md#CreateOrReplaceTier0Group_0) | **Put** /global-infra/tier-0s/{tier-0-id}/groups/{group-id} | Create or update Group under specified Tier-0
[**DeleteGroup**](PolicyInventoryApi.md#DeleteGroup) | **Delete** /infra/domains/{domain-id}/groups/{group-id} | Delete Group
[**DeleteGroupExternalIDExpression**](PolicyInventoryApi.md#DeleteGroupExternalIDExpression) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Delete Group External ID Expression
[**DeleteGroupExternalIDExpression_0**](PolicyInventoryApi.md#DeleteGroupExternalIDExpression_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Delete Group External ID Expression
[**DeleteGroupIPAddressExpression**](PolicyInventoryApi.md#DeleteGroupIPAddressExpression) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Delete Group IPAddressExpression
[**DeleteGroupIPAddressExpression_0**](PolicyInventoryApi.md#DeleteGroupIPAddressExpression_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Delete Group IPAddressExpression
[**DeleteGroupMACAddressExpression**](PolicyInventoryApi.md#DeleteGroupMACAddressExpression) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Delete Group MACAddressExpression
[**DeleteGroupMACAddressExpression_0**](PolicyInventoryApi.md#DeleteGroupMACAddressExpression_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Delete Group MACAddressExpression
[**DeleteGroupPathExpression**](PolicyInventoryApi.md#DeleteGroupPathExpression) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Delete Group Path Expression
[**DeleteGroupPathExpression_0**](PolicyInventoryApi.md#DeleteGroupPathExpression_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Delete Group Path Expression
[**DeleteGroup_0**](PolicyInventoryApi.md#DeleteGroup_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id} | Delete Group
[**DeletePolicyContextProfile**](PolicyInventoryApi.md#DeletePolicyContextProfile) | **Delete** /infra/context-profiles/{context-profile-id} | Delete Policy Context Profile
[**DeletePolicyContextProfile_0**](PolicyInventoryApi.md#DeletePolicyContextProfile_0) | **Delete** /global-infra/context-profiles/{context-profile-id} | Delete Policy Context Profile
[**DeleteServiceEntry**](PolicyInventoryApi.md#DeleteServiceEntry) | **Delete** /global-infra/services/{service-id}/service-entries/{service-entry-id} | Delete Service entry
[**DeleteServiceEntry_0**](PolicyInventoryApi.md#DeleteServiceEntry_0) | **Delete** /infra/services/{service-id}/service-entries/{service-entry-id} | Delete Service entry
[**DeleteServiceForTenant**](PolicyInventoryApi.md#DeleteServiceForTenant) | **Delete** /global-infra/services/{service-id} | Delete Service
[**DeleteServiceForTenant_0**](PolicyInventoryApi.md#DeleteServiceForTenant_0) | **Delete** /infra/services/{service-id} | Delete Service
[**DeleteTier0Group**](PolicyInventoryApi.md#DeleteTier0Group) | **Delete** /infra/tier-0s/{tier-0-id}/groups/{group-id} | Deletes Group under Tier-0
[**DeleteTier0Group_0**](PolicyInventoryApi.md#DeleteTier0Group_0) | **Delete** /global-infra/tier-0s/{tier-0-id}/groups/{group-id} | Deletes Group under Tier-0
[**GetGroupIPMembers**](PolicyInventoryApi.md#GetGroupIPMembers) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/ip-addresses | Get IP addresses that belong to this Group
[**GetGroupIPMembers_0**](PolicyInventoryApi.md#GetGroupIPMembers_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/ip-addresses | Get IP addresses that belong to this Group
[**GetGroupLPMembers**](PolicyInventoryApi.md#GetGroupLPMembers) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/logical-ports | Get logical ports that belong to this Group
[**GetGroupLPMembers_0**](PolicyInventoryApi.md#GetGroupLPMembers_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/logical-ports | Get logical ports that belong to this Group
[**GetGroupLSMembers**](PolicyInventoryApi.md#GetGroupLSMembers) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/logical-switches | Get logical switches that belong to this Group
[**GetGroupLSMembers_0**](PolicyInventoryApi.md#GetGroupLSMembers_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/logical-switches | Get logical switches that belong to this Group
[**GetGroupSegmentMembers**](PolicyInventoryApi.md#GetGroupSegmentMembers) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/segments | Get segments that belong to this Group
[**GetGroupSegmentMembers_0**](PolicyInventoryApi.md#GetGroupSegmentMembers_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/segments | Get segments that belong to this Group
[**GetGroupSegmentPortMembers**](PolicyInventoryApi.md#GetGroupSegmentPortMembers) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/segment-ports | Get segment ports that belong to this Group
[**GetGroupSegmentPortMembers_0**](PolicyInventoryApi.md#GetGroupSegmentPortMembers_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/segment-ports | Get segment ports that belong to this Group
[**GetGroupTags**](PolicyInventoryApi.md#GetGroupTags) | **Get** /infra/domains/{domain-id}/groups/{group-id}/tags | Get tags used to define conditions inside a Group
[**GetGroupTags_0**](PolicyInventoryApi.md#GetGroupTags_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/tags | Get tags used to define conditions inside a Group
[**GetGroupVIFMembers**](PolicyInventoryApi.md#GetGroupVIFMembers) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/vifs | Get Virtual Network Interface instances that belong to this Group
[**GetGroupVIFMembers_0**](PolicyInventoryApi.md#GetGroupVIFMembers_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/vifs | Get Virtual Network Interface instances that belong to this Group
[**GetGroupVMMembers**](PolicyInventoryApi.md#GetGroupVMMembers) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/members/virtual-machines | Get Virtual machines that belong to this Group
[**GetGroupVMMembers_0**](PolicyInventoryApi.md#GetGroupVMMembers_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/members/virtual-machines | Get Virtual machines that belong to this Group
[**GetGroupVMStatistics**](PolicyInventoryApi.md#GetGroupVMStatistics) | **Get** /infra/domains/{domain-id}/groups/{group-id}/statistics/virtual-machines | Get effective VMs for the Group
[**GetGroupVMStatistics_0**](PolicyInventoryApi.md#GetGroupVMStatistics_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/statistics/virtual-machines | Get effective VMs for the Group
[**GetGroupsForIPAddress**](PolicyInventoryApi.md#GetGroupsForIPAddress) | **Get** /global-infra/ip-address-group-associations | Get groups for which the given IP address is a member
[**GetGroupsForIPAddress_0**](PolicyInventoryApi.md#GetGroupsForIPAddress_0) | **Get** /infra/ip-address-group-associations | Get groups for which the given IP address is a member
[**GetGroupsForObject**](PolicyInventoryApi.md#GetGroupsForObject) | **Get** /infra/group-associations | Get groups for which the given object is a member
[**GetGroupsForObject_0**](PolicyInventoryApi.md#GetGroupsForObject_0) | **Get** /global-infra/group-associations | Get groups for which the given object is a member
[**GetGroupsForVIF**](PolicyInventoryApi.md#GetGroupsForVIF) | **Get** /infra/virtual-network-interface-group-associations | Get groups for which the given VIF is a member
[**GetGroupsForVIF_0**](PolicyInventoryApi.md#GetGroupsForVIF_0) | **Get** /global-infra/virtual-network-interface-group-associations | Get groups for which the given VIF is a member
[**GetGroupsForVM**](PolicyInventoryApi.md#GetGroupsForVM) | **Get** /global-infra/virtual-machine-group-associations | Get groups for which the given VM is a member
[**GetGroupsForVM_0**](PolicyInventoryApi.md#GetGroupsForVM_0) | **Get** /infra/virtual-machine-group-associations | Get groups for which the given VM is a member
[**GetMemberTypesForGroup**](PolicyInventoryApi.md#GetMemberTypesForGroup) | **Get** /infra/domains/{domain-id}/groups/{group-id}/member-types | Get member types for a given Group
[**GetMemberTypesForGroup_0**](PolicyInventoryApi.md#GetMemberTypesForGroup_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/member-types | Get member types for a given Group
[**GetPolicyContextProfile**](PolicyInventoryApi.md#GetPolicyContextProfile) | **Get** /infra/context-profiles/{context-profile-id} | Get PolicyContextProfile
[**GetPolicyContextProfile_0**](PolicyInventoryApi.md#GetPolicyContextProfile_0) | **Get** /global-infra/context-profiles/{context-profile-id} | Get PolicyContextProfile
[**GetProviderGroupIPMembers**](PolicyInventoryApi.md#GetProviderGroupIPMembers) | **Get** /global-infra/tier-0s/{tier-0-id}/groups/{group-id}/members/ip-addresses | Get IP addresses that belong to this Tier-0 Group
[**GetProviderGroupIPMembers_0**](PolicyInventoryApi.md#GetProviderGroupIPMembers_0) | **Get** /infra/tier-0s/{tier-0-id}/groups/{group-id}/members/ip-addresses | Get IP addresses that belong to this Tier-0 Group
[**GetProviderGroupVMMembers**](PolicyInventoryApi.md#GetProviderGroupVMMembers) | **Get** /global-infra/tier-0s/{tier-0-id}/groups/{group-id}/members/virtual-machines | Get Virtual machines that belong to this Tier-0 Group
[**GetProviderGroupVMMembers_0**](PolicyInventoryApi.md#GetProviderGroupVMMembers_0) | **Get** /infra/tier-0s/{tier-0-id}/groups/{group-id}/members/virtual-machines | Get Virtual machines that belong to this Tier-0 Group
[**ListGroupForDomain**](PolicyInventoryApi.md#ListGroupForDomain) | **Get** /global-infra/domains/{domain-id}/groups | List Groups for a domain
[**ListGroupForDomain_0**](PolicyInventoryApi.md#ListGroupForDomain_0) | **Get** /infra/domains/{domain-id}/groups | List Groups for a domain
[**ListPolicyContextProfiles**](PolicyInventoryApi.md#ListPolicyContextProfiles) | **Get** /infra/context-profiles | Get PolicyContextProfiles
[**ListPolicyContextProfiles_0**](PolicyInventoryApi.md#ListPolicyContextProfiles_0) | **Get** /global-infra/context-profiles | Get PolicyContextProfiles
[**ListProfileSupportedAttributes**](PolicyInventoryApi.md#ListProfileSupportedAttributes) | **Get** /infra/context-profiles/attributes | List Policy Context Profile supported attributes and sub-attributes
[**ListProfileSupportedAttributes_0**](PolicyInventoryApi.md#ListProfileSupportedAttributes_0) | **Get** /global-infra/context-profiles/attributes | List Policy Context Profile supported attributes and sub-attributes
[**ListServiceEntries**](PolicyInventoryApi.md#ListServiceEntries) | **Get** /infra/services/{service-id}/service-entries | List Service entries for the given service
[**ListServiceEntries_0**](PolicyInventoryApi.md#ListServiceEntries_0) | **Get** /global-infra/services/{service-id}/service-entries | List Service entries for the given service
[**ListServicesForTenant**](PolicyInventoryApi.md#ListServicesForTenant) | **Get** /global-infra/services | List Services for infra
[**ListServicesForTenant_0**](PolicyInventoryApi.md#ListServicesForTenant_0) | **Get** /infra/services | List Services for infra
[**ListTier0Group**](PolicyInventoryApi.md#ListTier0Group) | **Get** /global-infra/tier-0s/{tier-0-id}/groups | List Groups for Tier-0
[**ListTier0Group_0**](PolicyInventoryApi.md#ListTier0Group_0) | **Get** /infra/tier-0s/{tier-0-id}/groups | List Groups for Tier-0
[**PatchCreateOrUpdatePolicyContextProfile**](PolicyInventoryApi.md#PatchCreateOrUpdatePolicyContextProfile) | **Patch** /infra/context-profiles/{context-profile-id} | Create PolicyContextProfile
[**PatchCreateOrUpdatePolicyContextProfile_0**](PolicyInventoryApi.md#PatchCreateOrUpdatePolicyContextProfile_0) | **Patch** /global-infra/context-profiles/{context-profile-id} | Create PolicyContextProfile
[**PatchGroupExternalIDExpressionForDomain**](PolicyInventoryApi.md#PatchGroupExternalIDExpressionForDomain) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Patch a group external ID expression
[**PatchGroupExternalIDExpressionForDomain_0**](PolicyInventoryApi.md#PatchGroupExternalIDExpressionForDomain_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/external-id-expressions/{expression-id} | Patch a group external ID expression
[**PatchGroupForDomain**](PolicyInventoryApi.md#PatchGroupForDomain) | **Patch** /infra/domains/{domain-id}/groups/{group-id} | Patch a group
[**PatchGroupForDomain_0**](PolicyInventoryApi.md#PatchGroupForDomain_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id} | Patch a group
[**PatchGroupIPAddressExpressionForDomain**](PolicyInventoryApi.md#PatchGroupIPAddressExpressionForDomain) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Patch a group IP Address expression
[**PatchGroupIPAddressExpressionForDomain_0**](PolicyInventoryApi.md#PatchGroupIPAddressExpressionForDomain_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/ip-address-expressions/{expression-id} | Patch a group IP Address expression
[**PatchGroupMACAddressExpressionForDomain**](PolicyInventoryApi.md#PatchGroupMACAddressExpressionForDomain) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Patch a group MAC Address expression
[**PatchGroupMACAddressExpressionForDomain_0**](PolicyInventoryApi.md#PatchGroupMACAddressExpressionForDomain_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/mac-address-expressions/{expression-id} | Patch a group MAC Address expression
[**PatchGroupPathExpressionForDomain**](PolicyInventoryApi.md#PatchGroupPathExpressionForDomain) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Patch a group path expression
[**PatchGroupPathExpressionForDomain_0**](PolicyInventoryApi.md#PatchGroupPathExpressionForDomain_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/path-expressions/{expression-id} | Patch a group path expression
[**PatchServiceEntry**](PolicyInventoryApi.md#PatchServiceEntry) | **Patch** /global-infra/services/{service-id}/service-entries/{service-entry-id} | Patch a ServiceEntry
[**PatchServiceEntry_0**](PolicyInventoryApi.md#PatchServiceEntry_0) | **Patch** /infra/services/{service-id}/service-entries/{service-entry-id} | Patch a ServiceEntry
[**PatchServiceForTenant**](PolicyInventoryApi.md#PatchServiceForTenant) | **Patch** /global-infra/services/{service-id} | Patch a Service
[**PatchServiceForTenant_0**](PolicyInventoryApi.md#PatchServiceForTenant_0) | **Patch** /infra/services/{service-id} | Patch a Service
[**PatchTier0Group**](PolicyInventoryApi.md#PatchTier0Group) | **Patch** /infra/tier-0s/{tier-0-id}/groups/{group-id} | Create or update Group under specified Tier-0
[**PatchTier0Group_0**](PolicyInventoryApi.md#PatchTier0Group_0) | **Patch** /global-infra/tier-0s/{tier-0-id}/groups/{group-id} | Create or update Group under specified Tier-0
[**PutCreateOrUpdatePolicyContextProfile**](PolicyInventoryApi.md#PutCreateOrUpdatePolicyContextProfile) | **Put** /infra/context-profiles/{context-profile-id} | Create PolicyContextProfile
[**PutCreateOrUpdatePolicyContextProfile_0**](PolicyInventoryApi.md#PutCreateOrUpdatePolicyContextProfile_0) | **Put** /global-infra/context-profiles/{context-profile-id} | Create PolicyContextProfile
[**ReadGroupForDomain**](PolicyInventoryApi.md#ReadGroupForDomain) | **Get** /infra/domains/{domain-id}/groups/{group-id} | Read group
[**ReadGroupForDomain_0**](PolicyInventoryApi.md#ReadGroupForDomain_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id} | Read group
[**ReadServiceEntry**](PolicyInventoryApi.md#ReadServiceEntry) | **Get** /global-infra/services/{service-id}/service-entries/{service-entry-id} | Service entry
[**ReadServiceEntry_0**](PolicyInventoryApi.md#ReadServiceEntry_0) | **Get** /infra/services/{service-id}/service-entries/{service-entry-id} | Service entry
[**ReadServiceForTenant**](PolicyInventoryApi.md#ReadServiceForTenant) | **Get** /global-infra/services/{service-id} | Read a service
[**ReadServiceForTenant_0**](PolicyInventoryApi.md#ReadServiceForTenant_0) | **Get** /infra/services/{service-id} | Read a service
[**ReadTier0Group**](PolicyInventoryApi.md#ReadTier0Group) | **Get** /infra/tier-0s/{tier-0-id}/groups/{group-id} | Read Tier-0 Group
[**ReadTier0Group_0**](PolicyInventoryApi.md#ReadTier0Group_0) | **Get** /global-infra/tier-0s/{tier-0-id}/groups/{group-id} | Read Tier-0 Group
[**UpdateGroupForDomain**](PolicyInventoryApi.md#UpdateGroupForDomain) | **Put** /infra/domains/{domain-id}/groups/{group-id} | Create or update a group
[**UpdateGroupForDomain_0**](PolicyInventoryApi.md#UpdateGroupForDomain_0) | **Put** /global-infra/domains/{domain-id}/groups/{group-id} | Create or update a group
[**UpdateServiceEntry**](PolicyInventoryApi.md#UpdateServiceEntry) | **Put** /global-infra/services/{service-id}/service-entries/{service-entry-id} | Create or update a ServiceEntry
[**UpdateServiceEntry_0**](PolicyInventoryApi.md#UpdateServiceEntry_0) | **Put** /infra/services/{service-id}/service-entries/{service-entry-id} | Create or update a ServiceEntry
[**UpdateServiceForTenant**](PolicyInventoryApi.md#UpdateServiceForTenant) | **Put** /global-infra/services/{service-id} | Create or update a Service
[**UpdateServiceForTenant_0**](PolicyInventoryApi.md#UpdateServiceForTenant_0) | **Put** /infra/services/{service-id} | Create or update a Service

# **AddorRemoveGroupExternalIDMembers**
> AddorRemoveGroupExternalIDMembers(ctx, body, domainId, groupId, expressionId, action)
Add or Remove external id based members from/to a Group 

It will add or remove the specified members having external ID for a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMemberList**](GroupMemberList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupExternalIDMembers_0**
> AddorRemoveGroupExternalIDMembers_0(ctx, body, domainId, groupId, expressionId, action)
Add or Remove external id based members from/to a Group 

It will add or remove the specified members having external ID for a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMemberList**](GroupMemberList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupIPAddresses**
> AddorRemoveGroupIPAddresses(ctx, body, domainId, groupId, expressionId, action)
Add or Remove IP Addresses from/to a Group 

It will add or remove the specified IP Addresses from a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpAddressList**](IpAddressList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupIPAddresses_0**
> AddorRemoveGroupIPAddresses_0(ctx, body, domainId, groupId, expressionId, action)
Add or Remove IP Addresses from/to a Group 

It will add or remove the specified IP Addresses from a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpAddressList**](IpAddressList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupMACAddresses**
> AddorRemoveGroupMACAddresses(ctx, body, domainId, groupId, expressionId, action)
Add or Remove MAC Addresses from/to a Group 

It will add or remove the specified MAC Addresses from a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacAddressList**](MacAddressList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupMACAddresses_0**
> AddorRemoveGroupMACAddresses_0(ctx, body, domainId, groupId, expressionId, action)
Add or Remove MAC Addresses from/to a Group 

It will add or remove the specified MAC Addresses from a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacAddressList**](MacAddressList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupPathMembers**
> AddorRemoveGroupPathMembers(ctx, body, domainId, groupId, expressionId, action)
Add or Remove path based members from/to a Group 

It will add or remove the specified members having path for a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMemberList**](GroupMemberList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddorRemoveGroupPathMembers_0**
> AddorRemoveGroupPathMembers_0(ctx, body, domainId, groupId, expressionId, action)
Add or Remove path based members from/to a Group 

It will add or remove the specified members having path for a given expression of a group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMemberList**](GroupMemberList.md)|  | 
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
  **expressionId** | **string**|  | 
  **action** | **string**| Add or Remove group members. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceTier0Group**
> Group CreateOrReplaceTier0Group(ctx, body, tier0Id, groupId)
Create or update Group under specified Tier-0

If a Group with the group-id is not already present, create a new Group under the tier-0-id. Update if exists. The API valiates that Tier-0 is present before creating the Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceTier0Group_0**
> Group CreateOrReplaceTier0Group_0(ctx, body, tier0Id, groupId)
Create or update Group under specified Tier-0

If a Group with the group-id is not already present, create a new Group under the tier-0-id. Update if exists. The API valiates that Tier-0 is present before creating the Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, domainId, groupId, optional)
Delete Group

Delete Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
 **optional** | ***PolicyInventoryApiDeleteGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiDeleteGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failIfSubtreeExists** | **optional.Bool**| Do not delete if the group subtree has any entities | [default to false]
 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupExternalIDExpression**
> DeleteGroupExternalIDExpression(ctx, domainId, groupId, expressionId)
Delete Group External ID Expression

Delete Group External ID Expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| ExternalIDExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupExternalIDExpression_0**
> DeleteGroupExternalIDExpression_0(ctx, domainId, groupId, expressionId)
Delete Group External ID Expression

Delete Group External ID Expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| ExternalIDExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupIPAddressExpression**
> DeleteGroupIPAddressExpression(ctx, domainId, groupId, expressionId)
Delete Group IPAddressExpression

Delete Group IPAddressExpression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| IPAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupIPAddressExpression_0**
> DeleteGroupIPAddressExpression_0(ctx, domainId, groupId, expressionId)
Delete Group IPAddressExpression

Delete Group IPAddressExpression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| IPAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMACAddressExpression**
> DeleteGroupMACAddressExpression(ctx, domainId, groupId, expressionId)
Delete Group MACAddressExpression

Delete Group MACAddressExpression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| MACAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMACAddressExpression_0**
> DeleteGroupMACAddressExpression_0(ctx, domainId, groupId, expressionId)
Delete Group MACAddressExpression

Delete Group MACAddressExpression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| MACAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupPathExpression**
> DeleteGroupPathExpression(ctx, domainId, groupId, expressionId)
Delete Group Path Expression

Delete Group Path Expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| PathExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupPathExpression_0**
> DeleteGroupPathExpression_0(ctx, domainId, groupId, expressionId)
Delete Group Path Expression

Delete Group Path Expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| PathExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup_0**
> DeleteGroup_0(ctx, domainId, groupId, optional)
Delete Group

Delete Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
 **optional** | ***PolicyInventoryApiDeleteGroup_10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiDeleteGroup_10Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failIfSubtreeExists** | **optional.Bool**| Do not delete if the group subtree has any entities | [default to false]
 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyContextProfile**
> DeletePolicyContextProfile(ctx, contextProfileId, optional)
Delete Policy Context Profile

Deletes the specified Policy Context Profile. If the Policy Context Profile is consumed in a firewall rule, it won't get deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextProfileId** | **string**| Policy Context Profile Id | 
 **optional** | ***PolicyInventoryApiDeletePolicyContextProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiDeletePolicyContextProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyContextProfile_0**
> DeletePolicyContextProfile_0(ctx, contextProfileId, optional)
Delete Policy Context Profile

Deletes the specified Policy Context Profile. If the Policy Context Profile is consumed in a firewall rule, it won't get deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextProfileId** | **string**| Policy Context Profile Id | 
 **optional** | ***PolicyInventoryApiDeletePolicyContextProfile_11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiDeletePolicyContextProfile_11Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceEntry**
> DeleteServiceEntry(ctx, serviceId, serviceEntryId)
Delete Service entry

Delete Service entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceEntry_0**
> DeleteServiceEntry_0(ctx, serviceId, serviceEntryId)
Delete Service entry

Delete Service entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceForTenant**
> DeleteServiceForTenant(ctx, serviceId)
Delete Service

Delete Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceForTenant_0**
> DeleteServiceForTenant_0(ctx, serviceId)
Delete Service

Delete Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0Group**
> DeleteTier0Group(ctx, tier0Id, groupId)
Deletes Group under Tier-0

Delete the Group under Tier-0. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0Group_0**
> DeleteTier0Group_0(ctx, tier0Id, groupId)
Deletes Group under Tier-0

Delete the Group under Tier-0. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupIPMembers**
> PolicyGroupIpMembersListResult GetGroupIPMembers(ctx, domainId, groupId, optional)
Get IP addresses that belong to this Group

Get IP addresses that belong to this Group. This API is applicable for Groups containing either VirtualMachine, VIF, Segment ,Segment Port or IP Address member type.For Groups containing other member types,an empty list is returned 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupIPMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupIPMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupIpMembersListResult**](PolicyGroupIPMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupIPMembers_0**
> PolicyGroupIpMembersListResult GetGroupIPMembers_0(ctx, domainId, groupId, optional)
Get IP addresses that belong to this Group

Get IP addresses that belong to this Group. This API is applicable for Groups containing either VirtualMachine, VIF, Segment ,Segment Port or IP Address member type.For Groups containing other member types,an empty list is returned 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupIPMembers_15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupIPMembers_15Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupIpMembersListResult**](PolicyGroupIPMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupLPMembers**
> PolicyGroupMembersListResult GetGroupLPMembers(ctx, domainId, groupId, optional)
Get logical ports that belong to this Group

Get logical ports that belong to this Group This API is applicable for Groups containing either VirtualMachine, VIF, Segment or Segment Port member type.For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupLPMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupLPMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupLPMembers_0**
> PolicyGroupMembersListResult GetGroupLPMembers_0(ctx, domainId, groupId, optional)
Get logical ports that belong to this Group

Get logical ports that belong to this Group This API is applicable for Groups containing either VirtualMachine, VIF, Segment or Segment Port member type.For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupLPMembers_16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupLPMembers_16Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupLSMembers**
> PolicyGroupMembersListResult GetGroupLSMembers(ctx, domainId, groupId, optional)
Get logical switches that belong to this Group

Get logical switches that belong to this Group. This API is applicable for Groups containing Segment member type. For Groups containing other member types, an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupLSMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupLSMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupLSMembers_0**
> PolicyGroupMembersListResult GetGroupLSMembers_0(ctx, domainId, groupId, optional)
Get logical switches that belong to this Group

Get logical switches that belong to this Group. This API is applicable for Groups containing Segment member type. For Groups containing other member types, an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupLSMembers_17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupLSMembers_17Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupSegmentMembers**
> PolicyGroupMembersListResult GetGroupSegmentMembers(ctx, domainId, groupId, optional)
Get segments that belong to this Group

Get segments that belong to this Group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupSegmentMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupSegmentMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupSegmentMembers_0**
> PolicyGroupMembersListResult GetGroupSegmentMembers_0(ctx, domainId, groupId, optional)
Get segments that belong to this Group

Get segments that belong to this Group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupSegmentMembers_18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupSegmentMembers_18Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupSegmentPortMembers**
> PolicyGroupMembersListResult GetGroupSegmentPortMembers(ctx, domainId, groupId, optional)
Get segment ports that belong to this Group

Get segment ports that belong to this Group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupSegmentPortMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupSegmentPortMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupSegmentPortMembers_0**
> PolicyGroupMembersListResult GetGroupSegmentPortMembers_0(ctx, domainId, groupId, optional)
Get segment ports that belong to this Group

Get segment ports that belong to this Group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupSegmentPortMembers_19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupSegmentPortMembers_19Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupMembersListResult**](PolicyGroupMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTags**
> GroupTagsList GetGroupTags(ctx, domainId, groupId)
Get tags used to define conditions inside a Group

Get tags used to define conditions inside a Group. Also includes tags inside nested groups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 

### Return type

[**GroupTagsList**](GroupTagsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupTags_0**
> GroupTagsList GetGroupTags_0(ctx, domainId, groupId)
Get tags used to define conditions inside a Group

Get tags used to define conditions inside a Group. Also includes tags inside nested groups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 

### Return type

[**GroupTagsList**](GroupTagsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVIFMembers**
> VirtualNetworkInterfaceListResult GetGroupVIFMembers(ctx, domainId, groupId, optional)
Get Virtual Network Interface instances that belong to this Group

Get Virtual Network Interface instances that belong to this Group. This API is applicable for Groups containing VirtualNetworkInterface and VirtualMachine member types. For Groups containing other member types,an empty list is returned.target_id in response is external_id of VirtualNetworkInterface or VirtualMachine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVIFMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVIFMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualNetworkInterfaceListResult**](VirtualNetworkInterfaceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVIFMembers_0**
> VirtualNetworkInterfaceListResult GetGroupVIFMembers_0(ctx, domainId, groupId, optional)
Get Virtual Network Interface instances that belong to this Group

Get Virtual Network Interface instances that belong to this Group. This API is applicable for Groups containing VirtualNetworkInterface and VirtualMachine member types. For Groups containing other member types,an empty list is returned.target_id in response is external_id of VirtualNetworkInterface or VirtualMachine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVIFMembers_21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVIFMembers_21Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualNetworkInterfaceListResult**](VirtualNetworkInterfaceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVMMembers**
> RealizedVirtualMachineListResult GetGroupVMMembers(ctx, domainId, groupId, optional)
Get Virtual machines that belong to this Group

Get Virtual machines that belong to this Group. This API is applicable for Groups containing VirtualMachine,member type. For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVMMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVMMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVMMembers_0**
> RealizedVirtualMachineListResult GetGroupVMMembers_0(ctx, domainId, groupId, optional)
Get Virtual machines that belong to this Group

Get Virtual machines that belong to this Group. This API is applicable for Groups containing VirtualMachine,member type. For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVMMembers_22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVMMembers_22Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVMStatistics**
> RealizedVirtualMachineListResult GetGroupVMStatistics(ctx, domainId, groupId, optional)
Get effective VMs for the Group

Get the effective VM membership for the Group. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVMStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVMStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVMStatistics_0**
> RealizedVirtualMachineListResult GetGroupVMStatistics_0(ctx, domainId, groupId, optional)
Get effective VMs for the Group

Get the effective VM membership for the Group. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetGroupVMStatistics_23Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupVMStatistics_23Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForIPAddress**
> PolicyResourceReferenceForEpListResult GetGroupsForIPAddress(ctx, ipAddress, optional)
Get groups for which the given IP address is a member

Get policy groups for which the given IP address is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| IPAddress | 
 **optional** | ***PolicyInventoryApiGetGroupsForIPAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForIPAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForIPAddress_0**
> PolicyResourceReferenceForEpListResult GetGroupsForIPAddress_0(ctx, ipAddress, optional)
Get groups for which the given IP address is a member

Get policy groups for which the given IP address is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| IPAddress | 
 **optional** | ***PolicyInventoryApiGetGroupsForIPAddress_24Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForIPAddress_24Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForObject**
> PolicyResourceReferenceForEpListResult GetGroupsForObject(ctx, intentPath, optional)
Get groups for which the given object is a member

Get policy groups for which the given object is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String path of the intent object | 
 **optional** | ***PolicyInventoryApiGetGroupsForObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForObjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForObject_0**
> PolicyResourceReferenceForEpListResult GetGroupsForObject_0(ctx, intentPath, optional)
Get groups for which the given object is a member

Get policy groups for which the given object is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String path of the intent object | 
 **optional** | ***PolicyInventoryApiGetGroupsForObject_25Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForObject_25Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForVIF**
> PolicyResourceReferenceForEpListResult GetGroupsForVIF(ctx, vifExternalId, optional)
Get groups for which the given VIF is a member

Get policy groups for which the given VIF is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vifExternalId** | **string**| Virtual network interface external ID | 
 **optional** | ***PolicyInventoryApiGetGroupsForVIFOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForVIFOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForVIF_0**
> PolicyResourceReferenceForEpListResult GetGroupsForVIF_0(ctx, vifExternalId, optional)
Get groups for which the given VIF is a member

Get policy groups for which the given VIF is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vifExternalId** | **string**| Virtual network interface external ID | 
 **optional** | ***PolicyInventoryApiGetGroupsForVIF_26Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForVIF_26Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForVM**
> PolicyResourceReferenceForEpListResult GetGroupsForVM(ctx, vmExternalId, optional)
Get groups for which the given VM is a member

Get policy groups for which the given VM is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmExternalId** | **string**| Virtual machine external ID | 
 **optional** | ***PolicyInventoryApiGetGroupsForVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForVMOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupsForVM_0**
> PolicyResourceReferenceForEpListResult GetGroupsForVM_0(ctx, vmExternalId, optional)
Get groups for which the given VM is a member

Get policy groups for which the given VM is a member. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmExternalId** | **string**| Virtual machine external ID | 
 **optional** | ***PolicyInventoryApiGetGroupsForVM_27Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetGroupsForVM_27Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceForEpListResult**](PolicyResourceReferenceForEPListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberTypesForGroup**
> GroupMemberTypeListResult GetMemberTypesForGroup(ctx, domainId, groupId)
Get member types for a given Group

It retrieves member types for a given group. In case of nested groups, it calculates member types of child groups as well. Considers member type for members added via static members and dynamic membership criteria. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**GroupMemberTypeListResult**](GroupMemberTypeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberTypesForGroup_0**
> GroupMemberTypeListResult GetMemberTypesForGroup_0(ctx, domainId, groupId)
Get member types for a given Group

It retrieves member types for a given group. In case of nested groups, it calculates member types of child groups as well. Considers member type for members added via static members and dynamic membership criteria. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**GroupMemberTypeListResult**](GroupMemberTypeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyContextProfile**
> PolicyContextProfile GetPolicyContextProfile(ctx, contextProfileId)
Get PolicyContextProfile

Get a single PolicyContextProfile by id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextProfileId** | **string**|  | 

### Return type

[**PolicyContextProfile**](PolicyContextProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyContextProfile_0**
> PolicyContextProfile GetPolicyContextProfile_0(ctx, contextProfileId)
Get PolicyContextProfile

Get a single PolicyContextProfile by id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextProfileId** | **string**|  | 

### Return type

[**PolicyContextProfile**](PolicyContextProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProviderGroupIPMembers**
> PolicyGroupIpMembersListResult GetProviderGroupIPMembers(ctx, tier0Id, groupId, optional)
Get IP addresses that belong to this Tier-0 Group

Get IP addresses that belong to this Tier-0 Group. This API is applicable for Groups containing either VirtualMachine, VIF, Segment ,Segment Port or IP Address member type.For Groups containing other member types,an empty list is returned 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetProviderGroupIPMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetProviderGroupIPMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupIpMembersListResult**](PolicyGroupIPMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProviderGroupIPMembers_0**
> PolicyGroupIpMembersListResult GetProviderGroupIPMembers_0(ctx, tier0Id, groupId, optional)
Get IP addresses that belong to this Tier-0 Group

Get IP addresses that belong to this Tier-0 Group. This API is applicable for Groups containing either VirtualMachine, VIF, Segment ,Segment Port or IP Address member type.For Groups containing other member types,an empty list is returned 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetProviderGroupIPMembers_30Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetProviderGroupIPMembers_30Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyGroupIpMembersListResult**](PolicyGroupIPMembersListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProviderGroupVMMembers**
> RealizedVirtualMachineListResult GetProviderGroupVMMembers(ctx, tier0Id, groupId, optional)
Get Virtual machines that belong to this Tier-0 Group

Get Virtual machines that belong to this Tier-0 Group. This API is applicable for Groups containing VirtualMachine member type. For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetProviderGroupVMMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetProviderGroupVMMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProviderGroupVMMembers_0**
> RealizedVirtualMachineListResult GetProviderGroupVMMembers_0(ctx, tier0Id, groupId, optional)
Get Virtual machines that belong to this Tier-0 Group

Get Virtual machines that belong to this Tier-0 Group. This API is applicable for Groups containing VirtualMachine member type. For Groups containing other member types,an empty list is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **groupId** | **string**| Group Id | 
 **optional** | ***PolicyInventoryApiGetProviderGroupVMMembers_31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiGetProviderGroupVMMembers_31Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedVirtualMachineListResult**](RealizedVirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupForDomain**
> GroupListResult ListGroupForDomain(ctx, domainId, optional)
List Groups for a domain

List Groups for a domain. Groups can be filtered using member_types query parameter, which returns the groups that contains the specified member types. Multiple member types can be provided as comma separated values. The API also return groups having member type that are subset of provided member_types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
 **optional** | ***PolicyInventoryApiListGroupForDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListGroupForDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **memberTypes** | **optional.String**| Comma Seperated Member types | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupListResult**](GroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupForDomain_0**
> GroupListResult ListGroupForDomain_0(ctx, domainId, optional)
List Groups for a domain

List Groups for a domain. Groups can be filtered using member_types query parameter, which returns the groups that contains the specified member types. Multiple member types can be provided as comma separated values. The API also return groups having member type that are subset of provided member_types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
 **optional** | ***PolicyInventoryApiListGroupForDomain_32Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListGroupForDomain_32Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **memberTypes** | **optional.String**| Comma Seperated Member types | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupListResult**](GroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyContextProfiles**
> PolicyContextProfileListResult ListPolicyContextProfiles(ctx, optional)
Get PolicyContextProfiles

Get all PolicyContextProfiles 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListPolicyContextProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListPolicyContextProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyContextProfileListResult**](PolicyContextProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyContextProfiles_0**
> PolicyContextProfileListResult ListPolicyContextProfiles_0(ctx, optional)
Get PolicyContextProfiles

Get all PolicyContextProfiles 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListPolicyContextProfiles_33Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListPolicyContextProfiles_33Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyContextProfileListResult**](PolicyContextProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProfileSupportedAttributes**
> PolicyContextProfileListResult ListProfileSupportedAttributes(ctx, optional)
List Policy Context Profile supported attributes and sub-attributes

Returns supported attribute and sub-attributes for specified attribute key with their supported values, if provided in query/request parameter, else will fetch all supported attributes and sub-attributes for all supported attribute keys. Alternatively, to get a list of supported attributes and sub-attributes fire the following REST API GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListProfileSupportedAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListProfileSupportedAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeKey** | **optional.String**| Fetch attributes and sub-attributes for the given attribute key | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyContextProfileListResult**](PolicyContextProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProfileSupportedAttributes_0**
> PolicyContextProfileListResult ListProfileSupportedAttributes_0(ctx, optional)
List Policy Context Profile supported attributes and sub-attributes

Returns supported attribute and sub-attributes for specified attribute key with their supported values, if provided in query/request parameter, else will fetch all supported attributes and sub-attributes for all supported attribute keys. Alternatively, to get a list of supported attributes and sub-attributes fire the following REST API GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListProfileSupportedAttributes_34Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListProfileSupportedAttributes_34Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeKey** | **optional.String**| Fetch attributes and sub-attributes for the given attribute key | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyContextProfileListResult**](PolicyContextProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceEntries**
> ServiceEntryListResult ListServiceEntries(ctx, serviceId, optional)
List Service entries for the given service

Paginated list of Service entries for the given service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
 **optional** | ***PolicyInventoryApiListServiceEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListServiceEntriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceEntryListResult**](ServiceEntryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceEntries_0**
> ServiceEntryListResult ListServiceEntries_0(ctx, serviceId, optional)
List Service entries for the given service

Paginated list of Service entries for the given service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
 **optional** | ***PolicyInventoryApiListServiceEntries_35Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListServiceEntries_35Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceEntryListResult**](ServiceEntryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServicesForTenant**
> ServiceListResult ListServicesForTenant(ctx, optional)
List Services for infra

Paginated list of Services for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListServicesForTenantOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListServicesForTenantOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **defaultService** | **optional.Bool**| Fetch all default services | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceListResult**](ServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServicesForTenant_0**
> ServiceListResult ListServicesForTenant_0(ctx, optional)
List Services for infra

Paginated list of Services for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInventoryApiListServicesForTenant_36Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListServicesForTenant_36Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **defaultService** | **optional.Bool**| Fetch all default services | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceListResult**](ServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTier0Group**
> GroupListResult ListTier0Group(ctx, tier0Id, optional)
List Groups for Tier-0

Paginated list of all Groups for Tier-0. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
 **optional** | ***PolicyInventoryApiListTier0GroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListTier0GroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **memberTypes** | **optional.String**| Comma Seperated Member types | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupListResult**](GroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTier0Group_0**
> GroupListResult ListTier0Group_0(ctx, tier0Id, optional)
List Groups for Tier-0

Paginated list of all Groups for Tier-0. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
 **optional** | ***PolicyInventoryApiListTier0Group_37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInventoryApiListTier0Group_37Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **memberTypes** | **optional.String**| Comma Seperated Member types | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupListResult**](GroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCreateOrUpdatePolicyContextProfile**
> PatchCreateOrUpdatePolicyContextProfile(ctx, body, contextProfileId)
Create PolicyContextProfile

Creates/Updates a PolicyContextProfile, which encapsulates attribute and sub-attributes of network services. Rules for using attributes and sub-attributes in single PolicyContextProfile 1. One type of attribute can't have multiple occurrences. ( Eg. -    Attribute type APP_ID can be used only once per PolicyContextProfile.) 2. For specifying multiple values for an attribute, provide them in an array. 3. If sub-attribtes are mentioned for an attribute, then only single    value is allowed for that attribute. 4. To get a list of supported attributes and sub-attributes fire the following REST API    GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyContextProfile**](PolicyContextProfile.md)|  | 
  **contextProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCreateOrUpdatePolicyContextProfile_0**
> PatchCreateOrUpdatePolicyContextProfile_0(ctx, body, contextProfileId)
Create PolicyContextProfile

Creates/Updates a PolicyContextProfile, which encapsulates attribute and sub-attributes of network services. Rules for using attributes and sub-attributes in single PolicyContextProfile 1. One type of attribute can't have multiple occurrences. ( Eg. -    Attribute type APP_ID can be used only once per PolicyContextProfile.) 2. For specifying multiple values for an attribute, provide them in an array. 3. If sub-attribtes are mentioned for an attribute, then only single    value is allowed for that attribute. 4. To get a list of supported attributes and sub-attributes fire the following REST API    GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyContextProfile**](PolicyContextProfile.md)|  | 
  **contextProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupExternalIDExpressionForDomain**
> PatchGroupExternalIDExpressionForDomain(ctx, body, domainId, groupId, expressionId)
Patch a group external ID expression

If a group ExternalIDexpression with the expression-id is not already present, create a new ExternalIDexpresison. If it already exists, replace the existing ExternalIDexpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalIdExpression**](ExternalIdExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| ExternalIDExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupExternalIDExpressionForDomain_0**
> PatchGroupExternalIDExpressionForDomain_0(ctx, body, domainId, groupId, expressionId)
Patch a group external ID expression

If a group ExternalIDexpression with the expression-id is not already present, create a new ExternalIDexpresison. If it already exists, replace the existing ExternalIDexpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalIdExpression**](ExternalIdExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| ExternalIDExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupForDomain**
> PatchGroupForDomain(ctx, body, domainId, groupId)
Patch a group

If a group with the group-id is not already present, create a new group. If it already exists, patch the group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupForDomain_0**
> PatchGroupForDomain_0(ctx, body, domainId, groupId)
Patch a group

If a group with the group-id is not already present, create a new group. If it already exists, patch the group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupIPAddressExpressionForDomain**
> PatchGroupIPAddressExpressionForDomain(ctx, body, domainId, groupId, expressionId)
Patch a group IP Address expression

If a group IPAddressExpression with the expression-id is not already present, create a new IPAddressExpression. If it already exists, replace the existing IPAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpAddressExpression**](IpAddressExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| IPAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupIPAddressExpressionForDomain_0**
> PatchGroupIPAddressExpressionForDomain_0(ctx, body, domainId, groupId, expressionId)
Patch a group IP Address expression

If a group IPAddressExpression with the expression-id is not already present, create a new IPAddressExpression. If it already exists, replace the existing IPAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpAddressExpression**](IpAddressExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| IPAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupMACAddressExpressionForDomain**
> PatchGroupMACAddressExpressionForDomain(ctx, body, domainId, groupId, expressionId)
Patch a group MAC Address expression

If a group MACAddressExpression with the expression-id is not already present, create a new MACAddressExpression. If it already exists, replace the existing MACAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacAddressExpression**](MacAddressExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| MACAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupMACAddressExpressionForDomain_0**
> PatchGroupMACAddressExpressionForDomain_0(ctx, body, domainId, groupId, expressionId)
Patch a group MAC Address expression

If a group MACAddressExpression with the expression-id is not already present, create a new MACAddressExpression. If it already exists, replace the existing MACAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacAddressExpression**](MacAddressExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| MACAddressExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupPathExpressionForDomain**
> PatchGroupPathExpressionForDomain(ctx, body, domainId, groupId, expressionId)
Patch a group path expression

If a group path_expression with the expression-id is not already present, create a new pathexpresison. If it already exists, replace the existing pathexpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PathExpression**](PathExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| PathExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupPathExpressionForDomain_0**
> PatchGroupPathExpressionForDomain_0(ctx, body, domainId, groupId, expressionId)
Patch a group path expression

If a group path_expression with the expression-id is not already present, create a new pathexpresison. If it already exists, replace the existing pathexpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PathExpression**](PathExpression.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **expressionId** | **string**| PathExpression ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceEntry**
> PatchServiceEntry(ctx, body, serviceId, serviceEntryId)
Patch a ServiceEntry

If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, patch the service entry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceEntry**](ServiceEntry.md)|  | 
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceEntry_0**
> PatchServiceEntry_0(ctx, body, serviceId, serviceEntryId)
Patch a ServiceEntry

If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, patch the service entry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceEntry**](ServiceEntry.md)|  | 
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceForTenant**
> PatchServiceForTenant(ctx, body, serviceId)
Patch a Service

Create a new service if a service with the given ID does not already exist. Creates new service entries if populated in the service. If a service with the given ID already exists, patch the service including the nested service entries. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Service**](Service.md)|  | 
  **serviceId** | **string**| Service ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceForTenant_0**
> PatchServiceForTenant_0(ctx, body, serviceId)
Patch a Service

Create a new service if a service with the given ID does not already exist. Creates new service entries if populated in the service. If a service with the given ID already exists, patch the service including the nested service entries. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Service**](Service.md)|  | 
  **serviceId** | **string**| Service ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0Group**
> PatchTier0Group(ctx, body, tier0Id, groupId)
Create or update Group under specified Tier-0

If a Group with the group-id is not already present, create a new Group under the tier-0-id. Update if exists. The API valiates that Tier-0 is present before creating the Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0Group_0**
> PatchTier0Group_0(ctx, body, tier0Id, groupId)
Create or update Group under specified Tier-0

If a Group with the group-id is not already present, create a new Group under the tier-0-id. Update if exists. The API valiates that Tier-0 is present before creating the Group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreateOrUpdatePolicyContextProfile**
> PolicyContextProfile PutCreateOrUpdatePolicyContextProfile(ctx, body, contextProfileId)
Create PolicyContextProfile

Creates/Updates a PolicyContextProfile, which encapsulates attribute and sub-attributes of network services. Rules for using attributes and sub-attributes in single PolicyContextProfile 1. One type of attribute can't have multiple occurrences. ( Eg. -    Attribute type APP_ID can be used only once per PolicyContextProfile.) 2. For specifying multiple values for an attribute, provide them in an array. 3. If sub-attribtes are mentioned for an attribute, then only single    value is allowed for that attribute. 4. To get a list of supported attributes and sub-attributes fire the following REST API    GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyContextProfile**](PolicyContextProfile.md)|  | 
  **contextProfileId** | **string**|  | 

### Return type

[**PolicyContextProfile**](PolicyContextProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreateOrUpdatePolicyContextProfile_0**
> PolicyContextProfile PutCreateOrUpdatePolicyContextProfile_0(ctx, body, contextProfileId)
Create PolicyContextProfile

Creates/Updates a PolicyContextProfile, which encapsulates attribute and sub-attributes of network services. Rules for using attributes and sub-attributes in single PolicyContextProfile 1. One type of attribute can't have multiple occurrences. ( Eg. -    Attribute type APP_ID can be used only once per PolicyContextProfile.) 2. For specifying multiple values for an attribute, provide them in an array. 3. If sub-attribtes are mentioned for an attribute, then only single    value is allowed for that attribute. 4. To get a list of supported attributes and sub-attributes fire the following REST API    GET https://&lt;policy-mgr&gt;/policy/api/v1/infra/context-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyContextProfile**](PolicyContextProfile.md)|  | 
  **contextProfileId** | **string**|  | 

### Return type

[**PolicyContextProfile**](PolicyContextProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGroupForDomain**
> Group ReadGroupForDomain(ctx, domainId, groupId)
Read group

Read group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGroupForDomain_0**
> Group ReadGroupForDomain_0(ctx, domainId, groupId)
Read group

Read group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceEntry**
> ServiceEntry ReadServiceEntry(ctx, serviceId, serviceEntryId)
Service entry

Service entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

[**ServiceEntry**](ServiceEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceEntry_0**
> ServiceEntry ReadServiceEntry_0(ctx, serviceId, serviceEntryId)
Service entry

Service entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

[**ServiceEntry**](ServiceEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceForTenant**
> Service ReadServiceForTenant(ctx, serviceId)
Read a service

Read a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

[**Service**](Service.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceForTenant_0**
> Service ReadServiceForTenant_0(ctx, serviceId)
Read a service

Read a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**| Service ID | 

### Return type

[**Service**](Service.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTier0Group**
> Group ReadTier0Group(ctx, tier0Id, groupId)
Read Tier-0 Group

Read Tier-0 Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTier0Group_0**
> Group ReadTier0Group_0(ctx, tier0Id, groupId)
Read Tier-0 Group

Read Tier-0 Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupForDomain**
> Group UpdateGroupForDomain(ctx, body, domainId, groupId)
Create or update a group

If a group with the group-id is not already present, create a new group. If it already exists, update the group. Avoid creating groups with multiple MACAddressExpression and IPAddressExpression. In future releases, group will be restricted to contain a single MACAddressExpression and IPAddressExpression along with other expressions. To group IPAddresses or MACAddresses, use nested groups instead of multiple IPAddressExpressions/MACAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupForDomain_0**
> Group UpdateGroupForDomain_0(ctx, body, domainId, groupId)
Create or update a group

If a group with the group-id is not already present, create a new group. If it already exists, update the group. Avoid creating groups with multiple MACAddressExpression and IPAddressExpression. In future releases, group will be restricted to contain a single MACAddressExpression and IPAddressExpression along with other expressions. To group IPAddresses or MACAddresses, use nested groups instead of multiple IPAddressExpressions/MACAddressExpression. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceEntry**
> ServiceEntry UpdateServiceEntry(ctx, body, serviceId, serviceEntryId)
Create or update a ServiceEntry

If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, update the service entry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceEntry**](ServiceEntry.md)|  | 
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

[**ServiceEntry**](ServiceEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceEntry_0**
> ServiceEntry UpdateServiceEntry_0(ctx, body, serviceId, serviceEntryId)
Create or update a ServiceEntry

If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, update the service entry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceEntry**](ServiceEntry.md)|  | 
  **serviceId** | **string**| Service ID | 
  **serviceEntryId** | **string**| Service entry ID | 

### Return type

[**ServiceEntry**](ServiceEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceForTenant**
> Service UpdateServiceForTenant(ctx, body, serviceId)
Create or update a Service

Create a new service if a service with the given ID does not already exist. Creates new service entries if populated in the service. If a service with the given ID already exists, update the service including the nested service entries. This is a full replace. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Service**](Service.md)|  | 
  **serviceId** | **string**| Service ID | 

### Return type

[**Service**](Service.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceForTenant_0**
> Service UpdateServiceForTenant_0(ctx, body, serviceId)
Create or update a Service

Create a new service if a service with the given ID does not already exist. Creates new service entries if populated in the service. If a service with the given ID already exists, update the service including the nested service entries. This is a full replace. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Service**](Service.md)|  | 
  **serviceId** | **string**| Service ID | 

### Return type

[**Service**](Service.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

