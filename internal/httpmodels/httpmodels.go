package httpmodels

import (
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
)

type TestingRegisterRequest struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}

type TestingRegisterResponse struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}

type TestingLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TestingLoginResponse struct {
	Token string `json:"token"`
}

type TestingGetItemsRequest struct {
	SearchText string `json:"SearchText"`
}

type TestingGetItemsResponse struct {
	Items   []*Item `json:"items"`
	OrderID uint64  `json:"oderdID"`
}

type TestingGetItemByIDRequest struct {
	ID int64 `json:"id"`
}

type TestingGetItemByIDResponse struct {
	Item Item `json:"item"`
}

type TestingDeleteItemRequest struct {
	ID int64 `json:"id"`
}

type TestingPostItemRequest struct {
	Item Item `json:"item"`
}

type TestingPostItemResponse struct {
	Item Item `json:"item"`
}

type TestingPutItemRequset struct {
	Item Item `json:"item"`
}

type TestingGetUserByIDRequest struct {
	ID int64 `json:"id"`
}

type TestingGetUserByIDResponse struct {
	User User `json:"user"`
}

type TestingGetDraftRequestByIDRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type TestingGetDraftRequestByIDResponse struct {
	Request Request `json:"request"`
}

// type TestingPostItemToRequestRequest struct {
// 	UserID int64  `json:"userID"`
// 	Status string `json:"status"`
// }

type TestingPostItemToRequestResponse struct {
	RequestItem RequestItem `json:"requestItem"`
}

type TestingPostRequestRequest struct {
	Request Request `json:"request"`
}

type TestingPostRequestResponse struct {
	Request Request `json:"request"`
}

type TestingPostRequestItemRequest struct {
	RequestItem RequestItem `json:"requestItem"`
}

type TestingGetRequestsForAdminWithFiltersRequest struct {
	MinData time.Time `json:"minData"`
	MaxData time.Time `json:"maxData"`
	Status  string    `json:"status"`
	Login   string    `json:"login"`
}

type TestingGetRequestsForAdminWithFiltersResponse struct {
	Requests []*Request `json:"requests"`
}

type TestingGetRequestsRequest struct {
	CreatorID int64 `json:"creatorID"`
}

type TestingGetRequestsResponse struct {
	Requests []*Request `json:"requests"`
}

type TestingGetRequestItemsRequest struct {
	RequestID int64 `json:"requestID"`
}

type TestingGetRequestItemsResponse struct {
	RequestItems []*ItemInRequest `json:"requestItems"`
}

type TestingGetRequestByIDRequest struct {
	RequestID int64 `json:"requestID"`
}

type TestingGetRequestByIDResponse struct {
	Request Request `json:"request"`
}

type TestingPutRequestStatusRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type TestingDeleteRequestRequest struct {
	ID int64 `json:"id"`
}

type TestingDeleteDraftRequestItemsRequest struct {
	RequestID int64 `json:"requestID"`
	ItemID    int64 `json:"itemID"`
}

type TestingDeleteDraftRequestItemsResponse struct {
	RequestItems []*ItemInRequest `json:"requestItems"`
}
