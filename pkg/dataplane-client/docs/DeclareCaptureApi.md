# \DeclareCaptureApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeclareCapture**](DeclareCaptureApi.md#CreateDeclareCapture) | **Post** /services/haproxy/configuration/captures | Add a new declare capture
[**DeleteDeclareCapture**](DeclareCaptureApi.md#DeleteDeclareCapture) | **Delete** /services/haproxy/configuration/captures/{index} | Delete a declare capture
[**GetDeclareCapture**](DeclareCaptureApi.md#GetDeclareCapture) | **Get** /services/haproxy/configuration/captures/{index} | Return one declare capture
[**GetDeclareCaptures**](DeclareCaptureApi.md#GetDeclareCaptures) | **Get** /services/haproxy/configuration/captures | Return an array of declare captures
[**ReplaceDeclareCapture**](DeclareCaptureApi.md#ReplaceDeclareCapture) | **Put** /services/haproxy/configuration/captures/{index} | Replace a declare capture


# **CreateDeclareCapture**
> Capture CreateDeclareCapture(ctx, frontend, data, optional)
Add a new declare capture

Adds a new declare capture in the specified frontend in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frontend** | **string**| Parent frontend name | 
  **data** | [**Capture**](Capture.md)|  | 
 **optional** | ***DeclareCaptureApiCreateDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiCreateDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeclareCapture**
> DeleteDeclareCapture(ctx, index, frontend, optional)
Delete a declare capture

Deletes a declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **frontend** | **string**| Parent frontend name | 
 **optional** | ***DeclareCaptureApiDeleteDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiDeleteDeclareCaptureOpts struct

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

# **GetDeclareCapture**
> Capture GetDeclareCapture(ctx, index, frontend, optional)
Return one declare capture

Returns one declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **frontend** | **string**| Parent frontend name | 
 **optional** | ***DeclareCaptureApiGetDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiGetDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeclareCaptures**
> Captures GetDeclareCaptures(ctx, frontend, optional)
Return an array of declare captures

Returns an array of all declare capture records that are configured in specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **frontend** | **string**| Parent frontend name | 
 **optional** | ***DeclareCaptureApiGetDeclareCapturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiGetDeclareCapturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Captures**](captures.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDeclareCapture**
> Capture ReplaceDeclareCapture(ctx, index, frontend, data, optional)
Replace a declare capture

Replaces a declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **frontend** | **string**| Parent frontend name | 
  **data** | [**Capture**](Capture.md)|  | 
 **optional** | ***DeclareCaptureApiReplaceDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiReplaceDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

