# LbJwtAuthAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The property identifies the load balancer rule action type.  | [default to null]
**Tokens** | **[]string** | JWT is an open standard that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. Load balancer will search for every specified tokens one by one for the jwt message until found. This parameter is optional. In case not found or this field is not configured, load balancer searches the Bearer header by default in the http request \&quot;Authorization: Bearer &amp;lt;token&amp;gt;\&quot;.  | [optional] [default to null]
**PassJwtToPool** | **bool** | Specify whether to pass the JWT to backend server or remove it. By default, it is false which means will not pass the JWT to backend servers.  | [optional] [default to false]
**Realm** | **string** | A description of the protected area. If no realm is specified, clients often display a formatted hostname instead. The configured realm is returned when client request is rejected with 401 http status. In the response, it will be \&quot;WWW-Authentication: Bearer realm&#x3D;&amp;lt;realm&amp;gt;\&quot;.  | [optional] [default to null]
**Key** | [***LbJwtKey**](LBJwtKey.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

