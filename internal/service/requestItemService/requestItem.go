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

func (r *RequestItemService) GetRequestItems(ctx context.Context, req *httpmodels.TestingGetRequestItemsRequest) (*httpmodels.TestingGetRequestItemsResponse, error) {
	items, err := r.Repository.GetRequestItems(req.RequestID)
	if err != nil {
		return nil, errors.Wrap(err, "get request items")
	}

	convertItems := convertItemsToResponse(items)
	itemsMap := make(map[uint64]*httpmodels.ItemInRequest)

	for _, item := range convertItems {
		if _, ok := itemsMap[item.ID]; !ok {
			itemsMap[item.ID] = &httpmodels.ItemInRequest{
				Item:              *item,
				QuantityInRequest: 0,
			}
		}
		itemsMap[item.ID].QuantityInRequest += 1
	}

	res := make([]*httpmodels.ItemInRequest, 0, len(itemsMap))
	for _, v := range itemsMap {
		res = append(res, v)
	}

	resp := &httpmodels.TestingGetRequestItemsResponse{
		RequestItems: res,
	}

	return resp, nil
}

func (r *RequestItemService) DeleteDraftRequestItem(ctx context.Context, req *httpmodels.TestingDeleteDraftRequestItemsRequest) (*httpmodels.TestingDeleteDraftRequestItemsResponse, error) {
	items, err := r.Repository.GetRequestItems(req.RequestID)
	if err != nil {
		return nil, errors.Wrap(err, "get request items")
	}

	convertItems := convertItemsToResponse(items)
	itemsMap := make(map[uint64]*httpmodels.ItemInRequest)
	notDeleteFlag := true
	for _, item := range convertItems {
		if item.ID == uint64(req.ItemID) && notDeleteFlag {
			notDeleteFlag = false
			continue
		}
		if _, ok := itemsMap[item.ID]; !ok {
			itemsMap[item.ID] = &httpmodels.ItemInRequest{
				Item:              *item,
				QuantityInRequest: 0,
			}
		}
		itemsMap[item.ID].QuantityInRequest += 1
	}

	res := make([]*httpmodels.ItemInRequest, 0, len(itemsMap))
	for _, v := range itemsMap {
		res = append(res, v)
	}

	err = r.Repository.DeleteDraftRequestItem(req.RequestID, req.ItemID)
	if err != nil {
		return nil, errors.Wrap(err, "delete item from request")
	}
	resp := &httpmodels.TestingDeleteDraftRequestItemsResponse{
		RequestItems: res,
	}

	return resp, nil
}

func convertItemsToResponse(items []*ds.Item) []*httpmodels.Item {
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

// func convertRequestsToResponse(arr []*ds.Request) []*httpmodels.Request {
// 	result := make([]*httpmodels.Request, 0)
// 	for _, request := range arr {
// 		result = append(result, &httpmodels.Request{
// 			ID:             request.ID,
// 			Status:         request.Status,
// 			CreationDate:   request.CreationDate,
// 			FormationDate:  request.FormationDate,
// 			CompletionDate: request.CompletionDate,
// 			CreatorID:      request.CreatorID,
// 		})
// 	}
// 	return result
// }
