# \BindApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBind**](BindApi.md#CreateBind) | **Post** /services/haproxy/configuration/binds | Add a new bind
[**DeleteBind**](BindApi.md#DeleteBind) | **Delete** /services/haproxy/configuration/binds/{name} | Delete a bind
[**GetBind**](BindApi.md#GetBind) | **Get** /services/haproxy/configuration/binds/{name} | Return one bind
[**GetBinds**](BindApi.md#GetBinds) | **Get** /services/haproxy/configuration/binds | Return an array of binds
[**ReplaceBind**](BindApi.md#ReplaceBind) | **Put** /services/haproxy/configuration/binds/{name} | Replace a bind


# **CreateBind**
> Bind CreateBind(ctx, data, optional)
Add a new bind

Adds a new bind in the specified frontend in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Bind**](Bind.md)|  | 
 **optional** | ***BindApiCreateBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindApiCreateBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frontend** | **optional.String**| Parent frontend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Bind**](bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBind**
> DeleteBind(ctx, name, optional)
Delete a bind

Deletes a bind configuration by it's name in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
 **optional** | ***BindApiDeleteBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindApiDeleteBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frontend** | **optional.String**| Parent frontend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
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

# **GetBind**
> Bind GetBind(ctx, name, optional)
Return one bind

Returns one bind configuration by it's name in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
 **optional** | ***BindApiGetBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindApiGetBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frontend** | **optional.String**| Parent frontend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Bind**](bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBinds**
> Binds GetBinds(ctx, optional)
Return an array of binds

Returns an array of all binds that are configured in specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BindApiGetBindsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindApiGetBindsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frontend** | **optional.String**| Parent frontend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Binds**](binds.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBind**
> Bind ReplaceBind(ctx, name, data, optional)
Replace a bind

Replaces a bind configuration by it's name in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
  **data** | [**Bind**](Bind.md)|  | 
 **optional** | ***BindApiReplaceBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindApiReplaceBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **frontend** | **optional.String**| Parent frontend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Bind**](bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

