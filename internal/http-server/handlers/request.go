package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/gin-gonic/gin"
)

// GetRequests godoc
// @Summary      Get list of all orders
// @Tags         orders
// @Param        min_date    query     string  false  "min date"  Format(text)
// @Param        max_date    query     string  false  "max date"  Format(text)
// @Param        status      query     string  false  "order status"  Format(text)
// @Param        login       query     string  false  "order creator"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpmodels.TestingGetRequestsForAdminWithFiltersResponse
// @Router       /orders [get]
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
			minTime, _ = time.Parse("2006-01-02 15:04:05", min+" 00:00:00")
		}
		if max != "" {
			maxTime, _ = time.Parse("2006-01-02 15:04:05", max+" 23:59:59")
		} else {
			location, _ := time.LoadLocation("Europe/Moscow")

			day := time.Now().In(location)
			maxTime = time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, location)
		}
		if status == "" {
			status = "all"
		}

		req := httpmodels.TestingGetRequestsForAdminWithFiltersRequest{
			MinData: minTime,
			MaxData: maxTime,
			Status:  status,
		}
		data, err := h.RequestService.GetRequestsForAdminWithFilters(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetRequestById godoc
// @Summary      Get order by id
// @Tags         orders
// @Param        id    path     string  true  "order id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200  {object} httpmodels.UserRequest
// @Router       /orders/{id} [get]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req2 := httpmodels.TestingGetRequestByIDRequest{
		RequestID: requesId,
	}
	dataRequest, err := h.RequestService.GetRequestByID(ctx, &req2)
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

// PutRequestStatus godoc
// @Summary      Approve or decline order
// @Tags         orders
// @Param        status body httpmodels.RequestStatus true "Order status"
// @Param        id    path     string  true  "order id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /orders/{id}/approve [put]
func (h *Handler) PutRequestStatus(ctx *gin.Context) {
	id, role, err := h.getUserRole(ctx)
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
	userId, _ := strconv.ParseInt(id, 10, 64)
	requesId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error(),
		})
		return
	}

	status := httpmodels.RequestStatus{}

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req := httpmodels.TestingPutRequestStatusRequest{
		ID:      requesId,
		AdminId: userId,
		Status:  status.Status,
	}

	err = h.RequestService.PutRequestStatus(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// ConfirmRequest godoc
// @Summary      Confirm current order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /orders/make [put]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req := httpmodels.TestingPutRequestStatusRequest{
		ID:     int64(dataRequest.Request.ID),
		Status: "formed",
	}

	err = h.RequestService.ConfirmRequest(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// DeleteRequest godoc
// @Summary      Delete order
// @Tags         orders
// @Param        id body httpmodels.RequestID true "Order id"
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /orders/delete [delete]
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
	id := httpmodels.RequestID{}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// DeleteItemFromRequest godoc
// @Summary      Delete item from current order
// @Tags         orders
// @Param        id    path     string  true  "item id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200 {object} httpmodels.UserRequest
// @Router       /orders/items/{id} [delete]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
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
