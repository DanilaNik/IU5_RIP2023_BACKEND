package handlers

import (
	"net/http"
	"sort"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NewGetItems(ctx *gin.Context) {
	filter := ctx.Query("filter")
	searchText := ctx.Query("search")
	items, err := h.Repository.GetItems(searchText)
	if err != nil {
		h.Logger.Error(err.Error())
		h.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	filteredItems := filterItems(*items, filter)
	ctx.HTML(http.StatusOK, "items.tmpl", ds.ItemsData{Items: filteredItems, Filter: filter, SearchText: searchText})
	return
}

func (h *Handler) JSONGetItems(ctx *gin.Context) {
	searchText := ctx.Query("")
	items, err := h.Repository.GetItems(searchText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func filterItems(arr []ds.Item, f string) []ds.Item {
	switch f {
	case "min":
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i].Quantity < arr[j].Quantity
		})
	case "max":
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i].Quantity > arr[j].Quantity
		})
	}
	return arr
}
