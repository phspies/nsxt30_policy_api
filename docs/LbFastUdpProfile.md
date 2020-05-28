# LbFastUdpProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | An application profile can be bound to a virtual server to specify the application protocol characteristics. It is used to influence how load balancing is performed. Currently, three types of application profiles are supported: LBFastTCPProfile, LBFastUDPProfile and LBHttpProfile. LBFastTCPProfile or LBFastUDPProfile is typically used when the application is using a custom protocol or a standard protocol not supported by the load balancer. It is also used in cases where the user only wants L4 load balancing mainly because L4 load balancing has much higher performance and scalability, and/or supports connection mirroring. LBHttpProfile is used for both HTTP and HTTPS applications. Though application rules, if bound to the virtual server, can be used to accomplish the same goal, LBHttpProfile is intended to simplify enabling certain common use cases.  | [default to null]
**IdleTimeout** | **int64** | Though UDP is a connectionless protocol, for the purposes of load balancing, all UDP packets with the same flow signature (source and destination IP/ports and IP protocol) received within the idle timeout period are considered to belong to the same connection and are sent to the same backend server. If no packets are received for idle timeout period, the connection (association between flow signature and the selected server) is cleaned up.  | [optional] [default to 300]
**FlowMirroringEnabled** | **bool** | If flow mirroring is enabled, all the flows to the bounded virtual server are mirrored to the standby node.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

