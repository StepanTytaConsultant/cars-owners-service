package car

import (
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

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
