package owner

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cars-owners-service/internal/config"
	"github.com/cars-owners-service/internal/data/drivers/postgres"
	"github.com/cars-owners-service/internal/data/model"
	"github.com/cars-owners-service/internal/data/queriers"
)

const Owner = "owner"

var _ queriers.OwnerProvider = &ownerProvider{}

type ownerProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.OwnerProvider {
	return &ownerProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Owner),
	}
}

func (o ownerProvider) UpsertOwnerBatch(owners []model.Owner) error {
	if err := o.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&owners).Error; err != nil {
		return errors.Wrap(err, "failed to upsert owner")
	}
	return nil
}
