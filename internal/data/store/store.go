package store

import (
	"gorm.io/gorm"

	"github.com/cars-owners-service/internal/common"
	"github.com/cars-owners-service/internal/config"
)

const Store = "store"

type DataProvider interface {
}

type dataProvider struct {
	db *gorm.DB
}

func New(cfg config.Config) DataProvider {
	logging := cfg.Logging().WithField(common.SQLTag, cfg.Driver())
	logging.Info("DB connecting...")

	return &dataProvider{
		db: cfg.DB(),
	}
}
