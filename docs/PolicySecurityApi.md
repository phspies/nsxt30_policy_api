# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/policy/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateByodPolicyServiceInstance**](PolicySecurityApi.md#CreateByodPolicyServiceInstance) | **Put** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Create service instance
[**CreateByodPolicyServiceInstance_0**](PolicySecurityApi.md#CreateByodPolicyServiceInstance_0) | **Put** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Create service instance
[**CreateOrReplaceGatewayPolicyForDomain**](PolicySecurityApi.md#CreateOrReplaceGatewayPolicyForDomain) | **Put** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Update gateway policy
[**CreateOrReplaceGatewayPolicyForDomain_0**](PolicySecurityApi.md#CreateOrReplaceGatewayPolicyForDomain_0) | **Put** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Update gateway policy
[**CreateOrReplaceGatewayRule**](PolicySecurityApi.md#CreateOrReplaceGatewayRule) | **Put** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Update gateway rule
[**CreateOrReplaceGatewayRule_0**](PolicySecurityApi.md#CreateOrReplaceGatewayRule_0) | **Put** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Update gateway rule
[**CreateOrUpdateEndpointPolicy**](PolicySecurityApi.md#CreateOrUpdateEndpointPolicy) | **Put** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Create or update Endpoint policy
[**CreateOrUpdateEndpointPolicy_0**](PolicySecurityApi.md#CreateOrUpdateEndpointPolicy_0) | **Put** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Create or update Endpoint policy
[**CreateOrUpdateEndpointRule**](PolicySecurityApi.md#CreateOrUpdateEndpointRule) | **Put** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Update Endpoint rule
[**CreateOrUpdateEndpointRule_0**](PolicySecurityApi.md#CreateOrUpdateEndpointRule_0) | **Put** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Update Endpoint rule
[**CreateOrUpdateIdsClusterConfig**](PolicySecurityApi.md#CreateOrUpdateIdsClusterConfig) | **Put** /infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | create or update IDS config on cluster level
[**CreateOrUpdateIdsClusterConfig_0**](PolicySecurityApi.md#CreateOrUpdateIdsClusterConfig_0) | **Put** /global-infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | create or update IDS config on cluster level
[**CreateOrUpdateIdsProfile**](PolicySecurityApi.md#CreateOrUpdateIdsProfile) | **Put** /global-infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | create or update IDS profile
[**CreateOrUpdateIdsProfile_0**](PolicySecurityApi.md#CreateOrUpdateIdsProfile_0) | **Put** /infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | create or update IDS profile
[**CreateOrUpdateIdsRule**](PolicySecurityApi.md#CreateOrUpdateIdsRule) | **Put** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | create or update IDS rule
[**CreateOrUpdateIdsRule_0**](PolicySecurityApi.md#CreateOrUpdateIdsRule_0) | **Put** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | create or update IDS rule
[**CreateOrUpdateIdsSecurityPolicy**](PolicySecurityApi.md#CreateOrUpdateIdsSecurityPolicy) | **Put** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | create or update IDS security policy
[**CreateOrUpdateIdsSecurityPolicy_0**](PolicySecurityApi.md#CreateOrUpdateIdsSecurityPolicy_0) | **Put** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | create or update IDS security policy
[**CreateOrUpdateIdsStandaloneHostConfig**](PolicySecurityApi.md#CreateOrUpdateIdsStandaloneHostConfig) | **Put** /infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Create or update IDS configuration
[**CreateOrUpdateIdsStandaloneHostConfig_0**](PolicySecurityApi.md#CreateOrUpdateIdsStandaloneHostConfig_0) | **Put** /global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Create or update IDS configuration
[**CreateOrUpdateRedirectionPolicy**](PolicySecurityApi.md#CreateOrUpdateRedirectionPolicy) | **Put** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Create or update redirection policy
[**CreateOrUpdateRedirectionPolicy_0**](PolicySecurityApi.md#CreateOrUpdateRedirectionPolicy_0) | **Put** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Create or update redirection policy
[**CreateOrUpdateRedirectionRule**](PolicySecurityApi.md#CreateOrUpdateRedirectionRule) | **Put** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Update redirection rule
[**CreateOrUpdateRedirectionRule_0**](PolicySecurityApi.md#CreateOrUpdateRedirectionRule_0) | **Put** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Update redirection rule
[**CreateOrUpdateServiceReference**](PolicySecurityApi.md#CreateOrUpdateServiceReference) | **Put** /global-infra/service-references/{service-reference-id} | Create service reference
[**CreateOrUpdateServiceReference_0**](PolicySecurityApi.md#CreateOrUpdateServiceReference_0) | **Put** /infra/service-references/{service-reference-id} | Create service reference
[**CreateOrUpdateVirtualEndpoint**](PolicySecurityApi.md#CreateOrUpdateVirtualEndpoint) | **Put** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Create or update virtual endpoint
[**CreateOrUpdateVirtualEndpoint_0**](PolicySecurityApi.md#CreateOrUpdateVirtualEndpoint_0) | **Put** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Create or update virtual endpoint
[**CreatePolicyServiceInstance**](PolicySecurityApi.md#CreatePolicyServiceInstance) | **Put** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create service instance
[**CreatePolicyServiceInstance_0**](PolicySecurityApi.md#CreatePolicyServiceInstance_0) | **Put** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create service instance
[**CreatePolicyServiceProfile**](PolicySecurityApi.md#CreatePolicyServiceProfile) | **Put** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Create or update service profile
[**CreatePolicyServiceProfile_0**](PolicySecurityApi.md#CreatePolicyServiceProfile_0) | **Put** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Create or update service profile
[**CreateServiceChain**](PolicySecurityApi.md#CreateServiceChain) | **Put** /infra/service-chains/{service-chain-id} | Create  or update service chain
[**CreateServiceChain_0**](PolicySecurityApi.md#CreateServiceChain_0) | **Put** /global-infra/service-chains/{service-chain-id} | Create  or update service chain
[**CreateServiceDefinition**](PolicySecurityApi.md#CreateServiceDefinition) | **Post** /enforcement-points/{enforcement-point-id}/service-definitions | Create a Service Definition on given enforcement point.
[**CreateServiceInstanceEndpoint**](PolicySecurityApi.md#CreateServiceInstanceEndpoint) | **Put** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Create service instance endpoint
[**CreateServiceInstanceEndpoint_0**](PolicySecurityApi.md#CreateServiceInstanceEndpoint_0) | **Put** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Create service instance endpoint
[**CreateTier1PolicyServiceInstance**](PolicySecurityApi.md#CreateTier1PolicyServiceInstance) | **Put** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create Tier1 service instance
[**CreateTier1PolicyServiceInstance_0**](PolicySecurityApi.md#CreateTier1PolicyServiceInstance_0) | **Put** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create Tier1 service instance
[**DeleteByodPolicyServiceInstance**](PolicySecurityApi.md#DeleteByodPolicyServiceInstance) | **Delete** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Delete BYOD policy service instance
[**DeleteByodPolicyServiceInstance_0**](PolicySecurityApi.md#DeleteByodPolicyServiceInstance_0) | **Delete** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Delete BYOD policy service instance
[**DeleteCPUMemThresholdsProfile**](PolicySecurityApi.md#DeleteCPUMemThresholdsProfile) | **Delete** /infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Delete CPU and memory thresholds profile
[**DeleteCPUMemThresholdsProfile_0**](PolicySecurityApi.md#DeleteCPUMemThresholdsProfile_0) | **Delete** /global-infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Delete CPU and memory thresholds profile
[**DeleteCommunicationEntry**](PolicySecurityApi.md#DeleteCommunicationEntry) | **Delete** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Delete CommunicationEntry
[**DeleteCommunicationEntry_0**](PolicySecurityApi.md#DeleteCommunicationEntry_0) | **Delete** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Delete CommunicationEntry
[**DeleteCommunicationMapForDomain**](PolicySecurityApi.md#DeleteCommunicationMapForDomain) | **Delete** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id} | Deletes a communication map from this domain
[**DeleteCommunicationMapForDomain_0**](PolicySecurityApi.md#DeleteCommunicationMapForDomain_0) | **Delete** /infra/domains/{domain-id}/communication-maps/{communication-map-id} | Deletes a communication map from this domain
[**DeleteDnsSecurityProfile**](PolicySecurityApi.md#DeleteDnsSecurityProfile) | **Delete** /infra/dns-security-profiles/{profile-id} | Delete DNS security profile
[**DeleteDnsSecurityProfileBinding**](PolicySecurityApi.md#DeleteDnsSecurityProfileBinding) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Delete DNS security profile binding map
[**DeleteDnsSecurityProfileBinding_0**](PolicySecurityApi.md#DeleteDnsSecurityProfileBinding_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Delete DNS security profile binding map
[**DeleteDnsSecurityProfile_0**](PolicySecurityApi.md#DeleteDnsSecurityProfile_0) | **Delete** /global-infra/dns-security-profiles/{profile-id} | Delete DNS security profile
[**DeleteDraft**](PolicySecurityApi.md#DeleteDraft) | **Delete** /infra/drafts/{draft-id} | Delete a manual draft
[**DeleteDraft_0**](PolicySecurityApi.md#DeleteDraft_0) | **Delete** /global-infra/drafts/{draft-id} | Delete a manual draft
[**DeleteEndpointPolicy**](PolicySecurityApi.md#DeleteEndpointPolicy) | **Delete** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Delete Endpoint policy
[**DeleteEndpointPolicy_0**](PolicySecurityApi.md#DeleteEndpointPolicy_0) | **Delete** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Delete Endpoint policy
[**DeleteEndpointRule**](PolicySecurityApi.md#DeleteEndpointRule) | **Delete** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Delete EndpointRule
[**DeleteEndpointRule_0**](PolicySecurityApi.md#DeleteEndpointRule_0) | **Delete** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Delete EndpointRule
[**DeleteFloodProtectionProfile**](PolicySecurityApi.md#DeleteFloodProtectionProfile) | **Delete** /infra/flood-protection-profiles/{flood-protection-profile-id} | Delete Flood Protection Profile
[**DeleteFloodProtectionProfile_0**](PolicySecurityApi.md#DeleteFloodProtectionProfile_0) | **Delete** /global-infra/flood-protection-profiles/{flood-protection-profile-id} | Delete Flood Protection Profile
[**DeleteGatewayPolicy**](PolicySecurityApi.md#DeleteGatewayPolicy) | **Delete** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Delete GatewayPolicy
[**DeleteGatewayPolicy_0**](PolicySecurityApi.md#DeleteGatewayPolicy_0) | **Delete** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Delete GatewayPolicy
[**DeleteGatewayRule**](PolicySecurityApi.md#DeleteGatewayRule) | **Delete** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Delete rule
[**DeleteGatewayRule_0**](PolicySecurityApi.md#DeleteGatewayRule_0) | **Delete** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Delete rule
[**DeleteGroupMonitoringBinding**](PolicySecurityApi.md#DeleteGroupMonitoringBinding) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Delete Group Monitoring Profile Binding
[**DeleteGroupMonitoringBinding_0**](PolicySecurityApi.md#DeleteGroupMonitoringBinding_0) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Delete Group Monitoring Profile Binding
[**DeleteIdsProfile**](PolicySecurityApi.md#DeleteIdsProfile) | **Delete** /global-infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Delete IDS profile
[**DeleteIdsProfile_0**](PolicySecurityApi.md#DeleteIdsProfile_0) | **Delete** /infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Delete IDS profile
[**DeleteIdsRule**](PolicySecurityApi.md#DeleteIdsRule) | **Delete** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Delete IDS rule
[**DeleteIdsRule_0**](PolicySecurityApi.md#DeleteIdsRule_0) | **Delete** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Delete IDS rule
[**DeleteIdsSecurityPolicy**](PolicySecurityApi.md#DeleteIdsSecurityPolicy) | **Delete** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Delete IDS security policy
[**DeleteIdsSecurityPolicy_0**](PolicySecurityApi.md#DeleteIdsSecurityPolicy_0) | **Delete** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Delete IDS security policy
[**DeletePolicyFirewallCPUMemThresholdsProfileBindingMap**](PolicySecurityApi.md#DeletePolicyFirewallCPUMemThresholdsProfileBindingMap) | **Delete** /infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Delete Firewall CPU Memory Thresholds Profile Binding
[**DeletePolicyFirewallCPUMemThresholdsProfileBindingMap_0**](PolicySecurityApi.md#DeletePolicyFirewallCPUMemThresholdsProfileBindingMap_0) | **Delete** /global-infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Delete Firewall CPU Memory Thresholds Profile Binding
[**DeletePolicyFirewallFloodProtectionBinding**](PolicySecurityApi.md#DeletePolicyFirewallFloodProtectionBinding) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Delete Firewall Flood Protection Profile Binding
[**DeletePolicyFirewallFloodProtectionBinding_0**](PolicySecurityApi.md#DeletePolicyFirewallFloodProtectionBinding_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Delete Firewall Flood Protection Profile Binding
[**DeletePolicyFirewallScheduler**](PolicySecurityApi.md#DeletePolicyFirewallScheduler) | **Delete** /global-infra/firewall-schedulers/{firewall-scheduler-id} | Delete Policy Firewall Scheduler
[**DeletePolicyFirewallScheduler_0**](PolicySecurityApi.md#DeletePolicyFirewallScheduler_0) | **Delete** /infra/firewall-schedulers/{firewall-scheduler-id} | Delete Policy Firewall Scheduler
[**DeletePolicyFirewallSessionTimerBinding**](PolicySecurityApi.md#DeletePolicyFirewallSessionTimerBinding) | **Delete** /infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Delete Firewall Session Timer Profile Binding
[**DeletePolicyFirewallSessionTimerBinding_0**](PolicySecurityApi.md#DeletePolicyFirewallSessionTimerBinding_0) | **Delete** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Delete Firewall Session Timer Profile Binding
[**DeletePolicyFirewallSessionTimerProfile**](PolicySecurityApi.md#DeletePolicyFirewallSessionTimerProfile) | **Delete** /global-infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Delete Firewall Session Timer Profile
[**DeletePolicyFirewallSessionTimerProfile_0**](PolicySecurityApi.md#DeletePolicyFirewallSessionTimerProfile_0) | **Delete** /infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Delete Firewall Session Timer Profile
[**DeletePolicyServiceChain**](PolicySecurityApi.md#DeletePolicyServiceChain) | **Delete** /infra/service-chains/{service-chain-id} | Delete Service chain
[**DeletePolicyServiceChain_0**](PolicySecurityApi.md#DeletePolicyServiceChain_0) | **Delete** /global-infra/service-chains/{service-chain-id} | Delete Service chain
[**DeletePolicyServiceInstance**](PolicySecurityApi.md#DeletePolicyServiceInstance) | **Delete** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Delete policy service instance
[**DeletePolicyServiceInstanceEndpoint**](PolicySecurityApi.md#DeletePolicyServiceInstanceEndpoint) | **Delete** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Delete policy service instance endpoint
[**DeletePolicyServiceInstanceEndpoint_0**](PolicySecurityApi.md#DeletePolicyServiceInstanceEndpoint_0) | **Delete** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Delete policy service instance endpoint
[**DeletePolicyServiceInstance_0**](PolicySecurityApi.md#DeletePolicyServiceInstance_0) | **Delete** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Delete policy service instance
[**DeletePolicyServiceProfile**](PolicySecurityApi.md#DeletePolicyServiceProfile) | **Delete** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Delete Service profile
[**DeletePolicyServiceProfile_0**](PolicySecurityApi.md#DeletePolicyServiceProfile_0) | **Delete** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Delete Service profile
[**DeletePolicyUrlCategorizationConfig**](PolicySecurityApi.md#DeletePolicyUrlCategorizationConfig) | **Delete** /global-infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Delete PolicyUrlCategorizationConfig
[**DeletePolicyUrlCategorizationConfig_0**](PolicySecurityApi.md#DeletePolicyUrlCategorizationConfig_0) | **Delete** /infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Delete PolicyUrlCategorizationConfig
[**DeleteRedirectionPolicy**](PolicySecurityApi.md#DeleteRedirectionPolicy) | **Delete** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Delete redirection policy
[**DeleteRedirectionPolicy_0**](PolicySecurityApi.md#DeleteRedirectionPolicy_0) | **Delete** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Delete redirection policy
[**DeleteRedirectionRule**](PolicySecurityApi.md#DeleteRedirectionRule) | **Delete** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Delete RedirectionRule
[**DeleteRedirectionRule_0**](PolicySecurityApi.md#DeleteRedirectionRule_0) | **Delete** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Delete RedirectionRule
[**DeleteSecurityPolicyForDomain**](PolicySecurityApi.md#DeleteSecurityPolicyForDomain) | **Delete** /infra/domains/{domain-id}/security-policies/{security-policy-id} | Deletes a security policy from this domain
[**DeleteSecurityPolicyForDomain_0**](PolicySecurityApi.md#DeleteSecurityPolicyForDomain_0) | **Delete** /global-infra/domains/{domain-id}/security-policies/{security-policy-id} | Deletes a security policy from this domain
[**DeleteSecurityRule**](PolicySecurityApi.md#DeleteSecurityRule) | **Delete** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Delete rule
[**DeleteSecurityRule_0**](PolicySecurityApi.md#DeleteSecurityRule_0) | **Delete** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Delete rule
[**DeleteServiceDefinition**](PolicySecurityApi.md#DeleteServiceDefinition) | **Delete** /enforcement-points/{enforcement-point-id}/service-definitions/{service-definition-id} | Delete an existing Service Definition on the given enforcement point 
[**DeleteServiceReference**](PolicySecurityApi.md#DeleteServiceReference) | **Delete** /global-infra/service-references/{service-reference-id} | Delete Service Reference
[**DeleteServiceReference_0**](PolicySecurityApi.md#DeleteServiceReference_0) | **Delete** /infra/service-references/{service-reference-id} | Delete Service Reference
[**DeleteTier0FloodProtectionProfileBinding**](PolicySecurityApi.md#DeleteTier0FloodProtectionProfileBinding) | **Delete** /infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-0 Logical Router
[**DeleteTier0FloodProtectionProfileBinding_0**](PolicySecurityApi.md#DeleteTier0FloodProtectionProfileBinding_0) | **Delete** /global-infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-0 Logical Router
[**DeleteTier0LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#DeleteTier0LocaleServicesFloodProtectionProfileBinding) | **Delete** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices
[**DeleteTier0LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#DeleteTier0LocaleServicesFloodProtectionProfileBinding_0) | **Delete** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices
[**DeleteTier0LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#DeleteTier0LocaleServicesSessionTimerProfileBinding) | **Delete** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices
[**DeleteTier0LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#DeleteTier0LocaleServicesSessionTimerProfileBinding_0) | **Delete** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices
[**DeleteTier0SessionTimerProfileBinding**](PolicySecurityApi.md#DeleteTier0SessionTimerProfileBinding) | **Delete** /global-infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-0 Logical Router
[**DeleteTier0SessionTimerProfileBinding_0**](PolicySecurityApi.md#DeleteTier0SessionTimerProfileBinding_0) | **Delete** /infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-0 Logical Router
[**DeleteTier1FloodProtectionProfileBinding**](PolicySecurityApi.md#DeleteTier1FloodProtectionProfileBinding) | **Delete** /infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-1 Logical Router
[**DeleteTier1FloodProtectionProfileBinding_0**](PolicySecurityApi.md#DeleteTier1FloodProtectionProfileBinding_0) | **Delete** /global-infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-1 Logical Router
[**DeleteTier1LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#DeleteTier1LocaleServicesFloodProtectionProfileBinding) | **Delete** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices
[**DeleteTier1LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#DeleteTier1LocaleServicesFloodProtectionProfileBinding_0) | **Delete** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices
[**DeleteTier1LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#DeleteTier1LocaleServicesSessionTimerProfileBinding) | **Delete** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices
[**DeleteTier1LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#DeleteTier1LocaleServicesSessionTimerProfileBinding_0) | **Delete** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices
[**DeleteTier1PolicyServiceInstance**](PolicySecurityApi.md#DeleteTier1PolicyServiceInstance) | **Delete** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Delete Tier1 policy service instance
[**DeleteTier1PolicyServiceInstance_0**](PolicySecurityApi.md#DeleteTier1PolicyServiceInstance_0) | **Delete** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Delete Tier1 policy service instance
[**DeleteTier1SessionTimerProfileBinding**](PolicySecurityApi.md#DeleteTier1SessionTimerProfileBinding) | **Delete** /global-infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-1 Logical Router
[**DeleteTier1SessionTimerProfileBinding_0**](PolicySecurityApi.md#DeleteTier1SessionTimerProfileBinding_0) | **Delete** /infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Delete Session Timer Profile Binding for Tier-1 Logical Router
[**DeleteVirtualEndpoint**](PolicySecurityApi.md#DeleteVirtualEndpoint) | **Delete** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Delete virtual endpoint
[**DeleteVirtualEndpoint_0**](PolicySecurityApi.md#DeleteVirtualEndpoint_0) | **Delete** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Delete virtual endpoint
[**FilterFirewallExcludeListFilter**](PolicySecurityApi.md#FilterFirewallExcludeListFilter) | **Post** /global-infra/settings/firewall/security/exclude-list?action&#x3D;filter | Filter the firewall exclude list
[**FilterFirewallExcludeListFilter_0**](PolicySecurityApi.md#FilterFirewallExcludeListFilter_0) | **Post** /infra/settings/firewall/security/exclude-list?action&#x3D;filter | Filter the firewall exclude list
[**GetAggregatedConfigurationToBePublishedForDraft**](PolicySecurityApi.md#GetAggregatedConfigurationToBePublishedForDraft) | **Get** /infra/drafts/{draft-id}/aggregated | Get an aggregated configuration for the draft
[**GetAggregatedConfigurationToBePublishedForDraft_0**](PolicySecurityApi.md#GetAggregatedConfigurationToBePublishedForDraft_0) | **Get** /global-infra/drafts/{draft-id}/aggregated | Get an aggregated configuration for the draft
[**GetComputeClusterIdfwConfiguration**](PolicySecurityApi.md#GetComputeClusterIdfwConfiguration) | **Get** /infra/settings/firewall/idfw/cluster/{cluster-id} | Read compute cluster idfw configuration
[**GetComputeClusterIdfwConfiguration_0**](PolicySecurityApi.md#GetComputeClusterIdfwConfiguration_0) | **Get** /global-infra/settings/firewall/idfw/cluster/{cluster-id} | Read compute cluster idfw configuration
[**GetDfwFirewallConfiguration**](PolicySecurityApi.md#GetDfwFirewallConfiguration) | **Get** /infra/settings/firewall/security | Get dfw firewall configuration
[**GetDfwFirewallConfiguration_0**](PolicySecurityApi.md#GetDfwFirewallConfiguration_0) | **Get** /global-infra/settings/firewall/security | Get dfw firewall configuration
[**GetDnsSecurityProfileBinding**](PolicySecurityApi.md#GetDnsSecurityProfileBinding) | **Get** /infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Get DNS security profile binding map
[**GetDnsSecurityProfileBinding_0**](PolicySecurityApi.md#GetDnsSecurityProfileBinding_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Get DNS security profile binding map
[**GetFirewallExcludeList**](PolicySecurityApi.md#GetFirewallExcludeList) | **Get** /infra/settings/firewall/security/exclude-list | Read security policy exclude list
[**GetFirewallExcludeList_0**](PolicySecurityApi.md#GetFirewallExcludeList_0) | **Get** /global-infra/settings/firewall/security/exclude-list | Read security policy exclude list
[**GetFloodProtectionProfile**](PolicySecurityApi.md#GetFloodProtectionProfile) | **Get** /infra/flood-protection-profiles/{flood-protection-profile-id} | Get Flood Protection Profile
[**GetFloodProtectionProfile_0**](PolicySecurityApi.md#GetFloodProtectionProfile_0) | **Get** /global-infra/flood-protection-profiles/{flood-protection-profile-id} | Get Flood Protection Profile
[**GetGatewayPolicyStatistics**](PolicySecurityApi.md#GetGatewayPolicyStatistics) | **Get** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/statistics | Get gateway policy statistics
[**GetGatewayPolicyStatistics_0**](PolicySecurityApi.md#GetGatewayPolicyStatistics_0) | **Get** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/statistics | Get gateway policy statistics
[**GetGatewayRuleStatistics**](PolicySecurityApi.md#GetGatewayRuleStatistics) | **Get** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id}/statistics | Get gateway rule statistics
[**GetGatewayRuleStatistics_0**](PolicySecurityApi.md#GetGatewayRuleStatistics_0) | **Get** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id}/statistics | Get gateway rule statistics
[**GetGroupMonitoringBinding**](PolicySecurityApi.md#GetGroupMonitoringBinding) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Get Group Monitoring Profile Binding Map
[**GetGroupMonitoringBinding_0**](PolicySecurityApi.md#GetGroupMonitoringBinding_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Get Group Monitoring Profile Binding Map
[**GetIdsClusterConfig**](PolicySecurityApi.md#GetIdsClusterConfig) | **Get** /infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | Read IDS cluster config.
[**GetIdsClusterConfig_0**](PolicySecurityApi.md#GetIdsClusterConfig_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | Read IDS cluster config.
[**GetIdsProfile**](PolicySecurityApi.md#GetIdsProfile) | **Get** /global-infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Get IDS profile.
[**GetIdsProfile_0**](PolicySecurityApi.md#GetIdsProfile_0) | **Get** /infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Get IDS profile.
[**GetIdsRule**](PolicySecurityApi.md#GetIdsRule) | **Get** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Get IDS rule.
[**GetIdsRule_0**](PolicySecurityApi.md#GetIdsRule_0) | **Get** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Get IDS rule.
[**GetIdsSecurityPolicy**](PolicySecurityApi.md#GetIdsSecurityPolicy) | **Get** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Get IDS security policy.
[**GetIdsSecurityPolicy_0**](PolicySecurityApi.md#GetIdsSecurityPolicy_0) | **Get** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Get IDS security policy.
[**GetIdsSettings**](PolicySecurityApi.md#GetIdsSettings) | **Get** /infra/settings/firewall/security/intrusion-services | Get IDS system settings
[**GetIdsSettings_0**](PolicySecurityApi.md#GetIdsSettings_0) | **Get** /global-infra/settings/firewall/security/intrusion-services | Get IDS system settings
[**GetIdsSignatureStatus**](PolicySecurityApi.md#GetIdsSignatureStatus) | **Get** /infra/settings/firewall/security/intrusion-services/signatures/status | Get IDS signature status
[**GetIdsSignatureStatus_0**](PolicySecurityApi.md#GetIdsSignatureStatus_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/signatures/status | Get IDS signature status
[**GetIdsSignatureVersions**](PolicySecurityApi.md#GetIdsSignatureVersions) | **Get** /global-infra/settings/firewall/security/intrusion-services/signature-versions | Get IDS signature versions
[**GetIdsSignatureVersions_0**](PolicySecurityApi.md#GetIdsSignatureVersions_0) | **Get** /infra/settings/firewall/security/intrusion-services/signature-versions | Get IDS signature versions
[**GetIdsStandaloneHostConfig**](PolicySecurityApi.md#GetIdsStandaloneHostConfig) | **Get** /infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Read IDS config
[**GetIdsStandaloneHostConfig_0**](PolicySecurityApi.md#GetIdsStandaloneHostConfig_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Read IDS config
[**GetPolicyFirewallCPUMemThresholdsProfileBindingMap**](PolicySecurityApi.md#GetPolicyFirewallCPUMemThresholdsProfileBindingMap) | **Get** /infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Get Firewall CPU Memory Thresholds Profile Binding Map
[**GetPolicyFirewallCPUMemThresholdsProfileBindingMap_0**](PolicySecurityApi.md#GetPolicyFirewallCPUMemThresholdsProfileBindingMap_0) | **Get** /global-infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Get Firewall CPU Memory Thresholds Profile Binding Map
[**GetPolicyFirewallFloodProtectionBinding**](PolicySecurityApi.md#GetPolicyFirewallFloodProtectionBinding) | **Get** /infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Get Firewall Flood Protection Profile Binding Map
[**GetPolicyFirewallFloodProtectionBinding_0**](PolicySecurityApi.md#GetPolicyFirewallFloodProtectionBinding_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Get Firewall Flood Protection Profile Binding Map
[**GetPolicyFirewallScheduler**](PolicySecurityApi.md#GetPolicyFirewallScheduler) | **Get** /global-infra/firewall-schedulers/{firewall-scheduler-id} | Get PolicyFirewallScheduler
[**GetPolicyFirewallScheduler_0**](PolicySecurityApi.md#GetPolicyFirewallScheduler_0) | **Get** /infra/firewall-schedulers/{firewall-scheduler-id} | Get PolicyFirewallScheduler
[**GetPolicyFirewallSessionTimerBinding**](PolicySecurityApi.md#GetPolicyFirewallSessionTimerBinding) | **Get** /infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Get Firewall Session Timer Profile Binding Map
[**GetPolicyFirewallSessionTimerBinding_0**](PolicySecurityApi.md#GetPolicyFirewallSessionTimerBinding_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Get Firewall Session Timer Profile Binding Map
[**GetPolicyFirewallSessionTimerProfile**](PolicySecurityApi.md#GetPolicyFirewallSessionTimerProfile) | **Get** /global-infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Get Firewall Session Timer Profile
[**GetPolicyFirewallSessionTimerProfile_0**](PolicySecurityApi.md#GetPolicyFirewallSessionTimerProfile_0) | **Get** /infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Get Firewall Session Timer Profile
[**GetPolicyServiceInstanceStatistics**](PolicySecurityApi.md#GetPolicyServiceInstanceStatistics) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}/statistics | Get statistics for all runtimes associated with this PolicyServiceInstance
[**GetPolicyServiceInstanceStatistics_0**](PolicySecurityApi.md#GetPolicyServiceInstanceStatistics_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}/statistics | Get statistics for all runtimes associated with this PolicyServiceInstance
[**GetPolicyServiceProfileGroups**](PolicySecurityApi.md#GetPolicyServiceProfileGroups) | **Get** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id}/group-associations | Get Groups used in Redirection rules for a given Service Profile.
[**GetPolicyServiceProfileGroups_0**](PolicySecurityApi.md#GetPolicyServiceProfileGroups_0) | **Get** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id}/group-associations | Get Groups used in Redirection rules for a given Service Profile.
[**GetPolicyUrlCategorizationConfig**](PolicySecurityApi.md#GetPolicyUrlCategorizationConfig) | **Get** /global-infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Get PolicyUrlCategorizationConfig
[**GetPolicyUrlCategorizationConfig_0**](PolicySecurityApi.md#GetPolicyUrlCategorizationConfig_0) | **Get** /infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Get PolicyUrlCategorizationConfig
[**GetPreviewOfConfigurationAfterPublishOfDraft**](PolicySecurityApi.md#GetPreviewOfConfigurationAfterPublishOfDraft) | **Get** /infra/drafts/{draft-id}/complete | Get a preview of a configuration after publish of a draft
[**GetPreviewOfConfigurationAfterPublishOfDraft_0**](PolicySecurityApi.md#GetPreviewOfConfigurationAfterPublishOfDraft_0) | **Get** /global-infra/drafts/{draft-id}/complete | Get a preview of a configuration after publish of a draft
[**GetRuleStatistics**](PolicySecurityApi.md#GetRuleStatistics) | **Get** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id}/statistics | Get rule statistics
[**GetRuleStatistics_0**](PolicySecurityApi.md#GetRuleStatistics_0) | **Get** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id}/statistics | Get rule statistics
[**GetSecurityPolicyStatistics**](PolicySecurityApi.md#GetSecurityPolicyStatistics) | **Get** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/statistics | Get security policy statistics
[**GetSecurityPolicyStatistics_0**](PolicySecurityApi.md#GetSecurityPolicyStatistics_0) | **Get** /infra/domains/{domain-id}/security-policies/{security-policy-id}/statistics | Get security policy statistics
[**GetStandaloneHostIdfwConfiguration**](PolicySecurityApi.md#GetStandaloneHostIdfwConfiguration) | **Get** /infra/settings/firewall/idfw/standalone-host-switch-setting | Read idfw configuration for standalone host
[**GetStandaloneHostIdfwConfiguration_0**](PolicySecurityApi.md#GetStandaloneHostIdfwConfiguration_0) | **Get** /global-infra/settings/firewall/idfw/standalone-host-switch-setting | Read idfw configuration for standalone host
[**GetTier0FloodProtectionProfileBinding**](PolicySecurityApi.md#GetTier0FloodProtectionProfileBinding) | **Get** /infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-0 Logical Router
[**GetTier0FloodProtectionProfileBinding_0**](PolicySecurityApi.md#GetTier0FloodProtectionProfileBinding_0) | **Get** /global-infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-0 Logical Router
[**GetTier0LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#GetTier0LocaleServicesFloodProtectionProfileBinding) | **Get** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**GetTier0LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#GetTier0LocaleServicesFloodProtectionProfileBinding_0) | **Get** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**GetTier0LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#GetTier0LocaleServicesSessionTimerProfileBinding) | **Get** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**GetTier0LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#GetTier0LocaleServicesSessionTimerProfileBinding_0) | **Get** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**GetTier0SessionTimerProfileBinding**](PolicySecurityApi.md#GetTier0SessionTimerProfileBinding) | **Get** /global-infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-0 Logical Router
[**GetTier0SessionTimerProfileBinding_0**](PolicySecurityApi.md#GetTier0SessionTimerProfileBinding_0) | **Get** /infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-0 Logical Router
[**GetTier1FloodProtectionProfileBinding**](PolicySecurityApi.md#GetTier1FloodProtectionProfileBinding) | **Get** /infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-1 Logical Router
[**GetTier1FloodProtectionProfileBinding_0**](PolicySecurityApi.md#GetTier1FloodProtectionProfileBinding_0) | **Get** /global-infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-1 Logical Router
[**GetTier1LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#GetTier1LocaleServicesFloodProtectionProfileBinding) | **Get** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**GetTier1LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#GetTier1LocaleServicesFloodProtectionProfileBinding_0) | **Get** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**GetTier1LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#GetTier1LocaleServicesSessionTimerProfileBinding) | **Get** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**GetTier1LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#GetTier1LocaleServicesSessionTimerProfileBinding_0) | **Get** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**GetTier1PolicyServiceInstanceStatistics**](PolicySecurityApi.md#GetTier1PolicyServiceInstanceStatistics) | **Get** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}/statistics | Get statistics for all runtimes associated with this Tier1 PolicyServiceInstance
[**GetTier1PolicyServiceInstanceStatistics_0**](PolicySecurityApi.md#GetTier1PolicyServiceInstanceStatistics_0) | **Get** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}/statistics | Get statistics for all runtimes associated with this Tier1 PolicyServiceInstance
[**GetTier1SessionTimerProfileBinding**](PolicySecurityApi.md#GetTier1SessionTimerProfileBinding) | **Get** /global-infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-1 Logical Router
[**GetTier1SessionTimerProfileBinding_0**](PolicySecurityApi.md#GetTier1SessionTimerProfileBinding_0) | **Get** /infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Get Session Timer Profile Binding Map for Tier-1 Logical Router
[**ListByodPolicyServiceInstancesForTier0**](PolicySecurityApi.md#ListByodPolicyServiceInstancesForTier0) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances | Read all BYOD service instance objects under a tier-0
[**ListByodPolicyServiceInstancesForTier0_0**](PolicySecurityApi.md#ListByodPolicyServiceInstancesForTier0_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances | Read all BYOD service instance objects under a tier-0
[**ListCPUMemThresholdsProfiles**](PolicySecurityApi.md#ListCPUMemThresholdsProfiles) | **Get** /infra/settings/firewall/cpu-mem-thresholds-profiles | List all CPU and memory thresholds profiles
[**ListCPUMemThresholdsProfiles_0**](PolicySecurityApi.md#ListCPUMemThresholdsProfiles_0) | **Get** /global-infra/settings/firewall/cpu-mem-thresholds-profiles | List all CPU and memory thresholds profiles
[**ListCommunicationEntry**](PolicySecurityApi.md#ListCommunicationEntry) | **Get** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries | List CommunicationEntries
[**ListCommunicationEntry_0**](PolicySecurityApi.md#ListCommunicationEntry_0) | **Get** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries | List CommunicationEntries
[**ListCommunicationMapsForDomain**](PolicySecurityApi.md#ListCommunicationMapsForDomain) | **Get** /global-infra/domains/{domain-id}/communication-maps | List communication maps
[**ListCommunicationMapsForDomain_0**](PolicySecurityApi.md#ListCommunicationMapsForDomain_0) | **Get** /infra/domains/{domain-id}/communication-maps | List communication maps
[**ListComputeClusterIdfwConfiguration**](PolicySecurityApi.md#ListComputeClusterIdfwConfiguration) | **Get** /infra/settings/firewall/idfw/cluster | List compute cluster idfw Configuration
[**ListComputeClusterIdfwConfiguration_0**](PolicySecurityApi.md#ListComputeClusterIdfwConfiguration_0) | **Get** /global-infra/settings/firewall/idfw/cluster | List compute cluster idfw Configuration
[**ListDnsSecurityProfileBindings**](PolicySecurityApi.md#ListDnsSecurityProfileBindings) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps | Get DNS security profile binding map
[**ListDnsSecurityProfileBindings_0**](PolicySecurityApi.md#ListDnsSecurityProfileBindings_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps | Get DNS security profile binding map
[**ListDnsSecurityProfiles**](PolicySecurityApi.md#ListDnsSecurityProfiles) | **Get** /global-infra/dns-security-profiles | List all DNS security profiles
[**ListDnsSecurityProfiles_0**](PolicySecurityApi.md#ListDnsSecurityProfiles_0) | **Get** /infra/dns-security-profiles | List all DNS security profiles
[**ListDrafts**](PolicySecurityApi.md#ListDrafts) | **Get** /infra/drafts | List policy drafts
[**ListDrafts_0**](PolicySecurityApi.md#ListDrafts_0) | **Get** /global-infra/drafts | List policy drafts
[**ListEndpointPoliciesAcrossAllDomains**](PolicySecurityApi.md#ListEndpointPoliciesAcrossAllDomains) | **Get** /global-infra/domains/endpoint-policies | List Endpoint policies
[**ListEndpointPoliciesAcrossAllDomains_0**](PolicySecurityApi.md#ListEndpointPoliciesAcrossAllDomains_0) | **Get** /infra/domains/endpoint-policies | List Endpoint policies
[**ListEndpointRule**](PolicySecurityApi.md#ListEndpointRule) | **Get** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules | List Endpoint rules
[**ListEndpointRule_0**](PolicySecurityApi.md#ListEndpointRule_0) | **Get** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules | List Endpoint rules
[**ListFirewallFloodProtectionBindingsAcrossDomains**](PolicySecurityApi.md#ListFirewallFloodProtectionBindingsAcrossDomains) | **Get** /global-infra/domains/firewall-flood-protection-profile-binding-maps | List Firewall Flood Protection Profile Binding Maps for all domains
[**ListFirewallFloodProtectionBindingsAcrossDomains_0**](PolicySecurityApi.md#ListFirewallFloodProtectionBindingsAcrossDomains_0) | **Get** /infra/domains/firewall-flood-protection-profile-binding-maps | List Firewall Flood Protection Profile Binding Maps for all domains
[**ListFirewallSessionTimerBindingsAcrossDomains**](PolicySecurityApi.md#ListFirewallSessionTimerBindingsAcrossDomains) | **Get** /global-infra/domains/firewall-session-timer-profile-binding-maps | List Firewall Session Timer Profile Binding Maps for all domains
[**ListFirewallSessionTimerBindingsAcrossDomains_0**](PolicySecurityApi.md#ListFirewallSessionTimerBindingsAcrossDomains_0) | **Get** /infra/domains/firewall-session-timer-profile-binding-maps | List Firewall Session Timer Profile Binding Maps for all domains
[**ListFloodProtectionProfileBindings**](PolicySecurityApi.md#ListFloodProtectionProfileBindings) | **Get** /global-infra/flood-protection-profiles/{flood-protection-profile-id}/bindings | List Flood Protection Profiles
[**ListFloodProtectionProfileBindings_0**](PolicySecurityApi.md#ListFloodProtectionProfileBindings_0) | **Get** /infra/flood-protection-profiles/{flood-protection-profile-id}/bindings | List Flood Protection Profiles
[**ListFloodProtectionProfiles**](PolicySecurityApi.md#ListFloodProtectionProfiles) | **Get** /infra/flood-protection-profiles | List Flood Protection Profiles
[**ListFloodProtectionProfiles_0**](PolicySecurityApi.md#ListFloodProtectionProfiles_0) | **Get** /global-infra/flood-protection-profiles | List Flood Protection Profiles
[**ListGatewayPoliciesForDomain**](PolicySecurityApi.md#ListGatewayPoliciesForDomain) | **Get** /global-infra/domains/{domain-id}/gateway-policies | List gateway policies
[**ListGatewayPoliciesForDomain_0**](PolicySecurityApi.md#ListGatewayPoliciesForDomain_0) | **Get** /infra/domains/{domain-id}/gateway-policies | List gateway policies
[**ListGatewayRules**](PolicySecurityApi.md#ListGatewayRules) | **Get** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules | List rules
[**ListGatewayRules_0**](PolicySecurityApi.md#ListGatewayRules_0) | **Get** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules | List rules
[**ListGroupMonitoringBindings**](PolicySecurityApi.md#ListGroupMonitoringBindings) | **Get** /infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps | List Group Monitoring Profile Binding Maps
[**ListGroupMonitoringBindings_0**](PolicySecurityApi.md#ListGroupMonitoringBindings_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps | List Group Monitoring Profile Binding Maps
[**ListIdsClusterConfigs**](PolicySecurityApi.md#ListIdsClusterConfigs) | **Get** /infra/settings/firewall/security/intrusion-services/cluster-configs | List IDS cluster configs
[**ListIdsClusterConfigs_0**](PolicySecurityApi.md#ListIdsClusterConfigs_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/cluster-configs | List IDS cluster configs
[**ListIdsProfiles**](PolicySecurityApi.md#ListIdsProfiles) | **Get** /infra/settings/firewall/security/intrusion-services/profiles | List IDS profiles
[**ListIdsProfiles_0**](PolicySecurityApi.md#ListIdsProfiles_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/profiles | List IDS profiles
[**ListIdsRules**](PolicySecurityApi.md#ListIdsRules) | **Get** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules | List IDS rules
[**ListIdsRules_0**](PolicySecurityApi.md#ListIdsRules_0) | **Get** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules | List IDS rules
[**ListIdsSecurityPolicies**](PolicySecurityApi.md#ListIdsSecurityPolicies) | **Get** /global-infra/domains/{domain-id}/intrusion-service-policies | List IDS security policies
[**ListIdsSecurityPolicies_0**](PolicySecurityApi.md#ListIdsSecurityPolicies_0) | **Get** /infra/domains/{domain-id}/intrusion-service-policies | List IDS security policies
[**ListIdsSignatures**](PolicySecurityApi.md#ListIdsSignatures) | **Get** /infra/settings/firewall/security/intrusion-services/signature-versions/{version-id}/signatures | List IDS signatures
[**ListIdsSignatures_0**](PolicySecurityApi.md#ListIdsSignatures_0) | **Get** /global-infra/settings/firewall/security/intrusion-services/signature-versions/{version-id}/signatures | List IDS signatures
[**ListPolicyFirewallCPUMemThresholdsProfileBindingMaps**](PolicySecurityApi.md#ListPolicyFirewallCPUMemThresholdsProfileBindingMaps) | **Get** /global-infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps | List Firewall CPU Memory Thresholds Profile Binding Maps
[**ListPolicyFirewallCPUMemThresholdsProfileBindingMaps_0**](PolicySecurityApi.md#ListPolicyFirewallCPUMemThresholdsProfileBindingMaps_0) | **Get** /infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps | List Firewall CPU Memory Thresholds Profile Binding Maps
[**ListPolicyFirewallFloodProtectionBindings**](PolicySecurityApi.md#ListPolicyFirewallFloodProtectionBindings) | **Get** /infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps | List Firewall Flood Protection Profile Binding Maps
[**ListPolicyFirewallFloodProtectionBindings_0**](PolicySecurityApi.md#ListPolicyFirewallFloodProtectionBindings_0) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps | List Firewall Flood Protection Profile Binding Maps
[**ListPolicyFirewallSchedulers**](PolicySecurityApi.md#ListPolicyFirewallSchedulers) | **Get** /infra/firewall-schedulers | Get PolicyFirewallSchedulers
[**ListPolicyFirewallSchedulers_0**](PolicySecurityApi.md#ListPolicyFirewallSchedulers_0) | **Get** /global-infra/firewall-schedulers | Get PolicyFirewallSchedulers
[**ListPolicyFirewallSessionTimerBindings**](PolicySecurityApi.md#ListPolicyFirewallSessionTimerBindings) | **Get** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps | List Firewall Session Timer Profile Binding Maps
[**ListPolicyFirewallSessionTimerBindings_0**](PolicySecurityApi.md#ListPolicyFirewallSessionTimerBindings_0) | **Get** /infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps | List Firewall Session Timer Profile Binding Maps
[**ListPolicyFirewallSessionTimerProfiles**](PolicySecurityApi.md#ListPolicyFirewallSessionTimerProfiles) | **Get** /infra/firewall-session-timer-profiles | List Firewall Session Timer Profiles
[**ListPolicyFirewallSessionTimerProfiles_0**](PolicySecurityApi.md#ListPolicyFirewallSessionTimerProfiles_0) | **Get** /global-infra/firewall-session-timer-profiles | List Firewall Session Timer Profiles
[**ListPolicyServiceChainMappings**](PolicySecurityApi.md#ListPolicyServiceChainMappings) | **Get** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id}/service-chain-mappings | List all service chain mappings for given service profile.
[**ListPolicyServiceChainMappings_0**](PolicySecurityApi.md#ListPolicyServiceChainMappings_0) | **Get** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id}/service-chain-mappings | List all service chain mappings for given service profile.
[**ListPolicyServiceChains**](PolicySecurityApi.md#ListPolicyServiceChains) | **Get** /global-infra/service-chains | List service chains
[**ListPolicyServiceChains_0**](PolicySecurityApi.md#ListPolicyServiceChains_0) | **Get** /infra/service-chains | List service chains
[**ListPolicyServiceInstanceEndpoints**](PolicySecurityApi.md#ListPolicyServiceInstanceEndpoints) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints | List all service instance endpoint
[**ListPolicyServiceInstanceEndpoints_0**](PolicySecurityApi.md#ListPolicyServiceInstanceEndpoints_0) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints | List all service instance endpoint
[**ListPolicyServiceProfiles**](PolicySecurityApi.md#ListPolicyServiceProfiles) | **Get** /infra/service-references/{service-reference-id}/service-profiles | List service profiles
[**ListPolicyServiceProfiles_0**](PolicySecurityApi.md#ListPolicyServiceProfiles_0) | **Get** /global-infra/service-references/{service-reference-id}/service-profiles | List service profiles
[**ListPolicyUrlCategories**](PolicySecurityApi.md#ListPolicyUrlCategories) | **Get** /global-infra/url-categories | Get the list of URL categories.
[**ListPolicyUrlCategories_0**](PolicySecurityApi.md#ListPolicyUrlCategories_0) | **Get** /infra/url-categories | Get the list of URL categories.
[**ListPolicyUrlReputationSeverities**](PolicySecurityApi.md#ListPolicyUrlReputationSeverities) | **Get** /infra/url-reputation-severities | Get the list of reputation severity
[**ListPolicyUrlReputationSeverities_0**](PolicySecurityApi.md#ListPolicyUrlReputationSeverities_0) | **Get** /global-infra/url-reputation-severities | Get the list of reputation severity
[**ListRedirectionPolicies**](PolicySecurityApi.md#ListRedirectionPolicies) | **Get** /global-infra/domains/{domain-id}/redirection-policies | List redirection policys for a domain
[**ListRedirectionPoliciesAcrossAllDomains**](PolicySecurityApi.md#ListRedirectionPoliciesAcrossAllDomains) | **Get** /global-infra/domains/redirection-policies | List redirection policys
[**ListRedirectionPoliciesAcrossAllDomains_0**](PolicySecurityApi.md#ListRedirectionPoliciesAcrossAllDomains_0) | **Get** /infra/domains/redirection-policies | List redirection policys
[**ListRedirectionPolicies_0**](PolicySecurityApi.md#ListRedirectionPolicies_0) | **Get** /infra/domains/{domain-id}/redirection-policies | List redirection policys for a domain
[**ListRedirectionRules**](PolicySecurityApi.md#ListRedirectionRules) | **Get** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules | List rules
[**ListRedirectionRules_0**](PolicySecurityApi.md#ListRedirectionRules_0) | **Get** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules | List rules
[**ListSecurityPoliciesForDomain**](PolicySecurityApi.md#ListSecurityPoliciesForDomain) | **Get** /infra/domains/{domain-id}/security-policies | List security policies
[**ListSecurityPoliciesForDomain_0**](PolicySecurityApi.md#ListSecurityPoliciesForDomain_0) | **Get** /global-infra/domains/{domain-id}/security-policies | List security policies
[**ListSecurityRules**](PolicySecurityApi.md#ListSecurityRules) | **Get** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules | List rules
[**ListSecurityRules_0**](PolicySecurityApi.md#ListSecurityRules_0) | **Get** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules | List rules
[**ListServiceDefinitions**](PolicySecurityApi.md#ListServiceDefinitions) | **Get** /enforcement-points/{enforcement-point-id}/service-definitions | List all Service Definitions registered on given enforcement point.
[**ListServiceReferences**](PolicySecurityApi.md#ListServiceReferences) | **Get** /infra/service-references | List service references
[**ListServiceReferences_0**](PolicySecurityApi.md#ListServiceReferences_0) | **Get** /global-infra/service-references | List service references
[**ListSessionTimerProfileBindings**](PolicySecurityApi.md#ListSessionTimerProfileBindings) | **Get** /infra/session-timer-profiles/{session-timer-profile-id}/bindings | List Session Timer Profiles
[**ListSessionTimerProfileBindings_0**](PolicySecurityApi.md#ListSessionTimerProfileBindings_0) | **Get** /global-infra/session-timer-profiles/{session-timer-profile-id}/bindings | List Session Timer Profiles
[**ListVirtualEndpointsForTier0**](PolicySecurityApi.md#ListVirtualEndpointsForTier0) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints | List all virtual endpoints
[**ListVirtualEndpointsForTier0_0**](PolicySecurityApi.md#ListVirtualEndpointsForTier0_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints | List all virtual endpoints
[**MakeVersionAsActiveMakeActiveVersion**](PolicySecurityApi.md#MakeVersionAsActiveMakeActiveVersion) | **Post** /infra/settings/firewall/security/intrusion-services/signature-versions?action&#x3D;make_active_version | Change the state of IDS Signature Version
[**MakeVersionAsActiveMakeActiveVersion_0**](PolicySecurityApi.md#MakeVersionAsActiveMakeActiveVersion_0) | **Post** /global-infra/settings/firewall/security/intrusion-services/signature-versions?action&#x3D;make_active_version | Change the state of IDS Signature Version
[**PatchByodPolicyServiceInstance**](PolicySecurityApi.md#PatchByodPolicyServiceInstance) | **Patch** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Create service instance
[**PatchByodPolicyServiceInstance_0**](PolicySecurityApi.md#PatchByodPolicyServiceInstance_0) | **Patch** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Create service instance
[**PatchCPUMemThresholdsProfile**](PolicySecurityApi.md#PatchCPUMemThresholdsProfile) | **Patch** /infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Create or update CPU and memory thresholds profile
[**PatchCPUMemThresholdsProfile_0**](PolicySecurityApi.md#PatchCPUMemThresholdsProfile_0) | **Patch** /global-infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Create or update CPU and memory thresholds profile
[**PatchCommunicationEntry**](PolicySecurityApi.md#PatchCommunicationEntry) | **Patch** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Patch a CommunicationEntry
[**PatchCommunicationEntry_0**](PolicySecurityApi.md#PatchCommunicationEntry_0) | **Patch** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Patch a CommunicationEntry
[**PatchCommunicationMapForDomain**](PolicySecurityApi.md#PatchCommunicationMapForDomain) | **Patch** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id} | Patch communication map
[**PatchCommunicationMapForDomain_0**](PolicySecurityApi.md#PatchCommunicationMapForDomain_0) | **Patch** /infra/domains/{domain-id}/communication-maps/{communication-map-id} | Patch communication map
[**PatchComputeClusterIdfwConfiguration**](PolicySecurityApi.md#PatchComputeClusterIdfwConfiguration) | **Patch** /infra/settings/firewall/idfw/cluster/{cluster-id} | Patch compute cluster idfw configuration
[**PatchComputeClusterIdfwConfiguration_0**](PolicySecurityApi.md#PatchComputeClusterIdfwConfiguration_0) | **Patch** /global-infra/settings/firewall/idfw/cluster/{cluster-id} | Patch compute cluster idfw configuration
[**PatchDfwFirewallConfiguration**](PolicySecurityApi.md#PatchDfwFirewallConfiguration) | **Patch** /infra/settings/firewall/security | Update dfw firewall configuration
[**PatchDfwFirewallConfiguration_0**](PolicySecurityApi.md#PatchDfwFirewallConfiguration_0) | **Patch** /global-infra/settings/firewall/security | Update dfw firewall configuration
[**PatchDnsSecurityProfile**](PolicySecurityApi.md#PatchDnsSecurityProfile) | **Patch** /infra/dns-security-profiles/{profile-id} | Create or update DNS security profile
[**PatchDnsSecurityProfileBinding**](PolicySecurityApi.md#PatchDnsSecurityProfileBinding) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Create or update DNS security profile binding map
[**PatchDnsSecurityProfileBinding_0**](PolicySecurityApi.md#PatchDnsSecurityProfileBinding_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Create or update DNS security profile binding map
[**PatchDnsSecurityProfile_0**](PolicySecurityApi.md#PatchDnsSecurityProfile_0) | **Patch** /global-infra/dns-security-profiles/{profile-id} | Create or update DNS security profile
[**PatchDraft**](PolicySecurityApi.md#PatchDraft) | **Patch** /infra/drafts/{draft-id} | Patch a manual draft
[**PatchDraft_0**](PolicySecurityApi.md#PatchDraft_0) | **Patch** /global-infra/drafts/{draft-id} | Patch a manual draft
[**PatchEndpointPolicy**](PolicySecurityApi.md#PatchEndpointPolicy) | **Patch** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Create or update Endpoint policy
[**PatchEndpointPolicy_0**](PolicySecurityApi.md#PatchEndpointPolicy_0) | **Patch** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Create or update Endpoint policy
[**PatchEndpointRule**](PolicySecurityApi.md#PatchEndpointRule) | **Patch** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Update Endpoint rule
[**PatchEndpointRule_0**](PolicySecurityApi.md#PatchEndpointRule_0) | **Patch** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Update Endpoint rule
[**PatchExcludeList**](PolicySecurityApi.md#PatchExcludeList) | **Patch** /infra/settings/firewall/security/exclude-list | Patch exclusion list for security policy
[**PatchExcludeList_0**](PolicySecurityApi.md#PatchExcludeList_0) | **Patch** /global-infra/settings/firewall/security/exclude-list | Patch exclusion list for security policy
[**PatchFloodProtectionProfile**](PolicySecurityApi.md#PatchFloodProtectionProfile) | **Patch** /infra/flood-protection-profiles/{flood-protection-profile-id} | Create or update Flood Protection Profile
[**PatchFloodProtectionProfile_0**](PolicySecurityApi.md#PatchFloodProtectionProfile_0) | **Patch** /global-infra/flood-protection-profiles/{flood-protection-profile-id} | Create or update Flood Protection Profile
[**PatchGatewayPolicyForDomain**](PolicySecurityApi.md#PatchGatewayPolicyForDomain) | **Patch** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Update gateway policy
[**PatchGatewayPolicyForDomain_0**](PolicySecurityApi.md#PatchGatewayPolicyForDomain_0) | **Patch** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Update gateway policy
[**PatchGatewayRule**](PolicySecurityApi.md#PatchGatewayRule) | **Patch** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Update gateway rule
[**PatchGatewayRule_0**](PolicySecurityApi.md#PatchGatewayRule_0) | **Patch** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Update gateway rule
[**PatchGroupMonitoringBinding**](PolicySecurityApi.md#PatchGroupMonitoringBinding) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Create Group Monitoring Profile Binding Map
[**PatchGroupMonitoringBinding_0**](PolicySecurityApi.md#PatchGroupMonitoringBinding_0) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Create Group Monitoring Profile Binding Map
[**PatchIdsClusterConfig**](PolicySecurityApi.md#PatchIdsClusterConfig) | **Patch** /infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | Patch IDS config on cluster level
[**PatchIdsClusterConfig_0**](PolicySecurityApi.md#PatchIdsClusterConfig_0) | **Patch** /global-infra/settings/firewall/security/intrusion-services/cluster-configs/{cluster-id} | Patch IDS config on cluster level
[**PatchIdsProfile**](PolicySecurityApi.md#PatchIdsProfile) | **Patch** /global-infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Patch IDS profile
[**PatchIdsProfile_0**](PolicySecurityApi.md#PatchIdsProfile_0) | **Patch** /infra/settings/firewall/security/intrusion-services/profiles/{profile-id} | Patch IDS profile
[**PatchIdsRule**](PolicySecurityApi.md#PatchIdsRule) | **Patch** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Patch IDS rule
[**PatchIdsRule_0**](PolicySecurityApi.md#PatchIdsRule_0) | **Patch** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id} | Patch IDS rule
[**PatchIdsSecurityPolicy**](PolicySecurityApi.md#PatchIdsSecurityPolicy) | **Patch** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Patch IDS security policy
[**PatchIdsSecurityPolicy_0**](PolicySecurityApi.md#PatchIdsSecurityPolicy_0) | **Patch** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id} | Patch IDS security policy
[**PatchIdsSettings**](PolicySecurityApi.md#PatchIdsSettings) | **Patch** /infra/settings/firewall/security/intrusion-services | Patch Intrusion detection system settings
[**PatchIdsSettings_0**](PolicySecurityApi.md#PatchIdsSettings_0) | **Patch** /global-infra/settings/firewall/security/intrusion-services | Patch Intrusion detection system settings
[**PatchIdsStandaloneHostConfig**](PolicySecurityApi.md#PatchIdsStandaloneHostConfig) | **Patch** /infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Patch IDS configuration
[**PatchIdsStandaloneHostConfig_0**](PolicySecurityApi.md#PatchIdsStandaloneHostConfig_0) | **Patch** /global-infra/settings/firewall/security/intrusion-services/ids-standalone-host-config | Patch IDS configuration
[**PatchPolicyFirewallCPUMemThresholdsProfileBindingMap**](PolicySecurityApi.md#PatchPolicyFirewallCPUMemThresholdsProfileBindingMap) | **Patch** /infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Create or update Firewall CPU Memory Thresholds Profile Binding Map
[**PatchPolicyFirewallCPUMemThresholdsProfileBindingMap_0**](PolicySecurityApi.md#PatchPolicyFirewallCPUMemThresholdsProfileBindingMap_0) | **Patch** /global-infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Create or update Firewall CPU Memory Thresholds Profile Binding Map
[**PatchPolicyFirewallFloodProtectionProfileBindingMap**](PolicySecurityApi.md#PatchPolicyFirewallFloodProtectionProfileBindingMap) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Create or update Firewall Flood Protection Profile Binding Map
[**PatchPolicyFirewallFloodProtectionProfileBindingMap_0**](PolicySecurityApi.md#PatchPolicyFirewallFloodProtectionProfileBindingMap_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Create or update Firewall Flood Protection Profile Binding Map
[**PatchPolicyFirewallScheduler**](PolicySecurityApi.md#PatchPolicyFirewallScheduler) | **Patch** /global-infra/firewall-schedulers/{firewall-scheduler-id} | Create or Update PolicyFirewallScheduler
[**PatchPolicyFirewallScheduler_0**](PolicySecurityApi.md#PatchPolicyFirewallScheduler_0) | **Patch** /infra/firewall-schedulers/{firewall-scheduler-id} | Create or Update PolicyFirewallScheduler
[**PatchPolicyFirewallSessionTimerProfile**](PolicySecurityApi.md#PatchPolicyFirewallSessionTimerProfile) | **Patch** /global-infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Create or update Firewall Session Timer Profile
[**PatchPolicyFirewallSessionTimerProfileBindingMap**](PolicySecurityApi.md#PatchPolicyFirewallSessionTimerProfileBindingMap) | **Patch** /infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Create or update Firewall Session Timer Profile Binding Map
[**PatchPolicyFirewallSessionTimerProfileBindingMap_0**](PolicySecurityApi.md#PatchPolicyFirewallSessionTimerProfileBindingMap_0) | **Patch** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Create or update Firewall Session Timer Profile Binding Map
[**PatchPolicyFirewallSessionTimerProfile_0**](PolicySecurityApi.md#PatchPolicyFirewallSessionTimerProfile_0) | **Patch** /infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Create or update Firewall Session Timer Profile
[**PatchPolicyServiceInstance**](PolicySecurityApi.md#PatchPolicyServiceInstance) | **Patch** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create service instance
[**PatchPolicyServiceInstance_0**](PolicySecurityApi.md#PatchPolicyServiceInstance_0) | **Patch** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create service instance
[**PatchPolicyServiceProfile**](PolicySecurityApi.md#PatchPolicyServiceProfile) | **Patch** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Create service profile
[**PatchPolicyServiceProfile_0**](PolicySecurityApi.md#PatchPolicyServiceProfile_0) | **Patch** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Create service profile
[**PatchPolicyUrlCategorizationConfig**](PolicySecurityApi.md#PatchPolicyUrlCategorizationConfig) | **Patch** /global-infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Create or Update PolicyUrlCategorizationConfig
[**PatchPolicyUrlCategorizationConfig_0**](PolicySecurityApi.md#PatchPolicyUrlCategorizationConfig_0) | **Patch** /infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Create or Update PolicyUrlCategorizationConfig
[**PatchRedirectionPolicy**](PolicySecurityApi.md#PatchRedirectionPolicy) | **Patch** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Create or update redirection policy
[**PatchRedirectionPolicy_0**](PolicySecurityApi.md#PatchRedirectionPolicy_0) | **Patch** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Create or update redirection policy
[**PatchRedirectionRule**](PolicySecurityApi.md#PatchRedirectionRule) | **Patch** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Update redirection rule
[**PatchRedirectionRule_0**](PolicySecurityApi.md#PatchRedirectionRule_0) | **Patch** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Update redirection rule
[**PatchSecurityPolicyForDomain**](PolicySecurityApi.md#PatchSecurityPolicyForDomain) | **Patch** /infra/domains/{domain-id}/security-policies/{security-policy-id} | Patch security policy
[**PatchSecurityPolicyForDomain_0**](PolicySecurityApi.md#PatchSecurityPolicyForDomain_0) | **Patch** /global-infra/domains/{domain-id}/security-policies/{security-policy-id} | Patch security policy
[**PatchSecurityRule**](PolicySecurityApi.md#PatchSecurityRule) | **Patch** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Patch a rule
[**PatchSecurityRule_0**](PolicySecurityApi.md#PatchSecurityRule_0) | **Patch** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Patch a rule
[**PatchServiceChain**](PolicySecurityApi.md#PatchServiceChain) | **Patch** /infra/service-chains/{service-chain-id} | Create service chain
[**PatchServiceChain_0**](PolicySecurityApi.md#PatchServiceChain_0) | **Patch** /global-infra/service-chains/{service-chain-id} | Create service chain
[**PatchServiceInstanceEndpoint**](PolicySecurityApi.md#PatchServiceInstanceEndpoint) | **Patch** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Create service instance endpoint
[**PatchServiceInstanceEndpoint_0**](PolicySecurityApi.md#PatchServiceInstanceEndpoint_0) | **Patch** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Create service instance endpoint
[**PatchServiceReference**](PolicySecurityApi.md#PatchServiceReference) | **Patch** /global-infra/service-references/{service-reference-id} | Create service reference
[**PatchServiceReference_0**](PolicySecurityApi.md#PatchServiceReference_0) | **Patch** /infra/service-references/{service-reference-id} | Create service reference
[**PatchStandaloneHostIdfwConfiguration**](PolicySecurityApi.md#PatchStandaloneHostIdfwConfiguration) | **Patch** /infra/settings/firewall/idfw/standalone-host-switch-setting | Patch idfw configuration for standalone host
[**PatchStandaloneHostIdfwConfiguration_0**](PolicySecurityApi.md#PatchStandaloneHostIdfwConfiguration_0) | **Patch** /global-infra/settings/firewall/idfw/standalone-host-switch-setting | Patch idfw configuration for standalone host
[**PatchTier0FloodProtectionProfileBindingMap**](PolicySecurityApi.md#PatchTier0FloodProtectionProfileBindingMap) | **Patch** /infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router
[**PatchTier0FloodProtectionProfileBindingMap_0**](PolicySecurityApi.md#PatchTier0FloodProtectionProfileBindingMap_0) | **Patch** /global-infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router
[**PatchTier0LocalServicesSessionTimerProfileBindingMap**](PolicySecurityApi.md#PatchTier0LocalServicesSessionTimerProfileBindingMap) | **Patch** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**PatchTier0LocalServicesSessionTimerProfileBindingMap_0**](PolicySecurityApi.md#PatchTier0LocalServicesSessionTimerProfileBindingMap_0) | **Patch** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**PatchTier0LocaleServicesFloodProtectionProfileBindingMap**](PolicySecurityApi.md#PatchTier0LocaleServicesFloodProtectionProfileBindingMap) | **Patch** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**PatchTier0LocaleServicesFloodProtectionProfileBindingMap_0**](PolicySecurityApi.md#PatchTier0LocaleServicesFloodProtectionProfileBindingMap_0) | **Patch** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**PatchTier0SessionTimerProfileBindingMap**](PolicySecurityApi.md#PatchTier0SessionTimerProfileBindingMap) | **Patch** /global-infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router
[**PatchTier0SessionTimerProfileBindingMap_0**](PolicySecurityApi.md#PatchTier0SessionTimerProfileBindingMap_0) | **Patch** /infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router
[**PatchTier1FloodProtectionProfileBindingMap**](PolicySecurityApi.md#PatchTier1FloodProtectionProfileBindingMap) | **Patch** /infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router
[**PatchTier1FloodProtectionProfileBindingMap_0**](PolicySecurityApi.md#PatchTier1FloodProtectionProfileBindingMap_0) | **Patch** /global-infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router
[**PatchTier1LocaleServicesFloodProtectionProfileBindingMap**](PolicySecurityApi.md#PatchTier1LocaleServicesFloodProtectionProfileBindingMap) | **Patch** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**PatchTier1LocaleServicesFloodProtectionProfileBindingMap_0**](PolicySecurityApi.md#PatchTier1LocaleServicesFloodProtectionProfileBindingMap_0) | **Patch** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**PatchTier1LocaleServicesSessionTimerProfileBindingMap**](PolicySecurityApi.md#PatchTier1LocaleServicesSessionTimerProfileBindingMap) | **Patch** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**PatchTier1LocaleServicesSessionTimerProfileBindingMap_0**](PolicySecurityApi.md#PatchTier1LocaleServicesSessionTimerProfileBindingMap_0) | **Patch** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**PatchTier1PolicyServiceInstance**](PolicySecurityApi.md#PatchTier1PolicyServiceInstance) | **Patch** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create Tier1 service instance
[**PatchTier1PolicyServiceInstance_0**](PolicySecurityApi.md#PatchTier1PolicyServiceInstance_0) | **Patch** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Create Tier1 service instance
[**PatchTier1SessionTimerProfileBindingMap**](PolicySecurityApi.md#PatchTier1SessionTimerProfileBindingMap) | **Patch** /global-infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router
[**PatchTier1SessionTimerProfileBindingMap_0**](PolicySecurityApi.md#PatchTier1SessionTimerProfileBindingMap_0) | **Patch** /infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router
[**PatchVirtualEndpoint**](PolicySecurityApi.md#PatchVirtualEndpoint) | **Patch** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Create or update virtual endpoint
[**PatchVirtualEndpoint_0**](PolicySecurityApi.md#PatchVirtualEndpoint_0) | **Patch** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Create or update virtual endpoint
[**PublishDraftPublish**](PolicySecurityApi.md#PublishDraftPublish) | **Post** /infra/drafts/{draft-id}?action&#x3D;publish | Publish a draft
[**PublishDraftPublish_0**](PolicySecurityApi.md#PublishDraftPublish_0) | **Post** /global-infra/drafts/{draft-id}?action&#x3D;publish | Publish a draft
[**PutComputeClusterIdfwConfiguration**](PolicySecurityApi.md#PutComputeClusterIdfwConfiguration) | **Put** /infra/settings/firewall/idfw/cluster/{cluster-id} | Create or update compute cluster idfw configuration
[**PutComputeClusterIdfwConfiguration_0**](PolicySecurityApi.md#PutComputeClusterIdfwConfiguration_0) | **Put** /global-infra/settings/firewall/idfw/cluster/{cluster-id} | Create or update compute cluster idfw configuration
[**PutDfwFirewallConfiguration**](PolicySecurityApi.md#PutDfwFirewallConfiguration) | **Put** /infra/settings/firewall/security | Update dfw firewall configuration
[**PutDfwFirewallConfiguration_0**](PolicySecurityApi.md#PutDfwFirewallConfiguration_0) | **Put** /global-infra/settings/firewall/security | Update dfw firewall configuration
[**PutDraft**](PolicySecurityApi.md#PutDraft) | **Put** /infra/drafts/{draft-id} | Create or update a manual draft
[**PutDraft_0**](PolicySecurityApi.md#PutDraft_0) | **Put** /global-infra/drafts/{draft-id} | Create or update a manual draft
[**PutExcludeList**](PolicySecurityApi.md#PutExcludeList) | **Put** /infra/settings/firewall/security/exclude-list | Create or update exclusion list for security policy
[**PutExcludeList_0**](PolicySecurityApi.md#PutExcludeList_0) | **Put** /global-infra/settings/firewall/security/exclude-list | Create or update exclusion list for security policy
[**PutPolicyUrlCategorizationConfig**](PolicySecurityApi.md#PutPolicyUrlCategorizationConfig) | **Put** /global-infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Create or Update PolicyUrlCategorizationConfig
[**PutPolicyUrlCategorizationConfig_0**](PolicySecurityApi.md#PutPolicyUrlCategorizationConfig_0) | **Put** /infra/sites/{site-id}/enforcement-points/{enforcement-point-id}/edge-clusters/{edge-cluster-id}/url-categorization-configs/{url-categorization-config-id} | Create or Update PolicyUrlCategorizationConfig
[**PutStandaloneHostIdfwConfiguration**](PolicySecurityApi.md#PutStandaloneHostIdfwConfiguration) | **Put** /infra/settings/firewall/idfw/standalone-host-switch-setting | Create or update idfw configuration for standalone host
[**PutStandaloneHostIdfwConfiguration_0**](PolicySecurityApi.md#PutStandaloneHostIdfwConfiguration_0) | **Put** /global-infra/settings/firewall/idfw/standalone-host-switch-setting | Create or update idfw configuration for standalone host
[**ReadAllPolicyServiceInstancesForTier0**](PolicySecurityApi.md#ReadAllPolicyServiceInstancesForTier0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances | Read all service instance objects under a tier-0
[**ReadAllPolicyServiceInstancesForTier0_0**](PolicySecurityApi.md#ReadAllPolicyServiceInstancesForTier0_0) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances | Read all service instance objects under a tier-0
[**ReadAllPolicyServiceInstancesForTier1**](PolicySecurityApi.md#ReadAllPolicyServiceInstancesForTier1) | **Get** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances | Read all service instance objects under a tier-1
[**ReadAllPolicyServiceInstancesForTier1_0**](PolicySecurityApi.md#ReadAllPolicyServiceInstancesForTier1_0) | **Get** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances | Read all service instance objects under a tier-1
[**ReadByodPolicyServiceInstance**](PolicySecurityApi.md#ReadByodPolicyServiceInstance) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Read byod service instance
[**ReadByodPolicyServiceInstance_0**](PolicySecurityApi.md#ReadByodPolicyServiceInstance_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id} | Read byod service instance
[**ReadCPUMemThresholdsProfile**](PolicySecurityApi.md#ReadCPUMemThresholdsProfile) | **Get** /infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Read the CPU and memory thresholds profile
[**ReadCPUMemThresholdsProfile_0**](PolicySecurityApi.md#ReadCPUMemThresholdsProfile_0) | **Get** /global-infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Read the CPU and memory thresholds profile
[**ReadCommunicationEntry**](PolicySecurityApi.md#ReadCommunicationEntry) | **Get** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Read CommunicationEntry
[**ReadCommunicationEntry_0**](PolicySecurityApi.md#ReadCommunicationEntry_0) | **Get** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Read CommunicationEntry
[**ReadCommunicationMapForDomain**](PolicySecurityApi.md#ReadCommunicationMapForDomain) | **Get** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id} | Read communication-map
[**ReadCommunicationMapForDomain_0**](PolicySecurityApi.md#ReadCommunicationMapForDomain_0) | **Get** /infra/domains/{domain-id}/communication-maps/{communication-map-id} | Read communication-map
[**ReadDnsSecurityProfile**](PolicySecurityApi.md#ReadDnsSecurityProfile) | **Get** /infra/dns-security-profiles/{profile-id} | Read the DNS Forwarder for the given tier-0 instance
[**ReadDnsSecurityProfile_0**](PolicySecurityApi.md#ReadDnsSecurityProfile_0) | **Get** /global-infra/dns-security-profiles/{profile-id} | Read the DNS Forwarder for the given tier-0 instance
[**ReadDraft**](PolicySecurityApi.md#ReadDraft) | **Get** /infra/drafts/{draft-id} | Read draft
[**ReadDraft_0**](PolicySecurityApi.md#ReadDraft_0) | **Get** /global-infra/drafts/{draft-id} | Read draft
[**ReadEndpointPolicy**](PolicySecurityApi.md#ReadEndpointPolicy) | **Get** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Read Endpoint policy
[**ReadEndpointPolicy_0**](PolicySecurityApi.md#ReadEndpointPolicy_0) | **Get** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id} | Read Endpoint policy
[**ReadEndpointRule**](PolicySecurityApi.md#ReadEndpointRule) | **Get** /global-infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Read Endpoint rule
[**ReadEndpointRule_0**](PolicySecurityApi.md#ReadEndpointRule_0) | **Get** /infra/domains/{domain-id}/endpoint-policies/{endpoint-policy-id}/endpoint-rules/{endpoint-rule-id} | Read Endpoint rule
[**ReadGatewayPolicyForDomain**](PolicySecurityApi.md#ReadGatewayPolicyForDomain) | **Get** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Read gateway policy
[**ReadGatewayPolicyForDomain_0**](PolicySecurityApi.md#ReadGatewayPolicyForDomain_0) | **Get** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id} | Read gateway policy
[**ReadGatewayRule**](PolicySecurityApi.md#ReadGatewayRule) | **Get** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Read rule
[**ReadGatewayRule_0**](PolicySecurityApi.md#ReadGatewayRule_0) | **Get** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id} | Read rule
[**ReadPartnerService**](PolicySecurityApi.md#ReadPartnerService) | **Get** /global-infra/partner-services/{service-name} | Read partner service identified by provided name
[**ReadPartnerService_0**](PolicySecurityApi.md#ReadPartnerService_0) | **Get** /infra/partner-services/{service-name} | Read partner service identified by provided name
[**ReadPartnerServices**](PolicySecurityApi.md#ReadPartnerServices) | **Get** /global-infra/partner-services | Read partner services
[**ReadPartnerServices_0**](PolicySecurityApi.md#ReadPartnerServices_0) | **Get** /infra/partner-services | Read partner services
[**ReadPolicyServiceInstance**](PolicySecurityApi.md#ReadPolicyServiceInstance) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Read service instance
[**ReadPolicyServiceInstanceEndpoint**](PolicySecurityApi.md#ReadPolicyServiceInstanceEndpoint) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Read service instance endpoint
[**ReadPolicyServiceInstanceEndpoint_0**](PolicySecurityApi.md#ReadPolicyServiceInstanceEndpoint_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/byod-service-instances/{service-instance-id}/service-instance-endpoints/{service-instance-endpoint-id} | Read service instance endpoint
[**ReadPolicyServiceInstance_0**](PolicySecurityApi.md#ReadPolicyServiceInstance_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Read service instance
[**ReadPolicyServiceProfile**](PolicySecurityApi.md#ReadPolicyServiceProfile) | **Get** /infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Read service profile
[**ReadPolicyServiceProfile_0**](PolicySecurityApi.md#ReadPolicyServiceProfile_0) | **Get** /global-infra/service-references/{service-reference-id}/service-profiles/{service-profile-id} | Read service profile
[**ReadRedirectionPolicy**](PolicySecurityApi.md#ReadRedirectionPolicy) | **Get** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Read redirection policy
[**ReadRedirectionPolicy_0**](PolicySecurityApi.md#ReadRedirectionPolicy_0) | **Get** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id} | Read redirection policy
[**ReadRedirectionRule**](PolicySecurityApi.md#ReadRedirectionRule) | **Get** /infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Read rule
[**ReadRedirectionRule_0**](PolicySecurityApi.md#ReadRedirectionRule_0) | **Get** /global-infra/domains/{domain-id}/redirection-policies/{redirection-policy-id}/rules/{rule-id} | Read rule
[**ReadSecurityPolicyForDomain**](PolicySecurityApi.md#ReadSecurityPolicyForDomain) | **Get** /infra/domains/{domain-id}/security-policies/{security-policy-id} | Read security policy
[**ReadSecurityPolicyForDomain_0**](PolicySecurityApi.md#ReadSecurityPolicyForDomain_0) | **Get** /global-infra/domains/{domain-id}/security-policies/{security-policy-id} | Read security policy
[**ReadSecurityRule**](PolicySecurityApi.md#ReadSecurityRule) | **Get** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Read rule
[**ReadSecurityRule_0**](PolicySecurityApi.md#ReadSecurityRule_0) | **Get** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Read rule
[**ReadServiceChain**](PolicySecurityApi.md#ReadServiceChain) | **Get** /infra/service-chains/{service-chain-id} | Read service chain
[**ReadServiceChain_0**](PolicySecurityApi.md#ReadServiceChain_0) | **Get** /global-infra/service-chains/{service-chain-id} | Read service chain
[**ReadServiceDefinition**](PolicySecurityApi.md#ReadServiceDefinition) | **Get** /enforcement-points/{enforcement-point-id}/service-definitions/{service-definition-id} | Read Service Definition with given service-definition-id.
[**ReadServiceReference**](PolicySecurityApi.md#ReadServiceReference) | **Get** /global-infra/service-references/{service-reference-id} | Read service reference
[**ReadServiceReference_0**](PolicySecurityApi.md#ReadServiceReference_0) | **Get** /infra/service-references/{service-reference-id} | Read service reference
[**ReadTier1PolicyServiceInstance**](PolicySecurityApi.md#ReadTier1PolicyServiceInstance) | **Get** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Read Tier1 service instance
[**ReadTier1PolicyServiceInstance_0**](PolicySecurityApi.md#ReadTier1PolicyServiceInstance_0) | **Get** /infra/tier-1s/{tier-1-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id} | Read Tier1 service instance
[**ReadVirtualEndpoint**](PolicySecurityApi.md#ReadVirtualEndpoint) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Read virtual endpoint
[**ReadVirtualEndpoint_0**](PolicySecurityApi.md#ReadVirtualEndpoint_0) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/endpoints/virtual-endpoints/{virtual-endpoint-id} | Read virtual endpoint
[**RenewAuthenticationTokensForPolicyServiceInstanceReauth**](PolicySecurityApi.md#RenewAuthenticationTokensForPolicyServiceInstanceReauth) | **Post** /infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}?action&#x3D;reauth | Renew the authentication tokens
[**RenewAuthenticationTokensForPolicyServiceInstanceReauth_0**](PolicySecurityApi.md#RenewAuthenticationTokensForPolicyServiceInstanceReauth_0) | **Post** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-service-id}/service-instances/{service-instance-id}?action&#x3D;reauth | Renew the authentication tokens
[**ResetRuleStatsReset**](PolicySecurityApi.md#ResetRuleStatsReset) | **Post** /global-infra/settings/firewall/stats?action&#x3D;reset | Reset firewall rule statistics
[**ResetRuleStatsReset_0**](PolicySecurityApi.md#ResetRuleStatsReset_0) | **Post** /infra/settings/firewall/stats?action&#x3D;reset | Reset firewall rule statistics
[**ReviseCommunicationEntryRevise**](PolicySecurityApi.md#ReviseCommunicationEntryRevise) | **Post** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id}?action&#x3D;revise | Revise the positioning of communication entry
[**ReviseCommunicationEntryRevise_0**](PolicySecurityApi.md#ReviseCommunicationEntryRevise_0) | **Post** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id}?action&#x3D;revise | Revise the positioning of communication entry
[**ReviseCommunicationMapsRevise**](PolicySecurityApi.md#ReviseCommunicationMapsRevise) | **Post** /infra/domains/{domain-id}/communication-maps/{communication-map-id}?action&#x3D;revise | Revise the positioning of communication maps
[**ReviseCommunicationMapsRevise_0**](PolicySecurityApi.md#ReviseCommunicationMapsRevise_0) | **Post** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}?action&#x3D;revise | Revise the positioning of communication maps
[**ReviseGatewayPolicyRevise**](PolicySecurityApi.md#ReviseGatewayPolicyRevise) | **Post** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}?action&#x3D;revise | Revise the positioning of gateway policy
[**ReviseGatewayPolicyRevise_0**](PolicySecurityApi.md#ReviseGatewayPolicyRevise_0) | **Post** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}?action&#x3D;revise | Revise the positioning of gateway policy
[**ReviseGatewayRuleRevise**](PolicySecurityApi.md#ReviseGatewayRuleRevise) | **Post** /infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of gateway rule
[**ReviseGatewayRuleRevise_0**](PolicySecurityApi.md#ReviseGatewayRuleRevise_0) | **Post** /global-infra/domains/{domain-id}/gateway-policies/{gateway-policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of gateway rule
[**ReviseIdsRuleRevise**](PolicySecurityApi.md#ReviseIdsRuleRevise) | **Post** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of IDS rule
[**ReviseIdsRuleRevise_0**](PolicySecurityApi.md#ReviseIdsRuleRevise_0) | **Post** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of IDS rule
[**ReviseIdsSecurityPolicyRevise**](PolicySecurityApi.md#ReviseIdsSecurityPolicyRevise) | **Post** /global-infra/domains/{domain-id}/intrusion-service-policies/{policy-id}?action&#x3D;revise | Revise the positioning of IDS security policies
[**ReviseIdsSecurityPolicyRevise_0**](PolicySecurityApi.md#ReviseIdsSecurityPolicyRevise_0) | **Post** /infra/domains/{domain-id}/intrusion-service-policies/{policy-id}?action&#x3D;revise | Revise the positioning of IDS security policies
[**ReviseSecurityPoliciesRevise**](PolicySecurityApi.md#ReviseSecurityPoliciesRevise) | **Post** /infra/domains/{domain-id}/security-policies/{security-policy-id}?action&#x3D;revise | Revise the positioning of security policies
[**ReviseSecurityPoliciesRevise_0**](PolicySecurityApi.md#ReviseSecurityPoliciesRevise_0) | **Post** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}?action&#x3D;revise | Revise the positioning of security policies
[**ReviseSecurityRuleRevise**](PolicySecurityApi.md#ReviseSecurityRuleRevise) | **Post** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of rule
[**ReviseSecurityRuleRevise_0**](PolicySecurityApi.md#ReviseSecurityRuleRevise_0) | **Post** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id}?action&#x3D;revise | Revise the positioning of rule
[**UpdateCPUMemThresholdsProfile**](PolicySecurityApi.md#UpdateCPUMemThresholdsProfile) | **Put** /infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Create or update CPU and memory thresholds profile
[**UpdateCPUMemThresholdsProfile_0**](PolicySecurityApi.md#UpdateCPUMemThresholdsProfile_0) | **Put** /global-infra/settings/firewall/cpu-mem-thresholds-profiles/{profile-id} | Create or update CPU and memory thresholds profile
[**UpdateCommunicationEntry**](PolicySecurityApi.md#UpdateCommunicationEntry) | **Put** /infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Create or update a CommunicationEntry
[**UpdateCommunicationEntry_0**](PolicySecurityApi.md#UpdateCommunicationEntry_0) | **Put** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id}/communication-entries/{communication-entry-id} | Create or update a CommunicationEntry
[**UpdateCommunicationMapForDomain**](PolicySecurityApi.md#UpdateCommunicationMapForDomain) | **Put** /global-infra/domains/{domain-id}/communication-maps/{communication-map-id} | Create or Update communication map
[**UpdateCommunicationMapForDomain_0**](PolicySecurityApi.md#UpdateCommunicationMapForDomain_0) | **Put** /infra/domains/{domain-id}/communication-maps/{communication-map-id} | Create or Update communication map
[**UpdateDnsSecurityProfile**](PolicySecurityApi.md#UpdateDnsSecurityProfile) | **Put** /infra/dns-security-profiles/{profile-id} | Create or update DNS security profile
[**UpdateDnsSecurityProfileBinding**](PolicySecurityApi.md#UpdateDnsSecurityProfileBinding) | **Put** /infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Update DNS security profile binding map
[**UpdateDnsSecurityProfileBinding_0**](PolicySecurityApi.md#UpdateDnsSecurityProfileBinding_0) | **Put** /global-infra/domains/{domain-id}/groups/{group-id}/dns-security-profile-binding-maps/{dns-security-profile-binding-map-id} | Update DNS security profile binding map
[**UpdateDnsSecurityProfile_0**](PolicySecurityApi.md#UpdateDnsSecurityProfile_0) | **Put** /global-infra/dns-security-profiles/{profile-id} | Create or update DNS security profile
[**UpdateFloodProtectionProfile**](PolicySecurityApi.md#UpdateFloodProtectionProfile) | **Put** /infra/flood-protection-profiles/{flood-protection-profile-id} | Update Firewall Flood Protection Profile
[**UpdateFloodProtectionProfile_0**](PolicySecurityApi.md#UpdateFloodProtectionProfile_0) | **Put** /global-infra/flood-protection-profiles/{flood-protection-profile-id} | Update Firewall Flood Protection Profile
[**UpdateGroupMonitoringBinding**](PolicySecurityApi.md#UpdateGroupMonitoringBinding) | **Put** /global-infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Update Group Monitoring Profile Binding Map
[**UpdateGroupMonitoringBinding_0**](PolicySecurityApi.md#UpdateGroupMonitoringBinding_0) | **Put** /infra/domains/{domain-id}/groups/{group-id}/group-monitoring-profile-binding-maps/{group-monitoring-profile-binding-map-id} | Update Group Monitoring Profile Binding Map
[**UpdateIdsSettings**](PolicySecurityApi.md#UpdateIdsSettings) | **Put** /infra/settings/firewall/security/intrusion-services | Update Intrusion detection system settings
[**UpdateIdsSettings_0**](PolicySecurityApi.md#UpdateIdsSettings_0) | **Put** /global-infra/settings/firewall/security/intrusion-services | Update Intrusion detection system settings
[**UpdateIdsSignaturesUpdateSignatures**](PolicySecurityApi.md#UpdateIdsSignaturesUpdateSignatures) | **Post** /global-infra/settings/firewall/security/intrusion-services/signatures?action&#x3D;update_signatures | Download and update IDS signatures
[**UpdateIdsSignaturesUpdateSignatures_0**](PolicySecurityApi.md#UpdateIdsSignaturesUpdateSignatures_0) | **Post** /infra/settings/firewall/security/intrusion-services/signatures?action&#x3D;update_signatures | Download and update IDS signatures
[**UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap**](PolicySecurityApi.md#UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap) | **Put** /infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Update Firewall CPU Memory Thresholds Profile Binding Map
[**UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap_0**](PolicySecurityApi.md#UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap_0) | **Put** /global-infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpu-mem-thresholds-profile-binding-map-id} | Update Firewall CPU Memory Thresholds Profile Binding Map
[**UpdatePolicyFirewallFloodProtectionBinding**](PolicySecurityApi.md#UpdatePolicyFirewallFloodProtectionBinding) | **Put** /infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Update Firewall Flood Protection Profile Binding Map
[**UpdatePolicyFirewallFloodProtectionBinding_0**](PolicySecurityApi.md#UpdatePolicyFirewallFloodProtectionBinding_0) | **Put** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-flood-protection-profile-binding-maps/{firewall-flood-protection-profile-binding-map-id} | Update Firewall Flood Protection Profile Binding Map
[**UpdatePolicyFirewallScheduler**](PolicySecurityApi.md#UpdatePolicyFirewallScheduler) | **Put** /global-infra/firewall-schedulers/{firewall-scheduler-id} | Create or Update PolicyFirewallScheduler
[**UpdatePolicyFirewallScheduler_0**](PolicySecurityApi.md#UpdatePolicyFirewallScheduler_0) | **Put** /infra/firewall-schedulers/{firewall-scheduler-id} | Create or Update PolicyFirewallScheduler
[**UpdatePolicyFirewallSessionTimerBinding**](PolicySecurityApi.md#UpdatePolicyFirewallSessionTimerBinding) | **Put** /infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Update Firewall Session Timer Profile Binding Map
[**UpdatePolicyFirewallSessionTimerBinding_0**](PolicySecurityApi.md#UpdatePolicyFirewallSessionTimerBinding_0) | **Put** /global-infra/domains/{domain-id}/groups/{group-id}/firewall-session-timer-profile-binding-maps/{firewall-session-timer-profile-binding-map-id} | Update Firewall Session Timer Profile Binding Map
[**UpdatePolicyFirewallSessionTimerProfile**](PolicySecurityApi.md#UpdatePolicyFirewallSessionTimerProfile) | **Put** /global-infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Update Firewall Session Timer Profile
[**UpdatePolicyFirewallSessionTimerProfile_0**](PolicySecurityApi.md#UpdatePolicyFirewallSessionTimerProfile_0) | **Put** /infra/firewall-session-timer-profiles/{firewall-session-timer-profile-id} | Update Firewall Session Timer Profile
[**UpdateSecurityPolicyForDomain**](PolicySecurityApi.md#UpdateSecurityPolicyForDomain) | **Put** /infra/domains/{domain-id}/security-policies/{security-policy-id} | Create or Update security policy
[**UpdateSecurityPolicyForDomain_0**](PolicySecurityApi.md#UpdateSecurityPolicyForDomain_0) | **Put** /global-infra/domains/{domain-id}/security-policies/{security-policy-id} | Create or Update security policy
[**UpdateSecurityRule**](PolicySecurityApi.md#UpdateSecurityRule) | **Put** /global-infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Create or update a rule
[**UpdateSecurityRule_0**](PolicySecurityApi.md#UpdateSecurityRule_0) | **Put** /infra/domains/{domain-id}/security-policies/{security-policy-id}/rules/{rule-id} | Create or update a rule
[**UpdateServiceDefinition**](PolicySecurityApi.md#UpdateServiceDefinition) | **Put** /enforcement-points/{enforcement-point-id}/service-definitions/{service-definition-id} | Update an existing Service Definition on the given enforcement point 
[**UpdateTier0FloodProtectionProfileBinding**](PolicySecurityApi.md#UpdateTier0FloodProtectionProfileBinding) | **Put** /infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router
[**UpdateTier0FloodProtectionProfileBinding_0**](PolicySecurityApi.md#UpdateTier0FloodProtectionProfileBinding_0) | **Put** /global-infra/tier-0s/{tier0-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router
[**UpdateTier0LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#UpdateTier0LocaleServicesFloodProtectionProfileBinding) | **Put** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**UpdateTier0LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#UpdateTier0LocaleServicesFloodProtectionProfileBinding_0) | **Put** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices
[**UpdateTier0LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#UpdateTier0LocaleServicesSessionTimerProfileBinding) | **Put** /global-infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**UpdateTier0LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#UpdateTier0LocaleServicesSessionTimerProfileBinding_0) | **Put** /infra/tier-0s/{tier0-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices
[**UpdateTier0SessionTimerProfileBinding**](PolicySecurityApi.md#UpdateTier0SessionTimerProfileBinding) | **Put** /global-infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router
[**UpdateTier0SessionTimerProfileBinding_0**](PolicySecurityApi.md#UpdateTier0SessionTimerProfileBinding_0) | **Put** /infra/tier-0s/{tier0-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-0 Logical Router
[**UpdateTier1FloodProtectionProfileBinding**](PolicySecurityApi.md#UpdateTier1FloodProtectionProfileBinding) | **Put** /infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router
[**UpdateTier1FloodProtectionProfileBinding_0**](PolicySecurityApi.md#UpdateTier1FloodProtectionProfileBinding_0) | **Put** /global-infra/tier-1s/{tier1-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router
[**UpdateTier1LocaleServicesFloodProtectionProfileBinding**](PolicySecurityApi.md#UpdateTier1LocaleServicesFloodProtectionProfileBinding) | **Put** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**UpdateTier1LocaleServicesFloodProtectionProfileBinding_0**](PolicySecurityApi.md#UpdateTier1LocaleServicesFloodProtectionProfileBinding_0) | **Put** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/flood-protection-profile-bindings/{flood-protection-profile-binding-id} | Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices
[**UpdateTier1LocaleServicesSessionTimerProfileBinding**](PolicySecurityApi.md#UpdateTier1LocaleServicesSessionTimerProfileBinding) | **Put** /infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**UpdateTier1LocaleServicesSessionTimerProfileBinding_0**](PolicySecurityApi.md#UpdateTier1LocaleServicesSessionTimerProfileBinding_0) | **Put** /global-infra/tier-1s/{tier1-id}/locale-services/{locale-services-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices
[**UpdateTier1SessionTimerProfileBinding**](PolicySecurityApi.md#UpdateTier1SessionTimerProfileBinding) | **Put** /global-infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router
[**UpdateTier1SessionTimerProfileBinding_0**](PolicySecurityApi.md#UpdateTier1SessionTimerProfileBinding_0) | **Put** /infra/tier-1s/{tier1-id}/session-timer-profile-bindings/{session-timer-profile-binding-id} | Create or update Session Timer Profile Binding Map for Tier-1 Logical Router
[**ViewTier0GatewayFirewall**](PolicySecurityApi.md#ViewTier0GatewayFirewall) | **Get** /infra/tier-0s/{tier-0-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-0 logical router. 
[**ViewTier0GatewayFirewall_0**](PolicySecurityApi.md#ViewTier0GatewayFirewall_0) | **Get** /global-infra/tier-0s/{tier-0-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-0 logical router. 
[**ViewTier0LocaleServicesGatewayFirewall**](PolicySecurityApi.md#ViewTier0LocaleServicesGatewayFirewall) | **Get** /infra/tier-0s/{tier-0-id}/locale-services/{locale-services-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-0 LocalServices. 
[**ViewTier0LocaleServicesGatewayFirewall_0**](PolicySecurityApi.md#ViewTier0LocaleServicesGatewayFirewall_0) | **Get** /global-infra/tier-0s/{tier-0-id}/locale-services/{locale-services-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-0 LocalServices. 
[**ViewTier1GatewayFirewall**](PolicySecurityApi.md#ViewTier1GatewayFirewall) | **Get** /global-infra/tier-1s/{tier-1-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-1. 
[**ViewTier1GatewayFirewall_0**](PolicySecurityApi.md#ViewTier1GatewayFirewall_0) | **Get** /infra/tier-1s/{tier-1-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-1. 
[**ViewTier1LocaleServicesGatewayFirewall**](PolicySecurityApi.md#ViewTier1LocaleServicesGatewayFirewall) | **Get** /global-infra/tier-1s/{tier-1-id}/locale-services/{locale-services-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-1 LocalServices. 
[**ViewTier1LocaleServicesGatewayFirewall_0**](PolicySecurityApi.md#ViewTier1LocaleServicesGatewayFirewall_0) | **Get** /infra/tier-1s/{tier-1-id}/locale-services/{locale-services-id}/gateway-firewall | Get list of gateway policies with rules that belong to the specific Tier-1 LocalServices. 

# **CreateByodPolicyServiceInstance**
> ByodPolicyServiceInstance CreateByodPolicyServiceInstance(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create BYOD Service Instance which represent instance of service definition created on manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Byod service instance id | 

### Return type

[**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateByodPolicyServiceInstance_0**
> ByodPolicyServiceInstance CreateByodPolicyServiceInstance_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create BYOD Service Instance which represent instance of service definition created on manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Byod service instance id | 

### Return type

[**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceGatewayPolicyForDomain**
> GatewayPolicy CreateOrReplaceGatewayPolicyForDomain(ctx, body, domainId, gatewayPolicyId)
Update gateway policy

Update the gateway policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceGatewayPolicyForDomain_0**
> GatewayPolicy CreateOrReplaceGatewayPolicyForDomain_0(ctx, body, domainId, gatewayPolicyId)
Update gateway policy

Update the gateway policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceGatewayRule**
> Rule CreateOrReplaceGatewayRule(ctx, body, domainId, gatewayPolicyId, ruleId)
Update gateway rule

Update the gateway rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrReplaceGatewayRule_0**
> Rule CreateOrReplaceGatewayRule_0(ctx, body, domainId, gatewayPolicyId, ruleId)
Update gateway rule

Update the gateway rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEndpointPolicy**
> EndpointPolicy CreateOrUpdateEndpointPolicy(ctx, body, domainId, endpointPolicyId)
Create or update Endpoint policy

Create or update the Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointPolicy**](EndpointPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

[**EndpointPolicy**](EndpointPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEndpointPolicy_0**
> EndpointPolicy CreateOrUpdateEndpointPolicy_0(ctx, body, domainId, endpointPolicyId)
Create or update Endpoint policy

Create or update the Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointPolicy**](EndpointPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

[**EndpointPolicy**](EndpointPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEndpointRule**
> EndpointRule CreateOrUpdateEndpointRule(ctx, body, domainId, endpointPolicyId, endpointRuleId)
Update Endpoint rule

Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointRule**](EndpointRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

[**EndpointRule**](EndpointRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateEndpointRule_0**
> EndpointRule CreateOrUpdateEndpointRule_0(ctx, body, domainId, endpointPolicyId, endpointRuleId)
Update Endpoint rule

Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointRule**](EndpointRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

[**EndpointRule**](EndpointRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsClusterConfig**
> IdsClusterConfig CreateOrUpdateIdsClusterConfig(ctx, body, clusterId)
create or update IDS config on cluster level

Update intrusion detection system on cluster level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsClusterConfig**](IdsClusterConfig.md)|  | 
  **clusterId** | **string**| User entered ID | 

### Return type

[**IdsClusterConfig**](IdsClusterConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsClusterConfig_0**
> IdsClusterConfig CreateOrUpdateIdsClusterConfig_0(ctx, body, clusterId)
create or update IDS config on cluster level

Update intrusion detection system on cluster level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsClusterConfig**](IdsClusterConfig.md)|  | 
  **clusterId** | **string**| User entered ID | 

### Return type

[**IdsClusterConfig**](IdsClusterConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsProfile**
> IdsProfile CreateOrUpdateIdsProfile(ctx, body, profileId)
create or update IDS profile

Update intrusion detection profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsProfile**](IdsProfile.md)|  | 
  **profileId** | **string**| Profile ID | 

### Return type

[**IdsProfile**](IdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsProfile_0**
> IdsProfile CreateOrUpdateIdsProfile_0(ctx, body, profileId)
create or update IDS profile

Update intrusion detection profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsProfile**](IdsProfile.md)|  | 
  **profileId** | **string**| Profile ID | 

### Return type

[**IdsProfile**](IdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsRule**
> IdsRule CreateOrUpdateIdsRule(ctx, body, domainId, policyId, ruleId)
create or update IDS rule

Update intrusion detection system rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsRule_0**
> IdsRule CreateOrUpdateIdsRule_0(ctx, body, domainId, policyId, ruleId)
create or update IDS rule

Update intrusion detection system rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsSecurityPolicy**
> IdsSecurityPolicy CreateOrUpdateIdsSecurityPolicy(ctx, body, domainId, policyId)
create or update IDS security policy

Update intrusion detection system security policy for a domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsSecurityPolicy_0**
> IdsSecurityPolicy CreateOrUpdateIdsSecurityPolicy_0(ctx, body, domainId, policyId)
create or update IDS security policy

Update intrusion detection system security policy for a domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsStandaloneHostConfig**
> IdsStandaloneHostConfig CreateOrUpdateIdsStandaloneHostConfig(ctx, body)
Create or update IDS configuration

Update intrusion detection system configuration on standalone hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)|  | 

### Return type

[**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateIdsStandaloneHostConfig_0**
> IdsStandaloneHostConfig CreateOrUpdateIdsStandaloneHostConfig_0(ctx, body)
Create or update IDS configuration

Update intrusion detection system configuration on standalone hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)|  | 

### Return type

[**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateRedirectionPolicy**
> RedirectionPolicy CreateOrUpdateRedirectionPolicy(ctx, body, domainId, redirectionPolicyId)
Create or update redirection policy

Create or update the redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionPolicy**](RedirectionPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

[**RedirectionPolicy**](RedirectionPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateRedirectionPolicy_0**
> RedirectionPolicy CreateOrUpdateRedirectionPolicy_0(ctx, body, domainId, redirectionPolicyId)
Create or update redirection policy

Create or update the redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionPolicy**](RedirectionPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

[**RedirectionPolicy**](RedirectionPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateRedirectionRule**
> RedirectionRule CreateOrUpdateRedirectionRule(ctx, body, domainId, redirectionPolicyId, ruleId)
Update redirection rule

Create a rule with the rule-id is not already present, otherwise update the rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionRule**](RedirectionRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
  **ruleId** | **string**| Rule id | 

### Return type

[**RedirectionRule**](RedirectionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateRedirectionRule_0**
> RedirectionRule CreateOrUpdateRedirectionRule_0(ctx, body, domainId, redirectionPolicyId, ruleId)
Update redirection rule

Create a rule with the rule-id is not already present, otherwise update the rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionRule**](RedirectionRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
  **ruleId** | **string**| Rule id | 

### Return type

[**RedirectionRule**](RedirectionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateServiceReference**
> ServiceReference CreateOrUpdateServiceReference(ctx, body, serviceReferenceId)
Create service reference

Create Service Reference representing the intent to consume a given 3rd party service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceReference**](ServiceReference.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 

### Return type

[**ServiceReference**](ServiceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateServiceReference_0**
> ServiceReference CreateOrUpdateServiceReference_0(ctx, body, serviceReferenceId)
Create service reference

Create Service Reference representing the intent to consume a given 3rd party service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceReference**](ServiceReference.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 

### Return type

[**ServiceReference**](ServiceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateVirtualEndpoint**
> VirtualEndpoint CreateOrUpdateVirtualEndpoint(ctx, body, tier0Id, localeServiceId, virtualEndpointId)
Create or update virtual endpoint

Create or update virtual endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualEndpoint**](VirtualEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

[**VirtualEndpoint**](VirtualEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateVirtualEndpoint_0**
> VirtualEndpoint CreateOrUpdateVirtualEndpoint_0(ctx, body, tier0Id, localeServiceId, virtualEndpointId)
Create or update virtual endpoint

Create or update virtual endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualEndpoint**](VirtualEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

[**VirtualEndpoint**](VirtualEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyServiceInstance**
> PolicyServiceInstance CreatePolicyServiceInstance(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create service instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyServiceInstance_0**
> PolicyServiceInstance CreatePolicyServiceInstance_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create service instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyServiceProfile**
> PolicyServiceProfile CreatePolicyServiceProfile(ctx, body, serviceReferenceId, serviceProfileId)
Create or update service profile

Create or update Service profile to specify vendor temp- late attributes for a given 3rd party service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceProfile**](PolicyServiceProfile.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

[**PolicyServiceProfile**](PolicyServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyServiceProfile_0**
> PolicyServiceProfile CreatePolicyServiceProfile_0(ctx, body, serviceReferenceId, serviceProfileId)
Create or update service profile

Create or update Service profile to specify vendor temp- late attributes for a given 3rd party service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceProfile**](PolicyServiceProfile.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

[**PolicyServiceProfile**](PolicyServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceChain**
> PolicyServiceChain CreateServiceChain(ctx, body, serviceChainId)
Create  or update service chain

Create or update Service chain representing the sequence in which 3rd party services must be consumed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceChain**](PolicyServiceChain.md)|  | 
  **serviceChainId** | **string**| Service chain id | 

### Return type

[**PolicyServiceChain**](PolicyServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceChain_0**
> PolicyServiceChain CreateServiceChain_0(ctx, body, serviceChainId)
Create  or update service chain

Create or update Service chain representing the sequence in which 3rd party services must be consumed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceChain**](PolicyServiceChain.md)|  | 
  **serviceChainId** | **string**| Service chain id | 

### Return type

[**PolicyServiceChain**](PolicyServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceDefinition**
> ServiceDefinition CreateServiceDefinition(ctx, body, enforcementPointId)
Create a Service Definition on given enforcement point.

Create a Service Definition on given enforcement point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDefinition**](ServiceDefinition.md)|  | 
  **enforcementPointId** | **string**| Enforcement point id | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceInstanceEndpoint**
> ServiceInstanceEndpoint CreateServiceInstanceEndpoint(ctx, body, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Create service instance endpoint

Create service instance endpoint with given request if not exist. Modification of service instance endpoint is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

[**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceInstanceEndpoint_0**
> ServiceInstanceEndpoint CreateServiceInstanceEndpoint_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Create service instance endpoint

Create service instance endpoint with given request if not exist. Modification of service instance endpoint is not allowed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

[**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTier1PolicyServiceInstance**
> PolicyServiceInstance CreateTier1PolicyServiceInstance(ctx, body, tier1Id, localeServiceId, serviceInstanceId)
Create Tier1 service instance

Create Tier1 service instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTier1PolicyServiceInstance_0**
> PolicyServiceInstance CreateTier1PolicyServiceInstance_0(ctx, body, tier1Id, localeServiceId, serviceInstanceId)
Create Tier1 service instance

Create Tier1 service instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteByodPolicyServiceInstance**
> DeleteByodPolicyServiceInstance(ctx, tier0Id, localeServiceId, serviceInstanceId)
Delete BYOD policy service instance

Delete BYOD policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteByodPolicyServiceInstance_0**
> DeleteByodPolicyServiceInstance_0(ctx, tier0Id, localeServiceId, serviceInstanceId)
Delete BYOD policy service instance

Delete BYOD policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCPUMemThresholdsProfile**
> DeleteCPUMemThresholdsProfile(ctx, profileId)
Delete CPU and memory thresholds profile

Delete CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCPUMemThresholdsProfile_0**
> DeleteCPUMemThresholdsProfile_0(ctx, profileId)
Delete CPU and memory thresholds profile

Delete CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommunicationEntry**
> DeleteCommunicationEntry(ctx, domainId, communicationMapId, communicationEntryId)
Delete CommunicationEntry

Delete CommunicationEntry This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommunicationEntry_0**
> DeleteCommunicationEntry_0(ctx, domainId, communicationMapId, communicationEntryId)
Delete CommunicationEntry

Delete CommunicationEntry This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommunicationMapForDomain**
> DeleteCommunicationMapForDomain(ctx, domainId, communicationMapId)
Deletes a communication map from this domain

Deletes the communication map along with all the communication entries This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommunicationMapForDomain_0**
> DeleteCommunicationMapForDomain_0(ctx, domainId, communicationMapId)
Deletes a communication map from this domain

Deletes the communication map along with all the communication entries This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDnsSecurityProfile**
> DeleteDnsSecurityProfile(ctx, profileId)
Delete DNS security profile

Delete DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDnsSecurityProfileBinding**
> DeleteDnsSecurityProfileBinding(ctx, domainId, groupId, dnsSecurityProfileBindingMapId)
Delete DNS security profile binding map

API will delete DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDnsSecurityProfileBinding_0**
> DeleteDnsSecurityProfileBinding_0(ctx, domainId, groupId, dnsSecurityProfileBindingMapId)
Delete DNS security profile binding map

API will delete DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDnsSecurityProfile_0**
> DeleteDnsSecurityProfile_0(ctx, profileId)
Delete DNS security profile

Delete DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDraft**
> DeleteDraft(ctx, draftId)
Delete a manual draft

Delete a manual draft.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDraft_0**
> DeleteDraft_0(ctx, draftId)
Delete a manual draft

Delete a manual draft.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointPolicy**
> DeleteEndpointPolicy(ctx, domainId, endpointPolicyId)
Delete Endpoint policy

Delete Endpoint policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointPolicy_0**
> DeleteEndpointPolicy_0(ctx, domainId, endpointPolicyId)
Delete Endpoint policy

Delete Endpoint policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointRule**
> DeleteEndpointRule(ctx, domainId, endpointPolicyId, endpointRuleId)
Delete EndpointRule

Delete EndpointRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **endpointPolicyId** | **string**| EndpointPolicy ID | 
  **endpointRuleId** | **string**| EndpointRule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointRule_0**
> DeleteEndpointRule_0(ctx, domainId, endpointPolicyId, endpointRuleId)
Delete EndpointRule

Delete EndpointRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **endpointPolicyId** | **string**| EndpointPolicy ID | 
  **endpointRuleId** | **string**| EndpointRule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFloodProtectionProfile**
> DeleteFloodProtectionProfile(ctx, floodProtectionProfileId)
Delete Flood Protection Profile

API will delete Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFloodProtectionProfile_0**
> DeleteFloodProtectionProfile_0(ctx, floodProtectionProfileId)
Delete Flood Protection Profile

API will delete Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayPolicy**
> DeleteGatewayPolicy(ctx, domainId, gatewayPolicyId)
Delete GatewayPolicy

Delete GatewayPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayPolicy_0**
> DeleteGatewayPolicy_0(ctx, domainId, gatewayPolicyId)
Delete GatewayPolicy

Delete GatewayPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayRule**
> DeleteGatewayRule(ctx, domainId, gatewayPolicyId, ruleId)
Delete rule

Delete rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayRule_0**
> DeleteGatewayRule_0(ctx, domainId, gatewayPolicyId, ruleId)
Delete rule

Delete rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMonitoringBinding**
> DeleteGroupMonitoringBinding(ctx, domainId, groupId, groupMonitoringProfileBindingMapId)
Delete Group Monitoring Profile Binding

API will delete Group Monitoring Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMonitoringBinding_0**
> DeleteGroupMonitoringBinding_0(ctx, domainId, groupId, groupMonitoringProfileBindingMapId)
Delete Group Monitoring Profile Binding

API will delete Group Monitoring Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsProfile**
> DeleteIdsProfile(ctx, profileId)
Delete IDS profile

Delete intrusion detection profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsProfile_0**
> DeleteIdsProfile_0(ctx, profileId)
Delete IDS profile

Delete intrusion detection profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsRule**
> DeleteIdsRule(ctx, domainId, policyId, ruleId)
Delete IDS rule

Delete intrusion detection rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsRule_0**
> DeleteIdsRule_0(ctx, domainId, policyId, ruleId)
Delete IDS rule

Delete intrusion detection rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsSecurityPolicy**
> DeleteIdsSecurityPolicy(ctx, domainId, policyId)
Delete IDS security policy

Delete intrusion detection system security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdsSecurityPolicy_0**
> DeleteIdsSecurityPolicy_0(ctx, domainId, policyId)
Delete IDS security policy

Delete intrusion detection system security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallCPUMemThresholdsProfileBindingMap**
> DeletePolicyFirewallCPUMemThresholdsProfileBindingMap(ctx, cpuMemThresholdsProfileBindingMapId)
Delete Firewall CPU Memory Thresholds Profile Binding

API will delete Firewall CPU Memory Thresholds Profile Binding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallCPUMemThresholdsProfileBindingMap_0**
> DeletePolicyFirewallCPUMemThresholdsProfileBindingMap_0(ctx, cpuMemThresholdsProfileBindingMapId)
Delete Firewall CPU Memory Thresholds Profile Binding

API will delete Firewall CPU Memory Thresholds Profile Binding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallFloodProtectionBinding**
> DeletePolicyFirewallFloodProtectionBinding(ctx, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Delete Firewall Flood Protection Profile Binding

API will delete Firewall Flood Protection Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallFloodProtectionBinding_0**
> DeletePolicyFirewallFloodProtectionBinding_0(ctx, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Delete Firewall Flood Protection Profile Binding

API will delete Firewall Flood Protection Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallScheduler**
> DeletePolicyFirewallScheduler(ctx, firewallSchedulerId, optional)
Delete Policy Firewall Scheduler

Deletes the specified PolicyFirewallScheduler. If scheduler is consumed in a security policy, it won't get deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSchedulerId** | **string**|  | 
 **optional** | ***PolicySecurityApiDeletePolicyFirewallSchedulerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiDeletePolicyFirewallSchedulerOpts struct
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

# **DeletePolicyFirewallScheduler_0**
> DeletePolicyFirewallScheduler_0(ctx, firewallSchedulerId, optional)
Delete Policy Firewall Scheduler

Deletes the specified PolicyFirewallScheduler. If scheduler is consumed in a security policy, it won't get deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSchedulerId** | **string**|  | 
 **optional** | ***PolicySecurityApiDeletePolicyFirewallScheduler_38Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiDeletePolicyFirewallScheduler_38Opts struct
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

# **DeletePolicyFirewallSessionTimerBinding**
> DeletePolicyFirewallSessionTimerBinding(ctx, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Delete Firewall Session Timer Profile Binding

API will delete Firewall Session Timer Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallSessionTimerBinding_0**
> DeletePolicyFirewallSessionTimerBinding_0(ctx, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Delete Firewall Session Timer Profile Binding

API will delete Firewall Session Timer Profile Binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallSessionTimerProfile**
> DeletePolicyFirewallSessionTimerProfile(ctx, firewallSessionTimerProfileId)
Delete Firewall Session Timer Profile

API will delete Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyFirewallSessionTimerProfile_0**
> DeletePolicyFirewallSessionTimerProfile_0(ctx, firewallSessionTimerProfileId)
Delete Firewall Session Timer Profile

API will delete Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceChain**
> DeletePolicyServiceChain(ctx, serviceChainId)
Delete Service chain

This API can be user to delete service chain with given service-chain-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**| Id of Service chain | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceChain_0**
> DeletePolicyServiceChain_0(ctx, serviceChainId)
Delete Service chain

This API can be user to delete service chain with given service-chain-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**| Id of Service chain | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceInstance**
> DeletePolicyServiceInstance(ctx, tier0Id, localeServiceId, serviceInstanceId)
Delete policy service instance

Delete policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceInstanceEndpoint**
> DeletePolicyServiceInstanceEndpoint(ctx, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Delete policy service instance endpoint

Delete policy service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceInstanceEndpoint_0**
> DeletePolicyServiceInstanceEndpoint_0(ctx, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Delete policy service instance endpoint

Delete policy service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceInstance_0**
> DeletePolicyServiceInstance_0(ctx, tier0Id, localeServiceId, serviceInstanceId)
Delete policy service instance

Delete policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceProfile**
> DeletePolicyServiceProfile(ctx, serviceReferenceId, serviceProfileId)
Delete Service profile

This API can be used to delete service profile with given service-profile-id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyServiceProfile_0**
> DeletePolicyServiceProfile_0(ctx, serviceReferenceId, serviceProfileId)
Delete Service profile

This API can be used to delete service profile with given service-profile-id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyUrlCategorizationConfig**
> DeletePolicyUrlCategorizationConfig(ctx, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Delete PolicyUrlCategorizationConfig

Delete PolicyUrlCategorizationConfig. If deleted, the URL categorization will be disabled for that edge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyUrlCategorizationConfig_0**
> DeletePolicyUrlCategorizationConfig_0(ctx, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Delete PolicyUrlCategorizationConfig

Delete PolicyUrlCategorizationConfig. If deleted, the URL categorization will be disabled for that edge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectionPolicy**
> DeleteRedirectionPolicy(ctx, domainId, redirectionPolicyId)
Delete redirection policy

Delete redirection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectionPolicy_0**
> DeleteRedirectionPolicy_0(ctx, domainId, redirectionPolicyId)
Delete redirection policy

Delete redirection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectionRule**
> DeleteRedirectionRule(ctx, domainId, redirectionPolicyId, ruleId)
Delete RedirectionRule

Delete RedirectionRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **redirectionPolicyId** | **string**| Redirection Map ID | 
  **ruleId** | **string**| RedirectionRule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectionRule_0**
> DeleteRedirectionRule_0(ctx, domainId, redirectionPolicyId, ruleId)
Delete RedirectionRule

Delete RedirectionRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **redirectionPolicyId** | **string**| Redirection Map ID | 
  **ruleId** | **string**| RedirectionRule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityPolicyForDomain**
> DeleteSecurityPolicyForDomain(ctx, domainId, securityPolicyId)
Deletes a security policy from this domain

Deletes the security policy along with all the rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityPolicyForDomain_0**
> DeleteSecurityPolicyForDomain_0(ctx, domainId, securityPolicyId)
Deletes a security policy from this domain

Deletes the security policy along with all the rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityRule**
> DeleteSecurityRule(ctx, domainId, securityPolicyId, ruleId)
Delete rule

Delete rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecurityRule_0**
> DeleteSecurityRule_0(ctx, domainId, securityPolicyId, ruleId)
Delete rule

Delete rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceDefinition**
> DeleteServiceDefinition(ctx, enforcementPointId, serviceDefinitionId)
Delete an existing Service Definition on the given enforcement point 

Delete an existing Service Definition on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointId** | **string**| Enforcement point id | 
  **serviceDefinitionId** | **string**| Id of service definition | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceReference**
> DeleteServiceReference(ctx, serviceReferenceId, optional)
Delete Service Reference

This API can be used to delete a service reference with the given service-reference-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
 **optional** | ***PolicySecurityApiDeleteServiceReferenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiDeleteServiceReferenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete all children associated with service reference | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceReference_0**
> DeleteServiceReference_0(ctx, serviceReferenceId, optional)
Delete Service Reference

This API can be used to delete a service reference with the given service-reference-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
 **optional** | ***PolicySecurityApiDeleteServiceReference_50Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiDeleteServiceReference_50Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete all children associated with service reference | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0FloodProtectionProfileBinding**
> DeleteTier0FloodProtectionProfileBinding(ctx, tier0Id, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-0 Logical Router

API will delete Flood Protection Profile Binding for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0FloodProtectionProfileBinding_0**
> DeleteTier0FloodProtectionProfileBinding_0(ctx, tier0Id, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-0 Logical Router

API will delete Flood Protection Profile Binding for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0LocaleServicesFloodProtectionProfileBinding**
> DeleteTier0LocaleServicesFloodProtectionProfileBinding(ctx, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices

API will delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0LocaleServicesFloodProtectionProfileBinding_0**
> DeleteTier0LocaleServicesFloodProtectionProfileBinding_0(ctx, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices

API will delete Flood Protection Profile Binding for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0LocaleServicesSessionTimerProfileBinding**
> DeleteTier0LocaleServicesSessionTimerProfileBinding(ctx, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices

API will delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0LocaleServicesSessionTimerProfileBinding_0**
> DeleteTier0LocaleServicesSessionTimerProfileBinding_0(ctx, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices

API will delete Session Timer Profile Binding for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0SessionTimerProfileBinding**
> DeleteTier0SessionTimerProfileBinding(ctx, tier0Id, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-0 Logical Router

API will delete Session Timer Profile Binding for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier0SessionTimerProfileBinding_0**
> DeleteTier0SessionTimerProfileBinding_0(ctx, tier0Id, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-0 Logical Router

API will delete Session Timer Profile Binding for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1FloodProtectionProfileBinding**
> DeleteTier1FloodProtectionProfileBinding(ctx, tier1Id, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-1 Logical Router

API will delete Flood Protection Profile Binding for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1FloodProtectionProfileBinding_0**
> DeleteTier1FloodProtectionProfileBinding_0(ctx, tier1Id, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-1 Logical Router

API will delete Flood Protection Profile Binding for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1LocaleServicesFloodProtectionProfileBinding**
> DeleteTier1LocaleServicesFloodProtectionProfileBinding(ctx, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices

API will delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1LocaleServicesFloodProtectionProfileBinding_0**
> DeleteTier1LocaleServicesFloodProtectionProfileBinding_0(ctx, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices

API will delete Flood Protection Profile Binding for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1LocaleServicesSessionTimerProfileBinding**
> DeleteTier1LocaleServicesSessionTimerProfileBinding(ctx, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices

API will delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1LocaleServicesSessionTimerProfileBinding_0**
> DeleteTier1LocaleServicesSessionTimerProfileBinding_0(ctx, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices

API will delete Session Timer Profile Binding for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1PolicyServiceInstance**
> DeleteTier1PolicyServiceInstance(ctx, tier1Id, localeServiceId, serviceInstanceId)
Delete Tier1 policy service instance

Delete Tier1 policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1PolicyServiceInstance_0**
> DeleteTier1PolicyServiceInstance_0(ctx, tier1Id, localeServiceId, serviceInstanceId)
Delete Tier1 policy service instance

Delete Tier1 policy service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1SessionTimerProfileBinding**
> DeleteTier1SessionTimerProfileBinding(ctx, tier1Id, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-1 Logical Router

API will delete Session Timer Profile Binding for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTier1SessionTimerProfileBinding_0**
> DeleteTier1SessionTimerProfileBinding_0(ctx, tier1Id, sessionTimerProfileBindingId)
Delete Session Timer Profile Binding for Tier-1 Logical Router

API will delete Session Timer Profile Binding for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualEndpoint**
> DeleteVirtualEndpoint(ctx, tier0Id, localeServiceId, virtualEndpointId)
Delete virtual endpoint

Delete virtual endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualEndpoint_0**
> DeleteVirtualEndpoint_0(ctx, tier0Id, localeServiceId, virtualEndpointId)
Delete virtual endpoint

Delete virtual endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterFirewallExcludeListFilter**
> PolicyResourceReference FilterFirewallExcludeListFilter(ctx, intentPath, optional)
Filter the firewall exclude list

Filter the firewall exclude list by the given object, to check whether the object is a member of this exclude list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Path of the intent object to be searched in the exclude list | 
 **optional** | ***PolicySecurityApiFilterFirewallExcludeListFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiFilterFirewallExcludeListFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deepCheck** | **optional.Bool**| Check all parents | [default to false]
 **enforcementPointPath** | **optional.String**| Path of the enforcement point | 

### Return type

[**PolicyResourceReference**](PolicyResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilterFirewallExcludeListFilter_0**
> PolicyResourceReference FilterFirewallExcludeListFilter_0(ctx, intentPath, optional)
Filter the firewall exclude list

Filter the firewall exclude list by the given object, to check whether the object is a member of this exclude list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intentPath** | **string**| Path of the intent object to be searched in the exclude list | 
 **optional** | ***PolicySecurityApiFilterFirewallExcludeListFilter_61Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiFilterFirewallExcludeListFilter_61Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deepCheck** | **optional.Bool**| Check all parents | [default to false]
 **enforcementPointPath** | **optional.String**| Path of the enforcement point | 

### Return type

[**PolicyResourceReference**](PolicyResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedConfigurationToBePublishedForDraft**
> Infra GetAggregatedConfigurationToBePublishedForDraft(ctx, draftId)
Get an aggregated configuration for the draft

Get an aggregated configuration that will get applied onto current configuration during publish of this draft. The response is a hierarichal payload containing the aggregated configuration differences from the latest auto draft till the specified draft. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedConfigurationToBePublishedForDraft_0**
> Infra GetAggregatedConfigurationToBePublishedForDraft_0(ctx, draftId)
Get an aggregated configuration for the draft

Get an aggregated configuration that will get applied onto current configuration during publish of this draft. The response is a hierarichal payload containing the aggregated configuration differences from the latest auto draft till the specified draft. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeClusterIdfwConfiguration**
> ComputeClusterIdfwConfiguration GetComputeClusterIdfwConfiguration(ctx, clusterId)
Read compute cluster idfw configuration

Read compute cluster identity firewall configuration 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| Cluster ID | 

### Return type

[**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeClusterIdfwConfiguration_0**
> ComputeClusterIdfwConfiguration GetComputeClusterIdfwConfiguration_0(ctx, clusterId)
Read compute cluster idfw configuration

Read compute cluster identity firewall configuration 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| Cluster ID | 

### Return type

[**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDfwFirewallConfiguration**
> DfwFirewallConfiguration GetDfwFirewallConfiguration(ctx, )
Get dfw firewall configuration

Get the current dfw firewall configurations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDfwFirewallConfiguration_0**
> DfwFirewallConfiguration GetDfwFirewallConfiguration_0(ctx, )
Get dfw firewall configuration

Get the current dfw firewall configurations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsSecurityProfileBinding**
> DnsSecurityProfileBindingMap GetDnsSecurityProfileBinding(ctx, domainId, groupId, dnsSecurityProfileBindingMapId)
Get DNS security profile binding map

API will get DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

[**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsSecurityProfileBinding_0**
> DnsSecurityProfileBindingMap GetDnsSecurityProfileBinding_0(ctx, domainId, groupId, dnsSecurityProfileBindingMapId)
Get DNS security profile binding map

API will get DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

[**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallExcludeList**
> PolicyExcludeList GetFirewallExcludeList(ctx, )
Read security policy exclude list

Read exclude list for firewall 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PolicyExcludeList**](PolicyExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallExcludeList_0**
> PolicyExcludeList GetFirewallExcludeList_0(ctx, )
Read security policy exclude list

Read exclude list for firewall 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PolicyExcludeList**](PolicyExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFloodProtectionProfile**
> FloodProtectionProfile GetFloodProtectionProfile(ctx, floodProtectionProfileId)
Get Flood Protection Profile

API will get Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

[**FloodProtectionProfile**](FloodProtectionProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFloodProtectionProfile_0**
> FloodProtectionProfile GetFloodProtectionProfile_0(ctx, floodProtectionProfileId)
Get Flood Protection Profile

API will get Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

[**FloodProtectionProfile**](FloodProtectionProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPolicyStatistics**
> SecurityPolicyStatisticsListResult GetGatewayPolicyStatistics(ctx, domainId, gatewayPolicyId, optional)
Get gateway policy statistics

Get statistics of a gateay policy. - no enforcement point path specified: Stats will be evaluated on each enforcement. point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiGetGatewayPolicyStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetGatewayPolicyStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**SecurityPolicyStatisticsListResult**](SecurityPolicyStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPolicyStatistics_0**
> SecurityPolicyStatisticsListResult GetGatewayPolicyStatistics_0(ctx, domainId, gatewayPolicyId, optional)
Get gateway policy statistics

Get statistics of a gateay policy. - no enforcement point path specified: Stats will be evaluated on each enforcement. point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiGetGatewayPolicyStatistics_68Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetGatewayPolicyStatistics_68Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**SecurityPolicyStatisticsListResult**](SecurityPolicyStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayRuleStatistics**
> RuleStatisticsListResult GetGatewayRuleStatistics(ctx, domainId, gatewayPolicyId, ruleId, optional)
Get gateway rule statistics

Get statistics of a gateway rule. - no enforcement point path specified: Stats will be evaluated on each enforcement. point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiGetGatewayRuleStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetGatewayRuleStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**RuleStatisticsListResult**](RuleStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayRuleStatistics_0**
> RuleStatisticsListResult GetGatewayRuleStatistics_0(ctx, domainId, gatewayPolicyId, ruleId, optional)
Get gateway rule statistics

Get statistics of a gateway rule. - no enforcement point path specified: Stats will be evaluated on each enforcement. point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiGetGatewayRuleStatistics_69Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetGatewayRuleStatistics_69Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**RuleStatisticsListResult**](RuleStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupMonitoringBinding**
> GroupMonitoringProfileBindingMap GetGroupMonitoringBinding(ctx, domainId, groupId, groupMonitoringProfileBindingMapId)
Get Group Monitoring Profile Binding Map

API will get Group Monitoring Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

[**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupMonitoringBinding_0**
> GroupMonitoringProfileBindingMap GetGroupMonitoringBinding_0(ctx, domainId, groupId, groupMonitoringProfileBindingMapId)
Get Group Monitoring Profile Binding Map

API will get Group Monitoring Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

[**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsClusterConfig**
> IdsClusterConfig GetIdsClusterConfig(ctx, clusterId)
Read IDS cluster config.

Read intrusion detection system cluster config 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| User entered ID | 

### Return type

[**IdsClusterConfig**](IdsClusterConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsClusterConfig_0**
> IdsClusterConfig GetIdsClusterConfig_0(ctx, clusterId)
Read IDS cluster config.

Read intrusion detection system cluster config 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| User entered ID | 

### Return type

[**IdsClusterConfig**](IdsClusterConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsProfile**
> IdsProfile GetIdsProfile(ctx, profileId)
Get IDS profile.

Read intrusion detection profile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Profile ID | 

### Return type

[**IdsProfile**](IdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsProfile_0**
> IdsProfile GetIdsProfile_0(ctx, profileId)
Get IDS profile.

Read intrusion detection profile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Profile ID | 

### Return type

[**IdsProfile**](IdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsRule**
> IdsRule GetIdsRule(ctx, domainId, policyId, ruleId)
Get IDS rule.

Read intrusion detection rule 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsRule_0**
> IdsRule GetIdsRule_0(ctx, domainId, policyId, ruleId)
Get IDS rule.

Read intrusion detection rule 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSecurityPolicy**
> IdsSecurityPolicy GetIdsSecurityPolicy(ctx, domainId, policyId)
Get IDS security policy.

Read intrusion detection system security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSecurityPolicy_0**
> IdsSecurityPolicy GetIdsSecurityPolicy_0(ctx, domainId, policyId)
Get IDS security policy.

Read intrusion detection system security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSettings**
> IdsSettings GetIdsSettings(ctx, )
Get IDS system settings

Intrusion detection system settings. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsSettings**](IdsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSettings_0**
> IdsSettings GetIdsSettings_0(ctx, )
Get IDS system settings

Intrusion detection system settings. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsSettings**](IdsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSignatureStatus**
> IdsSignatureStatus GetIdsSignatureStatus(ctx, )
Get IDS signature status

Intrusion detection system signatures status. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsSignatureStatus**](IdsSignatureStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSignatureStatus_0**
> IdsSignatureStatus GetIdsSignatureStatus_0(ctx, )
Get IDS signature status

Intrusion detection system signatures status. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsSignatureStatus**](IdsSignatureStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSignatureVersions**
> IdsSignatureVersionListResult GetIdsSignatureVersions(ctx, optional)
Get IDS signature versions

Intrusion detection system signature versions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiGetIdsSignatureVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetIdsSignatureVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSignatureVersionListResult**](IdsSignatureVersionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsSignatureVersions_0**
> IdsSignatureVersionListResult GetIdsSignatureVersions_0(ctx, optional)
Get IDS signature versions

Intrusion detection system signature versions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiGetIdsSignatureVersions_77Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetIdsSignatureVersions_77Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSignatureVersionListResult**](IdsSignatureVersionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsStandaloneHostConfig**
> IdsStandaloneHostConfig GetIdsStandaloneHostConfig(ctx, )
Read IDS config

Read intrusion detection system config of standalone hosts. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsStandaloneHostConfig_0**
> IdsStandaloneHostConfig GetIdsStandaloneHostConfig_0(ctx, )
Read IDS config

Read intrusion detection system config of standalone hosts. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallCPUMemThresholdsProfileBindingMap**
> PolicyFirewallCpuMemThresholdsProfileBindingMap GetPolicyFirewallCPUMemThresholdsProfileBindingMap(ctx, cpuMemThresholdsProfileBindingMapId)
Get Firewall CPU Memory Thresholds Profile Binding Map

API will get Firewall CPU Memory Thresholds Profile Binding Map. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCPUMemThresholdsProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallCPUMemThresholdsProfileBindingMap_0**
> PolicyFirewallCpuMemThresholdsProfileBindingMap GetPolicyFirewallCPUMemThresholdsProfileBindingMap_0(ctx, cpuMemThresholdsProfileBindingMapId)
Get Firewall CPU Memory Thresholds Profile Binding Map

API will get Firewall CPU Memory Thresholds Profile Binding Map. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCPUMemThresholdsProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallFloodProtectionBinding**
> PolicyFirewallFloodProtectionProfileBindingMap GetPolicyFirewallFloodProtectionBinding(ctx, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Get Firewall Flood Protection Profile Binding Map

API will get Firewall Flood Protection Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallFloodProtectionBinding_0**
> PolicyFirewallFloodProtectionProfileBindingMap GetPolicyFirewallFloodProtectionBinding_0(ctx, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Get Firewall Flood Protection Profile Binding Map

API will get Firewall Flood Protection Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallScheduler**
> PolicyFirewallScheduler GetPolicyFirewallScheduler(ctx, firewallSchedulerId)
Get PolicyFirewallScheduler

Get a PolicyFirewallScheduler by id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSchedulerId** | **string**|  | 

### Return type

[**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallScheduler_0**
> PolicyFirewallScheduler GetPolicyFirewallScheduler_0(ctx, firewallSchedulerId)
Get PolicyFirewallScheduler

Get a PolicyFirewallScheduler by id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSchedulerId** | **string**|  | 

### Return type

[**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallSessionTimerBinding**
> PolicyFirewallSessionTimerProfileBindingMap GetPolicyFirewallSessionTimerBinding(ctx, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Get Firewall Session Timer Profile Binding Map

API will get Firewall Session Timer Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallSessionTimerBinding_0**
> PolicyFirewallSessionTimerProfileBindingMap GetPolicyFirewallSessionTimerBinding_0(ctx, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Get Firewall Session Timer Profile Binding Map

API will get Firewall Session Timer Profile Binding Map 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain-ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallSessionTimerProfile**
> PolicyFirewallSessionTimerProfile GetPolicyFirewallSessionTimerProfile(ctx, firewallSessionTimerProfileId)
Get Firewall Session Timer Profile

API will get Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

[**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyFirewallSessionTimerProfile_0**
> PolicyFirewallSessionTimerProfile GetPolicyFirewallSessionTimerProfile_0(ctx, firewallSessionTimerProfileId)
Get Firewall Session Timer Profile

API will get Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

[**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyServiceInstanceStatistics**
> PolicyServiceInstanceStatistics GetPolicyServiceInstanceStatistics(ctx, tier0Id, localeServiceId, serviceInstanceId, optional)
Get statistics for all runtimes associated with this PolicyServiceInstance

Get statistics for all data NICs on all runtimes associated with this PolicyServiceInstance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
 **optional** | ***PolicySecurityApiGetPolicyServiceInstanceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetPolicyServiceInstanceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**PolicyServiceInstanceStatistics**](PolicyServiceInstanceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyServiceInstanceStatistics_0**
> PolicyServiceInstanceStatistics GetPolicyServiceInstanceStatistics_0(ctx, tier0Id, localeServiceId, serviceInstanceId, optional)
Get statistics for all runtimes associated with this PolicyServiceInstance

Get statistics for all data NICs on all runtimes associated with this PolicyServiceInstance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
 **optional** | ***PolicySecurityApiGetPolicyServiceInstanceStatistics_84Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetPolicyServiceInstanceStatistics_84Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**PolicyServiceInstanceStatistics**](PolicyServiceInstanceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyServiceProfileGroups**
> ServiceProfileGroups GetPolicyServiceProfileGroups(ctx, serviceReferenceId, serviceProfileId, optional)
Get Groups used in Redirection rules for a given Service Profile.

List of Groups used in Redirection rules for a given Service Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 
 **optional** | ***PolicySecurityApiGetPolicyServiceProfileGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetPolicyServiceProfileGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**ServiceProfileGroups**](ServiceProfileGroups.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyServiceProfileGroups_0**
> ServiceProfileGroups GetPolicyServiceProfileGroups_0(ctx, serviceReferenceId, serviceProfileId, optional)
Get Groups used in Redirection rules for a given Service Profile.

List of Groups used in Redirection rules for a given Service Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 
 **optional** | ***PolicySecurityApiGetPolicyServiceProfileGroups_85Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetPolicyServiceProfileGroups_85Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**ServiceProfileGroups**](ServiceProfileGroups.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyUrlCategorizationConfig**
> PolicyUrlCategorizationConfig GetPolicyUrlCategorizationConfig(ctx, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Get PolicyUrlCategorizationConfig

Gets a PolicyUrlCategorizationConfig. This returns the details of the config like whether the URL categorization is enabled or disabled, the id of the context profiles which are used to filter the categories, and the update frequency of the data from the cloud. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyUrlCategorizationConfig_0**
> PolicyUrlCategorizationConfig GetPolicyUrlCategorizationConfig_0(ctx, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Get PolicyUrlCategorizationConfig

Gets a PolicyUrlCategorizationConfig. This returns the details of the config like whether the URL categorization is enabled or disabled, the id of the context profiles which are used to filter the categories, and the update frequency of the data from the cloud. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreviewOfConfigurationAfterPublishOfDraft**
> Infra GetPreviewOfConfigurationAfterPublishOfDraft(ctx, draftId)
Get a preview of a configuration after publish of a draft

Get a preview of a configuration which will be present after publish of a specified draft. The response essentially is a hierarichal payload containing the configuration, which will be in active after a specified draft gets published onto current configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreviewOfConfigurationAfterPublishOfDraft_0**
> Infra GetPreviewOfConfigurationAfterPublishOfDraft_0(ctx, draftId)
Get a preview of a configuration after publish of a draft

Get a preview of a configuration which will be present after publish of a specified draft. The response essentially is a hierarichal payload containing the configuration, which will be in active after a specified draft gets published onto current configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**Infra**](Infra.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleStatistics**
> RuleStatisticsListResult GetRuleStatistics(ctx, domainId, securityPolicyId, ruleId, optional)
Get rule statistics

Get statistics of a rule. - no enforcement point path specified: Stats will be evaluated on each enforcement point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **securityPolicyId** | **string**| Security policy id | 
  **ruleId** | **string**| Rule id | 
 **optional** | ***PolicySecurityApiGetRuleStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetRuleStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**RuleStatisticsListResult**](RuleStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleStatistics_0**
> RuleStatisticsListResult GetRuleStatistics_0(ctx, domainId, securityPolicyId, ruleId, optional)
Get rule statistics

Get statistics of a rule. - no enforcement point path specified: Stats will be evaluated on each enforcement point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **securityPolicyId** | **string**| Security policy id | 
  **ruleId** | **string**| Rule id | 
 **optional** | ***PolicySecurityApiGetRuleStatistics_88Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetRuleStatistics_88Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**RuleStatisticsListResult**](RuleStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityPolicyStatistics**
> SecurityPolicyStatisticsListResult GetSecurityPolicyStatistics(ctx, domainId, securityPolicyId, optional)
Get security policy statistics

Get statistics of a security policy. - no enforcement point path specified: Stats will be evaluated on each enforcement point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **securityPolicyId** | **string**| Security policy id | 
 **optional** | ***PolicySecurityApiGetSecurityPolicyStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetSecurityPolicyStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**SecurityPolicyStatisticsListResult**](SecurityPolicyStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityPolicyStatistics_0**
> SecurityPolicyStatisticsListResult GetSecurityPolicyStatistics_0(ctx, domainId, securityPolicyId, optional)
Get security policy statistics

Get statistics of a security policy. - no enforcement point path specified: Stats will be evaluated on each enforcement point. - {enforcement_point_path}: Stats are evaluated only on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **securityPolicyId** | **string**| Security policy id | 
 **optional** | ***PolicySecurityApiGetSecurityPolicyStatistics_89Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetSecurityPolicyStatistics_89Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**SecurityPolicyStatisticsListResult**](SecurityPolicyStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandaloneHostIdfwConfiguration**
> StandaloneHostIdfwConfiguration GetStandaloneHostIdfwConfiguration(ctx, )
Read idfw configuration for standalone host

Read identity firewall configuration for standalone host 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandaloneHostIdfwConfiguration_0**
> StandaloneHostIdfwConfiguration GetStandaloneHostIdfwConfiguration_0(ctx, )
Read idfw configuration for standalone host

Read identity firewall configuration for standalone host 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0FloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap GetTier0FloodProtectionProfileBinding(ctx, tier0Id, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-0 Logical Router

API will get Flood Protection Profile Binding Map for Tier-0 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0FloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap GetTier0FloodProtectionProfileBinding_0(ctx, tier0Id, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-0 Logical Router

API will get Flood Protection Profile Binding Map for Tier-0 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0LocaleServicesFloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap GetTier0LocaleServicesFloodProtectionProfileBinding(ctx, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0LocaleServicesFloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap GetTier0LocaleServicesFloodProtectionProfileBinding_0(ctx, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will get Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0LocaleServicesSessionTimerProfileBinding**
> SessionTimerProfileBindingMap GetTier0LocaleServicesSessionTimerProfileBinding(ctx, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0LocaleServicesSessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap GetTier0LocaleServicesSessionTimerProfileBinding_0(ctx, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will get Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0SessionTimerProfileBinding**
> SessionTimerProfileBindingMap GetTier0SessionTimerProfileBinding(ctx, tier0Id, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-0 Logical Router

API will get Session Timer Profile Binding Map for Tier-0 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier0SessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap GetTier0SessionTimerProfileBinding_0(ctx, tier0Id, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-0 Logical Router

API will get Session Timer Profile Binding Map for Tier-0 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1FloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap GetTier1FloodProtectionProfileBinding(ctx, tier1Id, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-1 Logical Router

API will get Flood Protection Profile Binding Map for Tier-1 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1FloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap GetTier1FloodProtectionProfileBinding_0(ctx, tier1Id, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-1 Logical Router

API will get Flood Protection Profile Binding Map for Tier-1 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1LocaleServicesFloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap GetTier1LocaleServicesFloodProtectionProfileBinding(ctx, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1LocaleServicesFloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap GetTier1LocaleServicesFloodProtectionProfileBinding_0(ctx, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will get Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1LocaleServicesSessionTimerProfileBinding**
> SessionTimerProfileBindingMap GetTier1LocaleServicesSessionTimerProfileBinding(ctx, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1LocaleServicesSessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap GetTier1LocaleServicesSessionTimerProfileBinding_0(ctx, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will get Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1PolicyServiceInstanceStatistics**
> PolicyServiceInstanceStatistics GetTier1PolicyServiceInstanceStatistics(ctx, tier1Id, localeServiceId, serviceInstanceId, optional)
Get statistics for all runtimes associated with this Tier1 PolicyServiceInstance

Get statistics for all data NICs on all runtimes associated with this Tier1 PolicyServiceInstance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 
 **optional** | ***PolicySecurityApiGetTier1PolicyServiceInstanceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetTier1PolicyServiceInstanceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**PolicyServiceInstanceStatistics**](PolicyServiceInstanceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1PolicyServiceInstanceStatistics_0**
> PolicyServiceInstanceStatistics GetTier1PolicyServiceInstanceStatistics_0(ctx, tier1Id, localeServiceId, serviceInstanceId, optional)
Get statistics for all runtimes associated with this Tier1 PolicyServiceInstance

Get statistics for all data NICs on all runtimes associated with this Tier1 PolicyServiceInstance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 
 **optional** | ***PolicySecurityApiGetTier1PolicyServiceInstanceStatistics_98Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiGetTier1PolicyServiceInstanceStatistics_98Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**PolicyServiceInstanceStatistics**](PolicyServiceInstanceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1SessionTimerProfileBinding**
> SessionTimerProfileBindingMap GetTier1SessionTimerProfileBinding(ctx, tier1Id, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-1 Logical Router

API will get Session Timer Profile Binding Map for Tier-1 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTier1SessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap GetTier1SessionTimerProfileBinding_0(ctx, tier1Id, sessionTimerProfileBindingId)
Get Session Timer Profile Binding Map for Tier-1 Logical Router

API will get Session Timer Profile Binding Map for Tier-1 Logical Router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByodPolicyServiceInstancesForTier0**
> ByodPolicyServiceInstanceListResult ListByodPolicyServiceInstancesForTier0(ctx, tier0Id, localeServiceId, optional)
Read all BYOD service instance objects under a tier-0

Read all BYOD service instance objects under a tier-0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiListByodPolicyServiceInstancesForTier0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListByodPolicyServiceInstancesForTier0Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ByodPolicyServiceInstanceListResult**](ByodPolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByodPolicyServiceInstancesForTier0_0**
> ByodPolicyServiceInstanceListResult ListByodPolicyServiceInstancesForTier0_0(ctx, tier0Id, localeServiceId, optional)
Read all BYOD service instance objects under a tier-0

Read all BYOD service instance objects under a tier-0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiListByodPolicyServiceInstancesForTier0_100Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListByodPolicyServiceInstancesForTier0_100Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ByodPolicyServiceInstanceListResult**](ByodPolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCPUMemThresholdsProfiles**
> PolicyFirewallCpuMemThresholdsProfileListResult ListCPUMemThresholdsProfiles(ctx, optional)
List all CPU and memory thresholds profiles

List all CPU and memory thresholds profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListCPUMemThresholdsProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCPUMemThresholdsProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileListResult**](PolicyFirewallCpuMemThresholdsProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCPUMemThresholdsProfiles_0**
> PolicyFirewallCpuMemThresholdsProfileListResult ListCPUMemThresholdsProfiles_0(ctx, optional)
List all CPU and memory thresholds profiles

List all CPU and memory thresholds profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListCPUMemThresholdsProfiles_101Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCPUMemThresholdsProfiles_101Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileListResult**](PolicyFirewallCpuMemThresholdsProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommunicationEntry**
> CommunicationEntryListResult ListCommunicationEntry(ctx, domainId, communicationMapId, optional)
List CommunicationEntries

List CommunicationEntries This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
 **optional** | ***PolicySecurityApiListCommunicationEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCommunicationEntryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CommunicationEntryListResult**](CommunicationEntryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommunicationEntry_0**
> CommunicationEntryListResult ListCommunicationEntry_0(ctx, domainId, communicationMapId, optional)
List CommunicationEntries

List CommunicationEntries This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
 **optional** | ***PolicySecurityApiListCommunicationEntry_102Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCommunicationEntry_102Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CommunicationEntryListResult**](CommunicationEntryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommunicationMapsForDomain**
> CommunicationMapListResult ListCommunicationMapsForDomain(ctx, domainId, optional)
List communication maps

List all communication maps for a domain. This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListCommunicationMapsForDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCommunicationMapsForDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CommunicationMapListResult**](CommunicationMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommunicationMapsForDomain_0**
> CommunicationMapListResult ListCommunicationMapsForDomain_0(ctx, domainId, optional)
List communication maps

List all communication maps for a domain. This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListCommunicationMapsForDomain_103Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListCommunicationMapsForDomain_103Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CommunicationMapListResult**](CommunicationMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeClusterIdfwConfiguration**
> ComputeClusterIdfwConfigurationListResult ListComputeClusterIdfwConfiguration(ctx, optional)
List compute cluster idfw Configuration

API will list all compute cluster wise identity firewall configuration 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListComputeClusterIdfwConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListComputeClusterIdfwConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ComputeClusterIdfwConfigurationListResult**](ComputeClusterIdfwConfigurationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeClusterIdfwConfiguration_0**
> ComputeClusterIdfwConfigurationListResult ListComputeClusterIdfwConfiguration_0(ctx, optional)
List compute cluster idfw Configuration

API will list all compute cluster wise identity firewall configuration 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListComputeClusterIdfwConfiguration_104Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListComputeClusterIdfwConfiguration_104Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ComputeClusterIdfwConfigurationListResult**](ComputeClusterIdfwConfigurationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDnsSecurityProfileBindings**
> DnsSecurityProfileBindingMapListResult ListDnsSecurityProfileBindings(ctx, domainId, groupId, optional)
Get DNS security profile binding map

API will get DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListDnsSecurityProfileBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDnsSecurityProfileBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DnsSecurityProfileBindingMapListResult**](DnsSecurityProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDnsSecurityProfileBindings_0**
> DnsSecurityProfileBindingMapListResult ListDnsSecurityProfileBindings_0(ctx, domainId, groupId, optional)
Get DNS security profile binding map

API will get DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListDnsSecurityProfileBindings_105Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDnsSecurityProfileBindings_105Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DnsSecurityProfileBindingMapListResult**](DnsSecurityProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDnsSecurityProfiles**
> DnsSecurityProfileListResult ListDnsSecurityProfiles(ctx, optional)
List all DNS security profiles

List all DNS security profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListDnsSecurityProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDnsSecurityProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DnsSecurityProfileListResult**](DnsSecurityProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDnsSecurityProfiles_0**
> DnsSecurityProfileListResult ListDnsSecurityProfiles_0(ctx, optional)
List all DNS security profiles

List all DNS security profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListDnsSecurityProfiles_106Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDnsSecurityProfiles_106Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DnsSecurityProfileListResult**](DnsSecurityProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDrafts**
> PolicyDraftListResult ListDrafts(ctx, optional)
List policy drafts

List policy drafts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListDraftsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDraftsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoDrafts** | **optional.Bool**| Fetch list of draft based on is_auto_draft flag | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyDraftListResult**](PolicyDraftListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDrafts_0**
> PolicyDraftListResult ListDrafts_0(ctx, optional)
List policy drafts

List policy drafts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListDrafts_107Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListDrafts_107Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoDrafts** | **optional.Bool**| Fetch list of draft based on is_auto_draft flag | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyDraftListResult**](PolicyDraftListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEndpointPoliciesAcrossAllDomains**
> EndpointPolicyListResult ListEndpointPoliciesAcrossAllDomains(ctx, optional)
List Endpoint policies

List all Endpoint policies across all domains ordered by precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListEndpointPoliciesAcrossAllDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListEndpointPoliciesAcrossAllDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EndpointPolicyListResult**](EndpointPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEndpointPoliciesAcrossAllDomains_0**
> EndpointPolicyListResult ListEndpointPoliciesAcrossAllDomains_0(ctx, optional)
List Endpoint policies

List all Endpoint policies across all domains ordered by precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListEndpointPoliciesAcrossAllDomains_108Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListEndpointPoliciesAcrossAllDomains_108Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EndpointPolicyListResult**](EndpointPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEndpointRule**
> EndpointRuleListResult ListEndpointRule(ctx, domainId, endpointPolicyId, optional)
List Endpoint rules

List Endpoint rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
 **optional** | ***PolicySecurityApiListEndpointRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListEndpointRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EndpointRuleListResult**](EndpointRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEndpointRule_0**
> EndpointRuleListResult ListEndpointRule_0(ctx, domainId, endpointPolicyId, optional)
List Endpoint rules

List Endpoint rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
 **optional** | ***PolicySecurityApiListEndpointRule_109Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListEndpointRule_109Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EndpointRuleListResult**](EndpointRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallFloodProtectionBindingsAcrossDomains**
> PolicyFirewallFloodProtectionProfileBindingMapListResult ListFirewallFloodProtectionBindingsAcrossDomains(ctx, optional)
List Firewall Flood Protection Profile Binding Maps for all domains

API will list all Firewall Flood Protection Profile Binding Maps across all domains. This API returns the binding maps order by the sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFirewallFloodProtectionBindingsAcrossDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFirewallFloodProtectionBindingsAcrossDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMapListResult**](PolicyFirewallFloodProtectionProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallFloodProtectionBindingsAcrossDomains_0**
> PolicyFirewallFloodProtectionProfileBindingMapListResult ListFirewallFloodProtectionBindingsAcrossDomains_0(ctx, optional)
List Firewall Flood Protection Profile Binding Maps for all domains

API will list all Firewall Flood Protection Profile Binding Maps across all domains. This API returns the binding maps order by the sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFirewallFloodProtectionBindingsAcrossDomains_110Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFirewallFloodProtectionBindingsAcrossDomains_110Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMapListResult**](PolicyFirewallFloodProtectionProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallSessionTimerBindingsAcrossDomains**
> PolicyFirewallSessionTimerProfileBindingMapListResult ListFirewallSessionTimerBindingsAcrossDomains(ctx, optional)
List Firewall Session Timer Profile Binding Maps for all domains

API will list all Firewall Session Timer Profile Binding Maps across all domains. This API returns the binding maps order by the sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFirewallSessionTimerBindingsAcrossDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFirewallSessionTimerBindingsAcrossDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMapListResult**](PolicyFirewallSessionTimerProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallSessionTimerBindingsAcrossDomains_0**
> PolicyFirewallSessionTimerProfileBindingMapListResult ListFirewallSessionTimerBindingsAcrossDomains_0(ctx, optional)
List Firewall Session Timer Profile Binding Maps for all domains

API will list all Firewall Session Timer Profile Binding Maps across all domains. This API returns the binding maps order by the sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFirewallSessionTimerBindingsAcrossDomains_111Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFirewallSessionTimerBindingsAcrossDomains_111Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMapListResult**](PolicyFirewallSessionTimerProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFloodProtectionProfileBindings**
> FloodProtectionProfileBindingListResult ListFloodProtectionProfileBindings(ctx, floodProtectionProfileId, optional)
List Flood Protection Profiles

API will list all Flood Protection Profiles bindings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**|  | 
 **optional** | ***PolicySecurityApiListFloodProtectionProfileBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFloodProtectionProfileBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FloodProtectionProfileBindingListResult**](FloodProtectionProfileBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFloodProtectionProfileBindings_0**
> FloodProtectionProfileBindingListResult ListFloodProtectionProfileBindings_0(ctx, floodProtectionProfileId, optional)
List Flood Protection Profiles

API will list all Flood Protection Profiles bindings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **floodProtectionProfileId** | **string**|  | 
 **optional** | ***PolicySecurityApiListFloodProtectionProfileBindings_112Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFloodProtectionProfileBindings_112Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FloodProtectionProfileBindingListResult**](FloodProtectionProfileBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFloodProtectionProfiles**
> FloodProtectionProfileListResult ListFloodProtectionProfiles(ctx, optional)
List Flood Protection Profiles

API will list all Flood Protection Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFloodProtectionProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFloodProtectionProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FloodProtectionProfileListResult**](FloodProtectionProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFloodProtectionProfiles_0**
> FloodProtectionProfileListResult ListFloodProtectionProfiles_0(ctx, optional)
List Flood Protection Profiles

API will list all Flood Protection Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListFloodProtectionProfiles_113Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListFloodProtectionProfiles_113Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FloodProtectionProfileListResult**](FloodProtectionProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGatewayPoliciesForDomain**
> GatewayPolicyListResult ListGatewayPoliciesForDomain(ctx, domainId, optional)
List gateway policies

List all gateway policies for specified Domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGatewayPoliciesForDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGatewayPoliciesForDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGatewayPoliciesForDomain_0**
> GatewayPolicyListResult ListGatewayPoliciesForDomain_0(ctx, domainId, optional)
List gateway policies

List all gateway policies for specified Domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGatewayPoliciesForDomain_114Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGatewayPoliciesForDomain_114Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGatewayRules**
> RuleListResult ListGatewayRules(ctx, domainId, gatewayPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGatewayRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGatewayRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RuleListResult**](RuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGatewayRules_0**
> RuleListResult ListGatewayRules_0(ctx, domainId, gatewayPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGatewayRules_115Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGatewayRules_115Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RuleListResult**](RuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupMonitoringBindings**
> GroupMonitoringProfileBindingMapListResult ListGroupMonitoringBindings(ctx, domainId, groupId, optional)
List Group Monitoring Profile Binding Maps

API will list all Group Monitoring Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGroupMonitoringBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGroupMonitoringBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupMonitoringProfileBindingMapListResult**](GroupMonitoringProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupMonitoringBindings_0**
> GroupMonitoringProfileBindingMapListResult ListGroupMonitoringBindings_0(ctx, domainId, groupId, optional)
List Group Monitoring Profile Binding Maps

API will list all Group Monitoring Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListGroupMonitoringBindings_116Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListGroupMonitoringBindings_116Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**GroupMonitoringProfileBindingMapListResult**](GroupMonitoringProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsClusterConfigs**
> IdsClusterConfigListResult ListIdsClusterConfigs(ctx, optional)
List IDS cluster configs

List intrusion detection system cluster configs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListIdsClusterConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsClusterConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsClusterConfigListResult**](IdsClusterConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsClusterConfigs_0**
> IdsClusterConfigListResult ListIdsClusterConfigs_0(ctx, optional)
List IDS cluster configs

List intrusion detection system cluster configs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListIdsClusterConfigs_117Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsClusterConfigs_117Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsClusterConfigListResult**](IdsClusterConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsProfiles**
> IdsProfileListResult ListIdsProfiles(ctx, optional)
List IDS profiles

List intrusion detection profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListIdsProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsProfileListResult**](IdsProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsProfiles_0**
> IdsProfileListResult ListIdsProfiles_0(ctx, optional)
List IDS profiles

List intrusion detection profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListIdsProfiles_118Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsProfiles_118Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsProfileListResult**](IdsProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsRules**
> IdsRuleListResult ListIdsRules(ctx, domainId, policyId, optional)
List IDS rules

List intrusion detection rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
 **optional** | ***PolicySecurityApiListIdsRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsRuleListResult**](IdsRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsRules_0**
> IdsRuleListResult ListIdsRules_0(ctx, domainId, policyId, optional)
List IDS rules

List intrusion detection rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
 **optional** | ***PolicySecurityApiListIdsRules_119Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsRules_119Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsRuleListResult**](IdsRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsSecurityPolicies**
> IdsSecurityPolicyListResult ListIdsSecurityPolicies(ctx, domainId, optional)
List IDS security policies

List intrusion detection system security policies. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
 **optional** | ***PolicySecurityApiListIdsSecurityPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsSecurityPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSecurityPolicyListResult**](IdsSecurityPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsSecurityPolicies_0**
> IdsSecurityPolicyListResult ListIdsSecurityPolicies_0(ctx, domainId, optional)
List IDS security policies

List intrusion detection system security policies. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain ID | 
 **optional** | ***PolicySecurityApiListIdsSecurityPolicies_120Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsSecurityPolicies_120Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSecurityPolicyListResult**](IdsSecurityPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsSignatures**
> IdsSignatureListResult ListIdsSignatures(ctx, versionId, optional)
List IDS signatures

List intrusion detection system signatures. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**|  | 
 **optional** | ***PolicySecurityApiListIdsSignaturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsSignaturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSignatureListResult**](IdsSignatureListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdsSignatures_0**
> IdsSignatureListResult ListIdsSignatures_0(ctx, versionId, optional)
List IDS signatures

List intrusion detection system signatures. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**|  | 
 **optional** | ***PolicySecurityApiListIdsSignatures_121Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListIdsSignatures_121Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdsSignatureListResult**](IdsSignatureListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallCPUMemThresholdsProfileBindingMaps**
> PolicyFirewallCpuMemThresholdsProfileBindingMapListResult ListPolicyFirewallCPUMemThresholdsProfileBindingMaps(ctx, optional)
List Firewall CPU Memory Thresholds Profile Binding Maps

API will list all Firewall CPU Memory Thresholds Profile Binding Maps. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallCPUMemThresholdsProfileBindingMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallCPUMemThresholdsProfileBindingMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMapListResult**](PolicyFirewallCPUMemThresholdsProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallCPUMemThresholdsProfileBindingMaps_0**
> PolicyFirewallCpuMemThresholdsProfileBindingMapListResult ListPolicyFirewallCPUMemThresholdsProfileBindingMaps_0(ctx, optional)
List Firewall CPU Memory Thresholds Profile Binding Maps

API will list all Firewall CPU Memory Thresholds Profile Binding Maps. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallCPUMemThresholdsProfileBindingMaps_122Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallCPUMemThresholdsProfileBindingMaps_122Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMapListResult**](PolicyFirewallCPUMemThresholdsProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallFloodProtectionBindings**
> PolicyFirewallFloodProtectionProfileBindingMapListResult ListPolicyFirewallFloodProtectionBindings(ctx, domainId, groupId, optional)
List Firewall Flood Protection Profile Binding Maps

API will list all Firewall Flood Protection Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListPolicyFirewallFloodProtectionBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallFloodProtectionBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMapListResult**](PolicyFirewallFloodProtectionProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallFloodProtectionBindings_0**
> PolicyFirewallFloodProtectionProfileBindingMapListResult ListPolicyFirewallFloodProtectionBindings_0(ctx, domainId, groupId, optional)
List Firewall Flood Protection Profile Binding Maps

API will list all Firewall Flood Protection Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListPolicyFirewallFloodProtectionBindings_123Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallFloodProtectionBindings_123Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMapListResult**](PolicyFirewallFloodProtectionProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSchedulers**
> PolicyFirewallSchedulerListResult ListPolicyFirewallSchedulers(ctx, optional)
Get PolicyFirewallSchedulers

Get all PolicyFirewallSchedulers 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallSchedulersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSchedulersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSchedulerListResult**](PolicyFirewallSchedulerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSchedulers_0**
> PolicyFirewallSchedulerListResult ListPolicyFirewallSchedulers_0(ctx, optional)
Get PolicyFirewallSchedulers

Get all PolicyFirewallSchedulers 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallSchedulers_124Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSchedulers_124Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSchedulerListResult**](PolicyFirewallSchedulerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSessionTimerBindings**
> PolicyFirewallSessionTimerProfileBindingMapListResult ListPolicyFirewallSessionTimerBindings(ctx, domainId, groupId, optional)
List Firewall Session Timer Profile Binding Maps

API will list all Firewall Session Timer Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListPolicyFirewallSessionTimerBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSessionTimerBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMapListResult**](PolicyFirewallSessionTimerProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSessionTimerBindings_0**
> PolicyFirewallSessionTimerProfileBindingMapListResult ListPolicyFirewallSessionTimerBindings_0(ctx, domainId, groupId, optional)
List Firewall Session Timer Profile Binding Maps

API will list all Firewall Session Timer Profile Binding Maps in current group id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***PolicySecurityApiListPolicyFirewallSessionTimerBindings_125Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSessionTimerBindings_125Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMapListResult**](PolicyFirewallSessionTimerProfileBindingMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSessionTimerProfiles**
> PolicyFirewallSessionTimerProfileListResult ListPolicyFirewallSessionTimerProfiles(ctx, optional)
List Firewall Session Timer Profiles

API will list all Firewall Session Timer Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallSessionTimerProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSessionTimerProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileListResult**](PolicyFirewallSessionTimerProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyFirewallSessionTimerProfiles_0**
> PolicyFirewallSessionTimerProfileListResult ListPolicyFirewallSessionTimerProfiles_0(ctx, optional)
List Firewall Session Timer Profiles

API will list all Firewall Session Timer Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyFirewallSessionTimerProfiles_126Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyFirewallSessionTimerProfiles_126Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyFirewallSessionTimerProfileListResult**](PolicyFirewallSessionTimerProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceChainMappings**
> ServiceChainMappingListResult ListPolicyServiceChainMappings(ctx, serviceReferenceId, serviceProfileId, optional)
List all service chain mappings for given service profile.

List all service chain mappings in the system for the given service profile. If no explicit enforcement point is provided in the request, will return for default. Else, will return for specified points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 
 **optional** | ***PolicySecurityApiListPolicyServiceChainMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceChainMappingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**ServiceChainMappingListResult**](ServiceChainMappingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceChainMappings_0**
> ServiceChainMappingListResult ListPolicyServiceChainMappings_0(ctx, serviceReferenceId, serviceProfileId, optional)
List all service chain mappings for given service profile.

List all service chain mappings in the system for the given service profile. If no explicit enforcement point is provided in the request, will return for default. Else, will return for specified points. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 
 **optional** | ***PolicySecurityApiListPolicyServiceChainMappings_127Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceChainMappings_127Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enforcementPointPath** | **optional.String**| String Path of the enforcement point | 

### Return type

[**ServiceChainMappingListResult**](ServiceChainMappingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceChains**
> PolicyServiceChainListResult ListPolicyServiceChains(ctx, optional)
List service chains

List all the service chains available for service insertion 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyServiceChainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceChainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceChainListResult**](PolicyServiceChainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceChains_0**
> PolicyServiceChainListResult ListPolicyServiceChains_0(ctx, optional)
List service chains

List all the service chains available for service insertion 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyServiceChains_128Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceChains_128Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceChainListResult**](PolicyServiceChainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceInstanceEndpoints**
> ServiceInstanceEndpointListResult ListPolicyServiceInstanceEndpoints(ctx, tier0Id, localeServiceId, serviceInstanceId, optional)
List all service instance endpoint

List all service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
 **optional** | ***PolicySecurityApiListPolicyServiceInstanceEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceInstanceEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceInstanceEndpointListResult**](ServiceInstanceEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceInstanceEndpoints_0**
> ServiceInstanceEndpointListResult ListPolicyServiceInstanceEndpoints_0(ctx, tier0Id, localeServiceId, serviceInstanceId, optional)
List all service instance endpoint

List all service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
 **optional** | ***PolicySecurityApiListPolicyServiceInstanceEndpoints_129Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceInstanceEndpoints_129Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceInstanceEndpointListResult**](ServiceInstanceEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceProfiles**
> PolicyServiceProfileListResult ListPolicyServiceProfiles(ctx, serviceReferenceId, optional)
List service profiles

List all the service profiles available for given service reference 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
 **optional** | ***PolicySecurityApiListPolicyServiceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceProfileListResult**](PolicyServiceProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyServiceProfiles_0**
> PolicyServiceProfileListResult ListPolicyServiceProfiles_0(ctx, serviceReferenceId, optional)
List service profiles

List all the service profiles available for given service reference 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Service reference id | 
 **optional** | ***PolicySecurityApiListPolicyServiceProfiles_130Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyServiceProfiles_130Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceProfileListResult**](PolicyServiceProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyUrlCategories**
> PolicyUrlCategoryListResult ListPolicyUrlCategories(ctx, optional)
Get the list of URL categories.

Gets the list of categories. This will provide all the supported categories along with their ids. Few examples of these categories are Shopping, Social Networks, Streaming sites, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyUrlCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyUrlCategoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyUrlCategoryListResult**](PolicyUrlCategoryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyUrlCategories_0**
> PolicyUrlCategoryListResult ListPolicyUrlCategories_0(ctx, optional)
Get the list of URL categories.

Gets the list of categories. This will provide all the supported categories along with their ids. Few examples of these categories are Shopping, Social Networks, Streaming sites, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyUrlCategories_131Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyUrlCategories_131Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyUrlCategoryListResult**](PolicyUrlCategoryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyUrlReputationSeverities**
> PolicyUrlReputationSeverityListResult ListPolicyUrlReputationSeverities(ctx, optional)
Get the list of reputation severity

Gets the list of reputation severities. This will provide all the supported severities along with their ids, min and max reputaitons. The min_reputation and max_reputation specify the range of the reputations which belong to a particular severity. For instance, any reputation between 1 to 20 belongs to the severity 'High Risk'. Similary a reputation between 81 to 100 belong to the severity 'Trustworthy'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyUrlReputationSeveritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyUrlReputationSeveritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyUrlReputationSeverityListResult**](PolicyUrlReputationSeverityListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyUrlReputationSeverities_0**
> PolicyUrlReputationSeverityListResult ListPolicyUrlReputationSeverities_0(ctx, optional)
Get the list of reputation severity

Gets the list of reputation severities. This will provide all the supported severities along with their ids, min and max reputaitons. The min_reputation and max_reputation specify the range of the reputations which belong to a particular severity. For instance, any reputation between 1 to 20 belongs to the severity 'High Risk'. Similary a reputation between 81 to 100 belong to the severity 'Trustworthy'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListPolicyUrlReputationSeverities_132Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListPolicyUrlReputationSeverities_132Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyUrlReputationSeverityListResult**](PolicyUrlReputationSeverityListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionPolicies**
> RedirectionPolicyListResult ListRedirectionPolicies(ctx, domainId, optional)
List redirection policys for a domain

List redirection policys for a domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
 **optional** | ***PolicySecurityApiListRedirectionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionPolicyListResult**](RedirectionPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionPoliciesAcrossAllDomains**
> RedirectionPolicyListResult ListRedirectionPoliciesAcrossAllDomains(ctx, optional)
List redirection policys

List all redirection policys across all domains ordered by precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListRedirectionPoliciesAcrossAllDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionPoliciesAcrossAllDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionPolicyListResult**](RedirectionPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionPoliciesAcrossAllDomains_0**
> RedirectionPolicyListResult ListRedirectionPoliciesAcrossAllDomains_0(ctx, optional)
List redirection policys

List all redirection policys across all domains ordered by precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListRedirectionPoliciesAcrossAllDomains_133Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionPoliciesAcrossAllDomains_133Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionPolicyListResult**](RedirectionPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionPolicies_0**
> RedirectionPolicyListResult ListRedirectionPolicies_0(ctx, domainId, optional)
List redirection policys for a domain

List redirection policys for a domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
 **optional** | ***PolicySecurityApiListRedirectionPolicies_134Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionPolicies_134Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionPolicyListResult**](RedirectionPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionRules**
> RedirectionRuleListResult ListRedirectionRules(ctx, domainId, redirectionPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
 **optional** | ***PolicySecurityApiListRedirectionRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionRuleListResult**](RedirectionRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRedirectionRules_0**
> RedirectionRuleListResult ListRedirectionRules_0(ctx, domainId, redirectionPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
 **optional** | ***PolicySecurityApiListRedirectionRules_135Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListRedirectionRules_135Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RedirectionRuleListResult**](RedirectionRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityPoliciesForDomain**
> SecurityPolicyListResult ListSecurityPoliciesForDomain(ctx, domainId, optional)
List security policies

List all security policies for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSecurityPoliciesForDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSecurityPoliciesForDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SecurityPolicyListResult**](SecurityPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityPoliciesForDomain_0**
> SecurityPolicyListResult ListSecurityPoliciesForDomain_0(ctx, domainId, optional)
List security policies

List all security policies for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSecurityPoliciesForDomain_136Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSecurityPoliciesForDomain_136Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includeRuleCount** | **optional.Bool**| Include the count of rules in policy | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SecurityPolicyListResult**](SecurityPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityRules**
> RuleListResult ListSecurityRules(ctx, domainId, securityPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSecurityRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSecurityRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RuleListResult**](RuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecurityRules_0**
> RuleListResult ListSecurityRules_0(ctx, domainId, securityPolicyId, optional)
List rules

List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSecurityRules_137Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSecurityRules_137Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RuleListResult**](RuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceDefinitions**
> ServiceInsertionServiceListResult ListServiceDefinitions(ctx, enforcementPointId)
List all Service Definitions registered on given enforcement point.

List all Service Definitions registered on given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointId** | **string**| Enforcement point id | 

### Return type

[**ServiceInsertionServiceListResult**](ServiceInsertionServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceReferences**
> ServiceReferenceListResult ListServiceReferences(ctx, optional)
List service references

List all the partner service references available for service insertion 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListServiceReferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListServiceReferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceReferenceListResult**](ServiceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceReferences_0**
> ServiceReferenceListResult ListServiceReferences_0(ctx, optional)
List service references

List all the partner service references available for service insertion 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiListServiceReferences_138Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListServiceReferences_138Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceReferenceListResult**](ServiceReferenceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSessionTimerProfileBindings**
> SessionTimerProfileBindingListResult ListSessionTimerProfileBindings(ctx, sessionTimerProfileId, optional)
List Session Timer Profiles

API will list all Session Timer Profiles bindings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionTimerProfileId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSessionTimerProfileBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSessionTimerProfileBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SessionTimerProfileBindingListResult**](SessionTimerProfileBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSessionTimerProfileBindings_0**
> SessionTimerProfileBindingListResult ListSessionTimerProfileBindings_0(ctx, sessionTimerProfileId, optional)
List Session Timer Profiles

API will list all Session Timer Profiles bindings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionTimerProfileId** | **string**|  | 
 **optional** | ***PolicySecurityApiListSessionTimerProfileBindings_139Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListSessionTimerProfileBindings_139Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**SessionTimerProfileBindingListResult**](SessionTimerProfileBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualEndpointsForTier0**
> VirtualEndpointListResult ListVirtualEndpointsForTier0(ctx, tier0Id, localeServiceId, optional)
List all virtual endpoints

List all virtual endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiListVirtualEndpointsForTier0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListVirtualEndpointsForTier0Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualEndpointListResult**](VirtualEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualEndpointsForTier0_0**
> VirtualEndpointListResult ListVirtualEndpointsForTier0_0(ctx, tier0Id, localeServiceId, optional)
List all virtual endpoints

List all virtual endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiListVirtualEndpointsForTier0_140Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiListVirtualEndpointsForTier0_140Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualEndpointListResult**](VirtualEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeVersionAsActiveMakeActiveVersion**
> MakeVersionAsActiveMakeActiveVersion(ctx, body)
Change the state of IDS Signature Version

Make this IDS Signature version as ACTIVE version and other versions as NOTACTIVE. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSignatureVersion**](IdsSignatureVersion.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeVersionAsActiveMakeActiveVersion_0**
> MakeVersionAsActiveMakeActiveVersion_0(ctx, body)
Change the state of IDS Signature Version

Make this IDS Signature version as ACTIVE version and other versions as NOTACTIVE. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSignatureVersion**](IdsSignatureVersion.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchByodPolicyServiceInstance**
> PatchByodPolicyServiceInstance(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create BYOD Service Instance which represent instance of service definition created on manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchByodPolicyServiceInstance_0**
> PatchByodPolicyServiceInstance_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create BYOD Service Instance which represent instance of service definition created on manager. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCPUMemThresholdsProfile**
> PatchCPUMemThresholdsProfile(ctx, body, profileId)
Create or update CPU and memory thresholds profile

Create or update CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCPUMemThresholdsProfile_0**
> PatchCPUMemThresholdsProfile_0(ctx, body, profileId)
Create or update CPU and memory thresholds profile

Create or update CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCommunicationEntry**
> PatchCommunicationEntry(ctx, body, domainId, communicationMapId, communicationEntryId)
Patch a CommunicationEntry

Patch the CommunicationEntry. If a communication entry for the given communication-entry-id is not present, the object will get created and if it is present it will be updated. This is a full replace  This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCommunicationEntry_0**
> PatchCommunicationEntry_0(ctx, body, domainId, communicationMapId, communicationEntryId)
Patch a CommunicationEntry

Patch the CommunicationEntry. If a communication entry for the given communication-entry-id is not present, the object will get created and if it is present it will be updated. This is a full replace  This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCommunicationMapForDomain**
> PatchCommunicationMapForDomain(ctx, body, domainId, communicationMapId)
Patch communication map

Patch the communication map for a domain. If a communication map for the given communication-map-id is not present, the object will get created and if it is present it will be updated. This is a full replace This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCommunicationMapForDomain_0**
> PatchCommunicationMapForDomain_0(ctx, body, domainId, communicationMapId)
Patch communication map

Patch the communication map for a domain. If a communication map for the given communication-map-id is not present, the object will get created and if it is present it will be updated. This is a full replace This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchComputeClusterIdfwConfiguration**
> PatchComputeClusterIdfwConfiguration(ctx, body, clusterId)
Patch compute cluster idfw configuration

Patch compute cluster identity firewall configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)|  | 
  **clusterId** | **string**| Cluster ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchComputeClusterIdfwConfiguration_0**
> PatchComputeClusterIdfwConfiguration_0(ctx, body, clusterId)
Patch compute cluster idfw configuration

Patch compute cluster identity firewall configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)|  | 
  **clusterId** | **string**| Cluster ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDfwFirewallConfiguration**
> PatchDfwFirewallConfiguration(ctx, body)
Update dfw firewall configuration

Update dfw firewall related configurations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDfwFirewallConfiguration_0**
> PatchDfwFirewallConfiguration_0(ctx, body)
Update dfw firewall configuration

Update dfw firewall related configurations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDnsSecurityProfile**
> PatchDnsSecurityProfile(ctx, body, profileId)
Create or update DNS security profile

Create or update DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfile**](DnsSecurityProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDnsSecurityProfileBinding**
> PatchDnsSecurityProfileBinding(ctx, body, domainId, groupId, dnsSecurityProfileBindingMapId)
Create or update DNS security profile binding map

API will create or update DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDnsSecurityProfileBinding_0**
> PatchDnsSecurityProfileBinding_0(ctx, body, domainId, groupId, dnsSecurityProfileBindingMapId)
Create or update DNS security profile binding map

API will create or update DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDnsSecurityProfile_0**
> PatchDnsSecurityProfile_0(ctx, body, profileId)
Create or update DNS security profile

Create or update DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfile**](DnsSecurityProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDraft**
> PatchDraft(ctx, body, draftId)
Patch a manual draft

Create a new manual draft if the specified draft id does not correspond to an existing draft. Update the manual draft otherwise. Auto draft can not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyDraft**](PolicyDraft.md)|  | 
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDraft_0**
> PatchDraft_0(ctx, body, draftId)
Patch a manual draft

Create a new manual draft if the specified draft id does not correspond to an existing draft. Update the manual draft otherwise. Auto draft can not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyDraft**](PolicyDraft.md)|  | 
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEndpointPolicy**
> PatchEndpointPolicy(ctx, body, domainId, endpointPolicyId)
Create or update Endpoint policy

Create or update the Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointPolicy**](EndpointPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEndpointPolicy_0**
> PatchEndpointPolicy_0(ctx, body, domainId, endpointPolicyId)
Create or update Endpoint policy

Create or update the Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointPolicy**](EndpointPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEndpointRule**
> PatchEndpointRule(ctx, body, domainId, endpointPolicyId, endpointRuleId)
Update Endpoint rule

Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointRule**](EndpointRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEndpointRule_0**
> PatchEndpointRule_0(ctx, body, domainId, endpointPolicyId, endpointRuleId)
Update Endpoint rule

Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointRule**](EndpointRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchExcludeList**
> PatchExcludeList(ctx, body)
Patch exclusion list for security policy

Patch exclusion list for security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyExcludeList**](PolicyExcludeList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchExcludeList_0**
> PatchExcludeList_0(ctx, body)
Patch exclusion list for security policy

Patch exclusion list for security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyExcludeList**](PolicyExcludeList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFloodProtectionProfile**
> PatchFloodProtectionProfile(ctx, body, floodProtectionProfileId)
Create or update Flood Protection Profile

API will create/update Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfile**](FloodProtectionProfile.md)|  | 
  **floodProtectionProfileId** | **string**| Firewall Flood Protection Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFloodProtectionProfile_0**
> PatchFloodProtectionProfile_0(ctx, body, floodProtectionProfileId)
Create or update Flood Protection Profile

API will create/update Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfile**](FloodProtectionProfile.md)|  | 
  **floodProtectionProfileId** | **string**| Firewall Flood Protection Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGatewayPolicyForDomain**
> PatchGatewayPolicyForDomain(ctx, body, domainId, gatewayPolicyId)
Update gateway policy

Update the gateway policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGatewayPolicyForDomain_0**
> PatchGatewayPolicyForDomain_0(ctx, body, domainId, gatewayPolicyId)
Update gateway policy

Update the gateway policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGatewayRule**
> PatchGatewayRule(ctx, body, domainId, gatewayPolicyId, ruleId)
Update gateway rule

Update the gateway rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGatewayRule_0**
> PatchGatewayRule_0(ctx, body, domainId, gatewayPolicyId, ruleId)
Update gateway rule

Update the gateway rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupMonitoringBinding**
> PatchGroupMonitoringBinding(ctx, body, domainId, groupId, groupMonitoringProfileBindingMapId)
Create Group Monitoring Profile Binding Map

API will create group monitoring profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchGroupMonitoringBinding_0**
> PatchGroupMonitoringBinding_0(ctx, body, domainId, groupId, groupMonitoringProfileBindingMapId)
Create Group Monitoring Profile Binding Map

API will create group monitoring profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsClusterConfig**
> PatchIdsClusterConfig(ctx, body, clusterId)
Patch IDS config on cluster level

Patch intrusion detection system on cluster level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsClusterConfig**](IdsClusterConfig.md)|  | 
  **clusterId** | **string**| User entered ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsClusterConfig_0**
> PatchIdsClusterConfig_0(ctx, body, clusterId)
Patch IDS config on cluster level

Patch intrusion detection system on cluster level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsClusterConfig**](IdsClusterConfig.md)|  | 
  **clusterId** | **string**| User entered ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsProfile**
> PatchIdsProfile(ctx, body, profileId)
Patch IDS profile

Patch intrusion detection system profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsProfile**](IdsProfile.md)|  | 
  **profileId** | **string**| Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsProfile_0**
> PatchIdsProfile_0(ctx, body, profileId)
Patch IDS profile

Patch intrusion detection system profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsProfile**](IdsProfile.md)|  | 
  **profileId** | **string**| Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsRule**
> PatchIdsRule(ctx, body, domainId, policyId, ruleId)
Patch IDS rule

Patch intrusion detection system rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsRule_0**
> PatchIdsRule_0(ctx, body, domainId, policyId, ruleId)
Patch IDS rule

Patch intrusion detection system rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 
  **ruleId** | **string**| Rule ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsSecurityPolicy**
> PatchIdsSecurityPolicy(ctx, body, domainId, policyId)
Patch IDS security policy

Patch intrusion detection system security policy for a domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsSecurityPolicy_0**
> PatchIdsSecurityPolicy_0(ctx, body, domainId, policyId)
Patch IDS security policy

Patch intrusion detection system security policy for a domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**| Domain ID | 
  **policyId** | **string**| Policy ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsSettings**
> PatchIdsSettings(ctx, body)
Patch Intrusion detection system settings

Intrusion detection system settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSettings**](IdsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsSettings_0**
> PatchIdsSettings_0(ctx, body)
Patch Intrusion detection system settings

Intrusion detection system settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSettings**](IdsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsStandaloneHostConfig**
> PatchIdsStandaloneHostConfig(ctx, body)
Patch IDS configuration

Patch intrusion detection system configuration on standalone hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIdsStandaloneHostConfig_0**
> PatchIdsStandaloneHostConfig_0(ctx, body)
Patch IDS configuration

Patch intrusion detection system configuration on standalone hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsStandaloneHostConfig**](IdsStandaloneHostConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallCPUMemThresholdsProfileBindingMap**
> PatchPolicyFirewallCPUMemThresholdsProfileBindingMap(ctx, body, cpuMemThresholdsProfileBindingMapId)
Create or update Firewall CPU Memory Thresholds Profile Binding Map

API will create or update Firewall CPU Memory Thresholds Profile binding map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCpuMemThresholdsProfileBindingMap.md)|  | 
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallCPUMemThresholdsProfileBindingMap_0**
> PatchPolicyFirewallCPUMemThresholdsProfileBindingMap_0(ctx, body, cpuMemThresholdsProfileBindingMapId)
Create or update Firewall CPU Memory Thresholds Profile Binding Map

API will create or update Firewall CPU Memory Thresholds Profile binding map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCpuMemThresholdsProfileBindingMap.md)|  | 
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallFloodProtectionProfileBindingMap**
> PatchPolicyFirewallFloodProtectionProfileBindingMap(ctx, body, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Create or update Firewall Flood Protection Profile Binding Map

API will create or update Firewall Flood Protection profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallFloodProtectionProfileBindingMap_0**
> PatchPolicyFirewallFloodProtectionProfileBindingMap_0(ctx, body, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Create or update Firewall Flood Protection Profile Binding Map

API will create or update Firewall Flood Protection profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallScheduler**
> PatchPolicyFirewallScheduler(ctx, body, firewallSchedulerId)
Create or Update PolicyFirewallScheduler

Creates/Updates a PolicyFirewallScheduler, which can be set at security policy. Note that at least one property out of \"days\", \"start_date\", \"time_interval\", \"end_date\" is required if \"recurring\" field is true. Also \"start_time\" and \"end_time\" should not be present. And if \"recurring\" field is false then \"start_date\" and \"end_date\" is mandatory, \"start_time\" and \"end_time\" is optional. Also the fields \"days\" and \"time_interval\" should not be present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)|  | 
  **firewallSchedulerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallScheduler_0**
> PatchPolicyFirewallScheduler_0(ctx, body, firewallSchedulerId)
Create or Update PolicyFirewallScheduler

Creates/Updates a PolicyFirewallScheduler, which can be set at security policy. Note that at least one property out of \"days\", \"start_date\", \"time_interval\", \"end_date\" is required if \"recurring\" field is true. Also \"start_time\" and \"end_time\" should not be present. And if \"recurring\" field is false then \"start_date\" and \"end_date\" is mandatory, \"start_time\" and \"end_time\" is optional. Also the fields \"days\" and \"time_interval\" should not be present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)|  | 
  **firewallSchedulerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallSessionTimerProfile**
> PatchPolicyFirewallSessionTimerProfile(ctx, body, firewallSessionTimerProfileId)
Create or update Firewall Session Timer Profile

API will create/update Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)|  | 
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallSessionTimerProfileBindingMap**
> PatchPolicyFirewallSessionTimerProfileBindingMap(ctx, body, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Create or update Firewall Session Timer Profile Binding Map

API will create or update Firewall Session Timer profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallSessionTimerProfileBindingMap_0**
> PatchPolicyFirewallSessionTimerProfileBindingMap_0(ctx, body, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Create or update Firewall Session Timer Profile Binding Map

API will create or update Firewall Session Timer profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyFirewallSessionTimerProfile_0**
> PatchPolicyFirewallSessionTimerProfile_0(ctx, body, firewallSessionTimerProfileId)
Create or update Firewall Session Timer Profile

API will create/update Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)|  | 
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyServiceInstance**
> PatchPolicyServiceInstance(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create Service Instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyServiceInstance_0**
> PatchPolicyServiceInstance_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId)
Create service instance

Create Service Instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyServiceProfile**
> PatchPolicyServiceProfile(ctx, body, serviceReferenceId, serviceProfileId)
Create service profile

Create Service profile to specify vendor template attri- butes for a given 3rd party service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceProfile**](PolicyServiceProfile.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyServiceProfile_0**
> PatchPolicyServiceProfile_0(ctx, body, serviceReferenceId, serviceProfileId)
Create service profile

Create Service profile to specify vendor template attri- butes for a given 3rd party service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceProfile**](PolicyServiceProfile.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyUrlCategorizationConfig**
> PolicyUrlCategorizationConfig PatchPolicyUrlCategorizationConfig(ctx, body, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Create or Update PolicyUrlCategorizationConfig

Creates/Updates a PolicyUrlCategorizationConfig. Creating or updating the PolicyUrlCategorizationConfig will enable or disable URL categorization for the given edge cluster. If the context_profiles field is empty, the edge cluster will detect all the categories of URLs. If context_profiles field has any context profiles, the edge cluster will detect only the categories listed within those context profiles. The context profiles should have attribute type URL_CATEGORY. The update_frequency specifies how frequently in minutes, the edge cluster will get updates about the URL data from the URL categorization cloud service. If the update_frequency is not specified, the default update frequency will be 30 min. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)|  | 
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPolicyUrlCategorizationConfig_0**
> PolicyUrlCategorizationConfig PatchPolicyUrlCategorizationConfig_0(ctx, body, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Create or Update PolicyUrlCategorizationConfig

Creates/Updates a PolicyUrlCategorizationConfig. Creating or updating the PolicyUrlCategorizationConfig will enable or disable URL categorization for the given edge cluster. If the context_profiles field is empty, the edge cluster will detect all the categories of URLs. If context_profiles field has any context profiles, the edge cluster will detect only the categories listed within those context profiles. The context profiles should have attribute type URL_CATEGORY. The update_frequency specifies how frequently in minutes, the edge cluster will get updates about the URL data from the URL categorization cloud service. If the update_frequency is not specified, the default update frequency will be 30 min. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)|  | 
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRedirectionPolicy**
> PatchRedirectionPolicy(ctx, body, domainId, redirectionPolicyId)
Create or update redirection policy

Create or update the redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionPolicy**](RedirectionPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRedirectionPolicy_0**
> PatchRedirectionPolicy_0(ctx, body, domainId, redirectionPolicyId)
Create or update redirection policy

Create or update the redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionPolicy**](RedirectionPolicy.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRedirectionRule**
> PatchRedirectionRule(ctx, body, domainId, redirectionPolicyId, ruleId)
Update redirection rule

Create a rule with the rule-id is not already present, otherwise update the rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionRule**](RedirectionRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| RedirectionPolicy id | 
  **ruleId** | **string**| rule id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRedirectionRule_0**
> PatchRedirectionRule_0(ctx, body, domainId, redirectionPolicyId, ruleId)
Update redirection rule

Create a rule with the rule-id is not already present, otherwise update the rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedirectionRule**](RedirectionRule.md)|  | 
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| RedirectionPolicy id | 
  **ruleId** | **string**| rule id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSecurityPolicyForDomain**
> PatchSecurityPolicyForDomain(ctx, body, domainId, securityPolicyId)
Patch security policy

Patch the security policy for a domain. If a security policy for the given security-policy-id is not present, the object will get created and if it is present it will be updated. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSecurityPolicyForDomain_0**
> PatchSecurityPolicyForDomain_0(ctx, body, domainId, securityPolicyId)
Patch security policy

Patch the security policy for a domain. If a security policy for the given security-policy-id is not present, the object will get created and if it is present it will be updated. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSecurityRule**
> PatchSecurityRule(ctx, body, domainId, securityPolicyId, ruleId)
Patch a rule

Patch the rule. If Rule corresponding to the the given rule-id is not present, the object will get created and if it is present it will be updated. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSecurityRule_0**
> PatchSecurityRule_0(ctx, body, domainId, securityPolicyId, ruleId)
Patch a rule

Patch the rule. If Rule corresponding to the the given rule-id is not present, the object will get created and if it is present it will be updated. This is a full replace 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceChain**
> PatchServiceChain(ctx, body, serviceChainId)
Create service chain

Create Service chain representing the sequence in which 3rd party services must be consumed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceChain**](PolicyServiceChain.md)|  | 
  **serviceChainId** | **string**| Service chain id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceChain_0**
> PatchServiceChain_0(ctx, body, serviceChainId)
Create service chain

Create Service chain representing the sequence in which 3rd party services must be consumed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceChain**](PolicyServiceChain.md)|  | 
  **serviceChainId** | **string**| Service chain id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceInstanceEndpoint**
> PatchServiceInstanceEndpoint(ctx, body, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Create service instance endpoint

Create Service instance endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceInstanceEndpoint_0**
> PatchServiceInstanceEndpoint_0(ctx, body, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Create service instance endpoint

Create Service instance endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceReference**
> PatchServiceReference(ctx, body, serviceReferenceId)
Create service reference

Create Service Reference representing the intent to consume a given 3rd party service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceReference**](ServiceReference.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchServiceReference_0**
> PatchServiceReference_0(ctx, body, serviceReferenceId)
Create service reference

Create Service Reference representing the intent to consume a given 3rd party service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceReference**](ServiceReference.md)|  | 
  **serviceReferenceId** | **string**| Service reference id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchStandaloneHostIdfwConfiguration**
> PatchStandaloneHostIdfwConfiguration(ctx, body)
Patch idfw configuration for standalone host

Patch identity firewall configuration for standalone host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchStandaloneHostIdfwConfiguration_0**
> PatchStandaloneHostIdfwConfiguration_0(ctx, body)
Patch idfw configuration for standalone host

Patch identity firewall configuration for standalone host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0FloodProtectionProfileBindingMap**
> PatchTier0FloodProtectionProfileBindingMap(ctx, body, tier0Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router

API will create or update Flood Protection profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0FloodProtectionProfileBindingMap_0**
> PatchTier0FloodProtectionProfileBindingMap_0(ctx, body, tier0Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router

API will create or update Flood Protection profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0LocalServicesSessionTimerProfileBindingMap**
> PatchTier0LocalServicesSessionTimerProfileBindingMap(ctx, body, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0LocalServicesSessionTimerProfileBindingMap_0**
> PatchTier0LocalServicesSessionTimerProfileBindingMap_0(ctx, body, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0LocaleServicesFloodProtectionProfileBindingMap**
> PatchTier0LocaleServicesFloodProtectionProfileBindingMap(ctx, body, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0LocaleServicesFloodProtectionProfileBindingMap_0**
> PatchTier0LocaleServicesFloodProtectionProfileBindingMap_0(ctx, body, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0SessionTimerProfileBindingMap**
> PatchTier0SessionTimerProfileBindingMap(ctx, body, tier0Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router

API will create or update Session Timer profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier0SessionTimerProfileBindingMap_0**
> PatchTier0SessionTimerProfileBindingMap_0(ctx, body, tier0Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router

API will create or update Session Timer profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1FloodProtectionProfileBindingMap**
> PatchTier1FloodProtectionProfileBindingMap(ctx, body, tier1Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router

API will create or update Flood Protection profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1FloodProtectionProfileBindingMap_0**
> PatchTier1FloodProtectionProfileBindingMap_0(ctx, body, tier1Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router

API will create or update Flood Protection profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1LocaleServicesFloodProtectionProfileBindingMap**
> PatchTier1LocaleServicesFloodProtectionProfileBindingMap(ctx, body, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1LocaleServicesFloodProtectionProfileBindingMap_0**
> PatchTier1LocaleServicesFloodProtectionProfileBindingMap_0(ctx, body, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1LocaleServicesSessionTimerProfileBindingMap**
> PatchTier1LocaleServicesSessionTimerProfileBindingMap(ctx, body, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1LocaleServicesSessionTimerProfileBindingMap_0**
> PatchTier1LocaleServicesSessionTimerProfileBindingMap_0(ctx, body, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1PolicyServiceInstance**
> PatchTier1PolicyServiceInstance(ctx, body, tier1Id, localeServiceId, serviceInstanceId)
Create Tier1 service instance

Create Tier1 Service Instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1PolicyServiceInstance_0**
> PatchTier1PolicyServiceInstance_0(ctx, body, tier1Id, localeServiceId, serviceInstanceId)
Create Tier1 service instance

Create Tier1 Service Instance. Please note that, only display_name, description and deployment_spec_name are allowed to be modified in an exisiting entity. If the deployment spec name is changed, it will trigger the upgrade operation for the SVMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyServiceInstance**](PolicyServiceInstance.md)|  | 
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Tier1 Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1SessionTimerProfileBindingMap**
> PatchTier1SessionTimerProfileBindingMap(ctx, body, tier1Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router

API will create or update Session Timer profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTier1SessionTimerProfileBindingMap_0**
> PatchTier1SessionTimerProfileBindingMap_0(ctx, body, tier1Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router

API will create or update Session Timer profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchVirtualEndpoint**
> PatchVirtualEndpoint(ctx, body, tier0Id, localeServiceId, virtualEndpointId)
Create or update virtual endpoint

Create or update virtual endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualEndpoint**](VirtualEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchVirtualEndpoint_0**
> PatchVirtualEndpoint_0(ctx, body, tier0Id, localeServiceId, virtualEndpointId)
Create or update virtual endpoint

Create or update virtual endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualEndpoint**](VirtualEndpoint.md)|  | 
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishDraftPublish**
> PublishDraftPublish(ctx, body, draftId)
Publish a draft

Read a draft and publish it by applying changes onto current configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishDraftPublish_0**
> PublishDraftPublish_0(ctx, body, draftId)
Publish a draft

Read a draft and publish it by applying changes onto current configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Infra**](Infra.md)|  | 
  **draftId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutComputeClusterIdfwConfiguration**
> ComputeClusterIdfwConfiguration PutComputeClusterIdfwConfiguration(ctx, body, clusterId)
Create or update compute cluster idfw configuration

Update the compute cluster idfw configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)|  | 
  **clusterId** | **string**| Cluster ID | 

### Return type

[**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutComputeClusterIdfwConfiguration_0**
> ComputeClusterIdfwConfiguration PutComputeClusterIdfwConfiguration_0(ctx, body, clusterId)
Create or update compute cluster idfw configuration

Update the compute cluster idfw configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)|  | 
  **clusterId** | **string**| Cluster ID | 

### Return type

[**ComputeClusterIdfwConfiguration**](ComputeClusterIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDfwFirewallConfiguration**
> DfwFirewallConfiguration PutDfwFirewallConfiguration(ctx, body)
Update dfw firewall configuration

Update dfw firewall related configurations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)|  | 

### Return type

[**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDfwFirewallConfiguration_0**
> DfwFirewallConfiguration PutDfwFirewallConfiguration_0(ctx, body)
Update dfw firewall configuration

Update dfw firewall related configurations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)|  | 

### Return type

[**DfwFirewallConfiguration**](DfwFirewallConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDraft**
> PolicyDraft PutDraft(ctx, body, draftId)
Create or update a manual draft

Create a new manual draft if the specified draft id does not correspond to an existing draft. Update the manual draft otherwise. Auto draft can not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyDraft**](PolicyDraft.md)|  | 
  **draftId** | **string**|  | 

### Return type

[**PolicyDraft**](PolicyDraft.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDraft_0**
> PolicyDraft PutDraft_0(ctx, body, draftId)
Create or update a manual draft

Create a new manual draft if the specified draft id does not correspond to an existing draft. Update the manual draft otherwise. Auto draft can not be updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyDraft**](PolicyDraft.md)|  | 
  **draftId** | **string**|  | 

### Return type

[**PolicyDraft**](PolicyDraft.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExcludeList**
> PolicyExcludeList PutExcludeList(ctx, body)
Create or update exclusion list for security policy

Update the exclusion list for security policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyExcludeList**](PolicyExcludeList.md)|  | 

### Return type

[**PolicyExcludeList**](PolicyExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExcludeList_0**
> PolicyExcludeList PutExcludeList_0(ctx, body)
Create or update exclusion list for security policy

Update the exclusion list for security policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyExcludeList**](PolicyExcludeList.md)|  | 

### Return type

[**PolicyExcludeList**](PolicyExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolicyUrlCategorizationConfig**
> PolicyUrlCategorizationConfig PutPolicyUrlCategorizationConfig(ctx, body, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Create or Update PolicyUrlCategorizationConfig

Creates/Updates a PolicyUrlCategorizationConfig. Creating or updating the PolicyUrlCategorizationConfig will enable or disable URL categorization for the given edge cluster. If the context_profiles field is empty, the edge cluster will detect all the categories of URLs. If context_profiles field has any context profiles, the edge cluster will detect only the categories listed within those context profiles. The context profiles should have attribute type URL_CATEGORY. The update_frequency specifies how frequently in minutes, the edge cluster will get updates about the URL data from the URL categorization cloud service. If the update_frequency is not specified, the default update frequency will be 30 min. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)|  | 
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolicyUrlCategorizationConfig_0**
> PolicyUrlCategorizationConfig PutPolicyUrlCategorizationConfig_0(ctx, body, siteId, enforcementPointId, edgeClusterId, urlCategorizationConfigId)
Create or Update PolicyUrlCategorizationConfig

Creates/Updates a PolicyUrlCategorizationConfig. Creating or updating the PolicyUrlCategorizationConfig will enable or disable URL categorization for the given edge cluster. If the context_profiles field is empty, the edge cluster will detect all the categories of URLs. If context_profiles field has any context profiles, the edge cluster will detect only the categories listed within those context profiles. The context profiles should have attribute type URL_CATEGORY. The update_frequency specifies how frequently in minutes, the edge cluster will get updates about the URL data from the URL categorization cloud service. If the update_frequency is not specified, the default update frequency will be 30 min. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)|  | 
  **siteId** | **string**|  | 
  **enforcementPointId** | **string**|  | 
  **edgeClusterId** | **string**|  | 
  **urlCategorizationConfigId** | **string**|  | 

### Return type

[**PolicyUrlCategorizationConfig**](PolicyUrlCategorizationConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStandaloneHostIdfwConfiguration**
> StandaloneHostIdfwConfiguration PutStandaloneHostIdfwConfiguration(ctx, body)
Create or update idfw configuration for standalone host

Update the idfw configuration for standalone host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)|  | 

### Return type

[**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStandaloneHostIdfwConfiguration_0**
> StandaloneHostIdfwConfiguration PutStandaloneHostIdfwConfiguration_0(ctx, body)
Create or update idfw configuration for standalone host

Update the idfw configuration for standalone host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)|  | 

### Return type

[**StandaloneHostIdfwConfiguration**](StandaloneHostIdfwConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAllPolicyServiceInstancesForTier0**
> PolicyServiceInstanceListResult ReadAllPolicyServiceInstancesForTier0(ctx, tier0Id, localeServiceId, optional)
Read all service instance objects under a tier-0

Read all service instance objects under a tier-0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiReadAllPolicyServiceInstancesForTier0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadAllPolicyServiceInstancesForTier0Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceInstanceListResult**](PolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAllPolicyServiceInstancesForTier0_0**
> PolicyServiceInstanceListResult ReadAllPolicyServiceInstancesForTier0_0(ctx, tier0Id, localeServiceId, optional)
Read all service instance objects under a tier-0

Read all service instance objects under a tier-0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiReadAllPolicyServiceInstancesForTier0_197Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadAllPolicyServiceInstancesForTier0_197Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceInstanceListResult**](PolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAllPolicyServiceInstancesForTier1**
> PolicyServiceInstanceListResult ReadAllPolicyServiceInstancesForTier1(ctx, tier1Id, localeServiceId, optional)
Read all service instance objects under a tier-1

Read all service instance objects under a tier-1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiReadAllPolicyServiceInstancesForTier1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadAllPolicyServiceInstancesForTier1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceInstanceListResult**](PolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAllPolicyServiceInstancesForTier1_0**
> PolicyServiceInstanceListResult ReadAllPolicyServiceInstancesForTier1_0(ctx, tier1Id, localeServiceId, optional)
Read all service instance objects under a tier-1

Read all service instance objects under a tier-1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
 **optional** | ***PolicySecurityApiReadAllPolicyServiceInstancesForTier1_198Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadAllPolicyServiceInstancesForTier1_198Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PolicyServiceInstanceListResult**](PolicyServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadByodPolicyServiceInstance**
> ByodPolicyServiceInstance ReadByodPolicyServiceInstance(ctx, tier0Id, localeServiceId, serviceInstanceId)
Read byod service instance

Read byod service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadByodPolicyServiceInstance_0**
> ByodPolicyServiceInstance ReadByodPolicyServiceInstance_0(ctx, tier0Id, localeServiceId, serviceInstanceId)
Read byod service instance

Read byod service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**ByodPolicyServiceInstance**](ByodPolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCPUMemThresholdsProfile**
> PolicyFirewallCpuMemThresholdsProfile ReadCPUMemThresholdsProfile(ctx, profileId)
Read the CPU and memory thresholds profile

Read the CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCPUMemThresholdsProfile_0**
> PolicyFirewallCpuMemThresholdsProfile ReadCPUMemThresholdsProfile_0(ctx, profileId)
Read the CPU and memory thresholds profile

Read the CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCommunicationEntry**
> CommunicationEntry ReadCommunicationEntry(ctx, domainId, communicationMapId, communicationEntryId)
Read CommunicationEntry

Read CommunicationEntry This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCommunicationEntry_0**
> CommunicationEntry ReadCommunicationEntry_0(ctx, domainId, communicationMapId, communicationEntryId)
Read CommunicationEntry

Read CommunicationEntry This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCommunicationMapForDomain**
> CommunicationMap ReadCommunicationMapForDomain(ctx, domainId, communicationMapId)
Read communication-map

Read communication-map for a domain. This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCommunicationMapForDomain_0**
> CommunicationMap ReadCommunicationMapForDomain_0(ctx, domainId, communicationMapId)
Read communication-map

Read communication-map for a domain. This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDnsSecurityProfile**
> DnsSecurityProfile ReadDnsSecurityProfile(ctx, profileId)
Read the DNS Forwarder for the given tier-0 instance

Read the DNS Forwarder for the given tier-0 instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**DnsSecurityProfile**](DnsSecurityProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDnsSecurityProfile_0**
> DnsSecurityProfile ReadDnsSecurityProfile_0(ctx, profileId)
Read the DNS Forwarder for the given tier-0 instance

Read the DNS Forwarder for the given tier-0 instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**DnsSecurityProfile**](DnsSecurityProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDraft**
> PolicyDraft ReadDraft(ctx, draftId)
Read draft

Read a draft for a given draft identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**PolicyDraft**](PolicyDraft.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDraft_0**
> PolicyDraft ReadDraft_0(ctx, draftId)
Read draft

Read a draft for a given draft identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftId** | **string**|  | 

### Return type

[**PolicyDraft**](PolicyDraft.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEndpointPolicy**
> EndpointPolicy ReadEndpointPolicy(ctx, domainId, endpointPolicyId)
Read Endpoint policy

Read Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

[**EndpointPolicy**](EndpointPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEndpointPolicy_0**
> EndpointPolicy ReadEndpointPolicy_0(ctx, domainId, endpointPolicyId)
Read Endpoint policy

Read Endpoint policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 

### Return type

[**EndpointPolicy**](EndpointPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEndpointRule**
> EndpointRule ReadEndpointRule(ctx, domainId, endpointPolicyId, endpointRuleId)
Read Endpoint rule

Read Endpoint rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

[**EndpointRule**](EndpointRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEndpointRule_0**
> EndpointRule ReadEndpointRule_0(ctx, domainId, endpointPolicyId, endpointRuleId)
Read Endpoint rule

Read Endpoint rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **endpointPolicyId** | **string**| Endpoint policy id | 
  **endpointRuleId** | **string**| Endpoint rule id | 

### Return type

[**EndpointRule**](EndpointRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGatewayPolicyForDomain**
> GatewayPolicy ReadGatewayPolicyForDomain(ctx, domainId, gatewayPolicyId)
Read gateway policy

Read gateway policy for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGatewayPolicyForDomain_0**
> GatewayPolicy ReadGatewayPolicyForDomain_0(ctx, domainId, gatewayPolicyId)
Read gateway policy

Read gateway policy for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGatewayRule**
> Rule ReadGatewayRule(ctx, domainId, gatewayPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadGatewayRule_0**
> Rule ReadGatewayRule_0(ctx, domainId, gatewayPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPartnerService**
> ServiceDefinition ReadPartnerService(ctx, serviceName)
Read partner service identified by provided name

Read the specific partner service identified by provided name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceName** | **string**| Name of the service | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPartnerService_0**
> ServiceDefinition ReadPartnerService_0(ctx, serviceName)
Read partner service identified by provided name

Read the specific partner service identified by provided name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceName** | **string**| Name of the service | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPartnerServices**
> ServiceInsertionServiceListResult ReadPartnerServices(ctx, optional)
Read partner services

Read all the partner services available for service insertion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiReadPartnerServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadPartnerServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceInsertionServiceListResult**](ServiceInsertionServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPartnerServices_0**
> ServiceInsertionServiceListResult ReadPartnerServices_0(ctx, optional)
Read partner services

Read all the partner services available for service insertion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicySecurityApiReadPartnerServices_210Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReadPartnerServices_210Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeMarkForDeleteObjects** | **optional.Bool**| Include objects that are marked for deletion in results | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceInsertionServiceListResult**](ServiceInsertionServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceInstance**
> PolicyServiceInstance ReadPolicyServiceInstance(ctx, tier0Id, localeServiceId, serviceInstanceId)
Read service instance

Read service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceInstanceEndpoint**
> ServiceInstanceEndpoint ReadPolicyServiceInstanceEndpoint(ctx, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Read service instance endpoint

Read service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

[**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceInstanceEndpoint_0**
> ServiceInstanceEndpoint ReadPolicyServiceInstanceEndpoint_0(ctx, tier0Id, localeServiceId, serviceInstanceId, serviceInstanceEndpointId)
Read service instance endpoint

Read service instance endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 
  **serviceInstanceEndpointId** | **string**| Service instance endpoint id | 

### Return type

[**ServiceInstanceEndpoint**](ServiceInstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceInstance_0**
> PolicyServiceInstance ReadPolicyServiceInstance_0(ctx, tier0Id, localeServiceId, serviceInstanceId)
Read service instance

Read service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceProfile**
> PolicyServiceProfile ReadPolicyServiceProfile(ctx, serviceReferenceId, serviceProfileId)
Read service profile

This API can be used to read service profile with given service-profile-id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

[**PolicyServiceProfile**](PolicyServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceProfile_0**
> PolicyServiceProfile ReadPolicyServiceProfile_0(ctx, serviceReferenceId, serviceProfileId)
Read service profile

This API can be used to read service profile with given service-profile-id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 
  **serviceProfileId** | **string**| Service profile id | 

### Return type

[**PolicyServiceProfile**](PolicyServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedirectionPolicy**
> RedirectionPolicy ReadRedirectionPolicy(ctx, domainId, redirectionPolicyId)
Read redirection policy

Read redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

[**RedirectionPolicy**](RedirectionPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedirectionPolicy_0**
> RedirectionPolicy ReadRedirectionPolicy_0(ctx, domainId, redirectionPolicyId)
Read redirection policy

Read redirection policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 

### Return type

[**RedirectionPolicy**](RedirectionPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedirectionRule**
> RedirectionRule ReadRedirectionRule(ctx, domainId, redirectionPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
  **ruleId** | **string**| Rule id | 

### Return type

[**RedirectionRule**](RedirectionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedirectionRule_0**
> RedirectionRule ReadRedirectionRule_0(ctx, domainId, redirectionPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Domain id | 
  **redirectionPolicyId** | **string**| Redirection map id | 
  **ruleId** | **string**| Rule id | 

### Return type

[**RedirectionRule**](RedirectionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityPolicyForDomain**
> SecurityPolicy ReadSecurityPolicyForDomain(ctx, domainId, securityPolicyId)
Read security policy

Read security policy for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityPolicyForDomain_0**
> SecurityPolicy ReadSecurityPolicyForDomain_0(ctx, domainId, securityPolicyId)
Read security policy

Read security policy for a domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityRule**
> Rule ReadSecurityRule(ctx, domainId, securityPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSecurityRule_0**
> Rule ReadSecurityRule_0(ctx, domainId, securityPolicyId, ruleId)
Read rule

Read rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceChain**
> PolicyServiceChain ReadServiceChain(ctx, serviceChainId)
Read service chain

This API can be used to read service chain with given service-chain-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**| Id of Service chain | 

### Return type

[**PolicyServiceChain**](PolicyServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceChain_0**
> PolicyServiceChain ReadServiceChain_0(ctx, serviceChainId)
Read service chain

This API can be used to read service chain with given service-chain-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**| Id of Service chain | 

### Return type

[**PolicyServiceChain**](PolicyServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceDefinition**
> ServiceDefinition ReadServiceDefinition(ctx, enforcementPointId, serviceDefinitionId)
Read Service Definition with given service-definition-id.

Read Service Definition with given service-definition-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enforcementPointId** | **string**| Enforcement point id | 
  **serviceDefinitionId** | **string**| Id of service definition | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceReference**
> ServiceReference ReadServiceReference(ctx, serviceReferenceId)
Read service reference

This API can be used to read service reference with the given service-reference-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 

### Return type

[**ServiceReference**](ServiceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceReference_0**
> ServiceReference ReadServiceReference_0(ctx, serviceReferenceId)
Read service reference

This API can be used to read service reference with the given service-reference-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceReferenceId** | **string**| Id of Service Reference | 

### Return type

[**ServiceReference**](ServiceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTier1PolicyServiceInstance**
> PolicyServiceInstance ReadTier1PolicyServiceInstance(ctx, tier1Id, localeServiceId, serviceInstanceId)
Read Tier1 service instance

Read Tier1 service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTier1PolicyServiceInstance_0**
> PolicyServiceInstance ReadTier1PolicyServiceInstance_0(ctx, tier1Id, localeServiceId, serviceInstanceId)
Read Tier1 service instance

Read Tier1 service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**| Tier-1 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

[**PolicyServiceInstance**](PolicyServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVirtualEndpoint**
> VirtualEndpoint ReadVirtualEndpoint(ctx, tier0Id, localeServiceId, virtualEndpointId)
Read virtual endpoint

Read virtual endpoint with given id under given Tier0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

[**VirtualEndpoint**](VirtualEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVirtualEndpoint_0**
> VirtualEndpoint ReadVirtualEndpoint_0(ctx, tier0Id, localeServiceId, virtualEndpointId)
Read virtual endpoint

Read virtual endpoint with given id under given Tier0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **virtualEndpointId** | **string**| Virtual endpoint id | 

### Return type

[**VirtualEndpoint**](VirtualEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewAuthenticationTokensForPolicyServiceInstanceReauth**
> RenewAuthenticationTokensForPolicyServiceInstanceReauth(ctx, tier0Id, localeServiceId, serviceInstanceId)
Renew the authentication tokens

Use this API when an alarm complaining JWT expiry is raised while deploying partner service VM. The OVF for partner service needs to be downloaded from partner services provider. It might be possible that the authentication token for this communication is expired when the service VM deployment starts. That will either require re-login through UI or use of this API. Certain authentication and authorization steps are internally processed in order to enable communication with partner service provider. This API offers the functionality to re-establish communication with partner services provider. This API needs open id and access token to be passed as headers. Those can be obtained from CSP authorize API. Please make sure to pass headers - Authorization:<Bearer ACCESS_TOKEN> and X-NSX-OpenId:<OPEN_ID>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewAuthenticationTokensForPolicyServiceInstanceReauth_0**
> RenewAuthenticationTokensForPolicyServiceInstanceReauth_0(ctx, tier0Id, localeServiceId, serviceInstanceId)
Renew the authentication tokens

Use this API when an alarm complaining JWT expiry is raised while deploying partner service VM. The OVF for partner service needs to be downloaded from partner services provider. It might be possible that the authentication token for this communication is expired when the service VM deployment starts. That will either require re-login through UI or use of this API. Certain authentication and authorization steps are internally processed in order to enable communication with partner service provider. This API offers the functionality to re-establish communication with partner services provider. This API needs open id and access token to be passed as headers. Those can be obtained from CSP authorize API. Please make sure to pass headers - Authorization:<Bearer ACCESS_TOKEN> and X-NSX-OpenId:<OPEN_ID>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**| Tier-0 id | 
  **localeServiceId** | **string**| Locale service id | 
  **serviceInstanceId** | **string**| Service instance id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetRuleStatsReset**
> ResetRuleStatsReset(ctx, category, optional)
Reset firewall rule statistics

Sets firewall rule statistics counter to zero. This operation is supported for given category, for example: DFW i.e. for all layer3 firewall (transport nodes only) rules or EDGE i.e. for all layer3 edge firewall (edge nodes only) rules. - no enforcement point path specified:   On global manager, it is mandatory to give an enforcement point path.   On local manager, reset of stats will be executed for each enforcement point. - {enforcement_point_path}: Reset of stats will be executed only for the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | **string**| Aggregation statistic category | 
 **optional** | ***PolicySecurityApiResetRuleStatsResetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiResetRuleStatsResetOpts struct
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

# **ResetRuleStatsReset_0**
> ResetRuleStatsReset_0(ctx, category, optional)
Reset firewall rule statistics

Sets firewall rule statistics counter to zero. This operation is supported for given category, for example: DFW i.e. for all layer3 firewall (transport nodes only) rules or EDGE i.e. for all layer3 edge firewall (edge nodes only) rules. - no enforcement point path specified:   On global manager, it is mandatory to give an enforcement point path.   On local manager, reset of stats will be executed for each enforcement point. - {enforcement_point_path}: Reset of stats will be executed only for the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | **string**| Aggregation statistic category | 
 **optional** | ***PolicySecurityApiResetRuleStatsReset_223Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiResetRuleStatsReset_223Opts struct
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

# **ReviseCommunicationEntryRevise**
> CommunicationEntry ReviseCommunicationEntryRevise(ctx, body, domainId, communicationMapId, communicationEntryId, optional)
Revise the positioning of communication entry

This is used to re-order a communictation entry within a communication map. This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id?action=revise 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseCommunicationEntryReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseCommunicationEntryReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The communication map/communication entry path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseCommunicationEntryRevise_0**
> CommunicationEntry ReviseCommunicationEntryRevise_0(ctx, body, domainId, communicationMapId, communicationEntryId, optional)
Revise the positioning of communication entry

This is used to re-order a communictation entry within a communication map. This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id?action=revise 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseCommunicationEntryRevise_224Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseCommunicationEntryRevise_224Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The communication map/communication entry path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseCommunicationMapsRevise**
> CommunicationMap ReviseCommunicationMapsRevise(ctx, body, domainId, communicationMapId, optional)
Revise the positioning of communication maps

This is used to set a precedence of a communication map w.r.t others. This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id?action=revise 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseCommunicationMapsReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseCommunicationMapsReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The communication map/communication entry path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseCommunicationMapsRevise_0**
> CommunicationMap ReviseCommunicationMapsRevise_0(ctx, body, domainId, communicationMapId, optional)
Revise the positioning of communication maps

This is used to set a precedence of a communication map w.r.t others. This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id?action=revise 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseCommunicationMapsRevise_225Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseCommunicationMapsRevise_225Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The communication map/communication entry path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseGatewayPolicyRevise**
> GatewayPolicy ReviseGatewayPolicyRevise(ctx, body, domainId, gatewayPolicyId, optional)
Revise the positioning of gateway policy

This is used to set a precedence of a gateway policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseGatewayPolicyReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseGatewayPolicyReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseGatewayPolicyRevise_0**
> GatewayPolicy ReviseGatewayPolicyRevise_0(ctx, body, domainId, gatewayPolicyId, optional)
Revise the positioning of gateway policy

This is used to set a precedence of a gateway policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayPolicy**](GatewayPolicy.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseGatewayPolicyRevise_226Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseGatewayPolicyRevise_226Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**GatewayPolicy**](GatewayPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseGatewayRuleRevise**
> Rule ReviseGatewayRuleRevise(ctx, body, domainId, gatewayPolicyId, ruleId, optional)
Revise the positioning of gateway rule

This is used to re-order a gateway rule within a gateway policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseGatewayRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseGatewayRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseGatewayRuleRevise_0**
> Rule ReviseGatewayRuleRevise_0(ctx, body, domainId, gatewayPolicyId, ruleId, optional)
Revise the positioning of gateway rule

This is used to re-order a gateway rule within a gateway policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **gatewayPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseGatewayRuleRevise_227Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseGatewayRuleRevise_227Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseIdsRuleRevise**
> IdsRule ReviseIdsRuleRevise(ctx, body, domainId, policyId, ruleId, optional)
Revise the positioning of IDS rule

This is used to re-order a rule within a security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseIdsRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseIdsRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseIdsRuleRevise_0**
> IdsRule ReviseIdsRuleRevise_0(ctx, body, domainId, policyId, ruleId, optional)
Revise the positioning of IDS rule

This is used to re-order a rule within a security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsRule**](IdsRule.md)|  | 
  **domainId** | **string**|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseIdsRuleRevise_228Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseIdsRuleRevise_228Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**IdsRule**](IdsRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseIdsSecurityPolicyRevise**
> IdsSecurityPolicy ReviseIdsSecurityPolicyRevise(ctx, body, domainId, policyId, optional)
Revise the positioning of IDS security policies

This is used to set a precedence of a security policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **policyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseIdsSecurityPolicyReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseIdsSecurityPolicyReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseIdsSecurityPolicyRevise_0**
> IdsSecurityPolicy ReviseIdsSecurityPolicyRevise_0(ctx, body, domainId, policyId, optional)
Revise the positioning of IDS security policies

This is used to set a precedence of a security policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSecurityPolicy**](IdsSecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **policyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseIdsSecurityPolicyRevise_229Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseIdsSecurityPolicyRevise_229Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**IdsSecurityPolicy**](IdsSecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSecurityPoliciesRevise**
> SecurityPolicy ReviseSecurityPoliciesRevise(ctx, body, domainId, securityPolicyId, optional)
Revise the positioning of security policies

This is used to set a precedence of a security policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseSecurityPoliciesReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseSecurityPoliciesReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSecurityPoliciesRevise_0**
> SecurityPolicy ReviseSecurityPoliciesRevise_0(ctx, body, domainId, securityPolicyId, optional)
Revise the positioning of security policies

This is used to set a precedence of a security policy w.r.t others. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseSecurityPoliciesRevise_230Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseSecurityPoliciesRevise_230Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSecurityRuleRevise**
> Rule ReviseSecurityRuleRevise(ctx, body, domainId, securityPolicyId, ruleId, optional)
Revise the positioning of rule

This is used to re-order a rule within a security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseSecurityRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseSecurityRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSecurityRuleRevise_0**
> Rule ReviseSecurityRuleRevise_0(ctx, body, domainId, securityPolicyId, ruleId, optional)
Revise the positioning of rule

This is used to re-order a rule within a security policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***PolicySecurityApiReviseSecurityRuleRevise_231Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicySecurityApiReviseSecurityRuleRevise_231Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **anchorPath** | **optional.**| The security policy/rule path if operation is &#x27;insert_after&#x27; or &#x27;insert_before&#x27;  | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCPUMemThresholdsProfile**
> PolicyFirewallCpuMemThresholdsProfile UpdateCPUMemThresholdsProfile(ctx, body, profileId)
Create or update CPU and memory thresholds profile

Create or update CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCPUMemThresholdsProfile_0**
> PolicyFirewallCpuMemThresholdsProfile UpdateCPUMemThresholdsProfile_0(ctx, body, profileId)
Create or update CPU and memory thresholds profile

Create or update CPU and memory thresholds profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfile**](PolicyFirewallCpuMemThresholdsProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommunicationEntry**
> CommunicationEntry UpdateCommunicationEntry(ctx, body, domainId, communicationMapId, communicationEntryId)
Create or update a CommunicationEntry

Update the CommunicationEntry. If a CommunicationEntry with the communication-entry-id is not already present, this API fails with a 404. Creation of CommunicationEntries is not allowed using this API. This API is deprecated. Please use the following API instead PUT /infra/domains/domain-id/security-policies/securit-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommunicationEntry_0**
> CommunicationEntry UpdateCommunicationEntry_0(ctx, body, domainId, communicationMapId, communicationEntryId)
Create or update a CommunicationEntry

Update the CommunicationEntry. If a CommunicationEntry with the communication-entry-id is not already present, this API fails with a 404. Creation of CommunicationEntries is not allowed using this API. This API is deprecated. Please use the following API instead PUT /infra/domains/domain-id/security-policies/securit-policy-id/rules/rule-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationEntry**](CommunicationEntry.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 
  **communicationEntryId** | **string**|  | 

### Return type

[**CommunicationEntry**](CommunicationEntry.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommunicationMapForDomain**
> CommunicationMap UpdateCommunicationMapForDomain(ctx, body, domainId, communicationMapId)
Create or Update communication map

Create or Update the communication map for a domain. This is a full replace. All the CommunicationEntries are replaced. This API is deprecated. Please use the following API instead. PUT /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommunicationMapForDomain_0**
> CommunicationMap UpdateCommunicationMapForDomain_0(ctx, body, domainId, communicationMapId)
Create or Update communication map

Create or Update the communication map for a domain. This is a full replace. All the CommunicationEntries are replaced. This API is deprecated. Please use the following API instead. PUT /infra/domains/domain-id/security-policies/security-policy-id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommunicationMap**](CommunicationMap.md)|  | 
  **domainId** | **string**|  | 
  **communicationMapId** | **string**|  | 

### Return type

[**CommunicationMap**](CommunicationMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsSecurityProfile**
> DnsSecurityProfile UpdateDnsSecurityProfile(ctx, body, profileId)
Create or update DNS security profile

Create or update DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfile**](DnsSecurityProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**DnsSecurityProfile**](DnsSecurityProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsSecurityProfileBinding**
> DnsSecurityProfileBindingMap UpdateDnsSecurityProfileBinding(ctx, body, domainId, groupId, dnsSecurityProfileBindingMapId)
Update DNS security profile binding map

API will update DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

[**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsSecurityProfileBinding_0**
> DnsSecurityProfileBindingMap UpdateDnsSecurityProfileBinding_0(ctx, body, domainId, groupId, dnsSecurityProfileBindingMapId)
Update DNS security profile binding map

API will update DNS security profile binding map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)|  | 
  **domainId** | **string**| Domain ID | 
  **groupId** | **string**| Group ID | 
  **dnsSecurityProfileBindingMapId** | **string**| DNS security profile binding map ID | 

### Return type

[**DnsSecurityProfileBindingMap**](DnsSecurityProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsSecurityProfile_0**
> DnsSecurityProfile UpdateDnsSecurityProfile_0(ctx, body, profileId)
Create or update DNS security profile

Create or update DNS security profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsSecurityProfile**](DnsSecurityProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**DnsSecurityProfile**](DnsSecurityProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFloodProtectionProfile**
> FloodProtectionProfile UpdateFloodProtectionProfile(ctx, body, floodProtectionProfileId)
Update Firewall Flood Protection Profile

API will update Firewall Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfile**](FloodProtectionProfile.md)|  | 
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

[**FloodProtectionProfile**](FloodProtectionProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFloodProtectionProfile_0**
> FloodProtectionProfile UpdateFloodProtectionProfile_0(ctx, body, floodProtectionProfileId)
Update Firewall Flood Protection Profile

API will update Firewall Flood Protection Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfile**](FloodProtectionProfile.md)|  | 
  **floodProtectionProfileId** | **string**| Flood Protection Profile ID | 

### Return type

[**FloodProtectionProfile**](FloodProtectionProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMonitoringBinding**
> GroupMonitoringProfileBindingMap UpdateGroupMonitoringBinding(ctx, body, domainId, groupId, groupMonitoringProfileBindingMapId)
Update Group Monitoring Profile Binding Map

API will update Group Monitoring Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

[**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupMonitoringBinding_0**
> GroupMonitoringProfileBindingMap UpdateGroupMonitoringBinding_0(ctx, body, domainId, groupId, groupMonitoringProfileBindingMapId)
Update Group Monitoring Profile Binding Map

API will update Group Monitoring Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **groupMonitoringProfileBindingMapId** | **string**| Group Monitoring Profile Binding Map ID | 

### Return type

[**GroupMonitoringProfileBindingMap**](GroupMonitoringProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdsSettings**
> IdsSettings UpdateIdsSettings(ctx, body)
Update Intrusion detection system settings

Intrusion detection system settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSettings**](IdsSettings.md)|  | 

### Return type

[**IdsSettings**](IdsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdsSettings_0**
> IdsSettings UpdateIdsSettings_0(ctx, body)
Update Intrusion detection system settings

Intrusion detection system settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsSettings**](IdsSettings.md)|  | 

### Return type

[**IdsSettings**](IdsSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdsSignaturesUpdateSignatures**
> UpdateIdsSignaturesUpdateSignatures(ctx, )
Download and update IDS signatures

Trigger the process to Download and update the IDS signatures manually. 

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

# **UpdateIdsSignaturesUpdateSignatures_0**
> UpdateIdsSignaturesUpdateSignatures_0(ctx, )
Download and update IDS signatures

Trigger the process to Download and update the IDS signatures manually. 

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

# **UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap**
> PolicyFirewallCpuMemThresholdsProfileBindingMap UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap(ctx, body, cpuMemThresholdsProfileBindingMapId)
Update Firewall CPU Memory Thresholds Profile Binding Map

API will update Firewall CPU Memory Thresholds Profile Binding Map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCpuMemThresholdsProfileBindingMap.md)|  | 
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCPUMemThresholdsProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap_0**
> PolicyFirewallCpuMemThresholdsProfileBindingMap UpdatePolicyFirewallCPUMemThresholdsProfileBindingMap_0(ctx, body, cpuMemThresholdsProfileBindingMapId)
Update Firewall CPU Memory Thresholds Profile Binding Map

API will update Firewall CPU Memory Thresholds Profile Binding Map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCpuMemThresholdsProfileBindingMap.md)|  | 
  **cpuMemThresholdsProfileBindingMapId** | **string**| Firewall CPU Memory Thresholds Profile Binding Map ID | 

### Return type

[**PolicyFirewallCpuMemThresholdsProfileBindingMap**](PolicyFirewallCPUMemThresholdsProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallFloodProtectionBinding**
> PolicyFirewallFloodProtectionProfileBindingMap UpdatePolicyFirewallFloodProtectionBinding(ctx, body, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Update Firewall Flood Protection Profile Binding Map

API will update Firewall Flood Protection Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallFloodProtectionBinding_0**
> PolicyFirewallFloodProtectionProfileBindingMap UpdatePolicyFirewallFloodProtectionBinding_0(ctx, body, domainId, groupId, firewallFloodProtectionProfileBindingMapId)
Update Firewall Flood Protection Profile Binding Map

API will update Firewall Flood Protection Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **firewallFloodProtectionProfileBindingMapId** | **string**| Firewall Flood Protection Profile Binding Map ID | 

### Return type

[**PolicyFirewallFloodProtectionProfileBindingMap**](PolicyFirewallFloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallScheduler**
> PolicyFirewallScheduler UpdatePolicyFirewallScheduler(ctx, body, firewallSchedulerId)
Create or Update PolicyFirewallScheduler

Updates a PolicyFirewallScheduler, which can be set at security policy. Note that at least one property out of \"days\", \"start_date\", \"time_interval\", \"end_date\" is required if \"recurring\" field is true. Also \"start_time\" and \"end_time\" should not be present. And if \"recurring\" field is false then \"start_date\" and \"end_date\" is mandatory, \"start_time\" and \"end_time\" is optional. Also the fields \"days\" and \"time_interval\" should not be present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)|  | 
  **firewallSchedulerId** | **string**|  | 

### Return type

[**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallScheduler_0**
> PolicyFirewallScheduler UpdatePolicyFirewallScheduler_0(ctx, body, firewallSchedulerId)
Create or Update PolicyFirewallScheduler

Updates a PolicyFirewallScheduler, which can be set at security policy. Note that at least one property out of \"days\", \"start_date\", \"time_interval\", \"end_date\" is required if \"recurring\" field is true. Also \"start_time\" and \"end_time\" should not be present. And if \"recurring\" field is false then \"start_date\" and \"end_date\" is mandatory, \"start_time\" and \"end_time\" is optional. Also the fields \"days\" and \"time_interval\" should not be present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)|  | 
  **firewallSchedulerId** | **string**|  | 

### Return type

[**PolicyFirewallScheduler**](PolicyFirewallScheduler.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallSessionTimerBinding**
> PolicyFirewallSessionTimerProfileBindingMap UpdatePolicyFirewallSessionTimerBinding(ctx, body, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Update Firewall Session Timer Profile Binding Map

API will update Firewall Session Timer Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallSessionTimerBinding_0**
> PolicyFirewallSessionTimerProfileBindingMap UpdatePolicyFirewallSessionTimerBinding_0(ctx, body, domainId, groupId, firewallSessionTimerProfileBindingMapId)
Update Firewall Session Timer Profile Binding Map

API will update Firewall Session Timer Profile Binding Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)|  | 
  **domainId** | **string**| DomainID | 
  **groupId** | **string**| Group ID | 
  **firewallSessionTimerProfileBindingMapId** | **string**| Firewall Session Timer Profile Binding Map ID | 

### Return type

[**PolicyFirewallSessionTimerProfileBindingMap**](PolicyFirewallSessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallSessionTimerProfile**
> PolicyFirewallSessionTimerProfile UpdatePolicyFirewallSessionTimerProfile(ctx, body, firewallSessionTimerProfileId)
Update Firewall Session Timer Profile

API will update Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)|  | 
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

[**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyFirewallSessionTimerProfile_0**
> PolicyFirewallSessionTimerProfile UpdatePolicyFirewallSessionTimerProfile_0(ctx, body, firewallSessionTimerProfileId)
Update Firewall Session Timer Profile

API will update Firewall Session Timer Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)|  | 
  **firewallSessionTimerProfileId** | **string**| Firewall Session Timer Profile ID | 

### Return type

[**PolicyFirewallSessionTimerProfile**](PolicyFirewallSessionTimerProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityPolicyForDomain**
> SecurityPolicy UpdateSecurityPolicyForDomain(ctx, body, domainId, securityPolicyId)
Create or Update security policy

Create or Update the security policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityPolicyForDomain_0**
> SecurityPolicy UpdateSecurityPolicyForDomain_0(ctx, body, domainId, securityPolicyId)
Create or Update security policy

Create or Update the security policy for a domain. This is a full replace. All the rules are replaced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityPolicy**](SecurityPolicy.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 

### Return type

[**SecurityPolicy**](SecurityPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityRule**
> Rule UpdateSecurityRule(ctx, body, domainId, securityPolicyId, ruleId)
Create or update a rule

Update the rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityRule_0**
> Rule UpdateSecurityRule_0(ctx, body, domainId, securityPolicyId, ruleId)
Create or update a rule

Update the rule. Create new rule if a rule with the rule-id is not already present. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **domainId** | **string**|  | 
  **securityPolicyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceDefinition**
> ServiceDefinition UpdateServiceDefinition(ctx, body, enforcementPointId, serviceDefinitionId)
Update an existing Service Definition on the given enforcement point 

Update an existing Service Definition on the given enforcement point. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDefinition**](ServiceDefinition.md)|  | 
  **enforcementPointId** | **string**| Enforcement point id | 
  **serviceDefinitionId** | **string**| Id of service definition | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0FloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap UpdateTier0FloodProtectionProfileBinding(ctx, body, tier0Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router

API will create or update Flood Protection profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0FloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap UpdateTier0FloodProtectionProfileBinding_0(ctx, body, tier0Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router

API will create or update Flood Protection profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0LocaleServicesFloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap UpdateTier0LocaleServicesFloodProtectionProfileBinding(ctx, body, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0LocaleServicesFloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap UpdateTier0LocaleServicesFloodProtectionProfileBinding_0(ctx, body, tier0Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0LocaleServicesSessionTimerProfileBinding**
> SessionTimerProfileBindingMap UpdateTier0LocaleServicesSessionTimerProfileBinding(ctx, body, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0LocaleServicesSessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap UpdateTier0LocaleServicesSessionTimerProfileBinding_0(ctx, body, tier0Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-0 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0SessionTimerProfileBinding**
> SessionTimerProfileBindingMap UpdateTier0SessionTimerProfileBinding(ctx, body, tier0Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router

API will create or update Session Timer profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier0SessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap UpdateTier0SessionTimerProfileBinding_0(ctx, body, tier0Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-0 Logical Router

API will create or update Session Timer profile binding map for Tier-0 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier0Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1FloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap UpdateTier1FloodProtectionProfileBinding(ctx, body, tier1Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router

API will create or update Flood Protection profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1FloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap UpdateTier1FloodProtectionProfileBinding_0(ctx, body, tier1Id, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router

API will create or update Flood Protection profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1LocaleServicesFloodProtectionProfileBinding**
> FloodProtectionProfileBindingMap UpdateTier1LocaleServicesFloodProtectionProfileBinding(ctx, body, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1LocaleServicesFloodProtectionProfileBinding_0**
> FloodProtectionProfileBindingMap UpdateTier1LocaleServicesFloodProtectionProfileBinding_0(ctx, body, tier1Id, localeServicesId, floodProtectionProfileBindingId)
Create or update Flood Protection Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Flood Protection profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **floodProtectionProfileBindingId** | **string**|  | 

### Return type

[**FloodProtectionProfileBindingMap**](FloodProtectionProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1LocaleServicesSessionTimerProfileBinding**
> SessionTimerProfileBindingMap UpdateTier1LocaleServicesSessionTimerProfileBinding(ctx, body, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1LocaleServicesSessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap UpdateTier1LocaleServicesSessionTimerProfileBinding_0(ctx, body, tier1Id, localeServicesId, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router LocaleServices

API will create or update Session Timer profile binding map for Tier-1 Logical Router LocaleServices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1SessionTimerProfileBinding**
> SessionTimerProfileBindingMap UpdateTier1SessionTimerProfileBinding(ctx, body, tier1Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router

API will create or update Session Timer profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTier1SessionTimerProfileBinding_0**
> SessionTimerProfileBindingMap UpdateTier1SessionTimerProfileBinding_0(ctx, body, tier1Id, sessionTimerProfileBindingId)
Create or update Session Timer Profile Binding Map for Tier-1 Logical Router

API will create or update Session Timer profile binding map for Tier-1 Logical Router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)|  | 
  **tier1Id** | **string**|  | 
  **sessionTimerProfileBindingId** | **string**|  | 

### Return type

[**SessionTimerProfileBindingMap**](SessionTimerProfileBindingMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier0GatewayFirewall**
> GatewayPolicyListResult ViewTier0GatewayFirewall(ctx, tier0Id)
Get list of gateway policies with rules that belong to the specific Tier-0 logical router. 

Get filtered view of gateway rules associated with the Tier-0. The gateay policies are returned in the order of category and precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier0GatewayFirewall_0**
> GatewayPolicyListResult ViewTier0GatewayFirewall_0(ctx, tier0Id)
Get list of gateway policies with rules that belong to the specific Tier-0 logical router. 

Get filtered view of gateway rules associated with the Tier-0. The gateay policies are returned in the order of category and precedence. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier0LocaleServicesGatewayFirewall**
> GatewayPolicyListResult ViewTier0LocaleServicesGatewayFirewall(ctx, tier0Id, localeServicesId)
Get list of gateway policies with rules that belong to the specific Tier-0 LocalServices. 

Get filtered view of Gateway Firewall rules associated with the Tier-0 Locale Services. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier0LocaleServicesGatewayFirewall_0**
> GatewayPolicyListResult ViewTier0LocaleServicesGatewayFirewall_0(ctx, tier0Id, localeServicesId)
Get list of gateway policies with rules that belong to the specific Tier-0 LocalServices. 

Get filtered view of Gateway Firewall rules associated with the Tier-0 Locale Services. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier0Id** | **string**|  | 
  **localeServicesId** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier1GatewayFirewall**
> GatewayPolicyListResult ViewTier1GatewayFirewall(ctx, tier1Id)
Get list of gateway policies with rules that belong to the specific Tier-1. 

Get filtered view of Gateway Firewall rules associated with the Tier-1. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier1GatewayFirewall_0**
> GatewayPolicyListResult ViewTier1GatewayFirewall_0(ctx, tier1Id)
Get list of gateway policies with rules that belong to the specific Tier-1. 

Get filtered view of Gateway Firewall rules associated with the Tier-1. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier1LocaleServicesGatewayFirewall**
> GatewayPolicyListResult ViewTier1LocaleServicesGatewayFirewall(ctx, tier1Id, localeServicesId)
Get list of gateway policies with rules that belong to the specific Tier-1 LocalServices. 

Get filtered view of Gateway Firewall rules associated with the Tier-1 Locale Services. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewTier1LocaleServicesGatewayFirewall_0**
> GatewayPolicyListResult ViewTier1LocaleServicesGatewayFirewall_0(ctx, tier1Id, localeServicesId)
Get list of gateway policies with rules that belong to the specific Tier-1 LocalServices. 

Get filtered view of Gateway Firewall rules associated with the Tier-1 Locale Services. The gateway policies are returned in the order of category and sequence number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tier1Id** | **string**|  | 
  **localeServicesId** | **string**|  | 

### Return type

[**GatewayPolicyListResult**](GatewayPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

