package utils

import (
	"encoding/json"
	"net/http"

	"github.com/cars-owners-service/internal/data/store"
)

type JSONError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type ServerRunInfo map[string]*JSONError

func ServiceRuns(r *http.Request) ServerRunInfo {
	return ServerRunInfo{
		store.Store: &JSONError{
			Code:    0,
			Message: "Test message",
		},
	}.clean()
}

func (m ServerRunInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*JSONError(m))
}

func (m ServerRunInfo) clean() ServerRunInfo {
	var keysToDelete []string
	for k, v := range m {
		if v == nil {
			keysToDelete = append(keysToDelete, k)
		}
	}
	for _, k := range keysToDelete {
		delete(m, k)
	}
	return m
}
