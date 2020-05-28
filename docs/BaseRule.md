# BaseRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**Disabled** | **bool** | Flag to disable the rule. Default is enabled. | [optional] [default to false]
**Direction** | **string** | Define direction of traffic.  | [optional] [default to DIRECTION.IN_OUT]
**IpProtocol** | **string** | Type of IP packet that should be matched while enforcing the rule. The value is set to IPV4_IPV6 for Layer3 rule if not specified. For Layer2/Ether rule the value must be null.  | [optional] [default to null]
**Notes** | **string** | Text for additional notes on changes. | [optional] [default to null]
**Logged** | **bool** | Flag to enable packet logging. Default is disabled. | [optional] [default to false]
**Profiles** | **[]string** | Holds the list of layer 7 service profile paths. These profiles accept attributes and sub-attributes of various network services (e.g. L4 AppId, encryption algorithm, domain name, etc) as key value pairs.  | [optional] [default to null]
**RuleId** | **int64** | This is a unique 4 byte positive number that is assigned by the system.  This rule id is passed all the way down to the data path. The first 1GB (1000 to 2^30) will be shared by GM and LM with zebra style striped number space. For E.g 1000 to (1Million -1) by LM, (1M - 2M-1) by GM and so on.  | [optional] [default to null]
**IsDefault** | **bool** | A flag to indicate whether rule is a default rule. | [optional] [default to null]
**Tag** | **string** | User level field which will be printed in CLI and packet logs.  | [optional] [default to null]
**SourceGroups** | **[]string** | We need paths as duplicate names may exist for groups under different domains. Along with paths we support IP Address of type IPv4 and IPv6. IP Address can be in one of the format(CIDR, IP Address, Range of IP Address). In order to specify all groups, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]
**DestinationGroups** | **[]string** | We need paths as duplicate names may exist for groups under different domains. Along with paths we support IP Address of type IPv4 and IPv6. IP Address can be in one of the format(CIDR, IP Address, Range of IP Address). In order to specify all groups, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]
**Services** | **[]string** | In order to specify all services, use the constant \&quot;ANY\&quot;. This is case insensitive. If \&quot;ANY\&quot; is used, it should be the ONLY element in the services array. Error will be thrown if ANY is used in conjunction with other values.  | [optional] [default to null]
**Scope** | **[]string** | The list of policy paths where the rule is applied LR/Edge/T0/T1/LRP etc. Note that a given rule can be applied on multiple LRs/LRPs.  | [optional] [default to null]
**ServiceEntries** | [**[]ServiceEntry**](ServiceEntry.md) | In order to specify raw services this can be used, along with services which contains path to services. This can be empty or null.  | [optional] [default to null]
**DestinationsExcluded** | **bool** | If set to true, the rule gets applied on all the groups that are NOT part of the destination groups. If false, the rule applies to the destination groups  | [optional] [default to false]
**SequenceNumber** | **int32** | This field is used to resolve conflicts between multiple Rules under Security or Gateway Policy for a Domain If no sequence number is specified in the payload, a value of 0 is assigned by default. If there are multiple rules with the same sequence number then their order is not deterministic. If a specific order of rules is desired, then one has to specify unique sequence numbers or use the POST request on the rule entity with a query parameter action&#x3D;revise to let the framework assign a sequence number  | [optional] [default to null]
**SourcesExcluded** | **bool** | If set to true, the rule gets applied on all the groups that are NOT part of the source groups. If false, the rule applies to the source groups  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

