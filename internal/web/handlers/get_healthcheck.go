package handlers

import (
	"net/http"

	"github.com/cars-owners-service/internal/web/render"
	"github.com/cars-owners-service/internal/web/responses"
	"github.com/cars-owners-service/resources"
)

func GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	render.Respond[*responses.EmptyDataResponse, *responses.EmptyMetaResponse](
		w,
		resources.NewStatus(http.StatusOK, "Server runs successfully"),
		nil,
		nil,
	)
}
