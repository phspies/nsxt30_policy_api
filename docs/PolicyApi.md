# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicyGroupServiceAssociations**](PolicyApi.md#GetPolicyGroupServiceAssociations) | **Get** /infra/group-service-associations | Get the list of services where the given group is consumed.
[**GetPolicyGroupServiceAssociations_0**](PolicyApi.md#GetPolicyGroupServiceAssociations_0) | **Get** /global-infra/group-service-associations | Get the list of services where the given group is consumed.

# **GetPolicyGroupServiceAssociations**
> PolicyResourceReferenceListResult GetPolicyGroupServiceAssociations(ctx, intentPath, optional)
Get the list of services where the given group is consumed.

The API returns all the services associated with the given Group. It also returns the services associated with the parent groups of the given group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Path of the entity | 
 **optional** | ***PolicyApiGetPolicyGroupServiceAssociationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiGetPolicyGroupServiceAssociationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceListResult**](PolicyResourceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyGroupServiceAssociations_0**
> PolicyResourceReferenceListResult GetPolicyGroupServiceAssociations_0(ctx, intentPath, optional)
Get the list of services where the given group is consumed.

The API returns all the services associated with the given Group. It also returns the services associated with the parent groups of the given group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Path of the entity | 
 **optional** | ***PolicyApiGetPolicyGroupServiceAssociations_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiGetPolicyGroupServiceAssociations_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyResourceReferenceListResult**](PolicyResourceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

