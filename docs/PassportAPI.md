# \PassportAPI

All URIs are relative to *https://dev.api.getsafepay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientPassportV1TokenPost**](PassportAPI.md#ClientPassportV1TokenPost) | **Post** /client/passport/v1/token | Generate Time-Based Token



## ClientPassportV1TokenPost

> TimeBasedTokenResponse ClientPassportV1TokenPost(ctx).Execute()

Generate Time-Based Token

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
	resp, r, err := apiClient.PassportAPI.ClientPassportV1TokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PassportAPI.ClientPassportV1TokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientPassportV1TokenPost`: TimeBasedTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PassportAPI.ClientPassportV1TokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClientPassportV1TokenPostRequest struct via the builder pattern


### Return type

[**TimeBasedTokenResponse**](TimeBasedTokenResponse.md)

### Authorization

[MerchantSecretAuth](../README.md#MerchantSecretAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

