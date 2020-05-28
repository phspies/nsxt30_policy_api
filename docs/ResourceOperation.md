# ResourceOperation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Event Source resource type.  | [default to null]
**ResourcePointer** | **string** | Regex path representing a regex expression on resources. This regex is used to identify the object(s) that is/are the source of the Event. For instance: specifying \&quot;Lb* | /infra/tier-0s/vmc/ipsec-vpn-services/default\&quot; as a source means that ANY resource starting with Lb or ANY resource with \&quot;/infra/tier-0s/vmc/ipsec-vpn-services/default\&quot; as path would be the source of the event in question.  | [default to null]
**OperationTypes** | **[]string** | Operation types. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

