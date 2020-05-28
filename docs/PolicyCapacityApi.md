# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicyCapacityUsage**](PolicyCapacityApi.md#GetPolicyCapacityUsage) | **Get** /global-infra/capacity/usage | Returns capacity usage data for NSX objects
[**GetPolicyCapacityUsage_0**](PolicyCapacityApi.md#GetPolicyCapacityUsage_0) | **Get** /infra/capacity/usage | Returns capacity usage data for NSX objects

# **GetPolicyCapacityUsage**
> PolicyCapacityUsageResponse GetPolicyCapacityUsage(ctx, optional)
Returns capacity usage data for NSX objects

Returns capacity usage data for NSX objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyCapacityApiGetPolicyCapacityUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyCapacityApiGetPolicyCapacityUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**|  | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyCapacityUsageResponse**](PolicyCapacityUsageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyCapacityUsage_0**
> PolicyCapacityUsageResponse GetPolicyCapacityUsage_0(ctx, optional)
Returns capacity usage data for NSX objects

Returns capacity usage data for NSX objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyCapacityApiGetPolicyCapacityUsage_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyCapacityApiGetPolicyCapacityUsage_1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**|  | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyCapacityUsageResponse**](PolicyCapacityUsageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

