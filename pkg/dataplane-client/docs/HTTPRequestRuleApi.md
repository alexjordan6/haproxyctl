# \HTTPRequestRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPRequestRule**](HTTPRequestRuleApi.md#CreateHTTPRequestRule) | **Post** /services/haproxy/configuration/http_request_rules | Add a new HTTP Request Rule
[**DeleteHTTPRequestRule**](HTTPRequestRuleApi.md#DeleteHTTPRequestRule) | **Delete** /services/haproxy/configuration/http_request_rules/{index} | Delete a HTTP Request Rule
[**GetHTTPRequestRule**](HTTPRequestRuleApi.md#GetHTTPRequestRule) | **Get** /services/haproxy/configuration/http_request_rules/{index} | Return one HTTP Request Rule
[**GetHTTPRequestRules**](HTTPRequestRuleApi.md#GetHTTPRequestRules) | **Get** /services/haproxy/configuration/http_request_rules | Return an array of all HTTP Request Rules
[**ReplaceHTTPRequestRule**](HTTPRequestRuleApi.md#ReplaceHTTPRequestRule) | **Put** /services/haproxy/configuration/http_request_rules/{index} | Replace a HTTP Request Rule


# **CreateHTTPRequestRule**
> HttpRequestRule CreateHTTPRequestRule(ctx, parentName, parentType, data, optional)
Add a new HTTP Request Rule

Adds a new HTTP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiCreateHTTPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiCreateHTTPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPRequestRule**
> DeleteHTTPRequestRule(ctx, index, parentName, parentType, optional)
Delete a HTTP Request Rule

Deletes a HTTP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPRequestRuleApiDeleteHTTPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiDeleteHTTPRequestRuleOpts struct

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

# **GetHTTPRequestRule**
> HttpRequestRule GetHTTPRequestRule(ctx, index, parentName, parentType, optional)
Return one HTTP Request Rule

Returns one HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPRequestRuleApiGetHTTPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetHTTPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPRequestRules**
> HttpRequestRules GetHTTPRequestRules(ctx, parentName, parentType, optional)
Return an array of all HTTP Request Rules

Returns all HTTP Request Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPRequestRuleApiGetHTTPRequestRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetHTTPRequestRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRules**](http_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPRequestRule**
> HttpRequestRule ReplaceHTTPRequestRule(ctx, index, parentName, parentType, data, optional)
Replace a HTTP Request Rule

Replaces a HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiReplaceHTTPRequestRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiReplaceHTTPRequestRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

