package responses

import (
	"encoding/json"
)

type JSONServerErrors map[string]error

func (m JSONServerErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]error(m))
}

func (m JSONServerErrors) Clean() JSONServerErrors {
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
