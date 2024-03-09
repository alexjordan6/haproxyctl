# \SitesApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSite**](SitesApi.md#CreateSite) | **Post** /services/haproxy/sites | Add a site
[**DeleteSite**](SitesApi.md#DeleteSite) | **Delete** /services/haproxy/sites/{name} | Delete a site
[**GetSite**](SitesApi.md#GetSite) | **Get** /services/haproxy/sites/{name} | Return a site
[**GetSites**](SitesApi.md#GetSites) | **Get** /services/haproxy/sites | Return an array of sites
[**ReplaceSite**](SitesApi.md#ReplaceSite) | **Put** /services/haproxy/sites/{name} | Replace a site


# **CreateSite**
> Site CreateSite(ctx, data, optional)
Add a site

Adds a new site to the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Site**](Site.md)|  | 
 **optional** | ***SitesApiCreateSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiCreateSiteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Site**](site.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSite**
> DeleteSite(ctx, name, optional)
Delete a site

Deletes a site from the configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Site frontend name | 
 **optional** | ***SitesApiDeleteSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiDeleteSiteOpts struct

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

# **GetSite**
> Site GetSite(ctx, name, optional)
Return a site

Returns one site configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Site frontend name | 
 **optional** | ***SitesApiGetSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiGetSiteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Site**](site.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSites**
> Sites GetSites(ctx, optional)
Return an array of sites

Returns an array of all configured sites.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SitesApiGetSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiGetSitesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Sites**](sites.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSite**
> Site ReplaceSite(ctx, name, data, optional)
Replace a site

Replaces a site configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Site frontend name | 
  **data** | [**Site**](Site.md)|  | 
 **optional** | ***SitesApiReplaceSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesApiReplaceSiteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Site**](site.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

