package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRequests(ctx *gin.Context) {
	id, role, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId, _ := strconv.ParseInt(id, 10, 64)

	if role == "Admin" {
		max := ctx.Query("max_date")
		min := ctx.Query("min_date")
		status := ctx.Query("status")

		var minTime, maxTime time.Time

		if min != "" {
			minTime, _ = time.Parse("2006-01-02", min)
		}
		if max != "" {
			maxTime, _ = time.Parse("2006-01-02", max)
		} else {
			maxTime = time.Now()
		}

		req := httpmodels.TestingGetRequestsForAdminWithFiltersRequest{
			MinData:   minTime,
			MaxData:   maxTime,
			Status:    status,
			CreatorID: userId,
		}
		data, err := h.RequestService.GetRequestsForAdminWithFilters(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
			return
		}

		ctx.JSON(http.StatusOK, data)
		return
	}

	req := httpmodels.TestingGetRequestsRequest{
		CreatorID: userId,
	}
	data, err := h.RequestService.GetRequests(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) GetRequestById(ctx *gin.Context) {
	_, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	requesId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	req1 := httpmodels.TestingGetRequestItemsRequest{
		RequestID: requesId,
	}

	dataReqestItems, err := h.RequestItemService.GetRequestItems(ctx, &req1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	req2 := httpmodels.TestingGetRequestByIDRequest{
		RequestID: requesId,
	}
	dataRequest, err := h.RequestService.GetRequestByID(ctx, &req2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	res := httpmodels.UserRequest{
		Request: dataRequest.Request,
		Items:   dataReqestItems.RequestItems,
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) PutRequestStatus(ctx *gin.Context) {
	_, role, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	if role != "Admin" {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}

	requesId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error(),
		})
		return
	}

	status := struct {
		Status string `json:"status"`
	}{}

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req := httpmodels.TestingPutRequestStatusRequest{
		ID:     requesId,
		Status: status.Status,
	}

	err = h.RequestService.PutRequestStatus(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) ConfirmRequest(ctx *gin.Context) {
	id, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, _ := strconv.ParseInt(id, 10, 64)

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, int(userID), "draft")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	req := httpmodels.TestingPutRequestStatusRequest{
		ID:     int64(dataRequest.Request.ID),
		Status: "formed",
	}

	err = h.RequestService.ConfirmRequest(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteRequest(ctx *gin.Context) {
	_, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	id := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(jsonData, &id)
	if err != nil || id.Id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req := httpmodels.TestingDeleteRequestRequest{
		ID: id.Id,
	}

	err = h.RequestService.DeleteRequest(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) DeleteItemFromRequest(ctx *gin.Context) {
	id, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	userID, _ := strconv.ParseInt(id, 10, 64)

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, int(userID), "draft")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	req := httpmodels.TestingDeleteDraftRequestItemsRequest{
		RequestID: int64(dataRequest.Request.ID),
		ItemID:    itemId,
	}
	dataReqestItems, err := h.RequestItemService.DeleteDraftRequestItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	res := httpmodels.UserRequest{
		Request: dataRequest.Request,
		Items:   dataReqestItems.RequestItems,
	}

	ctx.JSON(http.StatusOK, res)
}

// func (h *Handler) PutOrderStatus(ctx *gin.Context) {
// 	id, role, err := h.getUserRole(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusForbidden, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	if role != "Admin" {
// 		ctx.JSON(http.StatusForbidden, gin.H{})
// 		return
// 	}
// 	orderId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

// 	status := &struct {
// 		Status string `json:"status"`
// 	}{}

// 	jsonData, err := ctx.GetRawData()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error: ": err.Error(),
// 		})
// 		return
// 	}
// 	err = json.Unmarshal(jsonData, status)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
// 		return
// 	}

// 	order := &models.Order{}
// 	tx := s.db.DB.Where("deleted_at IS NULL").Where("id = ?", orderId).First(&order)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error: ": tx.Error.Error()})
// 		return
// 	}
// 	order.Status = status.Status
// 	order.AdminId, _ = strconv.ParseUint(id, 10, 64)
// 	tx = s.db.DB.Where("id = ?", orderId).Updates(order)
// 	if tx.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error: ": tx.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }
