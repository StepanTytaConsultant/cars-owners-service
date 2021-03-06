package render

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/cars-owners-service/internal/web/responses"
	"github.com/cars-owners-service/resources"
	"github.com/gocarina/gocsv"
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

func Error[M json.Marshaler](w http.ResponseWriter, error *resources.Error, serverErrors *responses.JSONServerErrors, meta M) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(error.GetCode()))

	res := responses.Response[*responses.EmptyDataResponse, M]{
		Errors: serverErrors,
		Meta:   meta,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to encode json"))
	}
}

func File(w http.ResponseWriter, in interface{}) {
	w.Header().Add("Content-Disposition", `attachment; filename="cars.csv"`)
	if err := gocsv.Marshal(in, w); err != nil {
		panic("failed to marshal csv")
	}
}
