# PolicyServiceInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerServiceName** | **string** | Unique name of Partner Service in the Marketplace | [default to null]
**TransportType** | **string** | Transport to be used while deploying Service-VM. | [optional] [default to TRANSPORT_TYPE.L2_BRIDGE]
**DeploymentMode** | **string** | Deployment mode specifies how the partner appliance will be deployed i.e. in HA or standalone mode. | [optional] [default to DEPLOYMENT_MODE.ACTIVE_STANDBY]
**PrimaryInterfaceMgmtIp** | **string** | Management IP Address of primary interface of the Service | [default to null]
**SecondaryPortgroupId** | **string** | Id of the standard or ditsributed port group for secondary management console. Please note that only 1 of the 2 values from 1. secondary_interface_network 2. secondary_portgroup_id are allowed to be passed. Both can&#x27;t be passed in the same request.  | [optional] [default to null]
**ContextId** | **string** | UUID of VCenter/Compute Manager as seen on NSX Manager, to which this service needs to be deployed. | [optional] [default to null]
**PrimaryPortgroupId** | **string** | Id of the standard or ditsributed port group for primary management console. Please note that only 1 of the 2 values from 1. primary_interface_network 2. primary_portgroup_id are allowed to be passed. Both can&#x27;t be passed in the same request.  | [optional] [default to null]
**SecondaryInterfaceMgmtIp** | **string** | Management IP Address of secondary interface of the Service | [optional] [default to null]
**ComputeId** | **string** | Id of the compute(ResourcePool) to which this service needs to be deployed. | [default to null]
**DeploymentSpecName** | **string** | Form factor for the deployment of partner service. | [default to null]
**DeploymentTemplateName** | **string** | Template for the deployment of partnet service. | [default to null]
**SecondaryGatewayAddress** | **string** | Gateway address for secondary management console. If the provided segment already has gateway, this field can be omitted. But if it is provided, it takes precedence always. However, if provided segment does not have gateway, this field must be provided.  | [optional] [default to null]
**StorageId** | **string** | Id of the storage(Datastore). VC moref of Datastore to which this service needs to be deployed. | [default to null]
**SecondarySubnetMask** | **string** | Subnet for secondary management console IP. If the provided segment already has subnet, this field can be omitted. But if it is provided, it takes precedence always. However, if provided segment does not have subnet, this field must be provided.  | [optional] [default to null]
**Attributes** | [**[]Attribute**](Attribute.md) | List of attributes specific to a partner for which the service is created. There attributes are passed on to the partner appliance. | [default to null]
**PrimarySubnetMask** | **string** | Subnet for primary management console IP. If the provided segment already has subnet, this field can be omitted. But if it is provided, it takes precedence always. However, if provided segment does not have subnet, this field must be provided.  | [optional] [default to null]
**PrimaryGatewayAddress** | **string** | Gateway address for primary management console. If the provided segment already has gateway, this field can be omitted. But if it is provided, it takes precedence always. However, if provided segment does not have gateway, this field must be provided.  | [optional] [default to null]
**PrimaryInterfaceNetwork** | **string** | Path of the segment to which primary interface of the Service VM needs to be connected | [optional] [default to null]
**SecondaryInterfaceNetwork** | **string** | Path of segment to which secondary interface of the Service VM needs to be connected | [optional] [default to null]
**FailurePolicy** | **string** | Failure policy for the Service VM. If this values is not provided, it will be defaulted to FAIL_CLOSE. | [optional] [default to FAILURE_POLICY.BLOCK]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

