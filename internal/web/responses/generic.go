package responses

import (
	"encoding/json"
	"github.com/cars-owners-service/resources"
)

type Response[D, M json.Marshaler] struct {
	Data   D                 `json:"data,omitempty"`
	Meta   M                 `json:"meta,omitempty"`
	Status *resources.Status `json:"status,omitempty"`
	Errors *JSONServerErrors `json:"errors,omitempty"`
}
