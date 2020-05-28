# FirewallIdentityStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**SelectiveSyncSettings** | [***FirewallIdentityStoreSelectiveSyncSettings**](FirewallIdentityStoreSelectiveSyncSettings.md) |  | [optional] [default to null]
**Name** | **string** | Directory domain name which best describes the domain. It could be unique fqdn name or it could also be descriptive. There is no unique contraint for domain name among different domains.  | [default to null]
**SyncSettings** | [***FirewallIdentityStoreSyncSettings**](FirewallIdentityStoreSyncSettings.md) |  | [optional] [default to null]
**LdapServers** | [**[]FirewallIdentityStoreLdapServer**](FirewallIdentityStoreLdapServer.md) | Directory domain LDAP servers&#x27; information including host, name, port, protocol and so on.  | [default to null]
**BaseDistinguishedName** | **string** | Each active directory domain has a domain naming context (NC), which contains domain-specific data. The root of this naming context is represented by a domain&#x27;s distinguished name (DN) and is typically referred to as the NC head.  | [default to null]
**NetbiosName** | **string** | NetBIOS names can contain all alphanumeric characters except for the certain disallowed characters. Names can contain a period, but names cannot start with a period. NetBIOS is similar to DNS in that it can serve as a directory service, but more limited as it has no provisions for a name hierarchy and names are limited to 15 characters. The netbios name is case insensitive and is stored in upper case regardless of input case.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

