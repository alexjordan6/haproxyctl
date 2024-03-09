# \TCPResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPResponseRule**](TCPResponseRuleApi.md#CreateTCPResponseRule) | **Post** /services/haproxy/configuration/tcp_response_rules | Add a new TCP Response Rule
[**DeleteTCPResponseRule**](TCPResponseRuleApi.md#DeleteTCPResponseRule) | **Delete** /services/haproxy/configuration/tcp_response_rules/{index} | Delete a TCP Response Rule
[**GetTCPResponseRule**](TCPResponseRuleApi.md#GetTCPResponseRule) | **Get** /services/haproxy/configuration/tcp_response_rules/{index} | Return one TCP Response Rule
[**GetTCPResponseRules**](TCPResponseRuleApi.md#GetTCPResponseRules) | **Get** /services/haproxy/configuration/tcp_response_rules | Return an array of all TCP Response Rules
[**ReplaceTCPResponseRule**](TCPResponseRuleApi.md#ReplaceTCPResponseRule) | **Put** /services/haproxy/configuration/tcp_response_rules/{index} | Replace a TCP Response Rule


# **CreateTCPResponseRule**
> TcpResponseRule CreateTCPResponseRule(ctx, backend, data, optional)
Add a new TCP Response Rule

Adds a new TCP Response Rule of the specified type in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backend** | **string**| Parent backend name | 
  **data** | [**TcpResponseRule**](TcpResponseRule.md)|  | 
 **optional** | ***TCPResponseRuleApiCreateTCPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiCreateTCPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPResponseRule**
> DeleteTCPResponseRule(ctx, index, backend, optional)
Delete a TCP Response Rule

Deletes a TCP Response Rule configuration by it's index from the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **backend** | **string**| Parent backend name | 
 **optional** | ***TCPResponseRuleApiDeleteTCPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiDeleteTCPResponseRuleOpts struct

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

# **GetTCPResponseRule**
> TcpResponseRule GetTCPResponseRule(ctx, index, backend, optional)
Return one TCP Response Rule

Returns one TCP Response Rule configuration by it's index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **backend** | **string**| Parent backend name | 
 **optional** | ***TCPResponseRuleApiGetTCPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiGetTCPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPResponseRules**
> TcpResponseRules GetTCPResponseRules(ctx, backend, optional)
Return an array of all TCP Response Rules

Returns all TCP Response Rules that are configured in specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backend** | **string**| Parent backend name | 
 **optional** | ***TCPResponseRuleApiGetTCPResponseRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiGetTCPResponseRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpResponseRules**](tcp_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPResponseRule**
> TcpResponseRule ReplaceTCPResponseRule(ctx, index, backend, data, optional)
Replace a TCP Response Rule

Replaces a TCP Response Rule configuration by it's Index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **backend** | **string**| Parent backend name | 
  **data** | [**TcpResponseRule**](TcpResponseRule.md)|  | 
 **optional** | ***TCPResponseRuleApiReplaceTCPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiReplaceTCPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

