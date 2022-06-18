package car

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/cars-owners-service/internal/config"
	"github.com/cars-owners-service/internal/data/drivers/postgres"
	"github.com/cars-owners-service/internal/data/model"
	"github.com/cars-owners-service/internal/data/queriers"
)

const Car = "car"

var _ queriers.CarProvider = &carProvider{}

type carProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.CarProvider {
	return &carProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Car),
	}
}

func (o carProvider) InsertCarBatch(cars []model.Car) error {
	if err := o.db.Create(&cars).Error; err != nil {
		return errors.Wrap(err, "failed to insert a car")
	}
	return nil
}

func (o carProvider) Paginate(page model.Page) queriers.CarProvider {
	o.db = o.db.Offset(page.Offset).Limit(page.Limit)
	return o
}

func (o carProvider) Search(search model.SearchCarsParams) queriers.CarProvider {
	if search.FirstName != nil {
		o.db = o.db.Where("owner.first_name ILIKE ?", fmt.Sprintf("%%%s%%", *search.FirstName))
	}

	if search.LastName != nil {
		o.db = o.db.Where("owner.last_name ILIKE ?", fmt.Sprintf("%%%s%%", *search.LastName))
	}
	return o
}

func (o carProvider) Filter(filter model.FilterParams) queriers.CarProvider {
	if len(filter.CarManufactur) > 0 {
		o.db = o.db.Where("car.car_manufactur IN ?", filter.CarManufactur)
	}

	if len(filter.City) > 0 {
		o.db = o.db.Where("owner.city IN ?", filter.City)
	}
	return o
}

func (o carProvider) Run() ([]model.Car, error) {
	var out []model.Car
	if err := o.db.Joins("LEFT JOIN owner on owner.id = car.owner_id").Model(&model.Car{}).Find(&out).Error; err != nil {
		return nil, errors.Wrap(err, "failed to run query")
	}
	return out, nil
}
