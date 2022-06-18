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

// OwnershipAddPost201Response struct for OwnershipAddPost201Response
type OwnershipAddPost201Response struct {
	Status *Status                          `json:"status,omitempty"`
	Data   *OwnershipAddPost201ResponseData `json:"data,omitempty"`
}

// NewOwnershipAddPost201Response instantiates a new OwnershipAddPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnershipAddPost201Response() *OwnershipAddPost201Response {
	this := OwnershipAddPost201Response{}
	return &this
}

// NewOwnershipAddPost201ResponseWithDefaults instantiates a new OwnershipAddPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnershipAddPost201ResponseWithDefaults() *OwnershipAddPost201Response {
	this := OwnershipAddPost201Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OwnershipAddPost201Response) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnershipAddPost201Response) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OwnershipAddPost201Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *OwnershipAddPost201Response) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OwnershipAddPost201Response) GetData() OwnershipAddPost201ResponseData {
	if o == nil || o.Data == nil {
		var ret OwnershipAddPost201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnershipAddPost201Response) GetDataOk() (*OwnershipAddPost201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OwnershipAddPost201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OwnershipAddPost201ResponseData and assigns it to the Data field.
func (o *OwnershipAddPost201Response) SetData(v OwnershipAddPost201ResponseData) {
	o.Data = &v
}

func (o OwnershipAddPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOwnershipAddPost201Response struct {
	value *OwnershipAddPost201Response
	isSet bool
}

func (v NullableOwnershipAddPost201Response) Get() *OwnershipAddPost201Response {
	return v.value
}

func (v *NullableOwnershipAddPost201Response) Set(val *OwnershipAddPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnershipAddPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnershipAddPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnershipAddPost201Response(val *OwnershipAddPost201Response) *NullableOwnershipAddPost201Response {
	return &NullableOwnershipAddPost201Response{value: val, isSet: true}
}

func (v NullableOwnershipAddPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnershipAddPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}