package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/swaggo/swag/example/celler/httputil"
)

// LoadS3 godoc
// @Summary      Upload s3 file
// @Tags         items
// @Param file formData file true "upload file"
// @Param metadata formData string false "metadata"
// @Accept       mpfd
// @Accept       json
// @Produce      json
// @Success      200  {object} httpmodels.ImageSwagger
// @Router       /items/image [post]
func (h *Handler) LoadS3(ctx *gin.Context) {
	var form httpmodels.FormSwagger
	err := ctx.ShouldBind(&form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(form.File.Filename)
	newFileName := uuid.New().String() + extension
	contentType := form.File.Header["Content-Type"][0]
	buffer, err := form.File.Open()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	h.Minio.PutObject(context.Background(), "warehouse", newFileName, buffer, form.File.Size, minio.PutObjectOptions{ContentType: contentType})
	reqParams := make(url.Values)
	link, err := h.Minio.PresignedGetObject(context.Background(), "warehouse", newFileName, 7*24*time.Hour, reqParams)
	if link != nil {
		ctx.JSON(http.StatusOK, httpmodels.ImageSwagger{
			Link:  link.String(),
			Error: "",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, httpmodels.ImageSwagger{
			Link:  "",
			Error: err.Error(),
		})
	}
}

// GetItems godoc
// @Summary      Get list of all items
// @Tags         items
// @Param        title    query     string  false  "filter by title"  Format(text)
// @Param        material    query     string  false  "filter by material"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpmodels.TestingGetItemsResponse
// @Router       /items [get]
func (h *Handler) GetItems(ctx *gin.Context) {
	userId, _, userErr := h.getUserRole(ctx)
	searchText := ctx.Query("title")
	material := ctx.Query("material")
	resp, err := h.ItemService.GetItems(ctx, searchText, material)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	if userErr != nil {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
	if err == nil {
		resp.OrderID = dataRequest.Request.ID
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetItemById godoc
// @Summary      Get item by id
// @Tags         items
// @Param        id    path     string  false  "item id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpmodels.TestingGetItemByIDResponse
// @Router       /items/{id} [get]
func (h *Handler) GetItemById(ctx *gin.Context) {
	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := httpmodels.TestingGetItemByIDRequest{
		ID: itemId,
	}

	resp, err := h.ItemService.GetItemByID(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// PostItem godoc
// @Summary      Create item
// @Tags         items
// @Param        itemPrototype body httpmodels.Item true "Item object"
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpmodels.Item
// @Router       /items/post [post]
func (h *Handler) PostItem(ctx *gin.Context) {
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

	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	item := httpmodels.Item{}
	err = json.Unmarshal(jsonData, &item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	item.Status = "enabled"

	req := httpmodels.TestingPostItemRequest{
		Item: item,
	}

	err = h.ItemService.PostItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

// Deleteitem godoc
// @Summary      Delete item by id
// @Tags         items
// @Param        id    path     int  true  "item id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /items/{id}/delete [delete]
func (h *Handler) DeleteItem(ctx *gin.Context) {
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

	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := httpmodels.TestingDeleteItemRequest{
		ID: itemId,
	}

	err = h.ItemService.DeleteItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// PutItem godoc
// @Summary      Change item
// @Tags         items
// @Param        itemPrototype body httpmodels.Item true "Item object"
// @Param        id    path     int  true  "item id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /items/{id}/put [put]
func (h *Handler) PutItem(ctx *gin.Context) {
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

	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
	}
	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	item := httpmodels.Item{}
	err = json.Unmarshal(jsonData, &item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	req := httpmodels.TestingPutItemRequset{
		Item: item,
	}

	err = h.ItemService.PutItem(ctx, &req, itemId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// PostItemToRequest godoc
// @Summary      Post item to current order
// @Tags         items
// @Param        id    path     int  true  "item id"  Format(text)
// @Accept       json
// @Produce      json
// @Success      200 {object} httpmodels.TestingGetDraftRequestByIDResponse
// @Router       /items/{id}/post [post]
func (h *Handler) PostItemToRequest(ctx *gin.Context) {
	userId, _, err := h.getUserRole(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	id, _ := strconv.Atoi(userId)

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
	if err != nil {
		location, _ := time.LoadLocation("Europe/Moscow")

		day := time.Now().In(location)

		req1 := &httpmodels.TestingPostRequestRequest{
			Request: httpmodels.Request{
				CreatorID:    uint64(id),
				Status:       "draft",
				CreationDate: day,
			},
		}
		err1 := h.RequestService.PostRequest(ctx, req1)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1.Error()})
			return
		}
		dataRequest, err1 = h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1.Error()})
			return
		}
	}

	req2 := &httpmodels.TestingPostRequestItemRequest{
		RequestItem: httpmodels.RequestItem{
			ItemID:    uint64(itemId),
			RequestID: dataRequest.Request.ID,
		},
	}
	err2 := h.RequestItemService.PostRequestItem(ctx, req2)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err2.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dataRequest)
}

// func filterItems(arr []ds.Item, f string) []ds.Item {
// 	switch f {
// 	case "min":
// 		sort.SliceStable(arr, func(i, j int) bool {
// 			return arr[i].Quantity < arr[j].Quantity
// 		})
// 	case "max":
// 		sort.SliceStable(arr, func(i, j int) bool {
// 			return arr[i].Quantity > arr[j].Quantity
// 		})
// 	}
// 	return arr
// }
