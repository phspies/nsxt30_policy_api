# ArpSnoopingConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpBindingLimit** | **int32** | Number of arp snooped IP addresses Indicates the number of arp snooped IP addresses to be remembered per LogicalPort. Decreasing this value, will retain the latest bindings from the existing list of address bindings. Increasing this value will retain existing bindings and also learn any new address bindings discovered on the port until the new limit is reached.  | [optional] [default to 1]
**ArpSnoopingEnabled** | **bool** | Indicates whether ARP snooping is enabled | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

