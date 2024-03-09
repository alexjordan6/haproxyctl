# \FilterApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilter**](FilterApi.md#CreateFilter) | **Post** /services/haproxy/configuration/filters | Add a new Filter
[**DeleteFilter**](FilterApi.md#DeleteFilter) | **Delete** /services/haproxy/configuration/filters/{index} | Delete a Filter
[**GetFilter**](FilterApi.md#GetFilter) | **Get** /services/haproxy/configuration/filters/{index} | Return one Filter
[**GetFilters**](FilterApi.md#GetFilters) | **Get** /services/haproxy/configuration/filters | Return an array of all Filters
[**ReplaceFilter**](FilterApi.md#ReplaceFilter) | **Put** /services/haproxy/configuration/filters/{index} | Replace a Filter


# **CreateFilter**
> Filter CreateFilter(ctx, parentName, parentType, data, optional)
Add a new Filter

Adds a new Filter of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**Filter**](Filter.md)|  | 
 **optional** | ***FilterApiCreateFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiCreateFilterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Filter**](filter.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFilter**
> DeleteFilter(ctx, index, parentName, parentType, optional)
Delete a Filter

Deletes a Filter configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Filter Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***FilterApiDeleteFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiDeleteFilterOpts struct

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

# **GetFilter**
> Filter GetFilter(ctx, index, parentName, parentType, optional)
Return one Filter

Returns one Filter configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Filter Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***FilterApiGetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiGetFilterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Filter**](filter.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilters**
> Filters GetFilters(ctx, parentName, parentType, optional)
Return an array of all Filters

Returns all Filters that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***FilterApiGetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiGetFiltersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Filters**](filters.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceFilter**
> Filter ReplaceFilter(ctx, index, parentName, parentType, data, optional)
Replace a Filter

Replaces a Filter configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Filter Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**Filter**](Filter.md)|  | 
 **optional** | ***FilterApiReplaceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterApiReplaceFilterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Filter**](filter.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

