package store

import (
	"github.com/cars-owners-service/internal/common"
	"github.com/cars-owners-service/internal/config"
	"github.com/cars-owners-service/internal/data/drivers/postgres/car"
	"github.com/cars-owners-service/internal/data/drivers/postgres/owner"
	"github.com/cars-owners-service/internal/data/queriers"
)

const Store = "store"

type DataProvider interface {
	CarProvider() queriers.CarProvider
	OwnerProvider() queriers.OwnerProvider
}

type dataProvider struct {
	cfg config.Config
}

func (d dataProvider) CarProvider() queriers.CarProvider {
	return car.New(d.cfg)
}

func (d dataProvider) OwnerProvider() queriers.OwnerProvider {
	return owner.New(d.cfg)
}

func New(cfg config.Config) DataProvider {
	logging := cfg.Logging().WithField(common.SQLTag, cfg.Driver())
	logging.Info("DB connecting...")

	return &dataProvider{
		cfg: cfg,
	}
}
