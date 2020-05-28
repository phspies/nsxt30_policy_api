# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptEULA**](SystemAdministrationApi.md#AcceptEULA) | **Post** /eula/accept | Accept end user license agreement 
[**AdvanceClusterRestoreAdvance**](SystemAdministrationApi.md#AdvanceClusterRestoreAdvance) | **Post** /cluster/restore?action&#x3D;advance | Advance any suspended restore operation
[**CancelClusterRestoreCancel**](SystemAdministrationApi.md#CancelClusterRestoreCancel) | **Post** /cluster/restore?action&#x3D;cancel | Cancel any running restore operation
[**ConfigureBackupConfig**](SystemAdministrationApi.md#ConfigureBackupConfig) | **Put** /cluster/backups/config | Configure backup
[**ConfigureRestoreConfig**](SystemAdministrationApi.md#ConfigureRestoreConfig) | **Put** /cluster/restore/config | Deprecated. Configure Restore SFTP server credentials
[**CreateOrUpdateLdapIdentitySource**](SystemAdministrationApi.md#CreateOrUpdateLdapIdentitySource) | **Put** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Update an existing LDAP identity source
[**CreateRegistrationToken**](SystemAdministrationApi.md#CreateRegistrationToken) | **Post** /aaa/registration-token | Create registration access token
[**CreateRoleBinding**](SystemAdministrationApi.md#CreateRoleBinding) | **Post** /aaa/role-bindings | Assign roles to User or Group
[**CreateView**](SystemAdministrationApi.md#CreateView) | **Post** /ui-views | Creates a new View.
[**CreateWidgetConfiguration**](SystemAdministrationApi.md#CreateWidgetConfiguration) | **Post** /ui-views/{view-id}/widgetconfigurations | Creates a new Widget Configuration.
[**DeletView**](SystemAdministrationApi.md#DeletView) | **Delete** /ui-views/{view-id} | Delete View
[**DeleteAllStaleRoleBindingsDeleteStaleBindings**](SystemAdministrationApi.md#DeleteAllStaleRoleBindingsDeleteStaleBindings) | **Post** /aaa/role-bindings?action&#x3D;delete_stale_bindings | Delete all stale role assignments
[**DeleteLdapIdentitySource**](SystemAdministrationApi.md#DeleteLdapIdentitySource) | **Delete** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Delete an LDAP identity source
[**DeleteRegistrationToken**](SystemAdministrationApi.md#DeleteRegistrationToken) | **Delete** /aaa/registration-token/{token} | Delete registration access token
[**DeleteRoleBinding**](SystemAdministrationApi.md#DeleteRoleBinding) | **Delete** /aaa/role-bindings/{binding-id} | Delete user/group&#x27;s roles assignment
[**DeleteWidgetConfiguration**](SystemAdministrationApi.md#DeleteWidgetConfiguration) | **Delete** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Delete Widget Configuration
[**FetchIdentitySourceLdapServerCertificateFetchCertificate**](SystemAdministrationApi.md#FetchIdentitySourceLdapServerCertificateFetchCertificate) | **Post** /aaa/ldap-identity-sources?action&#x3D;fetch_certificate | Fetch the server certificate of an LDAP server
[**GetAllRoleBindings**](SystemAdministrationApi.md#GetAllRoleBindings) | **Get** /aaa/role-bindings | Get all users and groups with their roles
[**GetAllRolesInfo**](SystemAdministrationApi.md#GetAllRolesInfo) | **Get** /aaa/roles | Get information about all roles
[**GetBackupConfig**](SystemAdministrationApi.md#GetBackupConfig) | **Get** /cluster/backups/config | Get backup configuration
[**GetBackupHistory**](SystemAdministrationApi.md#GetBackupHistory) | **Get** /cluster/backups/history | Get backup history
[**GetBackupOverview**](SystemAdministrationApi.md#GetBackupOverview) | **Get** /cluster/backups/overview | Get all backup related information for a site
[**GetBackupStatus**](SystemAdministrationApi.md#GetBackupStatus) | **Get** /cluster/backups/status | Get backup status
[**GetCurrentUserInfo**](SystemAdministrationApi.md#GetCurrentUserInfo) | **Get** /aaa/user-info | Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.
[**GetEULAAcceptance**](SystemAdministrationApi.md#GetEULAAcceptance) | **Get** /eula/acceptance | Return the acceptance status of end user license agreement 
[**GetEULAContent**](SystemAdministrationApi.md#GetEULAContent) | **Get** /eula/content | Return the content of end user license agreement 
[**GetErrorResolverInfo**](SystemAdministrationApi.md#GetErrorResolverInfo) | **Get** /error-resolver/{error_id} | Fetches metadata about the given error_id
[**GetGroupVidmSearchResult**](SystemAdministrationApi.md#GetGroupVidmSearchResult) | **Get** /aaa/vidm/groups | Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.
[**GetRegistrationToken**](SystemAdministrationApi.md#GetRegistrationToken) | **Get** /aaa/registration-token/{token} | Get registration access token
[**GetRestoreConfig**](SystemAdministrationApi.md#GetRestoreConfig) | **Get** /cluster/restore/config | Deprecated. Get Restore configuration
[**GetRoleBinding**](SystemAdministrationApi.md#GetRoleBinding) | **Get** /aaa/role-bindings/{binding-id} | Get user/group&#x27;s role information
[**GetRoleInfo**](SystemAdministrationApi.md#GetRoleInfo) | **Get** /aaa/roles/{role} | Get role information
[**GetSshFingerprintOfServerRetrieveSshFingerprint**](SystemAdministrationApi.md#GetSshFingerprintOfServerRetrieveSshFingerprint) | **Post** /cluster/backups?action&#x3D;retrieve_ssh_fingerprint | Get ssh fingerprint of remote(backup) server
[**GetUserVidmSearchResult**](SystemAdministrationApi.md#GetUserVidmSearchResult) | **Get** /aaa/vidm/users | Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.
[**GetVersionWhitelist**](SystemAdministrationApi.md#GetVersionWhitelist) | **Get** /upgrade/version-whitelist | Get the version whitelist
[**GetVersionWhitelistByComponent**](SystemAdministrationApi.md#GetVersionWhitelistByComponent) | **Get** /upgrade/version-whitelist/{component_type} | Get the version whitelist for the specified component
[**GetVidmSearchResult**](SystemAdministrationApi.md#GetVidmSearchResult) | **Post** /aaa/vidm/search | Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.
[**GetView**](SystemAdministrationApi.md#GetView) | **Get** /ui-views/{view-id} | Returns View Information
[**GetWidgetConfiguration**](SystemAdministrationApi.md#GetWidgetConfiguration) | **Get** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Returns Widget Configuration Information
[**InitiateClusterRestoreStart**](SystemAdministrationApi.md#InitiateClusterRestoreStart) | **Post** /cluster/restore?action&#x3D;start | Initiate a restore operation
[**ListClusterBackupTimestamps**](SystemAdministrationApi.md#ListClusterBackupTimestamps) | **Get** /cluster/restore/backuptimestamps | List timestamps of all available Cluster Backups.
[**ListErrorResolverInfo**](SystemAdministrationApi.md#ListErrorResolverInfo) | **Get** /error-resolver | Fetches a list of metadata for all the registered error resolvers
[**ListLdapIdentitySources**](SystemAdministrationApi.md#ListLdapIdentitySources) | **Get** /aaa/ldap-identity-sources | List LDAP identity sources
[**ListRestoreInstructionResources**](SystemAdministrationApi.md#ListRestoreInstructionResources) | **Get** /cluster/restore/instruction-resources | List resources for a given instruction, to be shown to/executed by users. 
[**ListRolesInfo**](SystemAdministrationApi.md#ListRolesInfo) | **Get** /aaa/roles-with-feature-permissions | Get information about all roles with features and their permissions
[**ListTasks**](SystemAdministrationApi.md#ListTasks) | **Get** /tasks | Get information about all tasks
[**ListViews**](SystemAdministrationApi.md#ListViews) | **Get** /ui-views | Returns the Views based on query criteria defined in ViewQueryParameters.
[**ListWidgetConfigurations**](SystemAdministrationApi.md#ListWidgetConfigurations) | **Get** /ui-views/{view-id}/widgetconfigurations | Returns the Widget Configurations based on query criteria defined in WidgetQueryParameters.
[**ProbeConfiguredLdapIdentitySourceProbe**](SystemAdministrationApi.md#ProbeConfiguredLdapIdentitySourceProbe) | **Post** /aaa/ldap-identity-sources/{ldap-identity-source-id}?action&#x3D;probe | Test the configuration of an existing LDAP identity source
[**ProbeIdentitySourceLdapServerProbeLdapServer**](SystemAdministrationApi.md#ProbeIdentitySourceLdapServerProbeLdapServer) | **Post** /aaa/ldap-identity-sources?action&#x3D;probe_ldap_server | Test an LDAP server
[**ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource**](SystemAdministrationApi.md#ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource) | **Post** /aaa/ldap-identity-sources?action&#x3D;probe_identity_source | Probe an LDAP identity source
[**ReadLdapIdentitySource**](SystemAdministrationApi.md#ReadLdapIdentitySource) | **Get** /aaa/ldap-identity-sources/{ldap-identity-source-id} | Read a single LDAP identity source
[**ReadManagementConfig**](SystemAdministrationApi.md#ReadManagementConfig) | **Get** /configs/management | Read NSX Management nodes global configuration.
[**ReadTaskProperties**](SystemAdministrationApi.md#ReadTaskProperties) | **Get** /tasks/{task-id} | Get information about the specified task
[**ReadTaskResult**](SystemAdministrationApi.md#ReadTaskResult) | **Get** /tasks/{task-id}/response | Get the response of a task
[**RegisterBatchRequest**](SystemAdministrationApi.md#RegisterBatchRequest) | **Post** /batch | Register a Collection of API Calls at a Single End Point
[**RequestOnetimeBackupBackupToRemote**](SystemAdministrationApi.md#RequestOnetimeBackupBackupToRemote) | **Post** /cluster?action&#x3D;backup_to_remote | Request one-time backup
[**RequestOnetimeInventorySummarySummarizeInventoryToRemote**](SystemAdministrationApi.md#RequestOnetimeInventorySummarySummarizeInventoryToRemote) | **Post** /cluster?action&#x3D;summarize_inventory_to_remote | Request one-time inventory summary.
[**ResolveErrorResolveError**](SystemAdministrationApi.md#ResolveErrorResolveError) | **Post** /error-resolver?action&#x3D;resolve_error | Resolves the error
[**RetryClusterRestoreRetry**](SystemAdministrationApi.md#RetryClusterRestoreRetry) | **Post** /cluster/restore?action&#x3D;retry | Retry any failed restore operation
[**SearchLdapIdentitySource**](SystemAdministrationApi.md#SearchLdapIdentitySource) | **Post** /aaa/ldap-identity-sources/{ldap-identity-source-id}/search | Search the LDAP identity source
[**SuspendClusterRestoreSuspend**](SystemAdministrationApi.md#SuspendClusterRestoreSuspend) | **Post** /cluster/restore?action&#x3D;suspend | Suspend any running restore operation
[**UpdateManagementConfig**](SystemAdministrationApi.md#UpdateManagementConfig) | **Put** /configs/management | Update NSX Management nodes global configuration
[**UpdateRoleBinding**](SystemAdministrationApi.md#UpdateRoleBinding) | **Put** /aaa/role-bindings/{binding-id} | Update User or Group&#x27;s roles
[**UpdateVersionWhitelist**](SystemAdministrationApi.md#UpdateVersionWhitelist) | **Put** /upgrade/version-whitelist/{component_type} | Update the version whitelist for the specified component type
[**UpdateView**](SystemAdministrationApi.md#UpdateView) | **Put** /ui-views/{view-id} | Update View
[**UpdateWidgetConfiguration**](SystemAdministrationApi.md#UpdateWidgetConfiguration) | **Put** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Update Widget Configuration

# **AcceptEULA**
> AcceptEULA(ctx, )
Accept end user license agreement 

Accept end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvanceClusterRestoreAdvance**
> ClusterRestoreStatus AdvanceClusterRestoreAdvance(ctx, body)
Advance any suspended restore operation

Advance any currently suspended restore operation. The operation might have been suspended because (1) the user had suspended it previously, or (2) the operation is waiting for user input, to be provided as a part of the POST request body. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED. Otherwise, a 409 response is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvanceClusterRestoreRequest**](AdvanceClusterRestoreRequest.md)|  | 

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelClusterRestoreCancel**
> ClusterRestoreStatus CancelClusterRestoreCancel(ctx, )
Cancel any running restore operation

This operation is only valid when a restore is in suspended state. The UI user can cancel any restore operation when the restore is suspended either due to an error, or for a user input. The API user would need to monitor the progression of a restore by calling periodically \"/api/v1/cluster/restore/status\" API. The response object (ClusterRestoreStatus), contains a field \"endpoints\". The API user can cancel the restore process if 'cancel' action is shown in the endpoint field. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureBackupConfig**
> BackupConfiguration ConfigureBackupConfig(ctx, body, optional)
Configure backup

Configure file server and timers for automated backup. If secret fields are omitted (password, passphrase) then use the previously set value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BackupConfiguration**](BackupConfiguration.md)|  | 
 **optional** | ***SystemAdministrationApiConfigureBackupConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiConfigureBackupConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameType** | **optional.**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **siteId** | **optional.**| Site ID | [default to localhost]

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureRestoreConfig**
> RestoreConfiguration ConfigureRestoreConfig(ctx, body)
Deprecated. Configure Restore SFTP server credentials

Deprecated. Please use API /cluster/backups/config, to configure remote file server(where backed-up files are stored) details during restore. In older versions - Configure file server where the backed-up files used for the Restore operation are available. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestoreConfiguration**](RestoreConfiguration.md)|  | 

### Return type

[**RestoreConfiguration**](RestoreConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateLdapIdentitySource**
> LdapIdentitySource CreateOrUpdateLdapIdentitySource(ctx, body, ldapIdentitySourceId)
Update an existing LDAP identity source

Update the configuration of an existing LDAP identity source. You may wish to verify the new configuration using the POST /aaa/ldap-identity-sources?action=probe API before changing the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LdapIdentitySource**](LdapIdentitySource.md)|  | 
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySource**](LdapIdentitySource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRegistrationToken**
> RegistrationToken CreateRegistrationToken(ctx, )
Create registration access token

The privileges of the registration token will be the same as the caller.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoleBinding**
> RoleBinding CreateRoleBinding(ctx, body)
Assign roles to User or Group

When assigning a user role, specify the user name with the same case as it appears in vIDM to access the NSX-T user interface. For example, if vIDM has the user name User1@example.com then the name attribute in the API call must be be User1@example.com and cannot be user1@example.com. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleBinding**](RoleBinding.md)|  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateView**
> View CreateView(ctx, body)
Creates a new View.

Creates a new View.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**View**](View.md)|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWidgetConfiguration**
> WidgetConfiguration CreateWidgetConfiguration(ctx, body, viewId)
Creates a new Widget Configuration.

Creates a new Widget Configuration and adds it to the specified view. Supported resource_types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration and ContainerConfiguration.  Note: Expressions should be given in a single line. If an expression spans   multiple lines, then form the expression in a single line. For label-value pairs, expressions are evaluated as follows:   a. First, render configurations are evaluated in their order of      appearance in the widget config. The 'field' is evaluated at the end.   b. Second, when render configuration is provided then the order of      evaluation is      1. If expressions provided in 'condition' and 'display value' are         well-formed and free of runtime-errors such as 'null pointers' and         evaluates to 'true'; Then remaining render configurations are not         evaluated, and the current render configuration's 'display value'         is taken as the final value.      2. If expression provided in 'condition' of render configuration is         false, then next render configuration is evaluated.      3. Finally, 'field' is evaluated only when every render configuration         evaluates to false and no error occurs during steps 1 and 2 above.  If an error occurs during evaluation of render configuration, then an   error message is shown. The display value corresponding to that label is   not shown and evaluation of the remaining render configurations continues   to collect and show all the error messages (marked with the 'Label' for   identification) as 'Error_Messages: {}'.  If during evaluation of expressions for any label-value pair an error   occurs, then it is marked with error. The errors are shown in the report,   along with the label value pairs that are error-free.  Important: For elements that take expressions, strings should be provided   by escaping them with a back-slash. These elements are - condition, field,   tooltip text and render_configuration's display_value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WidgetConfiguration**](WidgetConfiguration.md)|  | 
  **viewId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletView**
> DeletView(ctx, viewId)
Delete View

Delete View

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllStaleRoleBindingsDeleteStaleBindings**
> DeleteAllStaleRoleBindingsDeleteStaleBindings(ctx, )
Delete all stale role assignments

Delete all stale role assignments

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLdapIdentitySource**
> DeleteLdapIdentitySource(ctx, ldapIdentitySourceId)
Delete an LDAP identity source

Delete an LDAP identity source. Users defined in that source will no longer be able to access NSX.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegistrationToken**
> DeleteRegistrationToken(ctx, token)
Delete registration access token

Delete registration access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Registration token | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleBinding**
> DeleteRoleBinding(ctx, bindingId)
Delete user/group's roles assignment

Delete user/group's roles assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWidgetConfiguration**
> DeleteWidgetConfiguration(ctx, viewId, widgetconfigurationId)
Delete Widget Configuration

Detaches widget from a given view. If the widget is no longer part of any view, then it will be purged. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchIdentitySourceLdapServerCertificateFetchCertificate**
> PeerCertificateChain FetchIdentitySourceLdapServerCertificateFetchCertificate(ctx, body)
Fetch the server certificate of an LDAP server

Attempt to connect to an LDAP server and retrieve the server certificate it presents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentitySourceLdapServerEndpoint**](IdentitySourceLdapServerEndpoint.md)|  | 

### Return type

[**PeerCertificateChain**](PeerCertificateChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRoleBindings**
> RoleBindingListResult GetAllRoleBindings(ctx, optional)
Get all users and groups with their roles

Get all users and groups with their roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiGetAllRoleBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetAllRoleBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **identitySourceId** | **optional.String**| Identity source ID | 
 **identitySourceType** | **optional.String**| Identity source type | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **name** | **optional.String**| User/Group name | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **role** | **optional.String**| Role ID | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type | 

### Return type

[**RoleBindingListResult**](RoleBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRolesInfo**
> RoleListResult GetAllRolesInfo(ctx, )
Get information about all roles

Get information about all roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RoleListResult**](RoleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupConfig**
> BackupConfiguration GetBackupConfig(ctx, )
Get backup configuration

Get a configuration of a file server and timers for automated backup. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupHistory**
> BackupOperationHistory GetBackupHistory(ctx, )
Get backup history

Get history of previous backup operations 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupOperationHistory**](BackupOperationHistory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupOverview**
> BackupOverview GetBackupOverview(ctx, optional)
Get all backup related information for a site

Get a configuration of a file server, timers for automated backup, latest backup status, backups list for a site. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiGetBackupOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetBackupOverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **frameType** | **optional.String**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **showBackupsList** | **optional.Bool**| Need a list of backups | [default to true]
 **siteId** | **optional.String**| Site ID | [default to localhost]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BackupOverview**](BackupOverview.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupStatus**
> CurrentBackupOperationStatus GetBackupStatus(ctx, )
Get backup status

Get status of active backup operations 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentBackupOperationStatus**](CurrentBackupOperationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUserInfo**
> UserInfo GetCurrentUserInfo(ctx, )
Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.

Get information about logged-in user. The permissions parameter of the NsxRole has been deprecated.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEULAAcceptance**
> EulaAcceptance GetEULAAcceptance(ctx, )
Return the acceptance status of end user license agreement 

Return the acceptance status of end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EulaAcceptance**](EULAAcceptance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEULAContent**
> EulaContent GetEULAContent(ctx, optional)
Return the content of end user license agreement 

Return the content of end user license agreement in the specified format. By default, it's pure string without line break 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiGetEULAContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetEULAContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **valueFormat** | **optional.String**| End User License Agreement content output format | 

### Return type

[**EulaContent**](EULAContent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetErrorResolverInfo**
> ErrorResolverInfo GetErrorResolverInfo(ctx, errorId)
Fetches metadata about the given error_id

Returns some metadata about the given error_id. This includes information of whether there is a resolver present for the given error_id and its associated user input data 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **errorId** | **string**|  | 

### Return type

[**ErrorResolverInfo**](ErrorResolverInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupVidmSearchResult**
> VidmInfoListResult GetGroupVidmSearchResult(ctx, searchString, optional)
Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.

Get all the User Groups where vIDM display name matches the search key case insensitively. The search key is checked to be a substring of display name. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationApiGetGroupVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetGroupVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegistrationToken**
> RegistrationToken GetRegistrationToken(ctx, token)
Get registration access token

Get registration access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Registration token | 

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRestoreConfig**
> RestoreConfiguration GetRestoreConfig(ctx, )
Deprecated. Get Restore configuration

Deprecated. Please use API /cluster/backups/config, to get remote file server(where backuped-up files are stored) details durign restore. In older versions - Get configuration information for the file server used to store backed-up files. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestoreConfiguration**](RestoreConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleBinding**
> RoleBinding GetRoleBinding(ctx, bindingId)
Get user/group's role information

Get user/group's role information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleInfo**
> RoleWithFeatures GetRoleInfo(ctx, role)
Get role information

Get role information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| Role id | 

### Return type

[**RoleWithFeatures**](RoleWithFeatures.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSshFingerprintOfServerRetrieveSshFingerprint**
> RemoteServerFingerprint GetSshFingerprintOfServerRetrieveSshFingerprint(ctx, body)
Get ssh fingerprint of remote(backup) server

Get SHA256 fingerprint of ECDSA key of remote server. The caller should independently verify that the key is trusted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteServerFingerprintRequest**](RemoteServerFingerprintRequest.md)|  | 

### Return type

[**RemoteServerFingerprint**](RemoteServerFingerprint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserVidmSearchResult**
> VidmInfoListResult GetUserVidmSearchResult(ctx, searchString, optional)
Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.

Get all the users from vIDM whose userName, givenName or familyName matches the search key case insensitively. The search key is checked to be a substring of name or given name or family name. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationApiGetUserVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetUserVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionWhitelist**
> AcceptableComponentVersionList GetVersionWhitelist(ctx, )
Get the version whitelist

Get whitelist of versions for different components

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AcceptableComponentVersionList**](AcceptableComponentVersionList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionWhitelistByComponent**
> AcceptableComponentVersion GetVersionWhitelistByComponent(ctx, componentType)
Get the version whitelist for the specified component

Get whitelist of versions for a component. Component can include HOST, EDGE, CCP, MP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**|  | 

### Return type

[**AcceptableComponentVersion**](AcceptableComponentVersion.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVidmSearchResult**
> VidmInfoListResult GetVidmSearchResult(ctx, searchString, optional)
Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.

Get all the users and groups from vIDM matching the search key case insensitively. The search key is checked to be a substring of name or given name or family name of user and display name of group. This is a non paginated API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string to search for.  | 
 **optional** | ***SystemAdministrationApiGetVidmSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiGetVidmSearchResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VidmInfoListResult**](VidmInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetView**
> View GetView(ctx, viewId)
Returns View Information

Returns Information about a specific View. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetConfiguration**
> WidgetConfiguration GetWidgetConfiguration(ctx, viewId, widgetconfigurationId)
Returns Widget Configuration Information

Returns Information about a specific Widget Configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateClusterRestoreStart**
> ClusterRestoreStatus InitiateClusterRestoreStart(ctx, body)
Initiate a restore operation

Start the restore of an NSX cluster, from some previously backed-up configuration. This operation is only valid when a GET cluster/restore/status returns a status with value NOT_STARTED. Otherwise, a 409 response is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InitiateClusterRestoreRequest**](InitiateClusterRestoreRequest.md)|  | 

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterBackupTimestamps**
> ClusterBackupInfoListResult ListClusterBackupTimestamps(ctx, optional)
List timestamps of all available Cluster Backups.

Returns timestamps for all backup files that are available on the SFTP server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiListClusterBackupTimestampsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListClusterBackupTimestampsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ClusterBackupInfoListResult**](ClusterBackupInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListErrorResolverInfo**
> ErrorResolverInfoList ListErrorResolverInfo(ctx, )
Fetches a list of metadata for all the registered error resolvers

Returns a list of metadata for all the error resolvers registered. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ErrorResolverInfoList**](ErrorResolverInfoList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLdapIdentitySources**
> LdapIdentitySourceListResult ListLdapIdentitySources(ctx, optional)
List LDAP identity sources

Return a list of all configured LDAP identity sources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiListLdapIdentitySourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListLdapIdentitySourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LdapIdentitySourceListResult**](LdapIdentitySourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRestoreInstructionResources**
> ActionableResourceListResult ListRestoreInstructionResources(ctx, instructionId, optional)
List resources for a given instruction, to be shown to/executed by users. 

For restore operations requiring user input e.g. performing an action, accepting/rejecting an action, etc. the information to be conveyed to users is provided in this call. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instructionId** | **string**| Id of the instruction set whose instructions are to be returned | 
 **optional** | ***SystemAdministrationApiListRestoreInstructionResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListRestoreInstructionResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ActionableResourceListResult**](ActionableResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRolesInfo**
> RoleWithFeaturesListResult ListRolesInfo(ctx, optional)
Get information about all roles with features and their permissions

Get information about all roles with features and their permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiListRolesInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListRolesInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RoleWithFeaturesListResult**](RoleWithFeaturesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTasks**
> TaskListResult ListTasks(ctx, optional)
Get information about all tasks

Get information about all tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **requestUri** | **optional.String**| Request URI(s) to include in query result | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **status** | **optional.String**| Status(es) to include in query result | 
 **user** | **optional.String**| Names of users to include in query result | 

### Return type

[**TaskListResult**](TaskListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListViews**
> ViewList ListViews(ctx, optional)
Returns the Views based on query criteria defined in ViewQueryParameters.

If no query params are specified then all the views entitled for the user are returned. The views to which a user is entitled to include the views created by the user and the shared views. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiListViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListViewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **optional.String**| The tag for which associated views to be queried. | 
 **viewIds** | **optional.String**| Ids of the Views | 
 **widgetId** | **optional.String**| Id of widget configuration | 

### Return type

[**ViewList**](ViewList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWidgetConfigurations**
> WidgetConfigurationList ListWidgetConfigurations(ctx, viewId, optional)
Returns the Widget Configurations based on query criteria defined in WidgetQueryParameters.

If no query params are specified then all the Widget Configurations of the specified view are returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
 **optional** | ***SystemAdministrationApiListWidgetConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiListWidgetConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **optional.String**| Id of the container | 
 **widgetIds** | **optional.String**| Ids of the WidgetConfigurations | 

### Return type

[**WidgetConfigurationList**](WidgetConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeConfiguredLdapIdentitySourceProbe**
> LdapIdentitySourceProbeResults ProbeConfiguredLdapIdentitySourceProbe(ctx, ldapIdentitySourceId)
Test the configuration of an existing LDAP identity source

Attempt to connect to an existing LDAP identity source and report any errors encountered.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySourceProbeResults**](LdapIdentitySourceProbeResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeIdentitySourceLdapServerProbeLdapServer**
> IdentitySourceLdapServerProbeResult ProbeIdentitySourceLdapServerProbeLdapServer(ctx, body)
Test an LDAP server

Attempt to connect to an LDAP server and ensure that the server can be contacted using the given URL and authentication credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentitySourceLdapServer**](IdentitySourceLdapServer.md)|  | 

### Return type

[**IdentitySourceLdapServerProbeResult**](IdentitySourceLdapServerProbeResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource**
> LdapIdentitySourceProbeResults ProbeUnconfiguredLdapIdentitySourceProbeIdentitySource(ctx, body)
Probe an LDAP identity source

Verify that the configuration of an LDAP identity source is correct before actually creating the source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LdapIdentitySource**](LdapIdentitySource.md)|  | 

### Return type

[**LdapIdentitySourceProbeResults**](LdapIdentitySourceProbeResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLdapIdentitySource**
> LdapIdentitySource ReadLdapIdentitySource(ctx, ldapIdentitySourceId)
Read a single LDAP identity source

Return details about one LDAP identity source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 

### Return type

[**LdapIdentitySource**](LdapIdentitySource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadManagementConfig**
> ManagementConfig ReadManagementConfig(ctx, )
Read NSX Management nodes global configuration.

Returns the NSX Management nodes global configuration. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ManagementConfig**](ManagementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTaskProperties**
> TaskProperties ReadTaskProperties(ctx, taskId)
Get information about the specified task

Get information about the specified task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 

### Return type

[**TaskProperties**](TaskProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTaskResult**
> interface{} ReadTaskResult(ctx, taskId)
Get the response of a task

Get the response of a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterBatchRequest**
> BatchResponse RegisterBatchRequest(ctx, body, optional)
Register a Collection of API Calls at a Single End Point

Enables you to make multiple API requests using a single request. The batch API takes in an array of logical HTTP requests represented as JSON arrays. Each request has a method (GET, PUT, POST, or DELETE), a relative_url (the portion of the URL after https://&lt;nsx-mgr&gt;/api/), optional headers array (corresponding to HTTP headers) and an optional body (for POST and PUT requests). The batch API returns an array of logical HTTP responses represented as JSON arrays. Each response has a status code, an optional headers array and an optional body (which is a JSON-encoded string). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRequest**](BatchRequest.md)|  | 
 **optional** | ***SystemAdministrationApiRegisterBatchRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiRegisterBatchRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **atomic** | **optional.**| transactional atomicity for the batch of requests embedded in the batch list | [default to false]

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOnetimeBackupBackupToRemote**
> RequestOnetimeBackupBackupToRemote(ctx, optional)
Request one-time backup

Request one-time backup. The backup will be uploaded using the same server configuration as for automatic backup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationApiRequestOnetimeBackupBackupToRemoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiRequestOnetimeBackupBackupToRemoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameType** | **optional.String**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **siteId** | **optional.String**| Site ID | [default to localhost]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOnetimeInventorySummarySummarizeInventoryToRemote**
> RequestOnetimeInventorySummarySummarizeInventoryToRemote(ctx, )
Request one-time inventory summary.

Request one-time inventory summary. The backup will be uploaded using the same server configuration as for an automatic backup. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveErrorResolveError**
> ResolveErrorResolveError(ctx, body)
Resolves the error

Invokes the corresponding error resolver for the given error(s) present in the payload 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ErrorResolverMetadataList**](ErrorResolverMetadataList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryClusterRestoreRetry**
> ClusterRestoreStatus RetryClusterRestoreRetry(ctx, )
Retry any failed restore operation

Retry any currently in-progress, failed restore operation. Only the last step of the multi-step restore operation would have failed,and only that step is retried. This operation is only valid when a GET cluster/restore/status returns a status with value FAILED. Otherwise, a 409 response is returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLdapIdentitySource**
> LdapIdentitySourceSearchResultList SearchLdapIdentitySource(ctx, ldapIdentitySourceId, optional)
Search the LDAP identity source

Search the LDAP identity source for users and groups that match the given filter_value. In most cases, the LDAP source performs a case-insensitive search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapIdentitySourceId** | **string**|  | 
 **optional** | ***SystemAdministrationApiSearchLdapIdentitySourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationApiSearchLdapIdentitySourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterValue** | **optional.String**| Search filter value | 

### Return type

[**LdapIdentitySourceSearchResultList**](LdapIdentitySourceSearchResultList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendClusterRestoreSuspend**
> ClusterRestoreStatus SuspendClusterRestoreSuspend(ctx, )
Suspend any running restore operation

Suspend any currently running restore operation. The restore operation is made up of a number of steps. When this call is issued, any currently running step is allowed to finish (successfully or with errors), and the next step (and therefore the entire restore operation) is suspended until a subsequent resume or cancel call is issued. This operation is only valid when a GET cluster/restore/status returns a status with value RUNNING. Otherwise, a 409 response is returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateManagementConfig**
> ManagementConfig UpdateManagementConfig(ctx, body)
Update NSX Management nodes global configuration

Modifies the NSX Management nodes global configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ManagementConfig**](ManagementConfig.md)|  | 

### Return type

[**ManagementConfig**](ManagementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleBinding**
> RoleBinding UpdateRoleBinding(ctx, body, bindingId)
Update User or Group's roles

Update User or Group's roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleBinding**](RoleBinding.md)|  | 
  **bindingId** | **string**| User/Group&#x27;s id | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVersionWhitelist**
> UpdateVersionWhitelist(ctx, body, componentType)
Update the version whitelist for the specified component type

Update the version whitelist for the specified component type (HOST, EDGE, CCP, MP).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionList**](VersionList.md)|  | 
  **componentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateView**
> View UpdateView(ctx, body, viewId)
Update View

Update View

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**View**](View.md)|  | 
  **viewId** | **string**|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWidgetConfiguration**
> WidgetConfiguration UpdateWidgetConfiguration(ctx, body, viewId, widgetconfigurationId)
Update Widget Configuration

Updates the widget at the given view. If the widget is referenced by other views, then the widget will be updated in all the views that it is part of. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WidgetConfiguration**](WidgetConfiguration.md)|  | 
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

