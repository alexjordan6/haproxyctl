# \TCPCheckApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPCheck**](TCPCheckApi.md#CreateTCPCheck) | **Post** /services/haproxy/configuration/tcp_checks | Add a new TCP check
[**DeleteTCPCheck**](TCPCheckApi.md#DeleteTCPCheck) | **Delete** /services/haproxy/configuration/tcp_checks/{index} | Delete a TCP check
[**GetTCPCheck**](TCPCheckApi.md#GetTCPCheck) | **Get** /services/haproxy/configuration/tcp_checks/{index} | Return one TCP check
[**GetTCPChecks**](TCPCheckApi.md#GetTCPChecks) | **Get** /services/haproxy/configuration/tcp_checks | Return an array of TCP checks
[**ReplaceTCPCheck**](TCPCheckApi.md#ReplaceTCPCheck) | **Put** /services/haproxy/configuration/tcp_checks/{index} | Replace a TCP check


# **CreateTCPCheck**
> TcpCheck CreateTCPCheck(ctx, parentType, data, optional)
Add a new TCP check

Adds a new TCP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiCreateTCPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiCreateTCPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPCheck**
> DeleteTCPCheck(ctx, index, parentType, optional)
Delete a TCP check

Deletes a TCP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP check Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPCheckApiDeleteTCPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiDeleteTCPCheckOpts struct

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

# **GetTCPCheck**
> TcpCheck GetTCPCheck(ctx, index, parentType, optional)
Return one TCP check

Returns one TCP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPCheckApiGetTCPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetTCPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPChecks**
> TcpChecks GetTCPChecks(ctx, parentType, optional)
Return an array of TCP checks

Returns all TCP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentType** | **string**| Parent type | 
 **optional** | ***TCPCheckApiGetTCPChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetTCPChecksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpChecks**](tcp_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPCheck**
> TcpCheck ReplaceTCPCheck(ctx, index, parentType, data, optional)
Replace a TCP check

Replaces a TCP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentType** | **string**| Parent type | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiReplaceTCPCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiReplaceTCPCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentName** | **optional.String**| Parent name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

