package httpmodels

import (
	"mime/multipart"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
)

type Item struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
	Quantity uint64 `json:"quantity"`
	Height   uint64 `json:"height"`
	Width    uint64 `json:"width"`
	Depth    uint64 `json:"depth"`
	Barcode  uint64 `json:"barcode"`
}

type User struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}

type Request struct {
	ID             uint64    `json:"id"`
	Status         string    `json:"status"` //status in ('draft','deleted','formed','completed','rejected')
	CreationDate   time.Time `json:"creationDate"`
	FormationDate  time.Time `json:"formationDate"`
	CompletionDate time.Time `json:"completionDate"`
	CreatorID      uint64    `json:"creatorID"`
}

type RequestItem struct {
	ItemID    uint64 `json:"itemID"`
	RequestID uint64 `json:"requestID"`
}

type ItemInRequest struct {
	Item
	QuantityInRequest uint64 `json:"quantityInRequest"`
}

type UserRequest struct {
	Request Request          `json:"request"`
	Items   []*ItemInRequest `json:"items"`
}

type FormSwagger struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type ImageSwagger struct {
	Link  string `json:"link"`
	Error string `json:"error"`
}

type RequestStatus struct {
	Status string `json:"status"`
}

type RequestID struct {
	Id int64 `json:"id"`
}
