# \HTTPErrorRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPErrorRule**](HTTPErrorRuleApi.md#CreateHTTPErrorRule) | **Post** /services/haproxy/configuration/http_error_rules | Add a new HTTP Error Rule
[**DeleteHTTPErrorRule**](HTTPErrorRuleApi.md#DeleteHTTPErrorRule) | **Delete** /services/haproxy/configuration/http_error_rules/{index} | Delete a HTTP Error Rule
[**GetHTTPErrorRule**](HTTPErrorRuleApi.md#GetHTTPErrorRule) | **Get** /services/haproxy/configuration/http_error_rules/{index} | Return one HTTP Error Rule
[**GetHTTPErrorRules**](HTTPErrorRuleApi.md#GetHTTPErrorRules) | **Get** /services/haproxy/configuration/http_error_rules | Return an array of all HTTP Error Rules
[**ReplaceHTTPErrorRule**](HTTPErrorRuleApi.md#ReplaceHTTPErrorRule) | **Put** /services/haproxy/configuration/http_error_rules/{index} | Replace a HTTP Error Rule


# **CreateHTTPErrorRule**
> HttpErrorRule CreateHTTPErrorRule(ctx, parentType, data, optional)
Add a new HTTP Error Rule

Adds a new HTTP Error Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiCreateHTTPErrorRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiCreateHTTPErrorRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPErrorRule**
> DeleteHTTPErrorRule(ctx, index, parentType, optional)
Delete a HTTP Error Rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPErrorRuleApiDeleteHTTPErrorRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiDeleteHTTPErrorRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
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

# **GetHTTPErrorRule**
> HttpErrorRule GetHTTPErrorRule(ctx, index, parentType, optional)
Return one HTTP Error Rule

Returns one HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPErrorRuleApiGetHTTPErrorRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetHTTPErrorRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPErrorRules**
> HttpErrorRules GetHTTPErrorRules(ctx, parentType, optional)
Return an array of all HTTP Error Rules

Returns all HTTP Error Rules that are configured in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPErrorRuleApiGetHTTPErrorRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetHTTPErrorRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPErrorRule**
> HttpErrorRule ReplaceHTTPErrorRule(ctx, index, parentType, data, optional)
Replace a HTTP Error Rule

Replaces a HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceHTTPErrorRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceHTTPErrorRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

