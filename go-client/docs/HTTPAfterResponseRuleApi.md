# \HTTPAfterResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPAfterResponseRule**](HTTPAfterResponseRuleApi.md#CreateHTTPAfterResponseRule) | **Post** /services/haproxy/configuration/http_after_response_rules | Add a new HTTP After Response Rule
[**DeleteHTTPAfterResponseRule**](HTTPAfterResponseRuleApi.md#DeleteHTTPAfterResponseRule) | **Delete** /services/haproxy/configuration/http_after_response_rules/{index} | Delete a HTTP After Response Rule
[**GetHTTPAfterResponseRule**](HTTPAfterResponseRuleApi.md#GetHTTPAfterResponseRule) | **Get** /services/haproxy/configuration/http_after_response_rules/{index} | Return one HTTP After Response Rule
[**GetHTTPAfterResponseRules**](HTTPAfterResponseRuleApi.md#GetHTTPAfterResponseRules) | **Get** /services/haproxy/configuration/http_after_response_rules | Return an array of all HTTP After Response Rules
[**ReplaceHTTPAfterResponseRule**](HTTPAfterResponseRuleApi.md#ReplaceHTTPAfterResponseRule) | **Put** /services/haproxy/configuration/http_after_response_rules/{index} | Replace a HTTP After Response Rule


# **CreateHTTPAfterResponseRule**
> HttpAfterResponseRule CreateHTTPAfterResponseRule(ctx, parentName, parentType, data, optional)
Add a new HTTP After Response Rule

Adds a new HTTP After Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPAfterResponseRule**
> DeleteHTTPAfterResponseRule(ctx, index, parentName, parentType, optional)
Delete a HTTP After Response Rule

Deletes a HTTP After Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleOpts struct

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

# **GetHTTPAfterResponseRule**
> HttpAfterResponseRule GetHTTPAfterResponseRule(ctx, index, parentName, parentType, optional)
Return one HTTP After Response Rule

Returns one HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPAfterResponseRules**
> HttpAfterResponseRules GetHTTPAfterResponseRules(ctx, parentName, parentType, optional)
Return an array of all HTTP After Response Rules

Returns all HTTP After Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPAfterResponseRuleApiGetHTTPAfterResponseRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetHTTPAfterResponseRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRules**](http_after_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPAfterResponseRule**
> HttpAfterResponseRule ReplaceHTTPAfterResponseRule(ctx, index, parentName, parentType, data, optional)
Replace a HTTP After Response Rule

Replaces a HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

