package render

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/cars-owners-service/internal/web/responses"
	"github.com/cars-owners-service/resources"
)

// Respond valid json respond rendering
func Respond[D, M json.Marshaler](w http.ResponseWriter, status *resources.Status, data D, meta M) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(status.GetCode()))

	if status.GetCode() == http.StatusNoContent {
		return
	}

	res := responses.Response[D, M]{
		Status: status,
		Data:   data,
		Meta:   meta,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}
