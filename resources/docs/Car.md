# Car

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CarManufactur** | Pointer to **string** |  | [optional] 
**CarModel** | Pointer to **string** |  | [optional] 
**CarModelYear** | Pointer to **int32** |  | [optional] 
**OwnerId** | **int64** |  | 

## Methods

### NewCar

`func NewCar(id int64, ownerId int64, ) *Car`

NewCar instantiates a new Car object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarWithDefaults

`func NewCarWithDefaults() *Car`

NewCarWithDefaults instantiates a new Car object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Car) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Car) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Car) SetId(v int64)`

SetId sets Id field to given value.


### GetCarManufactur

`func (o *Car) GetCarManufactur() string`

GetCarManufactur returns the CarManufactur field if non-nil, zero value otherwise.

### GetCarManufacturOk

`func (o *Car) GetCarManufacturOk() (*string, bool)`

GetCarManufacturOk returns a tuple with the CarManufactur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarManufactur

`func (o *Car) SetCarManufactur(v string)`

SetCarManufactur sets CarManufactur field to given value.

### HasCarManufactur

`func (o *Car) HasCarManufactur() bool`

HasCarManufactur returns a boolean if a field has been set.

### GetCarModel

`func (o *Car) GetCarModel() string`

GetCarModel returns the CarModel field if non-nil, zero value otherwise.

### GetCarModelOk

`func (o *Car) GetCarModelOk() (*string, bool)`

GetCarModelOk returns a tuple with the CarModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarModel

`func (o *Car) SetCarModel(v string)`

SetCarModel sets CarModel field to given value.

### HasCarModel

`func (o *Car) HasCarModel() bool`

HasCarModel returns a boolean if a field has been set.

### GetCarModelYear

`func (o *Car) GetCarModelYear() int32`

GetCarModelYear returns the CarModelYear field if non-nil, zero value otherwise.

### GetCarModelYearOk

`func (o *Car) GetCarModelYearOk() (*int32, bool)`

GetCarModelYearOk returns a tuple with the CarModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarModelYear

`func (o *Car) SetCarModelYear(v int32)`

SetCarModelYear sets CarModelYear field to given value.

### HasCarModelYear

`func (o *Car) HasCarModelYear() bool`

HasCarModelYear returns a boolean if a field has been set.

### GetOwnerId

`func (o *Car) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Car) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Car) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


