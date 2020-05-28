# FirewallIdentityStoreLdapServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Username** | **string** | Directory LDAP server connection user name. | [optional] [default to null]
**Host** | **string** | Directory LDAP server DNS host name or ip address which is reachable by NSX manager to be connected and do object synchronization.  | [default to null]
**Protocol** | **string** | Directory LDAP server connection protocol which is either LDAP or LDAPS.  | [optional] [default to PROTOCOL.LDAP]
**Thumbprint** | **string** | Directory LDAP server certificate thumbprint used in secure LDAPS connection.  | [optional] [default to null]
**Password** | **string** | Directory LDAP server connection password. | [optional] [default to null]
**DomainName** | **string** | Directory domain name which best describes the domain. It could be unique fqdn name or it could also be descriptive. There is no unique constraint for domain name among different domains.  | [optional] [default to null]
**Port** | **int32** | Directory LDAP server connection TCP/UDP port. | [optional] [default to 389]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

