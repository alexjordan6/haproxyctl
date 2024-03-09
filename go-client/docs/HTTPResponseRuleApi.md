# \HTTPResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPResponseRule**](HTTPResponseRuleApi.md#CreateHTTPResponseRule) | **Post** /services/haproxy/configuration/http_response_rules | Add a new HTTP Response Rule
[**DeleteHTTPResponseRule**](HTTPResponseRuleApi.md#DeleteHTTPResponseRule) | **Delete** /services/haproxy/configuration/http_response_rules/{index} | Delete a HTTP Response Rule
[**GetHTTPResponseRule**](HTTPResponseRuleApi.md#GetHTTPResponseRule) | **Get** /services/haproxy/configuration/http_response_rules/{index} | Return one HTTP Response Rule
[**GetHTTPResponseRules**](HTTPResponseRuleApi.md#GetHTTPResponseRules) | **Get** /services/haproxy/configuration/http_response_rules | Return an array of all HTTP Response Rules
[**ReplaceHTTPResponseRule**](HTTPResponseRuleApi.md#ReplaceHTTPResponseRule) | **Put** /services/haproxy/configuration/http_response_rules/{index} | Replace a HTTP Response Rule


# **CreateHTTPResponseRule**
> HttpResponseRule CreateHTTPResponseRule(ctx, parentName, parentType, data, optional)
Add a new HTTP Response Rule

Adds a new HTTP Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiCreateHTTPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiCreateHTTPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPResponseRule**
> DeleteHTTPResponseRule(ctx, index, parentName, parentType, optional)
Delete a HTTP Response Rule

Deletes a HTTP Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPResponseRuleApiDeleteHTTPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiDeleteHTTPResponseRuleOpts struct

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

# **GetHTTPResponseRule**
> HttpResponseRule GetHTTPResponseRule(ctx, index, parentName, parentType, optional)
Return one HTTP Response Rule

Returns one HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPResponseRuleApiGetHTTPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetHTTPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPResponseRules**
> HttpResponseRules GetHTTPResponseRules(ctx, parentName, parentType, optional)
Return an array of all HTTP Response Rules

Returns all HTTP Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPResponseRuleApiGetHTTPResponseRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetHTTPResponseRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRules**](http_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPResponseRule**
> HttpResponseRule ReplaceHTTPResponseRule(ctx, index, parentName, parentType, data, optional)
Replace a HTTP Response Rule

Replaces a HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiReplaceHTTPResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiReplaceHTTPResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

