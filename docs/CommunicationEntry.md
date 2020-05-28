# CommunicationEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Direction** | **string** | Define direction of traffic.  | [optional] [default to DIRECTION.IN_OUT]
**Notes** | **string** | Text for additional notes on changes. | [optional] [default to null]
**Logged** | **bool** | Flag to enable packet logging. Default is disabled. | [optional] [default to false]
**Disabled** | **bool** | Flag to disable the rule. Default is enabled. | [optional] [default to false]
**Services** | **[]string** | In order to specify all services, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the services array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]
**Tag** | **string** | User level field which will be printed in CLI and packet logs.  | [optional] [default to null]
**DestinationGroups** | **[]string** | We need paths as duplicate names may exist for groups under different domains.In order to specify all groups, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]
**Action** | **string** | The action to be applied to all the services.  | [optional] [default to null]
**Scope** | **[]string** | The list of policy paths where the communication entry is applied Edge/LR/T0/T1/LRP/CGW/MGW/etc. Note that a given rule can be applied on multiple LRs/LRPs.  | [optional] [default to null]
**SequenceNumber** | **int32** | This field is used to resolve conflicts between multiple CommunicationEntries under CommunicationMap for a Domain If no sequence number is specified in the payload, a value of 0 is assigned by default. If there are multiple communication entries with the same sequence number then their order is not deterministic. If a specific order of communication entry is desired, then one has to specify unique sequence numbers or use the POST request on the communication entry entity with a query parameter action&#x3D;revise to let the framework assign a sequence number  | [optional] [default to null]
**SourceGroups** | **[]string** | We need paths as duplicate names may exist for groups under different domains. In order to specify all groups, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

