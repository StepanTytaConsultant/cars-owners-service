# \DefaultApi

All URIs are relative to *http://api.cars-owners.com.ua/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthcheckGet**](DefaultApi.md#HealthcheckGet) | **Get** /healthcheck | Returns success if listener runs successfully.
[**OwnershipAddPost**](DefaultApi.md#OwnershipAddPost) | **Post** /ownership/add | Add a new ownership
[**OwnershipCarsCsvGet**](DefaultApi.md#OwnershipCarsCsvGet) | **Get** /ownership/cars/csv | 
[**OwnershipCarsGet**](DefaultApi.md#OwnershipCarsGet) | **Get** /ownership/cars/ | 



## HealthcheckGet

> Status HealthcheckGet(ctx).Execute()

Returns success if listener runs successfully.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.HealthcheckGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthcheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthcheckGet`: Status
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckGetRequest struct via the builder pattern


### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnershipAddPost

> OwnershipAddPost201Response OwnershipAddPost(ctx).OwnershipAddPostRequest(ownershipAddPostRequest).Execute()

Add a new ownership

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ownershipAddPostRequest := *openapiclient.NewOwnershipAddPostRequest([]openapiclient.Ownership{*openapiclient.NewOwnership(int64(1))}) // OwnershipAddPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OwnershipAddPost(context.Background()).OwnershipAddPostRequest(ownershipAddPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OwnershipAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OwnershipAddPost`: OwnershipAddPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OwnershipAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOwnershipAddPostRequest struct via the builder pattern


| Name                        | Type                                                      | Description | Notes |
|-----------------------------|-----------------------------------------------------------|-------------|-------|
| **ownershipAddPostRequest** | [**OwnershipAddPostRequest**](OwnershipAddPostRequest.md) |             |       |

### Return type

[**OwnershipAddPost201Response**](OwnershipAddPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnershipCarsCsvGet

> *os.File OwnershipCarsCsvGet(ctx).PageOffset(pageOffset).PageLimit(pageLimit).SearchFirstName(searchFirstName).SearchLastName(searchLastName).FilterCarManufactur(filterCarManufactur).FilterCity(filterCity).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pageOffset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional)
    pageLimit := int32(56) // int32 | The numbers of items to return (optional)
    searchFirstName := "Jon" // string | First name of a person to return cars for (optional)
    searchLastName := "Doe" // string | Last name of a person to return cars for (optional)
    filterCarManufactur := "Lexus" // string | Car manufacturer to filter by (optional)
    filterCity := "Krzczonow" // string | City to filter by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OwnershipCarsCsvGet(context.Background()).PageOffset(pageOffset).PageLimit(pageLimit).SearchFirstName(searchFirstName).SearchLastName(searchLastName).FilterCarManufactur(filterCarManufactur).FilterCity(filterCity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OwnershipCarsCsvGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OwnershipCarsCsvGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OwnershipCarsCsvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOwnershipCarsCsvGetRequest struct via the builder pattern


| Name                       | Type         | Description                                                           | Notes                              |
|----------------------------|--------------|-----------------------------------------------------------------------|------------------------------------|
| **page[offset]**           | **int32**    | The number of items to skip before starting to collect the result set |                                    |
| **page[limit]**            | **int32**    | The numbers of items to return                                        |                                    |
| **search[first_name]**     | **string**   | First name of a person to return cars for                             |                                    |
| **search[last_name]**      | **string**   | Last name of a person to return cars for                              |                                    |
| **filter[car_manufactur]** | **[]string** | Car manufacturer to filter by                                         | filter[car_manufactur]=Volvo,Lexus |
| **filter[city]**           | **[]string** | City to filter by                                                     | filter[city]='San Jose,Krzczonow'  |

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnershipCarsGet

> OwnershipCarsGet200Response OwnershipCarsGet(ctx).PageOffset(pageOffset).PageLimit(pageLimit).SearchFirstName(searchFirstName).SearchLastName(searchLastName).FilterCarManufactur(filterCarManufactur).FilterCity(filterCity).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pageOffset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional)
    pageLimit := int32(56) // int32 | The numbers of items to return (optional)
    searchFirstName := "Jon" // string | First name of a person to return cars for (optional)
    searchLastName := "Doe" // string | Last name of a person to return cars for (optional)
    filterCarManufactur := "Lexus" // string | Car manufacturer to filter by (optional)
    filterCity := "Krzczonow" // string | City to filter by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OwnershipCarsGet(context.Background()).PageOffset(pageOffset).PageLimit(pageLimit).SearchFirstName(searchFirstName).SearchLastName(searchLastName).FilterCarManufactur(filterCarManufactur).FilterCity(filterCity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OwnershipCarsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OwnershipCarsGet`: OwnershipCarsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OwnershipCarsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOwnershipCarsGetRequest struct via the builder pattern


| Name                       | Type         | Description                                                           | Notes                              |
|----------------------------|--------------|-----------------------------------------------------------------------|------------------------------------|
| **page[offset]**           | **int32**    | The number of items to skip before starting to collect the result set |                                    |
| **page[limit]**            | **int32**    | The numbers of items to return                                        |                                    |
| **search[first_name]**     | **string**   | First name of a person to return cars for                             |                                    |
| **search[last_name]**      | **string**   | Last name of a person to return cars for                              |                                    |
| **filter[car_manufactur]** | **[]string** | Car manufacturer to filter by                                         | filter[car_manufactur]=Volvo,Lexus |
| **filter[city]**           | **[]string** | City to filter by                                                     | filter[city]='San Jose,Krzczonow'  |

### Return type

[**OwnershipCarsGet200Response**](OwnershipCarsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

