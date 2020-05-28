# NdSnoopingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NdSnoopingLimit** | **int32** | Maximum number of ND (Neighbor Discovery Protocol) snooped IPv6 addresses  | [optional] [default to 3]
**NdSnoopingEnabled** | **bool** | Enable this method will snoop the NS (Neighbor Solicitation) and NA (Neighbor Advertisement) messages in the ND (Neighbor Discovery Protocol) family of messages which are transmitted by a VM. From the NS messages, we will learn about the source which sent this NS message. From the NA message, we will learn the resolved address in the message which the VM is a recipient of. Addresses snooped by this method are subject to TOFU (Trust on First Use) policies as enforced by the system.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

