package responses

import "encoding/json"

type EmptyDataResponse struct {
	json.RawMessage
}

type EmptyMetaResponse struct {
	json.RawMessage
}
