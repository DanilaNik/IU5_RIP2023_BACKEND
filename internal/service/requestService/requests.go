package requestservice

import (
	"context"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/pkg/errors"
)

type RequestService struct {
	Repository *repository.Repository
	Config     *config.Config
}

func NewRequestService(repo *repository.Repository, cfg *config.Config) *RequestService {
	return &RequestService{
		Repository: repo,
		Config:     cfg,
	}
}

func (r *RequestService) GetDraftRequestByIdAndStatus(ctx context.Context, userID int, status string) (*httpmodels.TestingGetDraftRequestByIDResponse, error) {
	request, err := r.Repository.GetRequestByIDAndStatus(userID, status)
	if err != nil {
		return nil, errors.Wrap(err, "get online items")
	}

	resp := &httpmodels.TestingGetDraftRequestByIDResponse{
		Request: httpmodels.Request{
			ID:             request.ID,
			Status:         request.Status,
			CreationDate:   request.CreationDate,
			FormationDate:  request.FormationDate,
			CompletionDate: request.CompletionDate,
			CreatorID:      request.CreatorID,
		},
	}

	return resp, nil
}

func (r *RequestService) PostRequest(ctx context.Context, req *httpmodels.TestingPostRequestRequest) error {
	requestDB := ds.Request{
		Status:    req.Request.Status,
		CreatorID: req.Request.CreatorID,
	}

	err := r.Repository.PostRequest(requestDB)
	if err != nil {
		return errors.Wrap(err, "post item")
	}

	return nil
}
