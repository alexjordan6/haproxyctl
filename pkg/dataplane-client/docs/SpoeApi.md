# \SpoeApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpoe**](SpoeApi.md#CreateSpoe) | **Post** /services/haproxy/spoe/spoe_files | Creates SPOE file with its entries
[**CreateSpoeAgent**](SpoeApi.md#CreateSpoeAgent) | **Post** /services/haproxy/spoe/spoe_agents | Add a new spoe agent to scope
[**CreateSpoeGroup**](SpoeApi.md#CreateSpoeGroup) | **Post** /services/haproxy/spoe/spoe_groups | Add a new SPOE groups
[**CreateSpoeMessage**](SpoeApi.md#CreateSpoeMessage) | **Post** /services/haproxy/spoe/spoe_messages | Add a new spoe message to scope
[**CreateSpoeScope**](SpoeApi.md#CreateSpoeScope) | **Post** /services/haproxy/spoe/spoe_scopes | Add a new spoe scope
[**DeleteSpoeAgent**](SpoeApi.md#DeleteSpoeAgent) | **Delete** /services/haproxy/spoe/spoe_agents/{name} | Delete a SPOE agent
[**DeleteSpoeFile**](SpoeApi.md#DeleteSpoeFile) | **Delete** /services/haproxy/spoe/spoe_files/{name} | Delete SPOE file
[**DeleteSpoeGroup**](SpoeApi.md#DeleteSpoeGroup) | **Delete** /services/haproxy/spoe/spoe_groups/{name} | Delete a SPOE groups
[**DeleteSpoeMessage**](SpoeApi.md#DeleteSpoeMessage) | **Delete** /services/haproxy/spoe/spoe_messages/{name} | Delete a spoe message
[**DeleteSpoeScope**](SpoeApi.md#DeleteSpoeScope) | **Delete** /services/haproxy/spoe/spoe_scopes/{name} | Delete a SPOE scope
[**GetAllSpoeFiles**](SpoeApi.md#GetAllSpoeFiles) | **Get** /services/haproxy/spoe/spoe_files | Return all available SPOE files
[**GetOneSpoeFile**](SpoeApi.md#GetOneSpoeFile) | **Get** /services/haproxy/spoe/spoe_files/{name} | Return one SPOE file
[**GetSpoeAgent**](SpoeApi.md#GetSpoeAgent) | **Get** /services/haproxy/spoe/spoe_agents/{name} | Return a spoe agent
[**GetSpoeAgents**](SpoeApi.md#GetSpoeAgents) | **Get** /services/haproxy/spoe/spoe_agents | Return an array of spoe agents in one scope
[**GetSpoeConfigurationVersion**](SpoeApi.md#GetSpoeConfigurationVersion) | **Get** /services/haproxy/spoe/version | Return a SPOE configuration version
[**GetSpoeGroup**](SpoeApi.md#GetSpoeGroup) | **Get** /services/haproxy/spoe/spoe_groups/{name} | Return a SPOE groups
[**GetSpoeGroups**](SpoeApi.md#GetSpoeGroups) | **Get** /services/haproxy/spoe/spoe_groups | Return an array of SPOE groups
[**GetSpoeMessage**](SpoeApi.md#GetSpoeMessage) | **Get** /services/haproxy/spoe/spoe_messages/{name} | Return a spoe message
[**GetSpoeMessages**](SpoeApi.md#GetSpoeMessages) | **Get** /services/haproxy/spoe/spoe_messages | Return an array of spoe messages in one scope
[**GetSpoeScope**](SpoeApi.md#GetSpoeScope) | **Get** /services/haproxy/spoe/spoe_scopes/{name} | Return one SPOE scope
[**GetSpoeScopes**](SpoeApi.md#GetSpoeScopes) | **Get** /services/haproxy/spoe/spoe_scopes | Return an array of spoe scopes
[**ReplaceSpoeAgent**](SpoeApi.md#ReplaceSpoeAgent) | **Put** /services/haproxy/spoe/spoe_agents/{name} | Replace a SPOE agent
[**ReplaceSpoeGroup**](SpoeApi.md#ReplaceSpoeGroup) | **Put** /services/haproxy/spoe/spoe_groups/{name} | Replace a SPOE groups
[**ReplaceSpoeMessage**](SpoeApi.md#ReplaceSpoeMessage) | **Put** /services/haproxy/spoe/spoe_messages/{name} | Replace a spoe message


# **CreateSpoe**
> string CreateSpoe(ctx, optional)
Creates SPOE file with its entries

Creates SPOE file with its entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpoeApiCreateSpoeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileUpload** | **optional.Interface of *os.File**| The spoe file to upload | 

### Return type

**string**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeAgent**
> SpoeAgent CreateSpoeAgent(ctx, spoe, scope, data, optional)
Add a new spoe agent to scope

Adds a new spoe agent to the spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **data** | [**SpoeAgent**](SpoeAgent.md)|  | 
 **optional** | ***SpoeApiCreateSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeGroup**
> SpoeGroup CreateSpoeGroup(ctx, spoe, scope, data, optional)
Add a new SPOE groups

Adds a new SPOE groups to the SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **data** | [**SpoeGroup**](SpoeGroup.md)|  | 
 **optional** | ***SpoeApiCreateSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeMessage**
> SpoeMessage CreateSpoeMessage(ctx, spoe, scope, data, optional)
Add a new spoe message to scope

Adds a new spoe message to the spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **data** | [**SpoeMessage**](SpoeMessage.md)|  | 
 **optional** | ***SpoeApiCreateSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeScope**
> SpoeScope CreateSpoeScope(ctx, spoe, data, optional)
Add a new spoe scope

Adds a new spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **data** | [**SpoeScope**](SpoeScope.md)|  | 
 **optional** | ***SpoeApiCreateSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeScope**](spoe_scope.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeAgent**
> DeleteSpoeAgent(ctx, spoe, scope, name, optional)
Delete a SPOE agent

Deletes a SPOE agent from the configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe agent name | 
 **optional** | ***SpoeApiDeleteSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeFile**
> DeleteSpoeFile(ctx, name)
Delete SPOE file

Deletes SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SPOE file name | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeGroup**
> DeleteSpoeGroup(ctx, spoe, scope, name, optional)
Delete a SPOE groups

Deletes a SPOE groups from the one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe group name | 
 **optional** | ***SpoeApiDeleteSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeMessage**
> DeleteSpoeMessage(ctx, spoe, scope, name, optional)
Delete a spoe message

Deletes a spoe message from the SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe message name | 
 **optional** | ***SpoeApiDeleteSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeScope**
> DeleteSpoeScope(ctx, spoe, name, optional)
Delete a SPOE scope

Deletes a SPOE scope from the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **name** | **string**| Spoe scope name | 
 **optional** | ***SpoeApiDeleteSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeFiles**
> SpoeFiles GetAllSpoeFiles(ctx, )
Return all available SPOE files

Returns all available SPOE files.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SpoeFiles**](spoe_files.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneSpoeFile**
> InlineResponse2001 GetOneSpoeFile(ctx, name)
Return one SPOE file

Returns one SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SPOE file name | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeAgent**
> SpoeAgent GetSpoeAgent(ctx, spoe, scope, name, optional)
Return a spoe agent

Returns one spoe agent configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe agent name | 
 **optional** | ***SpoeApiGetSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeAgents**
> SpoeAgents GetSpoeAgents(ctx, spoe, scope, optional)
Return an array of spoe agents in one scope

Returns an array of all configured spoe agents in one scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
 **optional** | ***SpoeApiGetSpoeAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeAgentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeAgents**](spoe_agents.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeConfigurationVersion**
> int32 GetSpoeConfigurationVersion(ctx, spoe, optional)
Return a SPOE configuration version

Returns SPOE configuration version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
 **optional** | ***SpoeApiGetSpoeConfigurationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeConfigurationVersionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

**int32**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeGroup**
> SpoeGroup GetSpoeGroup(ctx, spoe, scope, name, optional)
Return a SPOE groups

Returns one SPOE groups configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe group name | 
 **optional** | ***SpoeApiGetSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeGroups**
> SpoeGroups GetSpoeGroups(ctx, spoe, scope, optional)
Return an array of SPOE groups

Returns an array of all configured SPOE groups in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
 **optional** | ***SpoeApiGetSpoeGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeGroups**](spoe_groups.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeMessage**
> SpoeMessage GetSpoeMessage(ctx, spoe, scope, name, optional)
Return a spoe message

Returns one spoe message configuration in SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe message name | 
 **optional** | ***SpoeApiGetSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeMessages**
> SpoeMessages GetSpoeMessages(ctx, spoe, scope, optional)
Return an array of spoe messages in one scope

Returns an array of all configured spoe messages in one scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
 **optional** | ***SpoeApiGetSpoeMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeMessages**](spoe_messages.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeScope**
> SpoeScope GetSpoeScope(ctx, spoe, name, optional)
Return one SPOE scope

Returns one SPOE scope in one SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **name** | **string**| Spoe scope | 
 **optional** | ***SpoeApiGetSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeScope**](spoe_scope.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeScopes**
> SpoeScopes GetSpoeScopes(ctx, spoe, optional)
Return an array of spoe scopes

Returns an array of all configured spoe scopes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
 **optional** | ***SpoeApiGetSpoeScopesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeScopesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeScopes**](spoe_scopes.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeAgent**
> SpoeAgent ReplaceSpoeAgent(ctx, spoe, scope, name, data, optional)
Replace a SPOE agent

Replaces a SPOE agent configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe agent name | 
  **data** | [**SpoeAgent**](SpoeAgent.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeGroup**
> SpoeGroup ReplaceSpoeGroup(ctx, spoe, scope, name, data, optional)
Replace a SPOE groups

Replaces a SPOE groups configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe group name | 
  **data** | [**SpoeGroup**](SpoeGroup.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeMessage**
> SpoeMessage ReplaceSpoeMessage(ctx, spoe, scope, name, data, optional)
Replace a spoe message

Replaces a spoe message configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spoe** | **string**| Spoe file name | 
  **scope** | **string**| Spoe scope | 
  **name** | **string**| Spoe message name | 
  **data** | [**SpoeMessage**](SpoeMessage.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

