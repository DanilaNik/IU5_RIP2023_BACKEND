package itemservice

import (
	"context"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/pkg/errors"
)

type ItemService struct {
	Repository *repository.Repository
	Config     *config.Config
}

func NewItemService(repo *repository.Repository, cfg *config.Config) *ItemService {
	return &ItemService{
		Repository: repo,
		Config:     cfg,
	}
}

func (i *ItemService) GetItems(ctx context.Context, searchText string) (*httpmodels.TestingGetItemsResponse, error) {
	items, err := i.Repository.GetItems(searchText, i.Config.ServiceHost)
	if err != nil {
		return nil, errors.Wrap(err, "get online items")
	}
	res := convertToResponse(items)
	resp := &httpmodels.TestingGetItemsResponse{
		Items: res,
	}

	return resp, nil
}

func (i *ItemService) GetItemByID(ctx context.Context, req *httpmodels.TestingGetItemByIDRequest) (*httpmodels.TestingGetItemByIDResponse, error) {
	item, err := i.Repository.GetItemByID(int(req.ID), i.Config.ServiceHost)
	if err != nil {
		return nil, errors.Wrap(err, "get item by id")
	}

	resp := &httpmodels.TestingGetItemByIDResponse{
		Item: httpmodels.Item{
			ID:       item.ID,
			Name:     item.Name,
			ImageURL: item.ImageURL,
			Status:   item.Status,
			Quantity: item.Quantity,
			Height:   item.Height,
			Width:    item.Width,
			Depth:    item.Depth,
			Barcode:  item.Barcode,
		},
	}

	return resp, nil
}

func (i *ItemService) PostItem(ctx context.Context, req *httpmodels.TestingPostItemRequest) error {
	itemDB := ds.Item{
		Name:     req.Item.Name,
		ImageURL: req.Item.ImageURL,
		Status:   req.Item.Status,
		Quantity: req.Item.Quantity,
		Height:   req.Item.Height,
		Width:    req.Item.Width,
		Depth:    req.Item.Depth,
		Barcode:  req.Item.Barcode,
	}

	err := i.Repository.PostItem(itemDB)
	if err != nil {
		return errors.Wrap(err, "post item")
	}

	return nil
}

func (i *ItemService) DeleteItem(ctx context.Context, req *httpmodels.TestingDeleteItemRequest) error {
	err := i.Repository.DeleteItem(int(req.ID))
	if err != nil {
		errors.Wrap(err, "delete item")
	}

	return nil
}

func (i *ItemService) PutItem(ctx context.Context, req *httpmodels.TestingPutItemRequset, id int64) error {
	itemDB := ds.Item{
		Name:     req.Item.Name,
		ImageURL: req.Item.ImageURL,
		Status:   req.Item.Status,
		Quantity: req.Item.Quantity,
		Height:   req.Item.Height,
		Width:    req.Item.Width,
		Depth:    req.Item.Depth,
		Barcode:  req.Item.Barcode,
	}
	err := i.Repository.PutItem(itemDB, id)
	if err != nil {
		return errors.Wrap(err, "put item")
	}

	return nil
}

func (i *ItemService) PostItemToRequest(ctx context.Context, req *httpmodels.TestingGetItemByIDRequest) {

}

func convertToResponse(items []*ds.Item) []*httpmodels.Item {
	result := make([]*httpmodels.Item, 0)
	for _, item := range items {
		result = append(result, &httpmodels.Item{
			ID:       item.ID,
			Name:     item.Name,
			ImageURL: item.ImageURL,
			Status:   item.Status,
			Quantity: item.Quantity,
			Height:   item.Height,
			Width:    item.Width,
			Depth:    item.Depth,
			Barcode:  item.Barcode,
		})
	}
	return result
}
