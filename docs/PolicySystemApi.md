# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartialPatchConfiguration**](PolicySystemApi.md#GetPartialPatchConfiguration) | **Get** /system-config/nsx-partial-patch-config | Fetch the policy partial patch configuration value.
[**UpdatePartialPatchConfig**](PolicySystemApi.md#UpdatePartialPatchConfig) | **Patch** /system-config/nsx-partial-patch-config | Saves the configuration for policy partial patch

# **GetPartialPatchConfiguration**
> PartialPatchConfig GetPartialPatchConfiguration(ctx, )
Fetch the policy partial patch configuration value.

Get Configuration values for nsx-partial-patch. By default partial patch is disbaled (i.e false). 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PartialPatchConfig**](PartialPatchConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePartialPatchConfig**
> UpdatePartialPatchConfig(ctx, body)
Saves the configuration for policy partial patch

Update partial patch configuration values. Only boolean value is allowed for enable_partial_patch 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PartialPatchConfig**](PartialPatchConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

