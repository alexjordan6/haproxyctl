# \SpoeTransactionsApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitSpoeTransaction**](SpoeTransactionsApi.md#CommitSpoeTransaction) | **Put** /services/haproxy/spoe_transactions/{id} | Commit transaction
[**DeleteSpoeTransaction**](SpoeTransactionsApi.md#DeleteSpoeTransaction) | **Delete** /services/haproxy/spoe_transactions/{id} | Delete a transaction
[**GetSpoeTransaction**](SpoeTransactionsApi.md#GetSpoeTransaction) | **Get** /services/haproxy/spoe_transactions/{id} | Return one SPOE configuration transactions
[**GetSpoeTransactions**](SpoeTransactionsApi.md#GetSpoeTransactions) | **Get** /services/haproxy/spoe_transactions | Return list of SPOE configuration transactions.
[**StartSpoeTransaction**](SpoeTransactionsApi.md#StartSpoeTransaction) | **Post** /services/haproxy/spoe_transactions | Start a new transaction


# **CommitSpoeTransaction**
> SpoeTransaction CommitSpoeTransaction(ctx, spoe, id, optional)
Commit transaction

Commit transaction, execute all operations in transaction and return msg

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **id** | **string**| Transaction id | 
 **optional** | ***SpoeTransactionsApiCommitSpoeTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeTransactionsApiCommitSpoeTransactionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**SpoeTransaction**](spoe_transaction.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeTransaction**
> DeleteSpoeTransaction(ctx, spoe, id)
Delete a transaction

Deletes a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **id** | **string**| Transaction id | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeTransaction**
> SpoeTransaction GetSpoeTransaction(ctx, spoe, id)
Return one SPOE configuration transactions

Returns one SPOE configuration transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **id** | **string**| Transaction id | 

### Return type

[**SpoeTransaction**](spoe_transaction.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeTransactions**
> SpoeTransactions GetSpoeTransactions(ctx, spoe, optional)
Return list of SPOE configuration transactions.

Returns a list of SPOE configuration transactions. Transactions can be filtered by their status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
 **optional** | ***SpoeTransactionsApiGetSpoeTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeTransactionsApiGetSpoeTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Filter by transaction status | 

### Return type

[**SpoeTransactions**](spoe_transactions.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSpoeTransaction**
> SpoeTransaction StartSpoeTransaction(ctx, spoe, version)
Start a new transaction

Starts a new transaction and returns it's id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **version** | **int32**| Configuration version on which to work on | 

### Return type

[**SpoeTransaction**](spoe_transaction.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

