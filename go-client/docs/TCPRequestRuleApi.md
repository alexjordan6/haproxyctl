# \TCPRequestRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPRequestRule**](TCPRequestRuleApi.md#CreateTCPRequestRule) | **Post** /services/haproxy/configuration/tcp_request_rules | Add a new TCP Request Rule
[**DeleteTCPRequestRule**](TCPRequestRuleApi.md#DeleteTCPRequestRule) | **Delete** /services/haproxy/configuration/tcp_request_rules/{index} | Delete a TCP Request Rule
[**GetTCPRequestRule**](TCPRequestRuleApi.md#GetTCPRequestRule) | **Get** /services/haproxy/configuration/tcp_request_rules/{index} | Return one TCP Request Rule
[**GetTCPRequestRules**](TCPRequestRuleApi.md#GetTCPRequestRules) | **Get** /services/haproxy/configuration/tcp_request_rules | Return an array of all TCP Request Rules
[**ReplaceTCPRequestRule**](TCPRequestRuleApi.md#ReplaceTCPRequestRule) | **Put** /services/haproxy/configuration/tcp_request_rules/{index} | Replace a TCP Request Rule


# **CreateTCPRequestRule**
> TcpRequestRule CreateTCPRequestRule(ctx, parentName, parentType, data, optional)
Add a new TCP Request Rule

Adds a new TCP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiCreateTCPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiCreateTCPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPRequestRule**
> DeleteTCPRequestRule(ctx, index, parentName, parentType, optional)
Delete a TCP Request Rule

Deletes a TCP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPRequestRuleApiDeleteTCPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiDeleteTCPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPRequestRule**
> TcpRequestRule GetTCPRequestRule(ctx, index, parentName, parentType, optional)
Return one TCP Request Rule

Returns one TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPRequestRuleApiGetTCPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetTCPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPRequestRules**
> TcpRequestRules GetTCPRequestRules(ctx, parentName, parentType, optional)
Return an array of all TCP Request Rules

Returns all TCP Request Rules that are configured in specified parent and parent type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPRequestRuleApiGetTCPRequestRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetTCPRequestRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRules**](tcp_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPRequestRule**
> TcpRequestRule ReplaceTCPRequestRule(ctx, index, parentName, parentType, data, optional)
Replace a TCP Request Rule

Replaces a TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiReplaceTCPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiReplaceTCPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

