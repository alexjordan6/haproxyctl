# \ServerApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRuntimeServer**](ServerApi.md#AddRuntimeServer) | **Post** /services/haproxy/runtime/servers | Adds a new server to a backend
[**CreateServer**](ServerApi.md#CreateServer) | **Post** /services/haproxy/configuration/servers | Add a new server
[**DeleteRuntimeServer**](ServerApi.md#DeleteRuntimeServer) | **Delete** /services/haproxy/runtime/servers/{name} | Deletes a server from a backend
[**DeleteServer**](ServerApi.md#DeleteServer) | **Delete** /services/haproxy/configuration/servers/{name} | Delete a server
[**GetRuntimeServer**](ServerApi.md#GetRuntimeServer) | **Get** /services/haproxy/runtime/servers/{name} | Return one server runtime settings
[**GetRuntimeServers**](ServerApi.md#GetRuntimeServers) | **Get** /services/haproxy/runtime/servers | Return an array of runtime servers&#39; settings
[**GetServer**](ServerApi.md#GetServer) | **Get** /services/haproxy/configuration/servers/{name} | Return one server
[**GetServers**](ServerApi.md#GetServers) | **Get** /services/haproxy/configuration/servers | Return an array of servers
[**ReplaceRuntimeServer**](ServerApi.md#ReplaceRuntimeServer) | **Put** /services/haproxy/runtime/servers/{name} | Replace server transient settings
[**ReplaceServer**](ServerApi.md#ReplaceServer) | **Put** /services/haproxy/configuration/servers/{name} | Replace a server


# **AddRuntimeServer**
> RuntimeAddServer AddRuntimeServer(ctx, backend, data)
Adds a new server to a backend

Adds a new server to the specified backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backend** | **string**| Parent backend name | 
  **data** | [**RuntimeAddServer**](RuntimeAddServer.md)|  | 

### Return type

[**RuntimeAddServer**](runtime_add_server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServer**
> Server CreateServer(ctx, data, optional)
Add a new server

Adds a new server in the specified backend in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Server**](Server.md)|  | 
 **optional** | ***ServerApiCreateServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiCreateServerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backend** | **optional.String**| Parent backend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Server**](server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRuntimeServer**
> DeleteRuntimeServer(ctx, name, backend)
Deletes a server from a backend

Deletes a server from the specified backend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
  **backend** | **string**| Parent backend name | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServer**
> DeleteServer(ctx, name, optional)
Delete a server

Deletes a server configuration by it's name in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
 **optional** | ***ServerApiDeleteServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiDeleteServerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backend** | **optional.String**| Parent backend name | 
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

# **GetRuntimeServer**
> RuntimeServer GetRuntimeServer(ctx, name, backend)
Return one server runtime settings

Returns one server runtime settings by it's name in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
  **backend** | **string**| Parent backend name | 

### Return type

[**RuntimeServer**](runtime_server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeServers**
> RuntimeServers GetRuntimeServers(ctx, backend)
Return an array of runtime servers' settings

Returns an array of all servers' runtime settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backend** | **string**| Parent backend name | 

### Return type

[**RuntimeServers**](runtime_servers.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServer**
> Server GetServer(ctx, name, optional)
Return one server

Returns one server configuration by it's name in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
 **optional** | ***ServerApiGetServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiGetServerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backend** | **optional.String**| Parent backend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Server**](server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServers**
> Servers GetServers(ctx, optional)
Return an array of servers

Returns an array of all servers that are configured in specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServerApiGetServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiGetServersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backend** | **optional.String**| Parent backend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Servers**](servers.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRuntimeServer**
> RuntimeServer ReplaceRuntimeServer(ctx, name, backend, data)
Replace server transient settings

Replaces a server transient settings by it's name in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
  **backend** | **string**| Parent backend name | 
  **data** | [**RuntimeServer**](RuntimeServer.md)|  | 

### Return type

[**RuntimeServer**](runtime_server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceServer**
> Server ReplaceServer(ctx, name, data, optional)
Replace a server

Replaces a server configuration by it's name in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Server name | 
  **data** | [**Server**](Server.md)|  | 
 **optional** | ***ServerApiReplaceServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiReplaceServerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backend** | **optional.String**| Parent backend name | 
 **parentName** | **optional.String**| Parent name | 
 **parentType** | **optional.String**| Parent type | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Server**](server.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

