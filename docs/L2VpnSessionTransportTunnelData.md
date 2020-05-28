# L2VpnSessionTransportTunnelData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalAddress** | **string** | IPv4 Address of local endpoint. | [default to null]
**PeerCode** | **string** | Peer code represents a base64 encoded string which has all the configuration for tunnel. E.g local/peer ips and protocol, encryption algorithm, etc. Peer code also contains PSK; be careful when sharing or storing it.  | [optional] [default to null]
**PeerAddress** | **string** | IPv4 Address of Peer endpoint on remote site. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

