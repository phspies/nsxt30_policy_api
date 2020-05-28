# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateInfraGlobalManager**](FederationInfraApi.md#CreateOrUpdateInfraGlobalManager) | **Put** /infra/global-managers/{global-manager-id} | Create or fully replace a Global Manager under infra
[**CreateOrUpdateInfraGlobalManager_0**](FederationInfraApi.md#CreateOrUpdateInfraGlobalManager_0) | **Put** /global-infra/global-managers/{global-manager-id} | Create or fully replace a Global Manager under infra
[**DeleteInfraGlobalManager**](FederationInfraApi.md#DeleteInfraGlobalManager) | **Delete** /infra/global-managers/{global-manager-id} | Delete a Global Manager under Infra
[**DeleteInfraGlobalManager_0**](FederationInfraApi.md#DeleteInfraGlobalManager_0) | **Delete** /global-infra/global-managers/{global-manager-id} | Delete a Global Manager under Infra
[**GetFederationUpgradeSummary**](FederationInfraApi.md#GetFederationUpgradeSummary) | **Get** /global-infra/upgrade-summary | Get Upgrade summary
[**GetFederationUpgradeSummary_0**](FederationInfraApi.md#GetFederationUpgradeSummary_0) | **Get** /infra/upgrade-summary | Get Upgrade summary
[**GetSiteOffboardingStatus**](FederationInfraApi.md#GetSiteOffboardingStatus) | **Get** /global-infra/site/offboarding-status | Get site offboarding status.
[**GetSiteOffboardingStatus_0**](FederationInfraApi.md#GetSiteOffboardingStatus_0) | **Get** /infra/site/offboarding-status | Get site offboarding status.
[**GetSpan**](FederationInfraApi.md#GetSpan) | **Get** /global-infra/span | Get span for an entity with specified path
[**GetSpan_0**](FederationInfraApi.md#GetSpan_0) | **Get** /infra/span | Get span for an entity with specified path
[**ListInfraGlobalManagers**](FederationInfraApi.md#ListInfraGlobalManagers) | **Get** /infra/global-managers | List Global Managers
[**ListInfraGlobalManagers_0**](FederationInfraApi.md#ListInfraGlobalManagers_0) | **Get** /global-infra/global-managers | List Global Managers
[**ListOverriddenResources**](FederationInfraApi.md#ListOverriddenResources) | **Get** /global-infra/overridden-resources | List overridden resources
[**ListOverriddenResources_0**](FederationInfraApi.md#ListOverriddenResources_0) | **Get** /infra/overridden-resources | List overridden resources
[**OffboardSiteOffboard**](FederationInfraApi.md#OffboardSiteOffboard) | **Post** /infra/site?action&#x3D;offboard | Submit site offboarding request.
[**OffboardSiteOffboard_0**](FederationInfraApi.md#OffboardSiteOffboard_0) | **Post** /global-infra/site?action&#x3D;offboard | Submit site offboarding request.
[**PatchInfraGlobalManager**](FederationInfraApi.md#PatchInfraGlobalManager) | **Patch** /infra/global-managers/{global-manager-id} | Create or patch a Global Manager
[**PatchInfraGlobalManager_0**](FederationInfraApi.md#PatchInfraGlobalManager_0) | **Patch** /global-infra/global-managers/{global-manager-id} | Create or patch a Global Manager
[**ReadFederationConfig**](FederationInfraApi.md#ReadFederationConfig) | **Get** /global-infra/federation-config | Read federation config
[**ReadFederationConfig_0**](FederationInfraApi.md#ReadFederationConfig_0) | **Get** /infra/federation-config | Read federation config
[**ReadInfraGlobalManager**](FederationInfraApi.md#ReadInfraGlobalManager) | **Get** /infra/global-managers/{global-manager-id} | Read a Global Manager
[**ReadInfraGlobalManager_0**](FederationInfraApi.md#ReadInfraGlobalManager_0) | **Get** /global-infra/global-managers/{global-manager-id} | Read a Global Manager
[**SwitchOverToStandBy**](FederationInfraApi.md#SwitchOverToStandBy) | **Post** /infra/global-managers | Switch over from Active to Standby Global Manager
[**SwitchOverToStandBy_0**](FederationInfraApi.md#SwitchOverToStandBy_0) | **Post** /global-infra/global-managers | Switch over from Active to Standby Global Manager

# **CreateOrUpdateInfraGlobalManager**
> GlobalManager CreateOrUpdateInfraGlobalManager(ctx, body, globalManagerId)
Create or fully replace a Global Manager under infra

Create or fully replace Global Manager under Infra. Revision is optional for creation and required for update. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManager**](GlobalManager.md)|  | 
  **globalManagerId** | **string**|  | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInfraGlobalManager_0**
> GlobalManager CreateOrUpdateInfraGlobalManager_0(ctx, body, globalManagerId)
Create or fully replace a Global Manager under infra

Create or fully replace Global Manager under Infra. Revision is optional for creation and required for update. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManager**](GlobalManager.md)|  | 
  **globalManagerId** | **string**|  | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraGlobalManager**
> DeleteInfraGlobalManager(ctx, globalManagerId)
Delete a Global Manager under Infra

Delete a particular global manager under Infra. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraGlobalManager_0**
> DeleteInfraGlobalManager_0(ctx, globalManagerId)
Delete a Global Manager under Infra

Delete a particular global manager under Infra. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFederationUpgradeSummary**
> FederationUpgradeSummaryListResult GetFederationUpgradeSummary(ctx, optional)
Get Upgrade summary

API will return high level summary of Upgrade across various sites. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiGetFederationUpgradeSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiGetFederationUpgradeSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentVersion** | **optional.String**| Filter on site current_version | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FederationUpgradeSummaryListResult**](FederationUpgradeSummaryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFederationUpgradeSummary_0**
> FederationUpgradeSummaryListResult GetFederationUpgradeSummary_0(ctx, optional)
Get Upgrade summary

API will return high level summary of Upgrade across various sites. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiGetFederationUpgradeSummary_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiGetFederationUpgradeSummary_3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentVersion** | **optional.String**| Filter on site current_version | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FederationUpgradeSummaryListResult**](FederationUpgradeSummaryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteOffboardingStatus**
> SiteOffBoardingState GetSiteOffboardingStatus(ctx, )
Get site offboarding status.

Get site offboarding status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiteOffBoardingState**](SiteOffBoardingState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteOffboardingStatus_0**
> SiteOffBoardingState GetSiteOffboardingStatus_0(ctx, )
Get site offboarding status.

Get site offboarding status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiteOffBoardingState**](SiteOffBoardingState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpan**
> Span GetSpan(ctx, intentPath, optional)
Get span for an entity with specified path

Get span for an entity with specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***FederationInfraApiGetSpanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiGetSpanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sitePath** | **optional.String**| Policy Path of the site | 

### Return type

[**Span**](Span.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpan_0**
> Span GetSpan_0(ctx, intentPath, optional)
Get span for an entity with specified path

Get span for an entity with specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***FederationInfraApiGetSpan_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiGetSpan_5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sitePath** | **optional.String**| Policy Path of the site | 

### Return type

[**Span**](Span.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfraGlobalManagers**
> GlobalManagerListResult ListInfraGlobalManagers(ctx, optional)
List Global Managers

List Global Managers under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiListInfraGlobalManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiListInfraGlobalManagersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GlobalManagerListResult**](GlobalManagerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfraGlobalManagers_0**
> GlobalManagerListResult ListInfraGlobalManagers_0(ctx, optional)
List Global Managers

List Global Managers under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiListInfraGlobalManagers_6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiListInfraGlobalManagers_6Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GlobalManagerListResult**](GlobalManagerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOverriddenResources**
> OverriddenResourceListResult ListOverriddenResources(ctx, optional)
List overridden resources

List overridden resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiListOverriddenResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiListOverriddenResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intentPath** | **optional.String**| Global resource path | 
 **sitePath** | **optional.String**| Site path | 

### Return type

[**OverriddenResourceListResult**](OverriddenResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOverriddenResources_0**
> OverriddenResourceListResult ListOverriddenResources_0(ctx, optional)
List overridden resources

List overridden resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FederationInfraApiListOverriddenResources_7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FederationInfraApiListOverriddenResources_7Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intentPath** | **optional.String**| Global resource path | 
 **sitePath** | **optional.String**| Site path | 

### Return type

[**OverriddenResourceListResult**](OverriddenResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OffboardSiteOffboard**
> SiteOffBoardingState OffboardSiteOffboard(ctx, )
Submit site offboarding request.

Submit site offboarding request.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiteOffBoardingState**](SiteOffBoardingState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OffboardSiteOffboard_0**
> SiteOffBoardingState OffboardSiteOffboard_0(ctx, )
Submit site offboarding request.

Submit site offboarding request.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiteOffBoardingState**](SiteOffBoardingState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfraGlobalManager**
> PatchInfraGlobalManager(ctx, body, globalManagerId)
Create or patch a Global Manager

Create or patch a Global Manager under Infra. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManager**](GlobalManager.md)|  | 
  **globalManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfraGlobalManager_0**
> PatchInfraGlobalManager_0(ctx, body, globalManagerId)
Create or patch a Global Manager

Create or patch a Global Manager under Infra. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManager**](GlobalManager.md)|  | 
  **globalManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFederationConfig**
> FederationConfig ReadFederationConfig(ctx, )
Read federation config

Read a federation config from Global Manager.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FederationConfig**](FederationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFederationConfig_0**
> FederationConfig ReadFederationConfig_0(ctx, )
Read federation config

Read a federation config from Global Manager.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FederationConfig**](FederationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInfraGlobalManager**
> GlobalManager ReadInfraGlobalManager(ctx, globalManagerId)
Read a Global Manager

Retrieve information about a particular configured global manager. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalManagerId** | **string**|  | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInfraGlobalManager_0**
> GlobalManager ReadInfraGlobalManager_0(ctx, globalManagerId)
Read a Global Manager

Retrieve information about a particular configured global manager. Global Manager id 'self' is reserved and can be used for referring to local logged in Global Manager. Example - /infra/global-managers/self 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalManagerId** | **string**|  | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchOverToStandBy**
> GlobalManager SwitchOverToStandBy(ctx, action)
Switch over from Active to Standby Global Manager

Switch over from Active to Standby Global Manager. This operation will fail if there is no Standby Global Manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **action** | **string**| Indicates whether it is managed switchover or due to failure | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchOverToStandBy_0**
> GlobalManager SwitchOverToStandBy_0(ctx, action)
Switch over from Active to Standby Global Manager

Switch over from Active to Standby Global Manager. This operation will fail if there is no Standby Global Manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **action** | **string**| Indicates whether it is managed switchover or due to failure | 

### Return type

[**GlobalManager**](GlobalManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

