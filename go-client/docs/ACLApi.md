# \ACLApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAcl**](ACLApi.md#CreateAcl) | **Post** /services/haproxy/configuration/acls | Add a new ACL line
[**DeleteAcl**](ACLApi.md#DeleteAcl) | **Delete** /services/haproxy/configuration/acls/{index} | Delete a ACL line
[**GetAcl**](ACLApi.md#GetAcl) | **Get** /services/haproxy/configuration/acls/{index} | Return one ACL line
[**GetAcls**](ACLApi.md#GetAcls) | **Get** /services/haproxy/configuration/acls | Return an array of all ACL lines
[**ReplaceAcl**](ACLApi.md#ReplaceAcl) | **Put** /services/haproxy/configuration/acls/{index} | Replace a ACL line


# **CreateAcl**
> Acl CreateAcl(ctx, parentName, parentType, data, optional)
Add a new ACL line

Adds a new ACL line of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**Acl**](Acl.md)|  | 
 **optional** | ***ACLApiCreateAclOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ACLApiCreateAclOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Acl**](acl.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAcl**
> DeleteAcl(ctx, index, parentName, parentType, optional)
Delete a ACL line

Deletes a ACL line configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| ACL line Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***ACLApiDeleteAclOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ACLApiDeleteAclOpts struct

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

# **GetAcl**
> Acl GetAcl(ctx, index, parentName, parentType, optional)
Return one ACL line

Returns one ACL line configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| ACL line Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***ACLApiGetAclOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ACLApiGetAclOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Acl**](acl.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAcls**
> Acls GetAcls(ctx, parentName, parentType, optional)
Return an array of all ACL lines

Returns all ACL lines that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
 **optional** | ***ACLApiGetAclsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ACLApiGetAclsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aclName** | **optional.String**| ACL name | 
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Acls**](acls.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAcl**
> Acl ReplaceAcl(ctx, index, parentName, parentType, data, optional)
Replace a ACL line

Replaces a ACL line configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| ACL line Index | 
  **parentName** | **string**| Parent name | 
  **parentType** | **string**| Parent type | 
  **data** | [**Acl**](Acl.md)|  | 
 **optional** | ***ACLApiReplaceAclOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ACLApiReplaceAclOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Acl**](acl.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

