# LbVariablePersistenceOnAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The property identifies the load balancer rule action type.  | [default to null]
**VariableHashEnabled** | **bool** | The property is used to enable a hash operation for variable value when composing the persistence key.  | [optional] [default to false]
**VariableName** | **string** | The property is the name of variable to be used. It specifies which variable&#x27;s value of a HTTP Request will be used in the key of persistence entry. The variable can be a built-in variable such as \&quot;_cookie_JSESSIONID\&quot;, a customized variable defined in LBVariableAssignmentAction or a captured variable in regular expression such as \&quot;article\&quot;. For the full list of built-in variables, please reference the NSX-T Administrator&#x27;s Guide.  | [default to null]
**PersistenceProfilePath** | **string** | If the persistence profile path is not specified, a default persistence table is created per virtual server. Currently, only LBGenericPersistenceProfile is supported.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

