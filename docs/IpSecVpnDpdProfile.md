# IpSecVpnDpdProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ChildPolicyConfigResource**](ChildPolicyConfigResource.md) | subtree for this type within policy tree containing nested elements.  | [optional] [default to null]
**Overridden** | **bool** | Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties.  | [optional] [default to false]
**MarkedForDelete** | **bool** | Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects.  | [optional] [default to false]
**RetryCount** | **int64** | Maximum number of DPD messages&#x27; retry attempts. This value is applicable for both dpd probe modes, periodic and on-demand.  | [optional] [default to 5]
**Enabled** | **bool** | If true, enable dead peer detection. | [optional] [default to true]
**DpdProbeMode** | **string** | DPD probe mode is used to query the liveliness of the peer. Two modes are possible: - PERIODIC: is used to query the liveliness of the peer at regular intervals (dpd_probe_interval). It does not take into consideration traffic coming from the peer. The benefit of this mode over the on-demand mode is earlier detection of dead peers. However, use of periodic DPD incurs extra overhead. When communicating to large numbers of peers, please consider using on-demand DPD instead. - ON_DEMAND: is used to query the liveliness of the peer by instructing the local endpoint to send DPD message to a peer if there is traffic to send to the peer AND the peer was idle for dpd_probe_interval seconds (i.e. there was no traffic from the peer for dpd_probe_interval seconds).  | [optional] [default to DPD_PROBE_MODE.PERIODIC]
**DpdProbeInterval** | **int64** | DPD probe interval defines an interval for DPD probes (in seconds). - When the DPD probe mode is periodic, this interval is the number of seconds between DPD messages. - When the DPD probe mode is on-demand, this interval is the number of seconds during which traffic is not received from the peer before DPD retry messages are sent if there is IPSec traffic to send. For PERIODIC Mode:  Minimum: 3  Maximum: 360  Default: 60 For ON_DEMAND Mode:  Minimum: 1  Maximum: 10  Default: 3  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

