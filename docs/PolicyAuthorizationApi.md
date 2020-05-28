# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectPermissions**](PolicyAuthorizationApi.md#DeleteObjectPermissions) | **Delete** /aaa/object-permissions | Delete object-permissions entries
[**GetObjectPermissions**](PolicyAuthorizationApi.md#GetObjectPermissions) | **Get** /aaa/object-permissions | Get list of Object-level RBAC entries.
[**GetPathPermissions**](PolicyAuthorizationApi.md#GetPathPermissions) | **Get** /aaa/effective-permissions | Get effective object permissions to object specified by path for current user.
[**UpdateObjectPermissions**](PolicyAuthorizationApi.md#UpdateObjectPermissions) | **Patch** /aaa/object-permissions | Create/update object permission mappings

# **DeleteObjectPermissions**
> DeleteObjectPermissions(ctx, optional)
Delete object-permissions entries

Delete object-permissions entries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyAuthorizationApiDeleteObjectPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAuthorizationApiDeleteObjectPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **inheritanceDisabled** | **optional.Bool**| Does children of this object inherit this rule | [default to false]
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **pathPrefix** | **optional.String**| Path prefix | 
 **roleName** | **optional.String**| Role name | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectPermissions**
> ObjectRolePermissionGroupListResult GetObjectPermissions(ctx, optional)
Get list of Object-level RBAC entries.

Get list of Object-level RBAC entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyAuthorizationApiGetObjectPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAuthorizationApiGetObjectPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **inheritanceDisabled** | **optional.Bool**| Does children of this object inherit this rule | [default to false]
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **pathPrefix** | **optional.String**| Path prefix | 
 **roleName** | **optional.String**| Role name | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ObjectRolePermissionGroupListResult**](ObjectRolePermissionGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPathPermissions**
> PathPermissionGroup GetPathPermissions(ctx, featureName, objectPath)
Get effective object permissions to object specified by path for current user.

Returns none if user doesn't have access or feature_name from required request parameter is empty/invalid/doesn't match with object-path provided. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**| Feature name | 
  **objectPath** | **string**| Exact object Policy path | 

### Return type

[**PathPermissionGroup**](PathPermissionGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateObjectPermissions**
> UpdateObjectPermissions(ctx, body)
Create/update object permission mappings

Create/update object permission mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObjectRolePermissionGroup**](ObjectRolePermissionGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

