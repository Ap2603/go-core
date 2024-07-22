# \DefaultAPI

All URIs are relative to *https://dev.api.getsafepay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthV1CompanyAuthenticatePost**](DefaultAPI.md#AuthV1CompanyAuthenticatePost) | **Post** /auth/v1/company/authenticate | Create JWT Token
[**UserV1GuestPost**](DefaultAPI.md#UserV1GuestPost) | **Post** /user/v1/guest/ | Create Guest JWT



## AuthV1CompanyAuthenticatePost

> AuthResponse AuthV1CompanyAuthenticatePost(ctx).AuthRequest(authRequest).Execute()

Create JWT Token



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
	authRequest := *openapiclient.NewAuthRequest("client_bb1d600f-f174-49dc-a34f-a79c77e237c8", "zparekh@getsafepay.com", "azazazazaz") // AuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AuthV1CompanyAuthenticatePost(context.Background()).AuthRequest(authRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AuthV1CompanyAuthenticatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthV1CompanyAuthenticatePost`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AuthV1CompanyAuthenticatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthV1CompanyAuthenticatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequest** | [**AuthRequest**](AuthRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserV1GuestPost

> GuestResponse UserV1GuestPost(ctx).GuestRequest(guestRequest).Execute()

Create Guest JWT



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
	guestRequest := *openapiclient.NewGuestRequest("hzaidi@getsafepay.com", "+923001234567", "PK") // GuestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserV1GuestPost(context.Background()).GuestRequest(guestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserV1GuestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserV1GuestPost`: GuestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserV1GuestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserV1GuestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestRequest** | [**GuestRequest**](GuestRequest.md) |  | 

### Return type

[**GuestResponse**](GuestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

