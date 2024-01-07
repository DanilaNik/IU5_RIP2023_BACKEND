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
)

// LoadS3 godoc
// @Summary      Upload s3 file
// @Tags         s3
// @Param file formData file true "upload file"
// @Param metadata formData string false "metadata"
// @Accept      mpfd
// @Accept       json
// @Produce      json
// @Success      200  {object}  Empty
// @Router       /s3/upload [post]
func (h *Handler) LoadS3(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	// filePath := "/files/" + newFileName
	contentType := file.Header["Content-Type"][0]
	buffer, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	h.Minio.PutObject(context.Background(), "cnc", newFileName, buffer, file.Size, minio.PutObjectOptions{ContentType: contentType})
	reqParams := make(url.Values)
	link, err := h.Minio.PresignedGetObject(context.Background(), "cnc", newFileName, 7*24*time.Hour, reqParams)
	if link == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"link": "",
			"err":  err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"link": link.String(),
		"err":  err,
	})
}

// @Summary GetItems
// @Description Get data about active items
// @Tags items
// @Accept  json
// @Produce  json
// @Success 201 {object} httpmodels.TestingGetItemsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items [get]
func (h *Handler) GetItems(ctx *gin.Context) {
	userId, _, userErr := h.getUserRole(ctx)
	searchText := ctx.Query("search")
	resp, err := h.ItemService.GetItems(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	if userErr != nil {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
	if err != nil {
		currentTime := time.Now()
		req1 := &httpmodels.TestingPostRequestRequest{
			Request: httpmodels.Request{
				CreatorID:    uint64(id),
				Status:       "draft",
				CreationDate: currentTime,
			},
		}
		err1 := h.RequestService.PostRequest(ctx, req1)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1})
			return
		}
		dataRequest, err1 = h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1})
			return
		}
	}

	resp.OrderID = dataRequest.Request.ID
	ctx.JSON(http.StatusOK, resp)
}

// @Summary GetItemById
// @Description Get data about item
// @Tags items
// @Accept  json
// @Produce  json
// @Success 201 {object} httpmodels.TestingGetItemByIDResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /item [get]
func (h *Handler) GetItemById(ctx *gin.Context) {
	idValue := ctx.Param("id")
	id, _ := strconv.Atoi(idValue)
	req := httpmodels.TestingGetItemByIDRequest{
		ID: int64(id),
	}

	resp, err := h.ItemService.GetItemByID(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) PostItem(ctx *gin.Context) {
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

	req := httpmodels.TestingPostItemRequest{
		Item: item,
	}

	err = h.ItemService.PostItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteItem(ctx *gin.Context) {
	itemId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := httpmodels.TestingDeleteItemRequest{
		ID: itemId,
	}

	err := h.ItemService.DeleteItem(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error: ": err.Error})
}

func (h *Handler) PutItem(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error: ": err.Error()})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

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
	// req := httpmodels.TestingPostItemToRequestRequest{
	// 	UserID: int64(id),
	// 	Status: "draft",
	// }

	dataRequest, err := h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
	if err != nil {
		currentTime := time.Now()
		req1 := &httpmodels.TestingPostRequestRequest{
			Request: httpmodels.Request{
				CreatorID:    uint64(id),
				Status:       "draft",
				CreationDate: currentTime,
			},
		}
		err1 := h.RequestService.PostRequest(ctx, req1)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1})
			return
		}
		dataRequest, err1 = h.RequestService.GetDraftRequestByIdAndStatus(ctx, id, "draft")
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err1})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err2})
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
