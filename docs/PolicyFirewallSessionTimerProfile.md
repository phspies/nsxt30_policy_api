# PolicyFirewallSessionTimerProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**TcpClosed** | **int64** | The timeout value of connection in seconds after one endpoint sends an RST. | [default to 20]
**TcpOpening** | **int64** | The timeout value of connection in seconds after a second packet has been transferred. | [default to 30]
**UdpSingle** | **int64** | The timeout value of connection in seconds if the source host sends more than one packet but the destination host has never sent one back. | [default to 30]
**TcpFinwait** | **int64** | The timeout value of connection in seconds after both FINs have been exchanged and connection is closed. | [default to 45]
**TcpFirstPacket** | **int64** | The timeout value of connection in seconds after the first packet has been sent. | [default to 120]
**TcpClosing** | **int64** | The timeout value of connection in seconds after the first FIN has been sent. | [default to 120]
**TcpEstablished** | **int64** | The timeout value of connection in seconds once the connection has become fully established. | [default to 43200]
**UdpMultiple** | **int64** | The timeout value of connection in seconds if both hosts have sent packets. | [default to 60]
**IcmpErrorReply** | **int64** | The timeout value for the connection after an ICMP error came back in response to an ICMP packet. | [default to 10]
**UdpFirstPacket** | **int64** | The timeout value of connection in seconds after the first packet. This will be the initial timeout for the new UDP flow. | [default to 60]
**IcmpFirstPacket** | **int64** | The timeout value of connection in seconds after the first packet. This will be the initial timeout for the new ICMP flow. | [default to 20]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

