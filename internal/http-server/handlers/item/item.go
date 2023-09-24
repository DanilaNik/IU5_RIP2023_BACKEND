package item

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/DanilaNik/IU5_RIP2023/internal/storage"
	"github.com/gin-gonic/gin"
)

func New(ctx *gin.Context) {
	filter := ctx.Query("filter")
	searchText := ctx.Query("search")
	data := storage.GetItems()

	if searchText != "" {
		var filteredItems []storage.Item
		for _, item := range data.Items {
			sarchBarcode, _ := strconv.Atoi(searchText)
			if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchText)) || item.Barcode == uint64(sarchBarcode) {
				filteredItems = append(filteredItems, item)
			}
		}
		filteredItems = filterItems(filteredItems, filter)
		ctx.HTML(http.StatusOK, "items.tmpl", storage.ItemsData{Items: filteredItems, Filter: filter, SearchText: searchText})
		return
	}

	res := filterItems(data.Items, filter)
	ctx.HTML(http.StatusOK, "items.tmpl", storage.ItemsData{Items: res, Filter: filter, SearchText: searchText})
}

func filterItems(arr []storage.Item, f string) []storage.Item {
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

//func New(ctx *gin.Context) {
//	filter := ctx.Query("filter")
//	searchText := ctx.Query("search")
//	data := storage.GetItems()
//
//	//var q1, q2 uint64
//	//if quantityMin == "" {
//	//	q1 = 0
//	//} else {
//	//	q1, err := strconv.ParseUint(quantityMin, 10, 64)
//	//	_ = q1
//	//	if err != nil {
//	//		return
//	//	}
//	//}
//	//if quantityMax == "" {
//	//	q2 = uint64(^uint32(0))
//	//} else {
//	//	q2, err := strconv.ParseUint(quantityMax, 10, 64)
//	//	_ = q2
//	//	if err != nil {
//	//		return
//	//	}
//	//}
//	//
//	//res := make([]*storage.Item, 0)
//	//
//	//for _, v := range data {
//	//	q := v.Quantity
//	//	if q < q1 || q >= q2 {
//	//		continue
//	//	}
//	//	if filter == "all" {
//	//		res = append(res, v)
//	//	} else if v.Status == filter {
//	//		res = append(res, v)
//	//	}
//	//}
//
//	if searchText != "" {
//		var filteredItems []storage.Item
//		for _, item := range data.Items {
//			if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchText)) {
//				filteredItems = append(filteredItems, item)
//			}
//		}
//		filteredItems = filterItems(filteredItems, filter)
//		ctx.HTML(http.StatusOK, "items.tmpl", storage.ItemsData{Items: filteredItems})
//		return
//	}
//	res := filterItems(data.Items, filter)
//	ctx.HTML(http.StatusOK, "items.tmpl", storage.ItemsData{Items: res})
//
//}
//
//func filterItems(arr []storage.Item, f string) []storage.Item {
//	switch f {
//	case "min":
//		sort.Slice(arr, func(i, j int) bool {
//			return arr[i].Quantity < arr[j].Quantity
//		})
//	case "max":
//		sort.Slice(arr, func(i, j int) bool {
//			return arr[i].Quantity > arr[j].Quantity
//		})
//	}
//	return arr
//}

//res := make([]*internal.Item, 0)
//if idStr := ctx.Query("item"); idStr != "" {
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		return
//	}
//	var currentItem storage.Item
//	for _, item := range data {
//		if uint(id) == item.ID {
//			currentItem = *item
//			break
//		}
//	}
//	if
//}
