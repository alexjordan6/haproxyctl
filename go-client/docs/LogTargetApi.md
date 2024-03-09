# \LogTargetApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogTarget**](LogTargetApi.md#CreateLogTarget) | **Post** /services/haproxy/configuration/log_targets | Add a new Log Target
[**DeleteLogTarget**](LogTargetApi.md#DeleteLogTarget) | **Delete** /services/haproxy/configuration/log_targets/{index} | Delete a Log Target
[**GetLogTarget**](LogTargetApi.md#GetLogTarget) | **Get** /services/haproxy/configuration/log_targets/{index} | Return one Log Target
[**GetLogTargets**](LogTargetApi.md#GetLogTargets) | **Get** /services/haproxy/configuration/log_targets | Return an array of all Log Targets
[**ReplaceLogTarget**](LogTargetApi.md#ReplaceLogTarget) | **Put** /services/haproxy/configuration/log_targets/{index} | Replace a Log Target


# **CreateLogTarget**
> LogTarget CreateLogTarget(ctx, parentType, data, optional)
Add a new Log Target

Adds a new Log Target of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
  **data** | [**LogTarget**](LogTarget.md)|  | 
 **optional** | ***LogTargetApiCreateLogTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogTargetApiCreateLogTargetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**LogTarget**](log_target.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogTarget**
> DeleteLogTarget(ctx, index, parentType, optional)
Delete a Log Target

Deletes a Log Target configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Log Target Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***LogTargetApiDeleteLogTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogTargetApiDeleteLogTargetOpts struct

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

# **GetLogTarget**
> LogTarget GetLogTarget(ctx, index, parentType, optional)
Return one Log Target

Returns one Log Target configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Log Target Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***LogTargetApiGetLogTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogTargetApiGetLogTargetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**LogTarget**](log_target.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogTargets**
> LogTargets GetLogTargets(ctx, parentType, optional)
Return an array of all Log Targets

Returns all Log Targets that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
 **optional** | ***LogTargetApiGetLogTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogTargetApiGetLogTargetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**LogTargets**](log_targets.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceLogTarget**
> LogTarget ReplaceLogTarget(ctx, index, parentType, data, optional)
Replace a Log Target

Replaces a Log Target configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Log Target Index | 
  **parentType** | **string**| Parent type | 
  **data** | [**LogTarget**](LogTarget.md)|  | 
 **optional** | ***LogTargetApiReplaceLogTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogTargetApiReplaceLogTargetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**LogTarget**](log_target.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

