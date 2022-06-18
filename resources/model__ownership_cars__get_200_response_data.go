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

// OwnershipCarsGet200ResponseData struct for OwnershipCarsGet200ResponseData
type OwnershipCarsGet200ResponseData struct {
	Cars []Car `json:"cars,omitempty"`
}

// NewOwnershipCarsGet200ResponseData instantiates a new OwnershipCarsGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnershipCarsGet200ResponseData() *OwnershipCarsGet200ResponseData {
	this := OwnershipCarsGet200ResponseData{}
	return &this
}

// NewOwnershipCarsGet200ResponseDataWithDefaults instantiates a new OwnershipCarsGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnershipCarsGet200ResponseDataWithDefaults() *OwnershipCarsGet200ResponseData {
	this := OwnershipCarsGet200ResponseData{}
	return &this
}

// GetCars returns the Cars field value if set, zero value otherwise.
func (o *OwnershipCarsGet200ResponseData) GetCars() []Car {
	if o == nil || o.Cars == nil {
		var ret []Car
		return ret
	}
	return o.Cars
}

// GetCarsOk returns a tuple with the Cars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnershipCarsGet200ResponseData) GetCarsOk() ([]Car, bool) {
	if o == nil || o.Cars == nil {
		return nil, false
	}
	return o.Cars, true
}

// HasCars returns a boolean if a field has been set.
func (o *OwnershipCarsGet200ResponseData) HasCars() bool {
	if o != nil && o.Cars != nil {
		return true
	}

	return false
}

// SetCars gets a reference to the given []Car and assigns it to the Cars field.
func (o *OwnershipCarsGet200ResponseData) SetCars(v []Car) {
	o.Cars = v
}

func (o OwnershipCarsGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cars != nil {
		toSerialize["cars"] = o.Cars
	}
	return json.Marshal(toSerialize)
}

type NullableOwnershipCarsGet200ResponseData struct {
	value *OwnershipCarsGet200ResponseData
	isSet bool
}

func (v NullableOwnershipCarsGet200ResponseData) Get() *OwnershipCarsGet200ResponseData {
	return v.value
}

func (v *NullableOwnershipCarsGet200ResponseData) Set(val *OwnershipCarsGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnershipCarsGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnershipCarsGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnershipCarsGet200ResponseData(val *OwnershipCarsGet200ResponseData) *NullableOwnershipCarsGet200ResponseData {
	return &NullableOwnershipCarsGet200ResponseData{value: val, isSet: true}
}

func (v NullableOwnershipCarsGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnershipCarsGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
