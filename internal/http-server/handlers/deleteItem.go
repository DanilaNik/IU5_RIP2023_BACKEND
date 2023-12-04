package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewDeleteItem(ctx *gin.Context) {
	idValue := ctx.Param("id")
	id, _ := strconv.Atoi(idValue)
	err := h.Repository.DeleteItem(id)
	if err != nil {
		h.Logger.Error(err.Error())
		h.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/items")
	return
}

func (h *Handler) JSONDeleteItem(ctx *gin.Context) {
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
	err = h.Repository.DeleteItem(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
}
