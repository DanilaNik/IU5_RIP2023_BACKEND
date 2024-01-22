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
		return nil, errors.Wrap(err, "get draft request")
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
		Status:       req.Request.Status,
		CreatorID:    req.Request.CreatorID,
		CreationDate: req.Request.CreationDate,
	}

	err := r.Repository.PostRequest(requestDB)
	if err != nil {
		return errors.Wrap(err, "post item")
	}

	return nil
}

func (r *RequestService) GetRequestsForAdminWithFilters(ctx context.Context, req *httpmodels.TestingGetRequestsForAdminWithFiltersRequest) (*httpmodels.TestingGetRequestsForAdminWithFiltersResponse, error) {
	requests, err := r.Repository.GetRequestsForAdminWithFilters(req.MinData, req.MaxData, req.Status)
	if err != nil {
		return nil, errors.Wrap(err, "get requests with filters")
	}
	result := convertRequestInfoToResponse(requests)

	resp := &httpmodels.TestingGetRequestsForAdminWithFiltersResponse{
		Requests: result,
	}

	return resp, nil
}

func (r *RequestService) GetRequests(ctx context.Context, req *httpmodels.TestingGetRequestsRequest) (*httpmodels.TestingGetRequestsForAdminWithFiltersResponse, error) {
	requests, err := r.Repository.GetRequests(req.CreatorID)
	if err != nil {
		return nil, errors.Wrap(err, "get requests with filters")
	}

	result := make([]*httpmodels.RequestInfo, 0)
	for _, request := range requests {
		reqInfo := httpmodels.RequestInfo{
			ID:             request.ID,
			Status:         request.Status,
			CreationDate:   request.CreationDate,
			FormationDate:  request.FormationDate,
			CompletionDate: request.CompletionDate,
			CreatorID:      request.CreatorID,
		}
		result = append(result, &reqInfo)
	}
	resp := &httpmodels.TestingGetRequestsForAdminWithFiltersResponse{
		Requests: result,
	}

	return resp, nil
}

func (r *RequestService) GetRequestByID(ctx context.Context, req *httpmodels.TestingGetRequestByIDRequest) (*httpmodels.TestingGetRequestByIDResponse, error) {
	request, err := r.Repository.GetRequestByID(req.RequestID)
	if err != nil {
		return nil, errors.Wrap(err, "get request by id")
	}

	resp := &httpmodels.TestingGetRequestByIDResponse{
		Request: httpmodels.Request{
			ID:             request.ID,
			Status:         request.Status,
			CreationDate:   request.CreationDate,
			FormationDate:  request.FormationDate,
			CompletionDate: request.CompletionDate,
			CreatorID:      request.CreatorID,
			AdminID:        request.AdminID,
		},
	}
	return resp, nil
}

func (r *RequestService) PutRequestStatus(ctx context.Context, req *httpmodels.TestingPutRequestStatusRequest) error {
	err := r.Repository.PutRequestStatus(req.ID, req.Status, req.AdminId)
	if err != nil {
		return errors.Wrap(err, "put aprove request status")
	}
	return nil
}

func (r *RequestService) ConfirmRequest(ctx context.Context, req *httpmodels.TestingPutRequestStatusRequest) error {
	err := r.Repository.ConfirmRequest(req.ID, req.Status)
	if err != nil {
		return errors.Wrap(err, "confirm request status")
	}
	return nil
}

func (r *RequestService) DeleteRequest(ctx context.Context, req *httpmodels.TestingDeleteRequestRequest) error {
	err := r.Repository.DeleteRequest(req.ID)
	if err != nil {
		return errors.Wrap(err, "delete request")
	}
	return nil
}

// func convertToResponse(arr []*ds.Request) []*httpmodels.Request {
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

func convertRequestInfoToResponse(arr []*ds.RequestInfo) []*httpmodels.RequestInfo {
	result := make([]*httpmodels.RequestInfo, 0)
	for _, request := range arr {
		result = append(result, &httpmodels.RequestInfo{
			ID:             request.ID,
			Status:         request.Status,
			CreationDate:   request.CreationDate,
			FormationDate:  request.FormationDate,
			CompletionDate: request.CompletionDate,
			CreatorID:      request.CreatorID,
			AdminID:        request.AdminID,
			UserEmail:      request.Email,
		})
	}
	return result
}
