package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRequests(ctx *gin.Context) {
	users, err := h.Repository.GetRequests()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GetRequestById(ctx *gin.Context) {
	idValue := ctx.Param("id")
	id, _ := strconv.Atoi(idValue)
	user, err := h.Repository.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteRequest(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	id := struct {
		Id uint64 `json:"id"`
	}{}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.DeleteRequest(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
}

func (h *Handler) UpdateRequestStatus(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	id := struct {
		Id uint64 `json:"id"`
	}{}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.DeleteRequest(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
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
