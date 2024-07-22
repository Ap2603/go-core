# GuestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GuestResponseData**](GuestResponseData.md) |  | [optional] 
**Status** | Pointer to [**AuthResponseStatus**](AuthResponseStatus.md) |  | [optional] 

## Methods

### NewGuestResponse

`func NewGuestResponse() *GuestResponse`

NewGuestResponse instantiates a new GuestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestResponseWithDefaults

`func NewGuestResponseWithDefaults() *GuestResponse`

NewGuestResponseWithDefaults instantiates a new GuestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GuestResponse) GetData() GuestResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GuestResponse) GetDataOk() (*GuestResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GuestResponse) SetData(v GuestResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GuestResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GuestResponse) GetStatus() AuthResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GuestResponse) GetStatusOk() (*AuthResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GuestResponse) SetStatus(v AuthResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GuestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


