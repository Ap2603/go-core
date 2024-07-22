/*
SafePay API

API for SafePay services

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthResponseStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthResponseStatus{}

// AuthResponseStatus struct for AuthResponseStatus
type AuthResponseStatus struct {
	Errors []string `json:"errors,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewAuthResponseStatus instantiates a new AuthResponseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthResponseStatus() *AuthResponseStatus {
	this := AuthResponseStatus{}
	return &this
}

// NewAuthResponseStatusWithDefaults instantiates a new AuthResponseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthResponseStatusWithDefaults() *AuthResponseStatus {
	this := AuthResponseStatus{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AuthResponseStatus) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponseStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AuthResponseStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AuthResponseStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthResponseStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthResponseStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthResponseStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthResponseStatus) SetMessage(v string) {
	o.Message = &v
}

func (o AuthResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthResponseStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAuthResponseStatus struct {
	value *AuthResponseStatus
	isSet bool
}

func (v NullableAuthResponseStatus) Get() *AuthResponseStatus {
	return v.value
}

func (v *NullableAuthResponseStatus) Set(val *AuthResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResponseStatus(val *AuthResponseStatus) *NullableAuthResponseStatus {
	return &NullableAuthResponseStatus{value: val, isSet: true}
}

func (v NullableAuthResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

