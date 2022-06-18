package handlers

import (
	"github.com/cars-owners-service/internal/web/ctx"
	"github.com/cars-owners-service/internal/web/render"
	"github.com/cars-owners-service/internal/web/requests"
	"github.com/cars-owners-service/internal/web/responses"
	"github.com/cars-owners-service/internal/web/utils/convert"
	"github.com/cars-owners-service/resources"
	"net/http"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	req, err := requests.NewGetCarsRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to get search cars request")
		render.InternalServerError[*responses.EmptyMetaResponse](w, nil)
		return
	}

	provider := ctx.Provider(r)

	cars, err := provider.CarProvider().
		Paginate(convert.FromResponsePage(req.PaginationParams)).
		Search(convert.FromResponseSearchCars(req.SearchCarsParams)).
		Filter(convert.FromResponseFilterCars(req.FilterParams)).
		Run()
	if err != nil {
		log.WithError(err).Error("failed to create new add ownership request")
		render.InternalServerError[*responses.EmptyMetaResponse](w, nil)
		return
	}

	render.Respond[*resources.OwnershipCarsGet200ResponseData, *responses.EmptyMetaResponse](
		w,
		resources.NewStatus(http.StatusOK, "Request result ok"),
		&resources.OwnershipCarsGet200ResponseData{
			Cars: convert.ToResponseCars(cars),
		},
		nil,
	)
}
