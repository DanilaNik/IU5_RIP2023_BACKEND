package ds

import "gorm.io/gorm"

type ItemsRequest struct {
	gorm.Model
	ID        uint64  `json:"id" gorm:"primary_key"`
	ItemID    uint64  `json:"itemID"`
	RequestID uint64  `json:"requestID"`
	Item      Item    `json:"item" gorm:"foreignkey:ItemID"`
	Request   Request `json:"request" gorm:"foreignkey:RequestID"`
}
