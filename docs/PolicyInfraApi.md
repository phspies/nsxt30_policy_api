# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTlsCertificate**](PolicyInfraApi.md#AddTlsCertificate) | **Put** /infra/certificates/{certificate-id} | Add a New Certificate
[**AddTlsCertificate_0**](PolicyInfraApi.md#AddTlsCertificate_0) | **Put** /global-infra/certificates/{certificate-id} | Add a New Certificate
[**CreateOrPatchInfraReaction**](PolicyInfraApi.md#CreateOrPatchInfraReaction) | **Patch** /global-infra/reactions/{reaction-id} | Create or patch a Reaction
[**CreateOrPatchInfraReaction_0**](PolicyInfraApi.md#CreateOrPatchInfraReaction_0) | **Patch** /infra/reactions/{reaction-id} | Create or patch a Reaction
[**CreateOrPatchTlsCrl**](PolicyInfraApi.md#CreateOrPatchTlsCrl) | **Patch** /infra/crls/{crl-id} | Create or patch a Certificate Revocation List
[**CreateOrPatchTlsCrl_0**](PolicyInfraApi.md#CreateOrPatchTlsCrl_0) | **Patch** /global-infra/crls/{crl-id} | Create or patch a Certificate Revocation List
[**CreateOrReplacePolicyLabelForInfra**](PolicyInfraApi.md#CreateOrReplacePolicyLabelForInfra) | **Put** /global-infra/labels/{label-id} | Create or replace label
[**CreateOrReplacePolicyLabelForInfra_0**](PolicyInfraApi.md#CreateOrReplacePolicyLabelForInfra_0) | **Put** /infra/labels/{label-id} | Create or replace label
[**CreateOrReplaceTenantConstraint**](PolicyInfraApi.md#CreateOrReplaceTenantConstraint) | **Put** /infra/constraints/{constraint-id} | Create or update tenant Constraint
[**CreateOrReplaceTenantConstraint_0**](PolicyInfraApi.md#CreateOrReplaceTenantConstraint_0) | **Put** /global-infra/constraints/{constraint-id} | Create or update tenant Constraint
[**CreateOrUpdateDomainDeploymentMapForInfra**](PolicyInfraApi.md#CreateOrUpdateDomainDeploymentMapForInfra) | **Put** /infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Create a new Domain Deployment Map under infra
[**CreateOrUpdateDomainDeploymentMapForInfra_0**](PolicyInfraApi.md#CreateOrUpdateDomainDeploymentMapForInfra_0) | **Put** /global-infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Create a new Domain Deployment Map under infra
[**CreateOrUpdateEnforcementPointForInfra**](PolicyInfraApi.md#CreateOrUpdateEnforcementPointForInfra) | **Put** /infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Create/update a new Enforcement Point under infra
[**CreateOrUpdateEnforcementPointForInfra_0**](PolicyInfraApi.md#CreateOrUpdateEnforcementPointForInfra_0) | **Put** /global-infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Create/update a new Enforcement Point under infra
[**CreateOrUpdateEnforcementPointForSite**](PolicyInfraApi.md#CreateOrUpdateEnforcementPointForSite) | **Put** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Create/update a new Enforcement Point under Site
[**CreateOrUpdateEnforcementPointForSite_0**](PolicyInfraApi.md#CreateOrUpdateEnforcementPointForSite_0) | **Put** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Create/update a new Enforcement Point under Site
[**CreateOrUpdateGlobalManagerConfig**](PolicyInfraApi.md#CreateOrUpdateGlobalManagerConfig) | **Put** /infra/global-manager-config | Create or fully replace Global Manager Config
[**CreateOrUpdateGlobalManagerConfig_0**](PolicyInfraApi.md#CreateOrUpdateGlobalManagerConfig_0) | **Put** /global-infra/global-manager-config | Create or fully replace Global Manager Config
[**CreateOrUpdateInfraReaction**](PolicyInfraApi.md#CreateOrUpdateInfraReaction) | **Put** /global-infra/reactions/{reaction-id} | Create or fully replace a Reaction
[**CreateOrUpdateInfraReaction_0**](PolicyInfraApi.md#CreateOrUpdateInfraReaction_0) | **Put** /infra/reactions/{reaction-id} | Create or fully replace a Reaction
[**CreateOrUpdateInfraSite**](PolicyInfraApi.md#CreateOrUpdateInfraSite) | **Put** /global-infra/sites/{site-id} | Create or fully replace a Site under infra
[**CreateOrUpdateInfraSite_0**](PolicyInfraApi.md#CreateOrUpdateInfraSite_0) | **Put** /infra/sites/{site-id} | Create or fully replace a Site under infra
[**CreateOrUpdateTlsCrl**](PolicyInfraApi.md#CreateOrUpdateTlsCrl) | **Put** /infra/crls/{crl-id} | Create or fully replace a Certificate Revocation List
[**CreateOrUpdateTlsCrl_0**](PolicyInfraApi.md#CreateOrUpdateTlsCrl_0) | **Put** /global-infra/crls/{crl-id} | Create or fully replace a Certificate Revocation List
[**CreateTlsCrlImport**](PolicyInfraApi.md#CreateTlsCrlImport) | **Post** /global-infra/crls/{crl-id}?action&#x3D;import | Create a new Certificate Revocation List
[**CreateTlsCrlImport_0**](PolicyInfraApi.md#CreateTlsCrlImport_0) | **Post** /infra/crls/{crl-id}?action&#x3D;import | Create a new Certificate Revocation List
[**DeleteDomain**](PolicyInfraApi.md#DeleteDomain) | **Delete** /global-infra/domains/{domain-id} | Delete Domain and all the entities contained by this domain
[**DeleteDomainDeploymentMap**](PolicyInfraApi.md#DeleteDomainDeploymentMap) | **Delete** /infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Delete Domain Deployment Map
[**DeleteDomainDeploymentMap_0**](PolicyInfraApi.md#DeleteDomainDeploymentMap_0) | **Delete** /global-infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Delete Domain Deployment Map
[**DeleteDomain_0**](PolicyInfraApi.md#DeleteDomain_0) | **Delete** /infra/domains/{domain-id} | Delete Domain and all the entities contained by this domain
[**DeleteEnforcementPoint**](PolicyInfraApi.md#DeleteEnforcementPoint) | **Delete** /infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Delete EnforcementPoint
[**DeleteEnforcementPointForSite**](PolicyInfraApi.md#DeleteEnforcementPointForSite) | **Delete** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Delete EnforcementPoint from Site
[**DeleteEnforcementPointForSite_0**](PolicyInfraApi.md#DeleteEnforcementPointForSite_0) | **Delete** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Delete EnforcementPoint from Site
[**DeleteEnforcementPoint_0**](PolicyInfraApi.md#DeleteEnforcementPoint_0) | **Delete** /global-infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Delete EnforcementPoint
[**DeleteInfraReaction**](PolicyInfraApi.md#DeleteInfraReaction) | **Delete** /global-infra/reactions/{reaction-id} | Delete Reaction
[**DeleteInfraReaction_0**](PolicyInfraApi.md#DeleteInfraReaction_0) | **Delete** /infra/reactions/{reaction-id} | Delete Reaction
[**DeleteInfraSite**](PolicyInfraApi.md#DeleteInfraSite) | **Delete** /global-infra/sites/{site-id} | Delete a site
[**DeleteInfraSite_0**](PolicyInfraApi.md#DeleteInfraSite_0) | **Delete** /infra/sites/{site-id} | Delete a site
[**DeletePolicyLabelForInfra**](PolicyInfraApi.md#DeletePolicyLabelForInfra) | **Delete** /global-infra/labels/{label-id} | Delete PolicyLabel object
[**DeletePolicyLabelForInfra_0**](PolicyInfraApi.md#DeletePolicyLabelForInfra_0) | **Delete** /infra/labels/{label-id} | Delete PolicyLabel object
[**DeleteTenantConstraint**](PolicyInfraApi.md#DeleteTenantConstraint) | **Delete** /infra/constraints/{constraint-id} | Delete tenant Constraint.
[**DeleteTenantConstraint_0**](PolicyInfraApi.md#DeleteTenantConstraint_0) | **Delete** /global-infra/constraints/{constraint-id} | Delete tenant Constraint.
[**DeleteTlsCertificate**](PolicyInfraApi.md#DeleteTlsCertificate) | **Delete** /infra/certificates/{certificate-id} | Delete Certificate for the Given Certificate ID
[**DeleteTlsCertificate_0**](PolicyInfraApi.md#DeleteTlsCertificate_0) | **Delete** /global-infra/certificates/{certificate-id} | Delete Certificate for the Given Certificate ID
[**DeleteTlsCrl**](PolicyInfraApi.md#DeleteTlsCrl) | **Delete** /infra/crls/{crl-id} | Delete a CRL
[**DeleteTlsCrl_0**](PolicyInfraApi.md#DeleteTlsCrl_0) | **Delete** /global-infra/crls/{crl-id} | Delete a CRL
[**FullSyncEnforcementPointForSiteFullSync**](PolicyInfraApi.md#FullSyncEnforcementPointForSiteFullSync) | **Post** /infra/sites/{site-id}/enforcement-points/{enforcement-point-id}?action&#x3D;full-sync | Full sync EnforcementPoint from Site
[**FullSyncEnforcementPointForSiteFullSync_0**](PolicyInfraApi.md#FullSyncEnforcementPointForSiteFullSync_0) | **Post** /global-infra/sites/{site-id}/enforcement-points/{enforcement-point-id}?action&#x3D;full-sync | Full sync EnforcementPoint from Site
[**GetInfraReaction**](PolicyInfraApi.md#GetInfraReaction) | **Get** /global-infra/reactions/{reaction-id} | Get Reaction
[**GetInfraReaction_0**](PolicyInfraApi.md#GetInfraReaction_0) | **Get** /infra/reactions/{reaction-id} | Get Reaction
[**GetInfraSiteListenerCertificate**](PolicyInfraApi.md#GetInfraSiteListenerCertificate) | **Get** /global-infra/sites/listener_certificate | Returns the certificate of the listener
[**GetInfraSiteListenerCertificate_0**](PolicyInfraApi.md#GetInfraSiteListenerCertificate_0) | **Get** /infra/sites/listener_certificate | Returns the certificate of the listener
[**GetTagBulkOperation**](PolicyInfraApi.md#GetTagBulkOperation) | **Get** /infra/tags/tag-operations/{operation-id} | Get details of tag bulk operation request
[**GetTagBulkOperationStatus**](PolicyInfraApi.md#GetTagBulkOperationStatus) | **Get** /global-infra/tags/tag-operations/{operation-id}/status | Get status of tag bulk operation
[**GetTagBulkOperationStatus_0**](PolicyInfraApi.md#GetTagBulkOperationStatus_0) | **Get** /infra/tags/tag-operations/{operation-id}/status | Get status of tag bulk operation
[**GetTagBulkOperation_0**](PolicyInfraApi.md#GetTagBulkOperation_0) | **Get** /global-infra/tags/tag-operations/{operation-id} | Get details of tag bulk operation request
[**GetTlsCertificate**](PolicyInfraApi.md#GetTlsCertificate) | **Get** /infra/certificates/{certificate-id} | Show Certificate Data for the Given Certificate ID
[**GetTlsCertificate_0**](PolicyInfraApi.md#GetTlsCertificate_0) | **Get** /global-infra/certificates/{certificate-id} | Show Certificate Data for the Given Certificate ID
[**GetTlsCrl**](PolicyInfraApi.md#GetTlsCrl) | **Get** /infra/crls/{crl-id} | Show CRL Data for the Given CRL id.
[**GetTlsCrl_0**](PolicyInfraApi.md#GetTlsCrl_0) | **Get** /global-infra/crls/{crl-id} | Show CRL Data for the Given CRL id.
[**ListAlarms**](PolicyInfraApi.md#ListAlarms) | **Get** /infra/realized-state/alarms | List All alarms in the system
[**ListAlarms_0**](PolicyInfraApi.md#ListAlarms_0) | **Get** /global-infra/realized-state/alarms | List All alarms in the system
[**ListAllTags**](PolicyInfraApi.md#ListAllTags) | **Get** /infra/tags | List all unique tags.
[**ListAllTags_0**](PolicyInfraApi.md#ListAllTags_0) | **Get** /global-infra/tags | List all unique tags.
[**ListAllUnAssociatedVirtualMachines**](PolicyInfraApi.md#ListAllUnAssociatedVirtualMachines) | **Get** /global-infra/realized-state/unassociated-virtual-machines | List all virtual machines which are not part of any group
[**ListAllUnAssociatedVirtualMachines_0**](PolicyInfraApi.md#ListAllUnAssociatedVirtualMachines_0) | **Get** /infra/realized-state/unassociated-virtual-machines | List all virtual machines which are not part of any group
[**ListAllVirtualMachines**](PolicyInfraApi.md#ListAllVirtualMachines) | **Get** /infra/realized-state/virtual-machines | List all virtual machines
[**ListAllVirtualMachines_0**](PolicyInfraApi.md#ListAllVirtualMachines_0) | **Get** /global-infra/realized-state/virtual-machines | List all virtual machines
[**ListDeploymentZonesForInfra**](PolicyInfraApi.md#ListDeploymentZonesForInfra) | **Get** /infra/deployment-zones | List Deployment Zones for infra
[**ListDeploymentZonesForInfra_0**](PolicyInfraApi.md#ListDeploymentZonesForInfra_0) | **Get** /global-infra/deployment-zones | List Deployment Zones for infra
[**ListDomainDeploymentMapsForInfra**](PolicyInfraApi.md#ListDomainDeploymentMapsForInfra) | **Get** /infra/domains/{domain-id}/domain-deployment-maps | List Domain Deployment maps for infra
[**ListDomainDeploymentMapsForInfra_0**](PolicyInfraApi.md#ListDomainDeploymentMapsForInfra_0) | **Get** /global-infra/domains/{domain-id}/domain-deployment-maps | List Domain Deployment maps for infra
[**ListDomainForInfra**](PolicyInfraApi.md#ListDomainForInfra) | **Get** /global-infra/domains | List domains for infra
[**ListDomainForInfra_0**](PolicyInfraApi.md#ListDomainForInfra_0) | **Get** /infra/domains | List domains for infra
[**ListEdgeClustersForEnforcementPoint**](PolicyInfraApi.md#ListEdgeClustersForEnforcementPoint) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters | List Edge Clusters under an Enforcement Point
[**ListEdgeClustersForEnforcementPoint_0**](PolicyInfraApi.md#ListEdgeClustersForEnforcementPoint_0) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters | List Edge Clusters under an Enforcement Point
[**ListEdgeNodesUnderEdgeClusterForEnforcementPoint**](PolicyInfraApi.md#ListEdgeNodesUnderEdgeClusterForEnforcementPoint) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id}/edge-nodes | List Edge Nodes under an Enforcement Point, Edge Cluster
[**ListEdgeNodesUnderEdgeClusterForEnforcementPoint_0**](PolicyInfraApi.md#ListEdgeNodesUnderEdgeClusterForEnforcementPoint_0) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id}/edge-nodes | List Edge Nodes under an Enforcement Point, Edge Cluster
[**ListEnforcementPointForInfra**](PolicyInfraApi.md#ListEnforcementPointForInfra) | **Get** /infra/deployment-zones/{deployment-zone-id}/enforcement-points | List enforcementpoints for infra
[**ListEnforcementPointForInfra_0**](PolicyInfraApi.md#ListEnforcementPointForInfra_0) | **Get** /global-infra/deployment-zones/{deployment-zone-id}/enforcement-points | List enforcementpoints for infra
[**ListEnforcementPointForSite**](PolicyInfraApi.md#ListEnforcementPointForSite) | **Get** /global-infra/sites/{site-id}/enforcement-points | List enforcementpoints under Site
[**ListEnforcementPointForSite_0**](PolicyInfraApi.md#ListEnforcementPointForSite_0) | **Get** /infra/sites/{site-id}/enforcement-points | List enforcementpoints under Site
[**ListEnforcementPointRealizedStates**](PolicyInfraApi.md#ListEnforcementPointRealizedStates) | **Get** /infra/realized-state/enforcement-points | List Enforcement Points
[**ListEnforcementPointRealizedStates_0**](PolicyInfraApi.md#ListEnforcementPointRealizedStates_0) | **Get** /global-infra/realized-state/enforcement-points | List Enforcement Points
[**ListFirewallSectionRealizedStates**](PolicyInfraApi.md#ListFirewallSectionRealizedStates) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/firewalls/firewall-sections | List Firewall Sections
[**ListFirewallSectionRealizedStates_0**](PolicyInfraApi.md#ListFirewallSectionRealizedStates_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/firewalls/firewall-sections | List Firewall Sections
[**ListIPSetRealizedStates**](PolicyInfraApi.md#ListIPSetRealizedStates) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/ip-sets/ip-sets-nsxt | List IPSets
[**ListIPSetRealizedStates_0**](PolicyInfraApi.md#ListIPSetRealizedStates_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/ip-sets/ip-sets-nsxt | List IPSets
[**ListInfraReactions**](PolicyInfraApi.md#ListInfraReactions) | **Get** /global-infra/reactions | Get Reaction list result
[**ListInfraReactions_0**](PolicyInfraApi.md#ListInfraReactions_0) | **Get** /infra/reactions | Get Reaction list result
[**ListMACSetRealizedStates**](PolicyInfraApi.md#ListMACSetRealizedStates) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/mac-sets/mac-sets-nsxt | List MACSets
[**ListMACSetRealizedStates_0**](PolicyInfraApi.md#ListMACSetRealizedStates_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/mac-sets/mac-sets-nsxt | List MACSets
[**ListNSGroupRealizedStates**](PolicyInfraApi.md#ListNSGroupRealizedStates) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/groups/nsgroups | List NS Groups
[**ListNSGroupRealizedStates_0**](PolicyInfraApi.md#ListNSGroupRealizedStates_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/groups/nsgroups | List NS Groups
[**ListNSServiceRealizedStates**](PolicyInfraApi.md#ListNSServiceRealizedStates) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/services/nsservices | List Realized NSServices
[**ListNSServiceRealizedStates_0**](PolicyInfraApi.md#ListNSServiceRealizedStates_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/services/nsservices | List Realized NSServices
[**ListPolicyLabelForInfra**](PolicyInfraApi.md#ListPolicyLabelForInfra) | **Get** /infra/labels | List labels for infra
[**ListPolicyLabelForInfra_0**](PolicyInfraApi.md#ListPolicyLabelForInfra_0) | **Get** /global-infra/labels | List labels for infra
[**ListRealizedEntities**](PolicyInfraApi.md#ListRealizedEntities) | **Get** /global-infra/realized-state/realized-entities | Get list of realized objects associated with intent object
[**ListRealizedEntities_0**](PolicyInfraApi.md#ListRealizedEntities_0) | **Get** /infra/realized-state/realized-entities | Get list of realized objects associated with intent object
[**ListSecurityGroupRealizedStates**](PolicyInfraApi.md#ListSecurityGroupRealizedStates) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/groups/securitygroups | List Security Groups
[**ListSecurityGroupRealizedStates_0**](PolicyInfraApi.md#ListSecurityGroupRealizedStates_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/groups/securitygroups | List Security Groups
[**ListSites**](PolicyInfraApi.md#ListSites) | **Get** /infra/sites | List Sites
[**ListSites_0**](PolicyInfraApi.md#ListSites_0) | **Get** /global-infra/sites | List Sites
[**ListTaggedObjects**](PolicyInfraApi.md#ListTaggedObjects) | **Get** /global-infra/tags/effective-resources | List all objects assigned with matching scope and tag values
[**ListTaggedObjects_0**](PolicyInfraApi.md#ListTaggedObjects_0) | **Get** /infra/tags/effective-resources | List all objects assigned with matching scope and tag values
[**ListTenantConstraints**](PolicyInfraApi.md#ListTenantConstraints) | **Get** /infra/constraints | List tenant Constraints.
[**ListTenantConstraints_0**](PolicyInfraApi.md#ListTenantConstraints_0) | **Get** /global-infra/constraints | List tenant Constraints.
[**ListTlsCertificates**](PolicyInfraApi.md#ListTlsCertificates) | **Get** /global-infra/certificates | Return All the User-Facing Components&#x27; Certificates
[**ListTlsCertificates_0**](PolicyInfraApi.md#ListTlsCertificates_0) | **Get** /infra/certificates | Return All the User-Facing Components&#x27; Certificates
[**ListTlsCrls**](PolicyInfraApi.md#ListTlsCrls) | **Get** /global-infra/crls | Return All Added CRLs
[**ListTlsCrls_0**](PolicyInfraApi.md#ListTlsCrls_0) | **Get** /infra/crls | Return All Added CRLs
[**ListTransportZonesForEnforcementPoint**](PolicyInfraApi.md#ListTransportZonesForEnforcementPoint) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/transport-zones | List Transport Zones under an Enforcement Point
[**ListTransportZonesForEnforcementPoint_0**](PolicyInfraApi.md#ListTransportZonesForEnforcementPoint_0) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/transport-zones | List Transport Zones under an Enforcement Point
[**ListVifsOnEnforcementPoint**](PolicyInfraApi.md#ListVifsOnEnforcementPoint) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/vifs | Listing of VIFs on the NSX Manager
[**ListVifsOnEnforcementPoint_0**](PolicyInfraApi.md#ListVifsOnEnforcementPoint_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/vifs | Listing of VIFs on the NSX Manager
[**ListVirtualMachinesOnEnforcementPoint**](PolicyInfraApi.md#ListVirtualMachinesOnEnforcementPoint) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines | Listing of Virtual machines on the NSX Manager
[**ListVirtualMachinesOnEnforcementPoint_0**](PolicyInfraApi.md#ListVirtualMachinesOnEnforcementPoint_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines | Listing of Virtual machines on the NSX Manager
[**PatchDomainDeploymentMapForInfra**](PolicyInfraApi.md#PatchDomainDeploymentMapForInfra) | **Patch** /infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Patch Domain Deployment Map under infra
[**PatchDomainDeploymentMapForInfra_0**](PolicyInfraApi.md#PatchDomainDeploymentMapForInfra_0) | **Patch** /global-infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Patch Domain Deployment Map under infra
[**PatchDomainForInfra**](PolicyInfraApi.md#PatchDomainForInfra) | **Patch** /global-infra/domains/{domain-id} | Patch a domain
[**PatchDomainForInfra_0**](PolicyInfraApi.md#PatchDomainForInfra_0) | **Patch** /infra/domains/{domain-id} | Patch a domain
[**PatchEnforcementPointForInfra**](PolicyInfraApi.md#PatchEnforcementPointForInfra) | **Patch** /infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Patch a new Enforcement Point under infra
[**PatchEnforcementPointForInfra_0**](PolicyInfraApi.md#PatchEnforcementPointForInfra_0) | **Patch** /global-infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Patch a new Enforcement Point under infra
[**PatchEnforcementPointForSite**](PolicyInfraApi.md#PatchEnforcementPointForSite) | **Patch** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Patch a new Enforcement Point under Site
[**PatchEnforcementPointForSite_0**](PolicyInfraApi.md#PatchEnforcementPointForSite_0) | **Patch** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Patch a new Enforcement Point under Site
[**PatchGlobalManagerConfig**](PolicyInfraApi.md#PatchGlobalManagerConfig) | **Patch** /infra/global-manager-config | Create or patch Global Manager Config
[**PatchGlobalManagerConfig_0**](PolicyInfraApi.md#PatchGlobalManagerConfig_0) | **Patch** /global-infra/global-manager-config | Create or patch Global Manager Config
[**PatchInfra**](PolicyInfraApi.md#PatchInfra) | **Patch** /infra | Update the infra including all the nested entities
[**PatchInfraSite**](PolicyInfraApi.md#PatchInfraSite) | **Patch** /global-infra/sites/{site-id} | Create or patch Site
[**PatchInfraSite_0**](PolicyInfraApi.md#PatchInfraSite_0) | **Patch** /infra/sites/{site-id} | Create or patch Site
[**PatchInfra_0**](PolicyInfraApi.md#PatchInfra_0) | **Patch** /global-infra | Update the infra including all the nested entities
[**PatchTenantConstraint**](PolicyInfraApi.md#PatchTenantConstraint) | **Patch** /infra/constraints/{constraint-id} | Create or update tenant Constraint
[**PatchTenantConstraint_0**](PolicyInfraApi.md#PatchTenantConstraint_0) | **Patch** /global-infra/constraints/{constraint-id} | Create or update tenant Constraint
[**PatchTlsCertificate**](PolicyInfraApi.md#PatchTlsCertificate) | **Patch** /infra/certificates/{certificate-id} | Add a New Certificate
[**PatchTlsCertificate_0**](PolicyInfraApi.md#PatchTlsCertificate_0) | **Patch** /global-infra/certificates/{certificate-id} | Add a New Certificate
[**ReadDeploymentZoneInfra**](PolicyInfraApi.md#ReadDeploymentZoneInfra) | **Get** /global-infra/deployment-zones/{deployment-zone-id} | Read a DeploymentZone
[**ReadDeploymentZoneInfra_0**](PolicyInfraApi.md#ReadDeploymentZoneInfra_0) | **Get** /infra/deployment-zones/{deployment-zone-id} | Read a DeploymentZone
[**ReadDomainDeploymentMapForInfra**](PolicyInfraApi.md#ReadDomainDeploymentMapForInfra) | **Get** /infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Read a DomainDeploymentMap
[**ReadDomainDeploymentMapForInfra_0**](PolicyInfraApi.md#ReadDomainDeploymentMapForInfra_0) | **Get** /global-infra/domains/{domain-id}/domain-deployment-maps/{domain-deployment-map-id} | Read a DomainDeploymentMap
[**ReadDomainForInfra**](PolicyInfraApi.md#ReadDomainForInfra) | **Get** /global-infra/domains/{domain-id} | Read domain
[**ReadDomainForInfra_0**](PolicyInfraApi.md#ReadDomainForInfra_0) | **Get** /infra/domains/{domain-id} | Read domain
[**ReadEdgeClusterForEnforcementPoint**](PolicyInfraApi.md#ReadEdgeClusterForEnforcementPoint) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id} | Read a Edge Cluster under an Enforcement Point
[**ReadEdgeClusterForEnforcementPoint_0**](PolicyInfraApi.md#ReadEdgeClusterForEnforcementPoint_0) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id} | Read a Edge Cluster under an Enforcement Point
[**ReadEdgeNodeUnderEdgeClusterForEnforcementPoint**](PolicyInfraApi.md#ReadEdgeNodeUnderEdgeClusterForEnforcementPoint) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id}/edge-nodes/{edge-node-id} | Read a Edge Node under an Enforcement Point, Edge Cluster
[**ReadEdgeNodeUnderEdgeClusterForEnforcementPoint_0**](PolicyInfraApi.md#ReadEdgeNodeUnderEdgeClusterForEnforcementPoint_0) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/edge-clusters/{edge-cluster-id}/edge-nodes/{edge-node-id} | Read a Edge Node under an Enforcement Point, Edge Cluster
[**ReadEnforcementPointForInfra**](PolicyInfraApi.md#ReadEnforcementPointForInfra) | **Get** /infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Read an Enforcement Point
[**ReadEnforcementPointForInfra_0**](PolicyInfraApi.md#ReadEnforcementPointForInfra_0) | **Get** /global-infra/deployment-zones/{deployment-zone-id}/enforcement-points/{enforcementpoint-id} | Read an Enforcement Point
[**ReadEnforcementPointForSite**](PolicyInfraApi.md#ReadEnforcementPointForSite) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Read an Enforcement Point under Infra/Site
[**ReadEnforcementPointForSite_0**](PolicyInfraApi.md#ReadEnforcementPointForSite_0) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id} | Read an Enforcement Point under Infra/Site
[**ReadEnforcementPointRealizedState**](PolicyInfraApi.md#ReadEnforcementPointRealizedState) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name} | Read Enforcement Point
[**ReadEnforcementPointRealizedState_0**](PolicyInfraApi.md#ReadEnforcementPointRealizedState_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name} | Read Enforcement Point
[**ReadFirewallSectionRealizedState**](PolicyInfraApi.md#ReadFirewallSectionRealizedState) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/firewalls/firewall-sections/{firewall-section-id} | Read Firewall
[**ReadFirewallSectionRealizedState_0**](PolicyInfraApi.md#ReadFirewallSectionRealizedState_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/firewalls/firewall-sections/{firewall-section-id} | Read Firewall
[**ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData**](PolicyInfraApi.md#ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData) | **Get** /infra/global-manager-config?action&#x3D;show-sensitive-data | Read Global Manager config along with sensitive data
[**ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData_0**](PolicyInfraApi.md#ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData_0) | **Get** /global-infra/global-manager-config?action&#x3D;show-sensitive-data | Read Global Manager config along with sensitive data
[**ReadIPSetRealizedState**](PolicyInfraApi.md#ReadIPSetRealizedState) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/ip-sets/ip-sets-nsxt/{ip-set-name} | Read IPSet Realized state
[**ReadIPSetRealizedState_0**](PolicyInfraApi.md#ReadIPSetRealizedState_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/ip-sets/ip-sets-nsxt/{ip-set-name} | Read IPSet Realized state
[**ReadInfra**](PolicyInfraApi.md#ReadInfra) | **Get** /infra | Read infra
[**ReadInfra_0**](PolicyInfraApi.md#ReadInfra_0) | **Get** /global-infra | Read infra
[**ReadIntentStatus**](PolicyInfraApi.md#ReadIntentStatus) | **Get** /infra/realized-state/status | Get consolidated status of an intent object
[**ReadIntentStatus_0**](PolicyInfraApi.md#ReadIntentStatus_0) | **Get** /global-infra/realized-state/status | Get consolidated status of an intent object
[**ReadMACSetRealizedState**](PolicyInfraApi.md#ReadMACSetRealizedState) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/mac-sets/mac-sets-nsxt/{mac-set-name} | Read MACSet Realized state
[**ReadMACSetRealizedState_0**](PolicyInfraApi.md#ReadMACSetRealizedState_0) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/mac-sets/mac-sets-nsxt/{mac-set-name} | Read MACSet Realized state
[**ReadNSGroupRealizedState**](PolicyInfraApi.md#ReadNSGroupRealizedState) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/groups/nsgroups/{nsgroup-name} | Read Group
[**ReadNSGroupRealizedState_0**](PolicyInfraApi.md#ReadNSGroupRealizedState_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/groups/nsgroups/{nsgroup-name} | Read Group
[**ReadNSServiceRealizedState**](PolicyInfraApi.md#ReadNSServiceRealizedState) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/services/nsservices/{nsservice-name} | Read NSService
[**ReadNSServiceRealizedState_0**](PolicyInfraApi.md#ReadNSServiceRealizedState_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/services/nsservices/{nsservice-name} | Read NSService
[**ReadPolicyLabelForInfra**](PolicyInfraApi.md#ReadPolicyLabelForInfra) | **Get** /global-infra/labels/{label-id} | Read lable
[**ReadPolicyLabelForInfra_0**](PolicyInfraApi.md#ReadPolicyLabelForInfra_0) | **Get** /infra/labels/{label-id} | Read lable
[**ReadRealizedEntity**](PolicyInfraApi.md#ReadRealizedEntity) | **Get** /global-infra/realized-state/realized-entity | Get realized entity uniquely identified by realized path
[**ReadRealizedEntity_0**](PolicyInfraApi.md#ReadRealizedEntity_0) | **Get** /infra/realized-state/realized-entity | Get realized entity uniquely identified by realized path
[**ReadSecurityGroupRealizedState**](PolicyInfraApi.md#ReadSecurityGroupRealizedState) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/groups/securitygroups/{securitygroup-name} | Read Group
[**ReadSecurityGroupRealizedState_0**](PolicyInfraApi.md#ReadSecurityGroupRealizedState_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/groups/securitygroups/{securitygroup-name} | Read Group
[**ReadSite**](PolicyInfraApi.md#ReadSite) | **Get** /global-infra/sites/{site-id} | Read a site
[**ReadSite_0**](PolicyInfraApi.md#ReadSite_0) | **Get** /infra/sites/{site-id} | Read a site
[**ReadTenantConstraint**](PolicyInfraApi.md#ReadTenantConstraint) | **Get** /infra/constraints/{constraint-id} | Read tenant Constraint.
[**ReadTenantConstraint_0**](PolicyInfraApi.md#ReadTenantConstraint_0) | **Get** /global-infra/constraints/{constraint-id} | Read tenant Constraint.
[**ReadTransportZoneForEnforcementPoint**](PolicyInfraApi.md#ReadTransportZoneForEnforcementPoint) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/transport-zones/{transport-zone-id} | Read a Transport Zone under an Enforcement Point
[**ReadTransportZoneForEnforcementPoint_0**](PolicyInfraApi.md#ReadTransportZoneForEnforcementPoint_0) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/transport-zones/{transport-zone-id} | Read a Transport Zone under an Enforcement Point
[**ReadVirtualMachineDetails**](PolicyInfraApi.md#ReadVirtualMachineDetails) | **Get** /infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines/{virtual-machine-id}/details | Read the details of a virtual machine on the NSX Manager
[**ReadVirtualMachineDetails_0**](PolicyInfraApi.md#ReadVirtualMachineDetails_0) | **Get** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines/{virtual-machine-id}/details | Read the details of a virtual machine on the NSX Manager
[**RefreshRealizedStateRefresh**](PolicyInfraApi.md#RefreshRealizedStateRefresh) | **Post** /global-infra/realized-state/realized-entity?action&#x3D;refresh | Refresh all realized entities associated with the intent-path
[**RefreshRealizedStateRefresh_0**](PolicyInfraApi.md#RefreshRealizedStateRefresh_0) | **Post** /infra/realized-state/realized-entity?action&#x3D;refresh | Refresh all realized entities associated with the intent-path
[**ReloadEnforcementPointForSiteReload**](PolicyInfraApi.md#ReloadEnforcementPointForSiteReload) | **Post** /global-infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}?action&#x3D;reload | Reload an Enforcement Point under Site
[**ReloadEnforcementPointForSiteReload_0**](PolicyInfraApi.md#ReloadEnforcementPointForSiteReload_0) | **Post** /infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}?action&#x3D;reload | Reload an Enforcement Point under Site
[**TagBulkUpdate**](PolicyInfraApi.md#TagBulkUpdate) | **Put** /infra/tags/tag-operations/{operation-id} | Assign or Unassign tag on multiple Virtual Machines.
[**TagBulkUpdate_0**](PolicyInfraApi.md#TagBulkUpdate_0) | **Put** /global-infra/tags/tag-operations/{operation-id} | Assign or Unassign tag on multiple Virtual Machines.
[**TagVirtualMachineUpdateTags**](PolicyInfraApi.md#TagVirtualMachineUpdateTags) | **Post** /global-infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines?action&#x3D;update_tags | Apply tags on virtual machine
[**TagVirtualMachineUpdateTags_0**](PolicyInfraApi.md#TagVirtualMachineUpdateTags_0) | **Post** /infra/realized-state/enforcement-points/{enforcement-point-name}/virtual-machines?action&#x3D;update_tags | Apply tags on virtual machine
[**UpdateDomainForInfra**](PolicyInfraApi.md#UpdateDomainForInfra) | **Put** /global-infra/domains/{domain-id} | Create or update a domain
[**UpdateDomainForInfra_0**](PolicyInfraApi.md#UpdateDomainForInfra_0) | **Put** /infra/domains/{domain-id} | Create or update a domain
[**UpdateInfra**](PolicyInfraApi.md#UpdateInfra) | **Put** /infra | Update the infra including all the nested entities
[**UpdateInfra_0**](PolicyInfraApi.md#UpdateInfra_0) | **Put** /global-infra | Update the infra including all the nested entities
[**UpdatePolicyLabelForInfra**](PolicyInfraApi.md#UpdatePolicyLabelForInfra) | **Patch** /global-infra/labels/{label-id} | Patch an existing label object
[**UpdatePolicyLabelForInfra_0**](PolicyInfraApi.md#UpdatePolicyLabelForInfra_0) | **Patch** /infra/labels/{label-id} | Patch an existing label object

# **AddTlsCertificate**
> TlsCertificate AddTlsCertificate(ctx, body, certificateId)
Add a New Certificate

Adds a new private-public certificate and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store.  A certificate chain will not be expanded into separate certificate instances for reference, but would be pushed to the enforcement point as a single certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsTrustData**](TlsTrustData.md)|  | 
  **certificateId** | **string**|  | 

### Return type

[**TlsCertificate**](TlsCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTlsCertificate_0**
> TlsCertificate AddTlsCertificate_0(ctx, body, certificateId)
Add a New Certificate

Adds a new private-public certificate and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store.  A certificate chain will not be expanded into separate certificate instances for reference, but would be pushed to the enforcement point as a single certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsTrustData**](TlsTrustData.md)|  | 
  **certificateId** | **string**|  | 

### Return type

[**TlsCertificate**](TlsCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrPatchInfraReaction**
> CreateOrPatchInfraReaction(ctx, body, reactionId)
Create or patch a Reaction

Create or patch a Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Reaction**](Reaction.md)|  | 
  **reactionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrPatchInfraReaction_0**
> CreateOrPatchInfraReaction_0(ctx, body, reactionId)
Create or patch a Reaction

Create or patch a Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Reaction**](Reaction.md)|  | 
  **reactionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrPatchTlsCrl**
> CreateOrPatchTlsCrl(ctx, body, crlId)
Create or patch a Certificate Revocation List

Create or patch a Certificate Revocation List for the given id. The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL must contain PEM data for a single CRL. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrPatchTlsCrl_0**
> CreateOrPatchTlsCrl_0(ctx, body, crlId)
Create or patch a Certificate Revocation List

Create or patch a Certificate Revocation List for the given id. The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL must contain PEM data for a single CRL. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplacePolicyLabelForInfra**
> PolicyLabel CreateOrReplacePolicyLabelForInfra(ctx, body, labelId)
Create or replace label

Create label if not exists, otherwise replaces the existing label. If label already exists then type attribute cannot be changed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyLabel**](PolicyLabel.md)|  | 
  **labelId** | **string**|  | 

### Return type

[**PolicyLabel**](PolicyLabel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplacePolicyLabelForInfra_0**
> PolicyLabel CreateOrReplacePolicyLabelForInfra_0(ctx, body, labelId)
Create or replace label

Create label if not exists, otherwise replaces the existing label. If label already exists then type attribute cannot be changed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyLabel**](PolicyLabel.md)|  | 
  **labelId** | **string**|  | 

### Return type

[**PolicyLabel**](PolicyLabel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceTenantConstraint**
> Constraint CreateOrReplaceTenantConstraint(ctx, body, constraintId)
Create or update tenant Constraint

Create tenant constraint if it does not exist, otherwise replace the existing constraint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Constraint**](Constraint.md)|  | 
  **constraintId** | **string**|  | 

### Return type

[**Constraint**](Constraint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceTenantConstraint_0**
> Constraint CreateOrReplaceTenantConstraint_0(ctx, body, constraintId)
Create or update tenant Constraint

Create tenant constraint if it does not exist, otherwise replace the existing constraint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Constraint**](Constraint.md)|  | 
  **constraintId** | **string**|  | 

### Return type

[**Constraint**](Constraint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateDomainDeploymentMapForInfra**
> DomainDeploymentMap CreateOrUpdateDomainDeploymentMapForInfra(ctx, body, domainId, domainDeploymentMapId)
Create a new Domain Deployment Map under infra

If the passed Domain Deployment Map does not already exist, create a new Domain Deployment Map. If it already exist, replace it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainDeploymentMap**](DomainDeploymentMap.md)|  | 
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

[**DomainDeploymentMap**](DomainDeploymentMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateDomainDeploymentMapForInfra_0**
> DomainDeploymentMap CreateOrUpdateDomainDeploymentMapForInfra_0(ctx, body, domainId, domainDeploymentMapId)
Create a new Domain Deployment Map under infra

If the passed Domain Deployment Map does not already exist, create a new Domain Deployment Map. If it already exist, replace it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainDeploymentMap**](DomainDeploymentMap.md)|  | 
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

[**DomainDeploymentMap**](DomainDeploymentMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEnforcementPointForInfra**
> EnforcementPoint CreateOrUpdateEnforcementPointForInfra(ctx, body, deploymentZoneId, enforcementpointId)
Create/update a new Enforcement Point under infra

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PUT /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEnforcementPointForInfra_0**
> EnforcementPoint CreateOrUpdateEnforcementPointForInfra_0(ctx, body, deploymentZoneId, enforcementpointId)
Create/update a new Enforcement Point under infra

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PUT /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEnforcementPointForSite**
> EnforcementPoint CreateOrUpdateEnforcementPointForSite(ctx, body, siteId, enforcementpointId)
Create/update a new Enforcement Point under Site

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEnforcementPointForSite_0**
> EnforcementPoint CreateOrUpdateEnforcementPointForSite_0(ctx, body, siteId, enforcementpointId)
Create/update a new Enforcement Point under Site

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, replace it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateGlobalManagerConfig**
> GlobalManagerConfig CreateOrUpdateGlobalManagerConfig(ctx, body)
Create or fully replace Global Manager Config

Create or fully replace a Global Manager Config. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManagerConfig**](GlobalManagerConfig.md)|  | 

### Return type

[**GlobalManagerConfig**](GlobalManagerConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateGlobalManagerConfig_0**
> GlobalManagerConfig CreateOrUpdateGlobalManagerConfig_0(ctx, body)
Create or fully replace Global Manager Config

Create or fully replace a Global Manager Config. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManagerConfig**](GlobalManagerConfig.md)|  | 

### Return type

[**GlobalManagerConfig**](GlobalManagerConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInfraReaction**
> Reaction CreateOrUpdateInfraReaction(ctx, body, reactionId)
Create or fully replace a Reaction

Create or fully replace a Reaction under Infra. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Reaction**](Reaction.md)|  | 
  **reactionId** | **string**|  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInfraReaction_0**
> Reaction CreateOrUpdateInfraReaction_0(ctx, body, reactionId)
Create or fully replace a Reaction

Create or fully replace a Reaction under Infra. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Reaction**](Reaction.md)|  | 
  **reactionId** | **string**|  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInfraSite**
> Site CreateOrUpdateInfraSite(ctx, body, siteId)
Create or fully replace a Site under infra

Create or fully replace a Site under Infra. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Site**](Site.md)|  | 
  **siteId** | **string**|  | 

### Return type

[**Site**](Site.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInfraSite_0**
> Site CreateOrUpdateInfraSite_0(ctx, body, siteId)
Create or fully replace a Site under infra

Create or fully replace a Site under Infra. Revision is optional for creation and required for update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Site**](Site.md)|  | 
  **siteId** | **string**|  | 

### Return type

[**Site**](Site.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateTlsCrl**
> TlsCrl CreateOrUpdateTlsCrl(ctx, body, crlId)
Create or fully replace a Certificate Revocation List

Create or replace a Certificate Revocation List for the given id. The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL must contain PEM data for a single CRL. Revision is required. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

[**TlsCrl**](TlsCrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateTlsCrl_0**
> TlsCrl CreateOrUpdateTlsCrl_0(ctx, body, crlId)
Create or fully replace a Certificate Revocation List

Create or replace a Certificate Revocation List for the given id. The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL must contain PEM data for a single CRL. Revision is required. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

[**TlsCrl**](TlsCrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTlsCrlImport**
> TlsCrlListResult CreateTlsCrlImport(ctx, body, crlId)
Create a new Certificate Revocation List

Adds a new certificate revocation list (CRLs). The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL can contain a single CRL or multiple CRLs depending on the PEM data. - Single CRL: a single CRL is created with the given id. - Composite CRL: multiple CRLs are generated. Each of the CRL is created with an id generated based on the given id. First CRL is created with crl-id, second with crl-id-1, third with crl-id-2, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

[**TlsCrlListResult**](TlsCrlListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTlsCrlImport_0**
> TlsCrlListResult CreateTlsCrlImport_0(ctx, body, crlId)
Create a new Certificate Revocation List

Adds a new certificate revocation list (CRLs). The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. The CRL can contain a single CRL or multiple CRLs depending on the PEM data. - Single CRL: a single CRL is created with the given id. - Composite CRL: multiple CRLs are generated. Each of the CRL is created with an id generated based on the given id. First CRL is created with crl-id, second with crl-id-1, third with crl-id-2, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsCrl**](TlsCrl.md)|  | 
  **crlId** | **string**|  | 

### Return type

[**TlsCrlListResult**](TlsCrlListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomain**
> DeleteDomain(ctx, domainId)
Delete Domain and all the entities contained by this domain

Delete the domain along with all the entities contained by this domain. The groups that are a part of this domain are also deleted along with the domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomainDeploymentMap**
> DeleteDomainDeploymentMap(ctx, domainId, domainDeploymentMapId)
Delete Domain Deployment Map

Delete Domain Deployment Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomainDeploymentMap_0**
> DeleteDomainDeploymentMap_0(ctx, domainId, domainDeploymentMapId)
Delete Domain Deployment Map

Delete Domain Deployment Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomain_0**
> DeleteDomain_0(ctx, domainId)
Delete Domain and all the entities contained by this domain

Delete the domain along with all the entities contained by this domain. The groups that are a part of this domain are also deleted along with the domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnforcementPoint**
> DeleteEnforcementPoint(ctx, deploymentZoneId, enforcementpointId)
Delete EnforcementPoint

Delete EnforcementPoint. This is a deprecated API. DeploymentZone has been renamed to Site. Use DELETE /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnforcementPointForSite**
> DeleteEnforcementPointForSite(ctx, siteId, enforcementpointId)
Delete EnforcementPoint from Site

Delete EnforcementPoint from Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnforcementPointForSite_0**
> DeleteEnforcementPointForSite_0(ctx, siteId, enforcementpointId)
Delete EnforcementPoint from Site

Delete EnforcementPoint from Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnforcementPoint_0**
> DeleteEnforcementPoint_0(ctx, deploymentZoneId, enforcementpointId)
Delete EnforcementPoint

Delete EnforcementPoint. This is a deprecated API. DeploymentZone has been renamed to Site. Use DELETE /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraReaction**
> DeleteInfraReaction(ctx, reactionId)
Delete Reaction

Delete a Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reactionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraReaction_0**
> DeleteInfraReaction_0(ctx, reactionId)
Delete Reaction

Delete a Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reactionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraSite**
> DeleteInfraSite(ctx, siteId, optional)
Delete a site

Delete a site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
 **optional** | ***PolicyInfraApiDeleteInfraSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiDeleteInfraSiteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInfraSite_0**
> DeleteInfraSite_0(ctx, siteId, optional)
Delete a site

Delete a site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
 **optional** | ***PolicyInfraApiDeleteInfraSite_19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiDeleteInfraSite_19Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyLabelForInfra**
> DeletePolicyLabelForInfra(ctx, labelId)
Delete PolicyLabel object

Delete PolicyLabel object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyLabelForInfra_0**
> DeletePolicyLabelForInfra_0(ctx, labelId)
Delete PolicyLabel object

Delete PolicyLabel object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenantConstraint**
> DeleteTenantConstraint(ctx, constraintId)
Delete tenant Constraint.

Delete tenant constraint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **constraintId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenantConstraint_0**
> DeleteTenantConstraint_0(ctx, constraintId)
Delete tenant Constraint.

Delete tenant constraint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **constraintId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTlsCertificate**
> DeleteTlsCertificate(ctx, certificateId)
Delete Certificate for the Given Certificate ID

Removes the specified certificate. The private key associated with the certificate is also deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| ID of certificate to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTlsCertificate_0**
> DeleteTlsCertificate_0(ctx, certificateId)
Delete Certificate for the Given Certificate ID

Removes the specified certificate. The private key associated with the certificate is also deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| ID of certificate to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTlsCrl**
> DeleteTlsCrl(ctx, crlId)
Delete a CRL

Deletes an existing CRL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTlsCrl_0**
> DeleteTlsCrl_0(ctx, crlId)
Delete a CRL

Deletes an existing CRL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FullSyncEnforcementPointForSiteFullSync**
> FullSyncEnforcementPointForSiteFullSync(ctx, siteId, enforcementPointId)
Full sync EnforcementPoint from Site

Full sync EnforcementPoint from Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FullSyncEnforcementPointForSiteFullSync_0**
> FullSyncEnforcementPointForSiteFullSync_0(ctx, siteId, enforcementPointId)
Full sync EnforcementPoint from Site

Full sync EnforcementPoint from Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraReaction**
> Reaction GetInfraReaction(ctx, reactionId)
Get Reaction

Get Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reactionId** | **string**|  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraReaction_0**
> Reaction GetInfraReaction_0(ctx, reactionId)
Get Reaction

Get Reaction under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reactionId** | **string**|  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraSiteListenerCertificate**
> TlsListenerCertificate GetInfraSiteListenerCertificate(ctx, address, port)
Returns the certificate of the listener

Connects to the given IP and port, and, if an SSL listener is present, returns the certificate of the listener. Intent of this API is \"Do you trust this certificate?\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Host name or IP address of TLS listener | 
  **port** | **int32**| TCP port number of the TLS listener | 

### Return type

[**TlsListenerCertificate**](TlsListenerCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfraSiteListenerCertificate_0**
> TlsListenerCertificate GetInfraSiteListenerCertificate_0(ctx, address, port)
Returns the certificate of the listener

Connects to the given IP and port, and, if an SSL listener is present, returns the certificate of the listener. Intent of this API is \"Do you trust this certificate?\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Host name or IP address of TLS listener | 
  **port** | **int32**| TCP port number of the TLS listener | 

### Return type

[**TlsListenerCertificate**](TlsListenerCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagBulkOperation**
> TagBulkOperation GetTagBulkOperation(ctx, operationId)
Get details of tag bulk operation request

Get details of tag bulk operation request with which tag is applied or removed on virtual machines. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperation**](TagBulkOperation.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagBulkOperationStatus**
> TagBulkOperationStatus GetTagBulkOperationStatus(ctx, operationId)
Get status of tag bulk operation

Get status of tag bulk operation with details of tag operation on each virtual machine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperationStatus**](TagBulkOperationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagBulkOperationStatus_0**
> TagBulkOperationStatus GetTagBulkOperationStatus_0(ctx, operationId)
Get status of tag bulk operation

Get status of tag bulk operation with details of tag operation on each virtual machine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperationStatus**](TagBulkOperationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagBulkOperation_0**
> TagBulkOperation GetTagBulkOperation_0(ctx, operationId)
Get details of tag bulk operation request

Get details of tag bulk operation request with which tag is applied or removed on virtual machines. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperation**](TagBulkOperation.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTlsCertificate**
> TlsCertificate GetTlsCertificate(ctx, certificateId, optional)
Show Certificate Data for the Given Certificate ID

Returns information for the specified certificate ID, including the certificate's id; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| ID of certificate to read | 
 **optional** | ***PolicyInfraApiGetTlsCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiGetTlsCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**TlsCertificate**](TlsCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTlsCertificate_0**
> TlsCertificate GetTlsCertificate_0(ctx, certificateId, optional)
Show Certificate Data for the Given Certificate ID

Returns information for the specified certificate ID, including the certificate's id; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| ID of certificate to read | 
 **optional** | ***PolicyInfraApiGetTlsCertificate_29Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiGetTlsCertificate_29Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**TlsCertificate**](TlsCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTlsCrl**
> TlsCrl GetTlsCrl(ctx, crlId, optional)
Show CRL Data for the Given CRL id.

Returns information about the specified CRL. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 
 **optional** | ***PolicyInfraApiGetTlsCrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiGetTlsCrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**TlsCrl**](TlsCrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTlsCrl_0**
> TlsCrl GetTlsCrl_0(ctx, crlId, optional)
Show CRL Data for the Given CRL id.

Returns information about the specified CRL. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 
 **optional** | ***PolicyInfraApiGetTlsCrl_30Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiGetTlsCrl_30Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]

### Return type

[**TlsCrl**](TlsCrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlarms**
> PolicyAlarmResourceListResult ListAlarms(ctx, optional)
List All alarms in the system

Paginated list of all alarms. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyAlarmResourceListResult**](PolicyAlarmResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlarms_0**
> PolicyAlarmResourceListResult ListAlarms_0(ctx, optional)
List All alarms in the system

Paginated list of all alarms. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAlarms_31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAlarms_31Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyAlarmResourceListResult**](PolicyAlarmResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllTags**
> TagInfoListResult ListAllTags(ctx, optional)
List all unique tags.

Returns paginated list of all unique tags. Supports filtering by scope, tag and source from which tags are synched. Supports starts with, equals and contains operators on scope and tag values. To filter tags by starts with on scope or tag, use '*' as prefix before the value. To filter tags by ends with on scope or tag, use '*' as suffix after the value. To filter tags by contain on scope or tag, use '*' as prefix and suffix on the value. Below special characters in the filter value needs to be escaped with hex values. - Character '&' needs to be escaped as '%26' - Character '[' needs to be escaped as '%5B' - Character ']' needs to be escaped as '%5D' - Character '+' needs to be escaped as '%2B' - Character '#' needs to be escaped as '%23' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **scope** | **optional.String**| Tag scope | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Source from which tags are synced. | 
 **tag** | **optional.String**| Tag value | 

### Return type

[**TagInfoListResult**](TagInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllTags_0**
> TagInfoListResult ListAllTags_0(ctx, optional)
List all unique tags.

Returns paginated list of all unique tags. Supports filtering by scope, tag and source from which tags are synched. Supports starts with, equals and contains operators on scope and tag values. To filter tags by starts with on scope or tag, use '*' as prefix before the value. To filter tags by ends with on scope or tag, use '*' as suffix after the value. To filter tags by contain on scope or tag, use '*' as prefix and suffix on the value. Below special characters in the filter value needs to be escaped with hex values. - Character '&' needs to be escaped as '%26' - Character '[' needs to be escaped as '%5B' - Character ']' needs to be escaped as '%5D' - Character '+' needs to be escaped as '%2B' - Character '#' needs to be escaped as '%23' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllTags_32Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllTags_32Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **scope** | **optional.String**| Tag scope | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Source from which tags are synced. | 
 **tag** | **optional.String**| Tag value | 

### Return type

[**TagInfoListResult**](TagInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllUnAssociatedVirtualMachines**
> VirtualMachineListResult ListAllUnAssociatedVirtualMachines(ctx, optional)
List all virtual machines which are not part of any group

This API filters objects of type virtual machine which are not part of any group. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllUnAssociatedVirtualMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllUnAssociatedVirtualMachinesOpts struct
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

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllUnAssociatedVirtualMachines_0**
> VirtualMachineListResult ListAllUnAssociatedVirtualMachines_0(ctx, optional)
List all virtual machines which are not part of any group

This API filters objects of type virtual machine which are not part of any group. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllUnAssociatedVirtualMachines_33Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllUnAssociatedVirtualMachines_33Opts struct
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

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllVirtualMachines**
> VirtualMachineListResult ListAllVirtualMachines(ctx, optional)
List all virtual machines

This API filters objects of type virtual machine. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllVirtualMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllVirtualMachinesOpts struct
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

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllVirtualMachines_0**
> VirtualMachineListResult ListAllVirtualMachines_0(ctx, optional)
List all virtual machines

This API filters objects of type virtual machine. This API also gives some VM details such as VM name, IDs and the current state of the VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListAllVirtualMachines_34Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListAllVirtualMachines_34Opts struct
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

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeploymentZonesForInfra**
> DeploymentZoneListResult ListDeploymentZonesForInfra(ctx, optional)
List Deployment Zones for infra

Paginated list of all Deployment zones for infra. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListDeploymentZonesForInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDeploymentZonesForInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DeploymentZoneListResult**](DeploymentZoneListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeploymentZonesForInfra_0**
> DeploymentZoneListResult ListDeploymentZonesForInfra_0(ctx, optional)
List Deployment Zones for infra

Paginated list of all Deployment zones for infra. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListDeploymentZonesForInfra_35Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDeploymentZonesForInfra_35Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DeploymentZoneListResult**](DeploymentZoneListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomainDeploymentMapsForInfra**
> DomainDeploymentMapListResult ListDomainDeploymentMapsForInfra(ctx, domainId, optional)
List Domain Deployment maps for infra

Paginated list of all Domain Deployment Entries for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicyInfraApiListDomainDeploymentMapsForInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDomainDeploymentMapsForInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DomainDeploymentMapListResult**](DomainDeploymentMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomainDeploymentMapsForInfra_0**
> DomainDeploymentMapListResult ListDomainDeploymentMapsForInfra_0(ctx, domainId, optional)
List Domain Deployment maps for infra

Paginated list of all Domain Deployment Entries for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicyInfraApiListDomainDeploymentMapsForInfra_36Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDomainDeploymentMapsForInfra_36Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DomainDeploymentMapListResult**](DomainDeploymentMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomainForInfra**
> DomainListResult ListDomainForInfra(ctx, optional)
List domains for infra

Paginated list of all domains for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListDomainForInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDomainForInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DomainListResult**](DomainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomainForInfra_0**
> DomainListResult ListDomainForInfra_0(ctx, optional)
List domains for infra

Paginated list of all domains for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListDomainForInfra_37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListDomainForInfra_37Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DomainListResult**](DomainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEdgeClustersForEnforcementPoint**
> PolicyEdgeClusterListResult ListEdgeClustersForEnforcementPoint(ctx, siteId, enforcementpointId, optional)
List Edge Clusters under an Enforcement Point

Paginated list of all Edge Clusters under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEdgeClustersForEnforcementPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEdgeClustersForEnforcementPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyEdgeClusterListResult**](PolicyEdgeClusterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEdgeClustersForEnforcementPoint_0**
> PolicyEdgeClusterListResult ListEdgeClustersForEnforcementPoint_0(ctx, siteId, enforcementpointId, optional)
List Edge Clusters under an Enforcement Point

Paginated list of all Edge Clusters under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEdgeClustersForEnforcementPoint_38Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEdgeClustersForEnforcementPoint_38Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyEdgeClusterListResult**](PolicyEdgeClusterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEdgeNodesUnderEdgeClusterForEnforcementPoint**
> PolicyEdgeNodeListResult ListEdgeNodesUnderEdgeClusterForEnforcementPoint(ctx, siteId, enforcementpointId, edgeClusterId, optional)
List Edge Nodes under an Enforcement Point, Edge Cluster

Paginated list of all Edge Nodes under an Enforcement Point, Edge Cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEdgeNodesUnderEdgeClusterForEnforcementPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEdgeNodesUnderEdgeClusterForEnforcementPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyEdgeNodeListResult**](PolicyEdgeNodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEdgeNodesUnderEdgeClusterForEnforcementPoint_0**
> PolicyEdgeNodeListResult ListEdgeNodesUnderEdgeClusterForEnforcementPoint_0(ctx, siteId, enforcementpointId, edgeClusterId, optional)
List Edge Nodes under an Enforcement Point, Edge Cluster

Paginated list of all Edge Nodes under an Enforcement Point, Edge Cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEdgeNodesUnderEdgeClusterForEnforcementPoint_39Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEdgeNodesUnderEdgeClusterForEnforcementPoint_39Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyEdgeNodeListResult**](PolicyEdgeNodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointForInfra**
> EnforcementPointListResult ListEnforcementPointForInfra(ctx, deploymentZoneId, optional)
List enforcementpoints for infra

Paginated list of all enforcementpoints for infra. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEnforcementPointForInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointForInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EnforcementPointListResult**](EnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointForInfra_0**
> EnforcementPointListResult ListEnforcementPointForInfra_0(ctx, deploymentZoneId, optional)
List enforcementpoints for infra

Paginated list of all enforcementpoints for infra. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEnforcementPointForInfra_40Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointForInfra_40Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EnforcementPointListResult**](EnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointForSite**
> EnforcementPointListResult ListEnforcementPointForSite(ctx, siteId, optional)
List enforcementpoints under Site

Paginated list of all enforcementpoints under Site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEnforcementPointForSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointForSiteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EnforcementPointListResult**](EnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointForSite_0**
> EnforcementPointListResult ListEnforcementPointForSite_0(ctx, siteId, optional)
List enforcementpoints under Site

Paginated list of all enforcementpoints under Site. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
 **optional** | ***PolicyInfraApiListEnforcementPointForSite_41Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointForSite_41Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EnforcementPointListResult**](EnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointRealizedStates**
> RealizedEnforcementPointListResult ListEnforcementPointRealizedStates(ctx, optional)
List Enforcement Points

Paginated list of all enforcement points. Returns the populated enforcement points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListEnforcementPointRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedEnforcementPointListResult**](RealizedEnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnforcementPointRealizedStates_0**
> RealizedEnforcementPointListResult ListEnforcementPointRealizedStates_0(ctx, optional)
List Enforcement Points

Paginated list of all enforcement points. Returns the populated enforcement points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListEnforcementPointRealizedStates_42Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListEnforcementPointRealizedStates_42Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedEnforcementPointListResult**](RealizedEnforcementPointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallSectionRealizedStates**
> RealizedFirewallSectionListResult ListFirewallSectionRealizedStates(ctx, enforcementPointName, optional)
List Firewall Sections

Paginated list of all Firewalls. Returns populated Firewalls. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListFirewallSectionRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListFirewallSectionRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedFirewallSectionListResult**](RealizedFirewallSectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallSectionRealizedStates_0**
> RealizedFirewallSectionListResult ListFirewallSectionRealizedStates_0(ctx, enforcementPointName, optional)
List Firewall Sections

Paginated list of all Firewalls. Returns populated Firewalls. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListFirewallSectionRealizedStates_43Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListFirewallSectionRealizedStates_43Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedFirewallSectionListResult**](RealizedFirewallSectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSetRealizedStates**
> GenericPolicyRealizedResourceListResult ListIPSetRealizedStates(ctx, enforcementPointName, optional)
List IPSets

Paginated list of all Realized IPSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListIPSetRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListIPSetRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSetRealizedStates_0**
> GenericPolicyRealizedResourceListResult ListIPSetRealizedStates_0(ctx, enforcementPointName, optional)
List IPSets

Paginated list of all Realized IPSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListIPSetRealizedStates_44Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListIPSetRealizedStates_44Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfraReactions**
> ReactionListResult ListInfraReactions(ctx, optional)
Get Reaction list result

Get paginated list of all Reactions under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListInfraReactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListInfraReactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ReactionListResult**](ReactionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInfraReactions_0**
> ReactionListResult ListInfraReactions_0(ctx, optional)
Get Reaction list result

Get paginated list of all Reactions under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListInfraReactions_45Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListInfraReactions_45Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ReactionListResult**](ReactionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMACSetRealizedStates**
> GenericPolicyRealizedResourceListResult ListMACSetRealizedStates(ctx, enforcementPointName, optional)
List MACSets

Paginated list of all Realized MACSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListMACSetRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListMACSetRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMACSetRealizedStates_0**
> GenericPolicyRealizedResourceListResult ListMACSetRealizedStates_0(ctx, enforcementPointName, optional)
List MACSets

Paginated list of all Realized MACSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListMACSetRealizedStates_46Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListMACSetRealizedStates_46Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSGroupRealizedStates**
> GenericPolicyRealizedResourceListResult ListNSGroupRealizedStates(ctx, enforcementPointName, optional)
List NS Groups

Paginated list of all NSGroups. Returns populated NSGroups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListNSGroupRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListNSGroupRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSGroupRealizedStates_0**
> GenericPolicyRealizedResourceListResult ListNSGroupRealizedStates_0(ctx, enforcementPointName, optional)
List NS Groups

Paginated list of all NSGroups. Returns populated NSGroups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListNSGroupRealizedStates_47Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListNSGroupRealizedStates_47Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSServiceRealizedStates**
> GenericPolicyRealizedResourceListResult ListNSServiceRealizedStates(ctx, enforcementPointName, optional)
List Realized NSServices

Paginated list of all Realized NSService. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListNSServiceRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListNSServiceRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSServiceRealizedStates_0**
> GenericPolicyRealizedResourceListResult ListNSServiceRealizedStates_0(ctx, enforcementPointName, optional)
List Realized NSServices

Paginated list of all Realized NSService. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListNSServiceRealizedStates_48Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListNSServiceRealizedStates_48Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyLabelForInfra**
> PolicyLabelListResult ListPolicyLabelForInfra(ctx, optional)
List labels for infra

Paginated list of all labels for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListPolicyLabelForInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListPolicyLabelForInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyLabelListResult**](PolicyLabelListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyLabelForInfra_0**
> PolicyLabelListResult ListPolicyLabelForInfra_0(ctx, optional)
List labels for infra

Paginated list of all labels for infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListPolicyLabelForInfra_49Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListPolicyLabelForInfra_49Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyLabelListResult**](PolicyLabelListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRealizedEntities**
> GenericPolicyRealizedResourceListResult ListRealizedEntities(ctx, intentPath, optional)
Get list of realized objects associated with intent object

Get list of realized entities associated with intent object, specified by path in query parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***PolicyInfraApiListRealizedEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListRealizedEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sitePath** | **optional.String**| Policy Path of the site | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRealizedEntities_0**
> GenericPolicyRealizedResourceListResult ListRealizedEntities_0(ctx, intentPath, optional)
Get list of realized objects associated with intent object

Get list of realized entities associated with intent object, specified by path in query parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***PolicyInfraApiListRealizedEntities_50Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListRealizedEntities_50Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sitePath** | **optional.String**| Policy Path of the site | 

### Return type

[**GenericPolicyRealizedResourceListResult**](GenericPolicyRealizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityGroupRealizedStates**
> RealizedSecurityGroupListResult ListSecurityGroupRealizedStates(ctx, enforcementPointName, optional)
List Security Groups

Paginated list of all Security Groups. Returns populated Security Groups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListSecurityGroupRealizedStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListSecurityGroupRealizedStatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedSecurityGroupListResult**](RealizedSecurityGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityGroupRealizedStates_0**
> RealizedSecurityGroupListResult ListSecurityGroupRealizedStates_0(ctx, enforcementPointName, optional)
List Security Groups

Paginated list of all Security Groups. Returns populated Security Groups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
 **optional** | ***PolicyInfraApiListSecurityGroupRealizedStates_51Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListSecurityGroupRealizedStates_51Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RealizedSecurityGroupListResult**](RealizedSecurityGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSites**
> SiteListResult ListSites(ctx, optional)
List Sites

List Sites under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SiteListResult**](SiteListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSites_0**
> SiteListResult ListSites_0(ctx, optional)
List Sites

List Sites under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListSites_52Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListSites_52Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SiteListResult**](SiteListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaggedObjects**
> PolicyResourceReferenceListResult ListTaggedObjects(ctx, optional)
List all objects assigned with matching scope and tag values

Paginated list of all objects assigned with matching scope and tag values. Objects are represented in form of resource reference. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTaggedObjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTaggedObjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **filterText** | **optional.String**| Filter text to restrict tagged objects list with matching filter text. | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **scope** | **optional.String**| Tag scope | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **tag** | **optional.String**| Tag value | 

### Return type

[**PolicyResourceReferenceListResult**](PolicyResourceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTaggedObjects_0**
> PolicyResourceReferenceListResult ListTaggedObjects_0(ctx, optional)
List all objects assigned with matching scope and tag values

Paginated list of all objects assigned with matching scope and tag values. Objects are represented in form of resource reference. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTaggedObjects_53Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTaggedObjects_53Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **filterText** | **optional.String**| Filter text to restrict tagged objects list with matching filter text. | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **scope** | **optional.String**| Tag scope | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **tag** | **optional.String**| Tag value | 

### Return type

[**PolicyResourceReferenceListResult**](PolicyResourceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTenantConstraints**
> ConstraintListResult ListTenantConstraints(ctx, optional)
List tenant Constraints.

List tenant constraints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTenantConstraintsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTenantConstraintsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ConstraintListResult**](ConstraintListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTenantConstraints_0**
> ConstraintListResult ListTenantConstraints_0(ctx, optional)
List tenant Constraints.

List tenant constraints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTenantConstraints_54Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTenantConstraints_54Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ConstraintListResult**](ConstraintListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTlsCertificates**
> TlsCertificateList ListTlsCertificates(ctx, optional)
Return All the User-Facing Components' Certificates

Returns all certificate information viewable by the user, including each certificate's id; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTlsCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTlsCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**TlsCertificateList**](TlsCertificateList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTlsCertificates_0**
> TlsCertificateList ListTlsCertificates_0(ctx, optional)
Return All the User-Facing Components' Certificates

Returns all certificate information viewable by the user, including each certificate's id; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTlsCertificates_55Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTlsCertificates_55Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**TlsCertificateList**](TlsCertificateList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTlsCrls**
> TlsCrlListResult ListTlsCrls(ctx, optional)
Return All Added CRLs

Returns information about all CRLs. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTlsCrlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTlsCrlsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**TlsCrlListResult**](TlsCrlListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTlsCrls_0**
> TlsCrlListResult ListTlsCrls_0(ctx, optional)
Return All Added CRLs

Returns information about all CRLs. For additional information, include the ?details=true modifier at the end of the request URI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiListTlsCrls_56Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTlsCrls_56Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **details** | **optional.Bool**| whether to expand the pem data and show all its details | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Type of certificate to return | 

### Return type

[**TlsCrlListResult**](TlsCrlListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportZonesForEnforcementPoint**
> PolicyTransportZoneListResult ListTransportZonesForEnforcementPoint(ctx, siteId, enforcementpointId, optional)
List Transport Zones under an Enforcement Point

Paginated list of all Transport Zones under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
 **optional** | ***PolicyInfraApiListTransportZonesForEnforcementPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTransportZonesForEnforcementPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyTransportZoneListResult**](PolicyTransportZoneListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportZonesForEnforcementPoint_0**
> PolicyTransportZoneListResult ListTransportZonesForEnforcementPoint_0(ctx, siteId, enforcementpointId, optional)
List Transport Zones under an Enforcement Point

Paginated list of all Transport Zones under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
 **optional** | ***PolicyInfraApiListTransportZonesForEnforcementPoint_57Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListTransportZonesForEnforcementPoint_57Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyTransportZoneListResult**](PolicyTransportZoneListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVifsOnEnforcementPoint**
> VirtualNetworkInterfaceListResult ListVifsOnEnforcementPoint(ctx, enforcementPointName, optional)
Listing of VIFs on the NSX Manager

This API lists VIFs from the specified NSX Manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
 **optional** | ***PolicyInfraApiListVifsOnEnforcementPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListVifsOnEnforcementPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **lportAttachmentId** | **optional.String**| LPort attachment ID of the VIF. | 
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

# **ListVifsOnEnforcementPoint_0**
> VirtualNetworkInterfaceListResult ListVifsOnEnforcementPoint_0(ctx, enforcementPointName, optional)
Listing of VIFs on the NSX Manager

This API lists VIFs from the specified NSX Manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
 **optional** | ***PolicyInfraApiListVifsOnEnforcementPoint_58Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListVifsOnEnforcementPoint_58Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **lportAttachmentId** | **optional.String**| LPort attachment ID of the VIF. | 
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

# **ListVirtualMachinesOnEnforcementPoint**
> SearchResponse ListVirtualMachinesOnEnforcementPoint(ctx, enforcementPointName, optional)
Listing of Virtual machines on the NSX Manager

This API filters objects of type virtual machines from the specified NSX Manager. This API has been deprecated. Please use the new API GET /infra/realized-state/virtual-machines 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
 **optional** | ***PolicyInfraApiListVirtualMachinesOnEnforcementPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListVirtualMachinesOnEnforcementPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **dsl** | **optional.String**| Search DSL (domain specific language) query | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **query** | **optional.String**| Search query | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualMachinesOnEnforcementPoint_0**
> SearchResponse ListVirtualMachinesOnEnforcementPoint_0(ctx, enforcementPointName, optional)
Listing of Virtual machines on the NSX Manager

This API filters objects of type virtual machines from the specified NSX Manager. This API has been deprecated. Please use the new API GET /infra/realized-state/virtual-machines 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
 **optional** | ***PolicyInfraApiListVirtualMachinesOnEnforcementPoint_59Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiListVirtualMachinesOnEnforcementPoint_59Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **dsl** | **optional.String**| Search DSL (domain specific language) query | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **query** | **optional.String**| Search query | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDomainDeploymentMapForInfra**
> PatchDomainDeploymentMapForInfra(ctx, body, domainId, domainDeploymentMapId)
Patch Domain Deployment Map under infra

If the passed Domain Deployment Map does not already exist, create a new Domain Deployment Map. If it already exist, patch it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainDeploymentMap**](DomainDeploymentMap.md)|  | 
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDomainDeploymentMapForInfra_0**
> PatchDomainDeploymentMapForInfra_0(ctx, body, domainId, domainDeploymentMapId)
Patch Domain Deployment Map under infra

If the passed Domain Deployment Map does not already exist, create a new Domain Deployment Map. If it already exist, patch it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainDeploymentMap**](DomainDeploymentMap.md)|  | 
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDomainForInfra**
> PatchDomainForInfra(ctx, body, domainId)
Patch a domain

If a domain with the domain-id is not already present, create a new domain. If it already exists, patch the domain 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Domain**](Domain.md)|  | 
  **domainId** | **string**| Domain ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDomainForInfra_0**
> PatchDomainForInfra_0(ctx, body, domainId)
Patch a domain

If a domain with the domain-id is not already present, create a new domain. If it already exists, patch the domain 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Domain**](Domain.md)|  | 
  **domainId** | **string**| Domain ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnforcementPointForInfra**
> PatchEnforcementPointForInfra(ctx, body, deploymentZoneId, enforcementpointId)
Patch a new Enforcement Point under infra

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PATCH /infra/sites/site-1/enforcement-points/enforcementpoint-1. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnforcementPointForInfra_0**
> PatchEnforcementPointForInfra_0(ctx, body, deploymentZoneId, enforcementpointId)
Patch a new Enforcement Point under infra

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it. This is a deprecated API. DeploymentZone has been renamed to Site. Use PATCH /infra/sites/site-1/enforcement-points/enforcementpoint-1. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnforcementPointForSite**
> PatchEnforcementPointForSite(ctx, body, siteId, enforcementpointId)
Patch a new Enforcement Point under Site

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnforcementPointForSite_0**
> PatchEnforcementPointForSite_0(ctx, body, siteId, enforcementpointId)
Patch a new Enforcement Point under Site

If the passed Enforcement Point does not already exist, create a new Enforcement Point. If it already exists, patch it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnforcementPoint**](EnforcementPoint.md)|  | 
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGlobalManagerConfig**
> PatchGlobalManagerConfig(ctx, body)
Create or patch Global Manager Config

Create or patch a Global Manager Config 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManagerConfig**](GlobalManagerConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGlobalManagerConfig_0**
> PatchGlobalManagerConfig_0(ctx, body)
Create or patch Global Manager Config

Create or patch a Global Manager Config 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalManagerConfig**](GlobalManagerConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfra**
> PatchInfra(ctx, body, optional)
Update the infra including all the nested entities

Patch API at infra level can be used in two flavours 1. Like a regular API to update Infra object 2. Hierarchical API: To create/update/delete entire or part of intent    hierarchy Hierarchical API: Provides users a way to create entire or part of intent in single API invocation. Input is expressed in a tree format. Each node in tree can have multiple children of different types. System will resolve the dependecies of nodes within the intent tree and will create the model. Children for any node can be specified using ChildResourceReference or ChildPolicyConfigResource. If a resource is specified using ChildResourceReference then it will not be updated only its children will be updated. If Object is specified using ChildPolicyConfigResource, object along with its children will be updated. Hierarchical API can also be used to delete any sub-branch of entire tree. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 
 **optional** | ***PolicyInfraApiPatchInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiPatchInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enforceRevisionCheck** | **optional.**| Force revision check | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfraSite**
> PatchInfraSite(ctx, body, siteId)
Create or patch Site

Create or patch Site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Site**](Site.md)|  | 
  **siteId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfraSite_0**
> PatchInfraSite_0(ctx, body, siteId)
Create or patch Site

Create or patch Site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Site**](Site.md)|  | 
  **siteId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInfra_0**
> PatchInfra_0(ctx, body, optional)
Update the infra including all the nested entities

Patch API at infra level can be used in two flavours 1. Like a regular API to update Infra object 2. Hierarchical API: To create/update/delete entire or part of intent    hierarchy Hierarchical API: Provides users a way to create entire or part of intent in single API invocation. Input is expressed in a tree format. Each node in tree can have multiple children of different types. System will resolve the dependecies of nodes within the intent tree and will create the model. Children for any node can be specified using ChildResourceReference or ChildPolicyConfigResource. If a resource is specified using ChildResourceReference then it will not be updated only its children will be updated. If Object is specified using ChildPolicyConfigResource, object along with its children will be updated. Hierarchical API can also be used to delete any sub-branch of entire tree. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 
 **optional** | ***PolicyInfraApiPatchInfra_66Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiPatchInfra_66Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enforceRevisionCheck** | **optional.**| Force revision check | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTenantConstraint**
> PatchTenantConstraint(ctx, body, constraintId)
Create or update tenant Constraint

Create tenant constraint if not exists, otherwise update the existing constraint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Constraint**](Constraint.md)|  | 
  **constraintId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTenantConstraint_0**
> PatchTenantConstraint_0(ctx, body, constraintId)
Create or update tenant Constraint

Create tenant constraint if not exists, otherwise update the existing constraint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Constraint**](Constraint.md)|  | 
  **constraintId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTlsCertificate**
> PatchTlsCertificate(ctx, body, certificateId)
Add a New Certificate

Adds a new private-public certificate and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store.  A certificate chain will not be expanded into separate certificate instances for reference, but would be pushed to the enforcement point as a single certificate.  This patch method does not modify an existing certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsTrustData**](TlsTrustData.md)|  | 
  **certificateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTlsCertificate_0**
> PatchTlsCertificate_0(ctx, body, certificateId)
Add a New Certificate

Adds a new private-public certificate and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store.  A certificate chain will not be expanded into separate certificate instances for reference, but would be pushed to the enforcement point as a single certificate.  This patch method does not modify an existing certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TlsTrustData**](TlsTrustData.md)|  | 
  **certificateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDeploymentZoneInfra**
> DeploymentZone ReadDeploymentZoneInfra(ctx, deploymentZoneId)
Read a DeploymentZone

Read a Deployment Zone. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 

### Return type

[**DeploymentZone**](DeploymentZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDeploymentZoneInfra_0**
> DeploymentZone ReadDeploymentZoneInfra_0(ctx, deploymentZoneId)
Read a DeploymentZone

Read a Deployment Zone. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 

### Return type

[**DeploymentZone**](DeploymentZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDomainDeploymentMapForInfra**
> DomainDeploymentMap ReadDomainDeploymentMapForInfra(ctx, domainId, domainDeploymentMapId)
Read a DomainDeploymentMap

Read a Domain Deployment Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

[**DomainDeploymentMap**](DomainDeploymentMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDomainDeploymentMapForInfra_0**
> DomainDeploymentMap ReadDomainDeploymentMapForInfra_0(ctx, domainId, domainDeploymentMapId)
Read a DomainDeploymentMap

Read a Domain Deployment Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **domainDeploymentMapId** | **string**|  | 

### Return type

[**DomainDeploymentMap**](DomainDeploymentMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDomainForInfra**
> Domain ReadDomainForInfra(ctx, domainId)
Read domain

Read a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDomainForInfra_0**
> Domain ReadDomainForInfra_0(ctx, domainId)
Read domain

Read a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEdgeClusterForEnforcementPoint**
> PolicyEdgeCluster ReadEdgeClusterForEnforcementPoint(ctx, siteId, enforcementpointId, edgeClusterId)
Read a Edge Cluster under an Enforcement Point

Read a Edge Cluster under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 

### Return type

[**PolicyEdgeCluster**](PolicyEdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEdgeClusterForEnforcementPoint_0**
> PolicyEdgeCluster ReadEdgeClusterForEnforcementPoint_0(ctx, siteId, enforcementpointId, edgeClusterId)
Read a Edge Cluster under an Enforcement Point

Read a Edge Cluster under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 

### Return type

[**PolicyEdgeCluster**](PolicyEdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEdgeNodeUnderEdgeClusterForEnforcementPoint**
> PolicyEdgeNode ReadEdgeNodeUnderEdgeClusterForEnforcementPoint(ctx, siteId, enforcementpointId, edgeClusterId, edgeNodeId)
Read a Edge Node under an Enforcement Point, Edge Cluster

Read a Edge Node under an Enforcement Point, Edge Cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **edgeNodeId** | **string**|  | 

### Return type

[**PolicyEdgeNode**](PolicyEdgeNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEdgeNodeUnderEdgeClusterForEnforcementPoint_0**
> PolicyEdgeNode ReadEdgeNodeUnderEdgeClusterForEnforcementPoint_0(ctx, siteId, enforcementpointId, edgeClusterId, edgeNodeId)
Read a Edge Node under an Enforcement Point, Edge Cluster

Read a Edge Node under an Enforcement Point, Edge Cluster 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **edgeNodeId** | **string**|  | 

### Return type

[**PolicyEdgeNode**](PolicyEdgeNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointForInfra**
> EnforcementPoint ReadEnforcementPointForInfra(ctx, deploymentZoneId, enforcementpointId)
Read an Enforcement Point

Read an Enforcement Point. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointForInfra_0**
> EnforcementPoint ReadEnforcementPointForInfra_0(ctx, deploymentZoneId, enforcementpointId)
Read an Enforcement Point

Read an Enforcement Point. This is a deprecated API. DeploymentZone has been renamed to Site. Use GET /infra/sites/site-id/enforcement-points/enforcementpoint-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentZoneId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointForSite**
> EnforcementPoint ReadEnforcementPointForSite(ctx, siteId, enforcementpointId)
Read an Enforcement Point under Infra/Site

Read an Enforcement Point under Infra/Site 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointForSite_0**
> EnforcementPoint ReadEnforcementPointForSite_0(ctx, siteId, enforcementpointId)
Read an Enforcement Point under Infra/Site

Read an Enforcement Point under Infra/Site 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointRealizedState**
> RealizedEnforcementPoint ReadEnforcementPointRealizedState(ctx, enforcementPointName)
Read Enforcement Point

Read a Enforcement Point and the complete tree underneath. Returns the populated enforcement point object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 

### Return type

[**RealizedEnforcementPoint**](RealizedEnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEnforcementPointRealizedState_0**
> RealizedEnforcementPoint ReadEnforcementPointRealizedState_0(ctx, enforcementPointName)
Read Enforcement Point

Read a Enforcement Point and the complete tree underneath. Returns the populated enforcement point object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 

### Return type

[**RealizedEnforcementPoint**](RealizedEnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFirewallSectionRealizedState**
> RealizedFirewallSection ReadFirewallSectionRealizedState(ctx, enforcementPointName, firewallSectionId)
Read Firewall

Read a Firewall and the complete tree underneath. Returns the populated Firewall object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **firewallSectionId** | **string**| Firewall Section Id | 

### Return type

[**RealizedFirewallSection**](RealizedFirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFirewallSectionRealizedState_0**
> RealizedFirewallSection ReadFirewallSectionRealizedState_0(ctx, enforcementPointName, firewallSectionId)
Read Firewall

Read a Firewall and the complete tree underneath. Returns the populated Firewall object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **firewallSectionId** | **string**| Firewall Section Id | 

### Return type

[**RealizedFirewallSection**](RealizedFirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData**
> GlobalManagerConfig ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData(ctx, )
Read Global Manager config along with sensitive data

Read a Global Manager config along with sensitive data. For example - rtep_config.ibgp_password 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalManagerConfig**](GlobalManagerConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData_0**
> GlobalManagerConfig ReadGlobalManagerConfigWithSensitiveDataShowSensitiveData_0(ctx, )
Read Global Manager config along with sensitive data

Read a Global Manager config along with sensitive data. For example - rtep_config.ibgp_password 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalManagerConfig**](GlobalManagerConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIPSetRealizedState**
> GenericPolicyRealizedResource ReadIPSetRealizedState(ctx, enforcementPointName, ipSetName)
Read IPSet Realized state

Read an IPSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **ipSetName** | **string**| IPSet name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIPSetRealizedState_0**
> GenericPolicyRealizedResource ReadIPSetRealizedState_0(ctx, enforcementPointName, ipSetName)
Read IPSet Realized state

Read an IPSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **ipSetName** | **string**| IPSet name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInfra**
> Infra ReadInfra(ctx, optional)
Read infra

Read infra. Returns only the infra related properties. Inner object are not populated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiReadInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiReadInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string as java regex | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInfra_0**
> Infra ReadInfra_0(ctx, optional)
Read infra

Read infra. Returns only the infra related properties. Inner object are not populated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyInfraApiReadInfra_80Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiReadInfra_80Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string as java regex | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIntentStatus**
> ConsolidatedRealizedStatus ReadIntentStatus(ctx, intentPath, optional)
Get consolidated status of an intent object

Get Consolidated Status of an intent object (with or without enforcement specific status details). The request is evaluated as follows: - <intent_path>: the request is evaluated on all enforcement points for the given intent without enforcement point specific details. - <intent_path, include_enforced_status>: the request is evaluated on all enforcement points for the given intent with enforcement point specific details. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Policy Path of the intent object | 
 **optional** | ***PolicyInfraApiReadIntentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiReadIntentStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeEnforcedStatus** | **optional.Bool**| Include Enforced Status Flag | [default to false]
 **sitePath** | **optional.String**| Policy Path of the site from where the realization status needs to be fetched | 

### Return type

[**ConsolidatedRealizedStatus**](ConsolidatedRealizedStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIntentStatus_0**
> ConsolidatedRealizedStatus ReadIntentStatus_0(ctx, intentPath, optional)
Get consolidated status of an intent object

Get Consolidated Status of an intent object (with or without enforcement specific status details). The request is evaluated as follows: - <intent_path>: the request is evaluated on all enforcement points for the given intent without enforcement point specific details. - <intent_path, include_enforced_status>: the request is evaluated on all enforcement points for the given intent with enforcement point specific details. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Policy Path of the intent object | 
 **optional** | ***PolicyInfraApiReadIntentStatus_81Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiReadIntentStatus_81Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeEnforcedStatus** | **optional.Bool**| Include Enforced Status Flag | [default to false]
 **sitePath** | **optional.String**| Policy Path of the site from where the realization status needs to be fetched | 

### Return type

[**ConsolidatedRealizedStatus**](ConsolidatedRealizedStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMACSetRealizedState**
> GenericPolicyRealizedResource ReadMACSetRealizedState(ctx, enforcementPointName, macSetName)
Read MACSet Realized state

Read an MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **macSetName** | **string**| MACSet name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMACSetRealizedState_0**
> GenericPolicyRealizedResource ReadMACSetRealizedState_0(ctx, enforcementPointName, macSetName)
Read MACSet Realized state

Read an MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **macSetName** | **string**| MACSet name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSGroupRealizedState**
> GenericPolicyRealizedResource ReadNSGroupRealizedState(ctx, enforcementPointName, nsgroupName)
Read Group

Read a NSGroup and the complete tree underneath. Returns the populated NSgroup object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **nsgroupName** | **string**| Group Name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSGroupRealizedState_0**
> GenericPolicyRealizedResource ReadNSGroupRealizedState_0(ctx, enforcementPointName, nsgroupName)
Read Group

Read a NSGroup and the complete tree underneath. Returns the populated NSgroup object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **nsgroupName** | **string**| Group Name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSServiceRealizedState**
> GenericPolicyRealizedResource ReadNSServiceRealizedState(ctx, enforcementPointName, nsserviceName)
Read NSService

Read a NSService. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **nsserviceName** | **string**| NSService Name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSServiceRealizedState_0**
> GenericPolicyRealizedResource ReadNSServiceRealizedState_0(ctx, enforcementPointName, nsserviceName)
Read NSService

Read a NSService. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **nsserviceName** | **string**| NSService Name | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyLabelForInfra**
> PolicyLabel ReadPolicyLabelForInfra(ctx, labelId)
Read lable

Read a label. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelId** | **string**|  | 

### Return type

[**PolicyLabel**](PolicyLabel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyLabelForInfra_0**
> PolicyLabel ReadPolicyLabelForInfra_0(ctx, labelId)
Read lable

Read a label. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelId** | **string**|  | 

### Return type

[**PolicyLabel**](PolicyLabel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRealizedEntity**
> GenericPolicyRealizedResource ReadRealizedEntity(ctx, realizedPath)
Get realized entity uniquely identified by realized path

Get realized entity uniquely identified by realized path, specified by query parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **realizedPath** | **string**| String Path of the realized object | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRealizedEntity_0**
> GenericPolicyRealizedResource ReadRealizedEntity_0(ctx, realizedPath)
Get realized entity uniquely identified by realized path

Get realized entity uniquely identified by realized path, specified by query parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **realizedPath** | **string**| String Path of the realized object | 

### Return type

[**GenericPolicyRealizedResource**](GenericPolicyRealizedResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityGroupRealizedState**
> RealizedSecurityGroup ReadSecurityGroupRealizedState(ctx, enforcementPointName, securitygroupName)
Read Group

Read a Security Group and the complete tree underneath. Returns the populated Security Group object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **securitygroupName** | **string**| Group Name | 

### Return type

[**RealizedSecurityGroup**](RealizedSecurityGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityGroupRealizedState_0**
> RealizedSecurityGroup ReadSecurityGroupRealizedState_0(ctx, enforcementPointName, securitygroupName)
Read Group

Read a Security Group and the complete tree underneath. Returns the populated Security Group object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**| Enforcement Point Name | 
  **securitygroupName** | **string**| Group Name | 

### Return type

[**RealizedSecurityGroup**](RealizedSecurityGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSite**
> Site ReadSite(ctx, siteId)
Read a site

Read a site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 

### Return type

[**Site**](Site.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSite_0**
> Site ReadSite_0(ctx, siteId)
Read a site

Read a site under Infra. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 

### Return type

[**Site**](Site.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTenantConstraint**
> Constraint ReadTenantConstraint(ctx, constraintId)
Read tenant Constraint.

Read tenant constraint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **constraintId** | **string**|  | 

### Return type

[**Constraint**](Constraint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTenantConstraint_0**
> Constraint ReadTenantConstraint_0(ctx, constraintId)
Read tenant Constraint.

Read tenant constraint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **constraintId** | **string**|  | 

### Return type

[**Constraint**](Constraint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTransportZoneForEnforcementPoint**
> PolicyTransportZone ReadTransportZoneForEnforcementPoint(ctx, siteId, enforcementpointId, transportZoneId)
Read a Transport Zone under an Enforcement Point

Read a Transport Zone under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **transportZoneId** | **string**|  | 

### Return type

[**PolicyTransportZone**](PolicyTransportZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTransportZoneForEnforcementPoint_0**
> PolicyTransportZone ReadTransportZoneForEnforcementPoint_0(ctx, siteId, enforcementpointId, transportZoneId)
Read a Transport Zone under an Enforcement Point

Read a Transport Zone under an Enforcement Point 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 
  **transportZoneId** | **string**|  | 

### Return type

[**PolicyTransportZone**](PolicyTransportZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVirtualMachineDetails**
> VirtualMachineDetails ReadVirtualMachineDetails(ctx, enforcementPointName, virtualMachineId)
Read the details of a virtual machine on the NSX Manager

This API return optional details about a virtual machines (e.g. user login session) from the specified enforcement point. In case of NSXT, virtual-machine-id would be the value of the external_id of the virtual machine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
  **virtualMachineId** | **string**|  | 

### Return type

[**VirtualMachineDetails**](VirtualMachineDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVirtualMachineDetails_0**
> VirtualMachineDetails ReadVirtualMachineDetails_0(ctx, enforcementPointName, virtualMachineId)
Read the details of a virtual machine on the NSX Manager

This API return optional details about a virtual machines (e.g. user login session) from the specified enforcement point. In case of NSXT, virtual-machine-id would be the value of the external_id of the virtual machine. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointName** | **string**|  | 
  **virtualMachineId** | **string**|  | 

### Return type

[**VirtualMachineDetails**](VirtualMachineDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshRealizedStateRefresh**
> RefreshRealizedStateRefresh(ctx, intentPath, optional)
Refresh all realized entities associated with the intent-path

Refresh the status and statistics of all realized entities associated with given intent path synchronously. The vmw-async: True HTTP header cannot be used with this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***PolicyInfraApiRefreshRealizedStateRefreshOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiRefreshRealizedStateRefreshOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshRealizedStateRefresh_0**
> RefreshRealizedStateRefresh_0(ctx, intentPath, optional)
Refresh all realized entities associated with the intent-path

Refresh the status and statistics of all realized entities associated with given intent path synchronously. The vmw-async: True HTTP header cannot be used with this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| String Path of the intent object | 
 **optional** | ***PolicyInfraApiRefreshRealizedStateRefresh_92Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyInfraApiRefreshRealizedStateRefresh_92Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadEnforcementPointForSiteReload**
> EnforcementPoint ReloadEnforcementPointForSiteReload(ctx, siteId, enforcementpointId)
Reload an Enforcement Point under Site

Reload an Enforcement Point under Site. This will read and update fabric configs from enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadEnforcementPointForSiteReload_0**
> EnforcementPoint ReloadEnforcementPointForSiteReload_0(ctx, siteId, enforcementpointId)
Reload an Enforcement Point under Site

Reload an Enforcement Point under Site. This will read and update fabric configs from enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementpointId** | **string**|  | 

### Return type

[**EnforcementPoint**](EnforcementPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagBulkUpdate**
> TagBulkOperation TagBulkUpdate(ctx, body, operationId)
Assign or Unassign tag on multiple Virtual Machines.

Tag can be assigned or unassigned on multiple objects. Supported object type is  restricted to Virtual Machine for now and support for other objects will be added  later. Permissions for tag bulk operation would be similar to virtual machine tag permissions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagBulkOperation**](TagBulkOperation.md)|  | 
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperation**](TagBulkOperation.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagBulkUpdate_0**
> TagBulkOperation TagBulkUpdate_0(ctx, body, operationId)
Assign or Unassign tag on multiple Virtual Machines.

Tag can be assigned or unassigned on multiple objects. Supported object type is  restricted to Virtual Machine for now and support for other objects will be added  later. Permissions for tag bulk operation would be similar to virtual machine tag permissions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagBulkOperation**](TagBulkOperation.md)|  | 
  **operationId** | **string**|  | 

### Return type

[**TagBulkOperation**](TagBulkOperation.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagVirtualMachineUpdateTags**
> TagVirtualMachineUpdateTags(ctx, body, enforcementPointName)
Apply tags on virtual machine

Allows an admin to apply multiple tags to a virtual machine. This operation does not store the intent on the policy side. It applies the tag directly on the specified enforcement point. This operation will replace the existing tags on the virtual machine with the ones that have been passed. If the application of tag fails on the enforcement point, then an error is reported. The admin will have to retry the operation again. Policy framework does not perform a retry. Failure could occur due to multiple reasons. For e.g enforcement point is down, Enforcement point could not apply the tag due to constraints like max tags limit exceeded, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualMachineTagsUpdate**](VirtualMachineTagsUpdate.md)|  | 
  **enforcementPointName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagVirtualMachineUpdateTags_0**
> TagVirtualMachineUpdateTags_0(ctx, body, enforcementPointName)
Apply tags on virtual machine

Allows an admin to apply multiple tags to a virtual machine. This operation does not store the intent on the policy side. It applies the tag directly on the specified enforcement point. This operation will replace the existing tags on the virtual machine with the ones that have been passed. If the application of tag fails on the enforcement point, then an error is reported. The admin will have to retry the operation again. Policy framework does not perform a retry. Failure could occur due to multiple reasons. For e.g enforcement point is down, Enforcement point could not apply the tag due to constraints like max tags limit exceeded, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualMachineTagsUpdate**](VirtualMachineTagsUpdate.md)|  | 
  **enforcementPointName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDomainForInfra**
> Domain UpdateDomainForInfra(ctx, body, domainId)
Create or update a domain

If a domain with the domain-id is not already present, create a new domain. If it already exists, update the domain including the nested groups. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Domain**](Domain.md)|  | 
  **domainId** | **string**| Domain ID | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDomainForInfra_0**
> Domain UpdateDomainForInfra_0(ctx, body, domainId)
Create or update a domain

If a domain with the domain-id is not already present, create a new domain. If it already exists, update the domain including the nested groups. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Domain**](Domain.md)|  | 
  **domainId** | **string**| Domain ID | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInfra**
> Infra UpdateInfra(ctx, body)
Update the infra including all the nested entities

Update the infra including all the nested entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInfra_0**
> Infra UpdateInfra_0(ctx, body)
Update the infra including all the nested entities

Update the infra including all the nested entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyLabelForInfra**
> UpdatePolicyLabelForInfra(ctx, body, labelId)
Patch an existing label object

Create label if not exists, otherwise take the partial updates. Note, once the label is created type attribute can not be changed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyLabel**](PolicyLabel.md)|  | 
  **labelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyLabelForInfra_0**
> UpdatePolicyLabelForInfra_0(ctx, body, labelId)
Patch an existing label object

Create label if not exists, otherwise take the partial updates. Note, once the label is created type attribute can not be changed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyLabel**](PolicyLabel.md)|  | 
  **labelId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

