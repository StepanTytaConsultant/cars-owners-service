package requests

import (
	"encoding/json"
	"github.com/cars-owners-service/resources"
	"github.com/pkg/errors"
	"net/http"
)

type PostOwnershipRequest struct {
	OwnershipAddRequest resources.OwnershipAddPostRequest
}

func NewPostOwnershipRequest(r *http.Request) (*PostOwnershipRequest, error) {
	req := PostOwnershipRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req.OwnershipAddRequest); err != nil {
		return nil, errors.Wrap(err, "failed to decode add ownership request")
	}
	return &req, nil
}
