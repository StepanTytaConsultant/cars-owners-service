package requests

import (
	"github.com/cars-owners-service/internal/web/urlvals"
	"github.com/pkg/errors"
	"net/http"
)

type GetCarsRequest struct {
	urlvals.SearchCarsParams
	urlvals.FilterParams
	urlvals.PaginationParams
}

func NewGetCarsRequest(r *http.Request) (*GetCarsRequest, error) {
	req := GetCarsRequest{}
	if err := urlvals.PopulateParams(r.URL.Query(), &req); err != nil {
		return nil, errors.Wrap(err, "failed to populate params")
	}
	return &req, nil
}
