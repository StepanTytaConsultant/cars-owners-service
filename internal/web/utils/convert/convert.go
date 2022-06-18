package convert

import (
	"github.com/cars-owners-service/internal/data/model"
	"github.com/cars-owners-service/internal/web/urlvals"
	"github.com/cars-owners-service/internal/web/utils"
	"github.com/cars-owners-service/resources"
	"strings"
)

func SplitOwnerships(ownership []resources.Ownership) ([]model.Owner, []model.Car) {
	owners := make([]model.Owner, len(ownership))
	cars := make([]model.Car, len(ownership))
	for i, own := range ownership {
		owners[i], cars[i] = SplitOwnership(own)
	}
	return owners, cars
}

func SplitOwnership(ownership resources.Ownership) (model.Owner, model.Car) {
	return extractOwnerFromOwnership(ownership), extractCarFromOwnership(ownership)
}

func extractOwnerFromOwnership(ownership resources.Ownership) model.Owner {
	owner := model.Owner{
		ID: ownership.GetId(),
	}

	if v, ok := ownership.GetFirstNameOk(); ok {
		owner.FirstName = v
	}

	if v, ok := ownership.GetLastNameOk(); ok {
		owner.LastName = v
	}

	if v, ok := ownership.GetEmailOk(); ok {
		owner.Email = v
	}

	if v, ok := ownership.GetGenderOk(); ok {
		owner.Gender = v
		if owner.Gender != nil {
			owner.Gender = StringPtr(strings.ToLower(*owner.Gender))
		}
	}

	if v, ok := ownership.GetAddressOk(); ok {
		owner.Address = v
	}

	if owner.Address != nil {
		tokens := strings.Split(*owner.Address, ",")
		if len(tokens) == 2 {
			owner.City = StringPtr(strings.TrimSpace(tokens[1]))
		}
	}
	return owner
}

func extractCarFromOwnership(ownership resources.Ownership) model.Car {
	car := model.Car{
		OwnerID: ownership.Id,
	}

	if v, ok := ownership.GetCarManufacturOk(); ok {
		car.CarManufactur = v
	}

	if v, ok := ownership.GetCarModelOk(); ok {
		car.CarModel = v
	}

	if v, ok := ownership.GetCarModelYearOk(); ok {
		car.CarModelYear = v
	}
	return car
}

func ToResponseOwners(owners []model.Owner) []resources.Owner {
	responseOwners := make([]resources.Owner, len(owners))
	for i, owner := range owners {
		responseOwners[i] = resources.Owner{
			Id:        owner.ID,
			FirstName: owner.FirstName,
			LastName:  owner.LastName,
			Email:     owner.Email,
			Gender:    owner.Gender,
			Address:   owner.Address,
			City:      owner.City,
		}
	}
	return responseOwners
}

func ToResponseCars(cars []model.Car) []resources.Car {
	responseCars := make([]resources.Car, len(cars))
	for i, car := range cars {
		responseCars[i] = resources.Car{
			Id:            car.ID,
			CarManufactur: car.CarManufactur,
			CarModel:      car.CarModel,
			CarModelYear:  car.CarModelYear,
			OwnerId:       car.OwnerID,
		}
	}
	return responseCars
}

func StringPtr(s string) *string {
	return &s
}

func FromInt32Ptr(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func FromResponsePage(params urlvals.PaginationParams) model.Page {
	limit := utils.DefaultLimit
	if params.Limit != nil {
		limit = int(FromInt32Ptr(params.Limit))
	}
	return model.Page{
		Limit:  limit,
		Offset: int(FromInt32Ptr(params.Offset)),
	}
}

func FromResponseSearchCars(params urlvals.SearchCarsParams) model.SearchCarsParams {
	return model.SearchCarsParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}
}

func FromResponseFilterCars(params urlvals.FilterParams) model.FilterParams {
	return model.FilterParams{
		CarManufactur: params.CarManufactur,
		City:          params.City,
	}
}
