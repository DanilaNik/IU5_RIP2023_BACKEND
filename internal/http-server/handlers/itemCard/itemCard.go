package itemCard

import (
	"net/http"
	"strconv"

	"github.com/DanilaNik/IU5_RIP2023/internal/storage"
	"github.com/gin-gonic/gin"
)

func New(ctx *gin.Context) {
	id := ctx.Param("id")
	barcode, _ := strconv.Atoi(id)
	data := storage.GetItems()
	var status string
	var filteredItems []storage.Item
	for _, item := range data.Items {
		if item.Barcode == uint64(barcode) {
			filteredItems = append(filteredItems, item)
			status = "ok"
			break
		}
	}
	ctx.HTML(http.StatusOK, "item_card.tmpl", storage.ItemsData{Items: filteredItems, Status: status})
	return
}
