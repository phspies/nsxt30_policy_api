# VirtualMachineDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**ActiveSessions** | [**[]UserSession**](UserSession.md) | List of active (still logged in) user login/session data (no limit). | [optional] [default to null]
**ArchivedSessions** | [**[]UserSession**](UserSession.md) | Optional list of up to 5 most recent archived (previously logged in) user login/session data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

