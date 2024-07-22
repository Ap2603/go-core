# AuthResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthResponseStatus

`func NewAuthResponseStatus() *AuthResponseStatus`

NewAuthResponseStatus instantiates a new AuthResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponseStatusWithDefaults

`func NewAuthResponseStatusWithDefaults() *AuthResponseStatus`

NewAuthResponseStatusWithDefaults instantiates a new AuthResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *AuthResponseStatus) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AuthResponseStatus) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AuthResponseStatus) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AuthResponseStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *AuthResponseStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthResponseStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthResponseStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthResponseStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

