package handlers

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"net/http"
	"strconv"

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
