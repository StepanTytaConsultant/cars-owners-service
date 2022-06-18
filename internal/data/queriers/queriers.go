package queriers

import "github.com/cars-owners-service/internal/data/model"

type CarProvider interface {
	InsertCarBatch(cars []model.Car) error
	Paginate(page model.Page) CarProvider
	Search(search model.SearchCarsParams) CarProvider
	Filter(filter model.FilterParams) CarProvider
	Run() ([]model.Ownership, error)
}

type OwnerProvider interface {
	UpsertOwnerBatch(owners []model.Owner) error
}
