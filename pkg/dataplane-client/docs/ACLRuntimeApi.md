# \ACLRuntimeApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPayloadRuntimeACL**](ACLRuntimeApi.md#AddPayloadRuntimeACL) | **Put** /services/haproxy/runtime/acl_file_entries | Add a new ACL payload
[**ServicesHaproxyRuntimeAclFileEntriesGet**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclFileEntriesGet) | **Get** /services/haproxy/runtime/acl_file_entries | Return an ACL entries
[**ServicesHaproxyRuntimeAclFileEntriesIdDelete**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclFileEntriesIdDelete) | **Delete** /services/haproxy/runtime/acl_file_entries/{id} | Delete an ACL entry
[**ServicesHaproxyRuntimeAclFileEntriesIdGet**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclFileEntriesIdGet) | **Get** /services/haproxy/runtime/acl_file_entries/{id} | Return an ACL entry
[**ServicesHaproxyRuntimeAclFileEntriesPost**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclFileEntriesPost) | **Post** /services/haproxy/runtime/acl_file_entries | Add entry to an ACL file
[**ServicesHaproxyRuntimeAclsGet**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclsGet) | **Get** /services/haproxy/runtime/acls | Return an array of all ACL files
[**ServicesHaproxyRuntimeAclsIdGet**](ACLRuntimeApi.md#ServicesHaproxyRuntimeAclsIdGet) | **Get** /services/haproxy/runtime/acls/{id} | Return an ACL file


# **AddPayloadRuntimeACL**
> AclFilesEntries AddPayloadRuntimeACL(ctx, aclId, data)
Add a new ACL payload

Adds a new ACL payload.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aclId** | **string**| ACL ID | 
  **data** | [**AclFilesEntries**](AclFilesEntries.md)|  | 

### Return type

[**AclFilesEntries**](acl_files_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclFileEntriesGet**
> AclFilesEntries ServicesHaproxyRuntimeAclFileEntriesGet(ctx, aclId)
Return an ACL entries

Returns an ACL runtime setting using the runtime socket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aclId** | **string**| ACL ID | 

### Return type

[**AclFilesEntries**](acl_files_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclFileEntriesIdDelete**
> ServicesHaproxyRuntimeAclFileEntriesIdDelete(ctx, aclId, id)
Delete an ACL entry

Deletes the entry from the ACL by its value using the runtime socket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aclId** | **string**| ACL ID | 
  **id** | **string**| File entry ID | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclFileEntriesIdGet**
> AclFileEntry ServicesHaproxyRuntimeAclFileEntriesIdGet(ctx, aclId, id)
Return an ACL entry

Returns the ACL entry by its ID using the runtime socket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aclId** | **string**| ACL ID | 
  **id** | **string**| File entry ID | 

### Return type

[**AclFileEntry**](acl_file_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclFileEntriesPost**
> AclFileEntry ServicesHaproxyRuntimeAclFileEntriesPost(ctx, aclId, data)
Add entry to an ACL file

Adds an entry into the ACL file using the runtime socket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aclId** | **string**| ACL ID | 
  **data** | [**AclFileEntry**](AclFileEntry.md)|  | 

### Return type

[**AclFileEntry**](acl_file_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclsGet**
> AclFiles ServicesHaproxyRuntimeAclsGet(ctx, )
Return an array of all ACL files

Returns all ACL files using the runtime socket.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AclFiles**](acl_files.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesHaproxyRuntimeAclsIdGet**
> AclFile ServicesHaproxyRuntimeAclsIdGet(ctx, id)
Return an ACL file

Returns an ACL file by id using the runtime socket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ACL file entry ID | 

### Return type

[**AclFile**](acl_file.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

