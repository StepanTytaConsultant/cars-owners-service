package listener

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cars-owners-service/internal/common"
	"github.com/cars-owners-service/internal/config"
)

const (
	Listener = "listener"
	port     = "port"
)

type Service interface {
	Listen() error
}

type service struct {
	log    *logrus.Entry
	cfg    config.Config
	router chi.Router
}

func New(cfg config.Config) Service {
	l := &service{
		cfg: cfg,
		log: cfg.Logging().WithField(common.APITag, Listener),
	}
	l.setupRouter()

	return l
}

func (l *service) Listen() error {
	l.log.WithField(port, l.cfg.Address()).Info("Starting service...")

	if err := http.ListenAndServe(l.cfg.Address(), l.router); err != nil {
		return errors.Wrap(err, "service failed")
	}
	return nil
}
