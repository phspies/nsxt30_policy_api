# IpSecVpnSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Psk** | **string** | IPSec Pre-shared key. Maximum length of this field is 128 characters. | [optional] [default to null]
**DpdProfilePath** | **string** | Policy path referencing Dead Peer Detection (DPD) profile. Default is set to system default profile. | [optional] [default to null]
**IkeProfilePath** | **string** | Policy path referencing IKE profile to be used. Default is set according to system default profile. | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable IPSec VPN session. | [optional] [default to true]
**ConnectionInitiationMode** | **string** | Connection initiation mode used by local endpoint to establish ike connection with peer site. INITIATOR - In this mode local endpoint initiates tunnel setup and will also respond to incoming tunnel setup requests from peer gateway. RESPOND_ONLY - In this mode, local endpoint shall only respond to incoming tunnel setup requests. It shall not initiate the tunnel setup. ON_DEMAND - In this mode local endpoint will initiate tunnel creation once first packet matching the policy rule is received and will also respond to incoming initiation request.  | [optional] [default to CONNECTION_INITIATION_MODE.INITIATOR]
**LocalEndpointPath** | **string** | Policy path referencing Local endpoint. | [default to null]
**TunnelProfilePath** | **string** | Policy path referencing Tunnel profile to be used. Default is set to system default profile. | [optional] [default to null]
**ComplianceSuite** | **string** | Compliance suite.  | [optional] [default to null]
**TcpMssClamping** | [***TcpMaximumSegmentSizeClamping**](TcpMaximumSegmentSizeClamping.md) |  | [optional] [default to null]
**AuthenticationMode** | **string** | Peer authentication mode. PSK - In this mode a secret key shared between local and peer sites is to be used for authentication. The secret key can be a string with a maximum length of 128 characters. CERTIFICATE - In this mode a certificate defined at the global level is to be used for authentication.  | [optional] [default to AUTHENTICATION_MODE.PSK]
**PeerId** | **string** | Peer ID to uniquely identify the peer site. The peer ID is the public IP address of the remote device terminating the VPN tunnel. When NAT is configured for the peer, enter the private IP address of the peer. | [default to null]
**PeerAddress** | **string** | Public IPV4 address of the remote device terminating the VPN connection. | [default to null]
**ResourceType** | **string** | A Policy Based VPN requires to define protect rules that match   local and peer subnets. IPSec security associations is   negotiated for each pair of local and peer subnet. A Route Based VPN is more flexible, more powerful and recommended over   policy based VPN. IP Tunnel port is created and all traffic routed via   tunnel port is protected. Routes can be configured statically   or can be learned through BGP. A route based VPN is must for establishing   redundant VPN session to remote site.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

