# \HTTPCheckApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPCheck**](HTTPCheckApi.md#CreateHTTPCheck) | **Post** /services/haproxy/configuration/http_checks | Add a new HTTP check
[**DeleteHTTPCheck**](HTTPCheckApi.md#DeleteHTTPCheck) | **Delete** /services/haproxy/configuration/http_checks/{index} | Delete a HTTP check
[**GetHTTPCheck**](HTTPCheckApi.md#GetHTTPCheck) | **Get** /services/haproxy/configuration/http_checks/{index} | Return one HTTP check
[**GetHTTPChecks**](HTTPCheckApi.md#GetHTTPChecks) | **Get** /services/haproxy/configuration/http_checks | Return an array of HTTP checks
[**ReplaceHTTPCheck**](HTTPCheckApi.md#ReplaceHTTPCheck) | **Put** /services/haproxy/configuration/http_checks/{index} | Replace a HTTP check


# **CreateHTTPCheck**
> HttpCheck CreateHTTPCheck(ctx, parentType, data, optional)
Add a new HTTP check

Adds a new HTTP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiCreateHTTPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiCreateHTTPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPCheck**
> DeleteHTTPCheck(ctx, index, parentType, optional)
Delete a HTTP check

Deletes a HTTP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP check Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPCheckApiDeleteHTTPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiDeleteHTTPCheckOpts struct

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

# **GetHTTPCheck**
> HttpCheck GetHTTPCheck(ctx, index, parentType, optional)
Return one HTTP check

Returns one HTTP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPCheckApiGetHTTPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetHTTPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPChecks**
> HttpChecks GetHTTPChecks(ctx, parentType, optional)
Return an array of HTTP checks

Returns all HTTP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
 **optional** | ***HTTPCheckApiGetHTTPChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetHTTPChecksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpChecks**](http_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPCheck**
> HttpCheck ReplaceHTTPCheck(ctx, index, parentType, data, optional)
Replace a HTTP check

Replaces a HTTP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentType** | **string**| Parent type | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiReplaceHTTPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiReplaceHTTPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

