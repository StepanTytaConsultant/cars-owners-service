/*
Cars owners API

This is a REST API of the Cars owners App.

API version: 0.0.1-beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// OwnershipAddPost400Response struct for OwnershipAddPost400Response
type OwnershipAddPost400Response struct {
	Error *Error `json:"error,omitempty"`
}

// NewOwnershipAddPost400Response instantiates a new OwnershipAddPost400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnershipAddPost400Response() *OwnershipAddPost400Response {
	this := OwnershipAddPost400Response{}
	return &this
}

// NewOwnershipAddPost400ResponseWithDefaults instantiates a new OwnershipAddPost400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnershipAddPost400ResponseWithDefaults() *OwnershipAddPost400Response {
	this := OwnershipAddPost400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OwnershipAddPost400Response) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnershipAddPost400Response) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OwnershipAddPost400Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *OwnershipAddPost400Response) SetError(v Error) {
	o.Error = &v
}

func (o OwnershipAddPost400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableOwnershipAddPost400Response struct {
	value *OwnershipAddPost400Response
	isSet bool
}

func (v NullableOwnershipAddPost400Response) Get() *OwnershipAddPost400Response {
	return v.value
}

func (v *NullableOwnershipAddPost400Response) Set(val *OwnershipAddPost400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnershipAddPost400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnershipAddPost400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnershipAddPost400Response(val *OwnershipAddPost400Response) *NullableOwnershipAddPost400Response {
	return &NullableOwnershipAddPost400Response{value: val, isSet: true}
}

func (v NullableOwnershipAddPost400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnershipAddPost400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}