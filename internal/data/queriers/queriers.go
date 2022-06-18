package queriers

import "github.com/cars-owners-service/internal/data/model"

type CarProvider interface {
	InsertCarBatch(cars []model.Car) error
}

type OwnerProvider interface {
	UpsertOwnerBatch(owners []model.Owner) error
}
