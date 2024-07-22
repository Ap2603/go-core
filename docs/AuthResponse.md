# AuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AuthResponseData**](AuthResponseData.md) |  | [optional] 
**Status** | Pointer to [**AuthResponseStatus**](AuthResponseStatus.md) |  | [optional] 

## Methods

### NewAuthResponse

`func NewAuthResponse() *AuthResponse`

NewAuthResponse instantiates a new AuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponseWithDefaults

`func NewAuthResponseWithDefaults() *AuthResponse`

NewAuthResponseWithDefaults instantiates a new AuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthResponse) GetData() AuthResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthResponse) GetDataOk() (*AuthResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthResponse) SetData(v AuthResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *AuthResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *AuthResponse) GetStatus() AuthResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthResponse) GetStatusOk() (*AuthResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthResponse) SetStatus(v AuthResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


