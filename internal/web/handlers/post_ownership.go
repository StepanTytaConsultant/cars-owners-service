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

func PostOwnership(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)

	req, err := requests.NewPostOwnershipRequest(r)
	if err != nil {
		log.WithError(err).Error("failed to create new add ownership request")
		render.InternalServerError[*responses.EmptyMetaResponse](w, nil)
		return
	}

	owners, cars := convert.SplitOwnerships(req.OwnershipAddRequest.GetData())

	provider := ctx.Provider(r)

	if err := provider.OwnerProvider().UpsertOwnerBatch(owners); err != nil {
		log.WithError(err).Error("failed to upsert owners")
		render.InternalServerError[*responses.EmptyMetaResponse](w, nil)
		return
	}

	if err := provider.CarProvider().InsertCarBatch(cars); err != nil {
		log.WithError(err).Error("failed to insert cars")
		render.InternalServerError[*responses.EmptyMetaResponse](w, nil)
		return
	}

	render.Respond[*resources.OwnershipAddPost201ResponseData, *responses.EmptyMetaResponse](
		w,
		resources.NewStatus(http.StatusCreated, "Ownerships added"),
		&resources.OwnershipAddPost201ResponseData{
			Owners: convert.ToResponseOwners(owners),
			Cars:   convert.ToResponseCars(cars),
		},
		nil,
	)
}
