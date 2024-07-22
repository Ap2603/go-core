# Go API client for openapi

API for SafePay services

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://dev.api.getsafepay.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APISettingsAPI* | [**ClientApiSettingsV1Post**](docs/APISettingsAPI.md#clientapisettingsv1post) | **Post** /client/api-settings/v1/ | Create API Key
*DefaultAPI* | [**AuthV1CompanyAuthenticatePost**](docs/DefaultAPI.md#authv1companyauthenticatepost) | **Post** /auth/v1/company/authenticate | Create JWT Token
*DefaultAPI* | [**UserV1GuestPost**](docs/DefaultAPI.md#userv1guestpost) | **Post** /user/v1/guest/ | Create Guest JWT
*PassportAPI* | [**ClientPassportV1TokenPost**](docs/PassportAPI.md#clientpassportv1tokenpost) | **Post** /client/passport/v1/token | Generate Time-Based Token


## Documentation For Models

 - [ApiSettingsResponse](docs/ApiSettingsResponse.md)
 - [ApiSettingsResponseData](docs/ApiSettingsResponseData.md)
 - [ApiSettingsResponseStatus](docs/ApiSettingsResponseStatus.md)
 - [AuthRequest](docs/AuthRequest.md)
 - [AuthResponse](docs/AuthResponse.md)
 - [AuthResponseData](docs/AuthResponseData.md)
 - [AuthResponseStatus](docs/AuthResponseStatus.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GuestRequest](docs/GuestRequest.md)
 - [GuestResponse](docs/GuestResponse.md)
 - [GuestResponseData](docs/GuestResponseData.md)
 - [TimeBasedTokenResponse](docs/TimeBasedTokenResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### MerchantSecretAuth

- **Type**: API key
- **API key parameter name**: X-SFPY-MERCHANT-SECRET
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-SFPY-MERCHANT-SECRET and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"X-SFPY-MERCHANT-SECRET": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


