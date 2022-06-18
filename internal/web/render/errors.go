package render

import (
	"encoding/json"
	"github.com/cars-owners-service/resources"
	"net/http"
)

func InternalServerError[M json.Marshaler](w http.ResponseWriter, meta M) {
	Error[M](w, resources.NewError(http.StatusInternalServerError, "Something went wrong"), nil, meta)
}
