package requestitemservice

import (
	"context"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/pkg/errors"
)

type RequestItemService struct {
	Repository *repository.Repository
	Config     *config.Config
}

func NewRequestItemService(repo *repository.Repository, cfg *config.Config) *RequestItemService {
	return &RequestItemService{
		Repository: repo,
		Config:     cfg,
	}
}

func (r *RequestItemService) PostRequestItem(ctx context.Context, req *httpmodels.TestingPostRequestItemRequest) error {
	requestItemDB := ds.ItemsRequest{
		ItemID:    req.RequestItem.ItemID,
		RequestID: req.RequestItem.RequestID,
	}

	err := r.Repository.PostRequestItem(requestItemDB)
	if err != nil {
		return errors.Wrap(err, "post item")
	}

	return nil
}
