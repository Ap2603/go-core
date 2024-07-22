# \APISettingsAPI

All URIs are relative to *https://dev.api.getsafepay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientApiSettingsV1Post**](APISettingsAPI.md#ClientApiSettingsV1Post) | **Post** /client/api-settings/v1/ | Create API Key



## ClientApiSettingsV1Post

> ApiSettingsResponse ClientApiSettingsV1Post(ctx).Execute()

Create API Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APISettingsAPI.ClientApiSettingsV1Post(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APISettingsAPI.ClientApiSettingsV1Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientApiSettingsV1Post`: ApiSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `APISettingsAPI.ClientApiSettingsV1Post`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClientApiSettingsV1PostRequest struct via the builder pattern


### Return type

[**ApiSettingsResponse**](ApiSettingsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

