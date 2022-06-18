package listener

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/cars-owners-service/internal/data/store"
	"github.com/cars-owners-service/internal/listener/middlewares"
	"github.com/cars-owners-service/internal/web/ctx"
	"github.com/cars-owners-service/internal/web/handlers"
)

func (l *service) setupRouter() {
	l.router = chi.NewRouter()

	l.router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middlewares.Context(
			ctx.CtxLog(l.log),
			ctx.CtxProvider(store.New(l.cfg)),
		),
	)

	l.router.Get("/healthcheck", handlers.GetHealthcheck)
	l.router.Post("/ownership/add", handlers.PostOwnership)
	l.router.Get("/ownership/cars", handlers.GetCars)
	l.router.Get("/ownership/cars/csv", handlers.GetCarsCsv)
}
