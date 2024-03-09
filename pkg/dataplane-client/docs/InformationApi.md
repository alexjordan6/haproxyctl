# \InformationApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHaproxyProcessInfo**](InformationApi.md#GetHaproxyProcessInfo) | **Get** /services/haproxy/runtime/info | Return HAProxy process information
[**GetInfo**](InformationApi.md#GetInfo) | **Get** /info | Return API, hardware and OS information


# **GetHaproxyProcessInfo**
> ProcessInfos GetHaproxyProcessInfo(ctx, )
Return HAProxy process information

Return HAProxy process information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProcessInfos**](process_infos.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfo**
> Info GetInfo(ctx, )
Return API, hardware and OS information

Return API, hardware and OS information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Info**](info.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

