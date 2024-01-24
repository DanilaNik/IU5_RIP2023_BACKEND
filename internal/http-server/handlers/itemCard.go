package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewGetItemById(ctx *gin.Context) {
	idValue := ctx.Param("id")
	id, _ := strconv.Atoi(idValue)
	item, err := h.Repository.GetItemByID(id)
	if err != nil {
		h.Logger.Error(err.Error())
		h.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	res := make([]ds.Item, 0)
	res = append(res, *item)
	ctx.HTML(http.StatusOK, "item_card.tmpl", ds.ItemsData{Items: res, Status: item.Status})
	return
}

func (h *Handler) JSONGetItemById(ctx *gin.Context) {
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
	item, err := h.Repository.GetItemByID(int(id.Id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, item)
}