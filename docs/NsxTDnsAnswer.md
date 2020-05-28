# NsxTDnsAnswer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Resource type of the DNS forwarder nslookup answer.  | [default to null]
**EnforcementPointPath** | **string** | Policy path referencing the enforcement point from where the DNS forwarder nslookup answer is fetched.  | [optional] [default to null]
**AuthoritativeAnswers** | [**[]NsxTDnsQueryAnswer**](NsxTDnsQueryAnswer.md) | Authoritative answers | [optional] [default to null]
**EdgeNodeId** | **string** | ID of the edge node that performed the query.  | [default to null]
**DnsServer** | **string** | Dns server ip address and port, format is \&quot;ip address#port\&quot;.  | [default to null]
**NonAuthoritativeAnswers** | [**[]NsxTDnsQueryAnswer**](NsxTDnsQueryAnswer.md) | Non authoritative answers | [optional] [default to null]
**RawAnswer** | **string** | It can be NXDOMAIN or error message which is not consisted of authoritative_answer or non_authoritative_answer.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

