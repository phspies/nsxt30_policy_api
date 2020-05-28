# SegmentPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**InitState** | **string** | Set initial state when a new logical port is created. &#x27;UNBLOCKED_VLAN&#x27; means new port will be unblocked on traffic in creation, also VLAN will be set with corresponding logical switch setting. This port setting can only be configured at port creation, and cannot be modified.  | [optional] [default to INIT_STATE.VLAN]
**AdminState** | **string** | Represents desired state of the segment port | [optional] [default to ADMIN_STATE.UP]
**Attachment** | [***PortAttachment**](PortAttachment.md) |  | [optional] [default to null]
**ExtraConfigs** | [**[]SegmentExtraConfig**](SegmentExtraConfig.md) | This property could be used for vendor specific configuration in key value string pairs. Segment port setting will override segment setting if the same key was set on both segment and segment port.  | [optional] [default to null]
**IgnoredAddressBindings** | [**[]PortAddressBindingEntry**](PortAddressBindingEntry.md) | IP Discovery module uses various mechanisms to discover address bindings being used on each segment port. If a user would like to ignore any specific discovered address bindings or prevent the discovery of a particular set of discovered bindings, then those address bindings can be provided here. Currently IP range in CIDR format is not supported.  | [optional] [default to null]
**AddressBindings** | [**[]PortAddressBindingEntry**](PortAddressBindingEntry.md) | Static address binding used for the port. | [optional] [default to null]
**SourceSiteId** | **string** | This field will refer to the source site on which the segment port is discovered. This field is populated by GM, when it receives corresponding notification from LM.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

