package ds

import "gorm.io/gorm"

/*
	Item Status Code:
	enabled - 1
	deleted - 2
*/

type Item struct {
	gorm.Model
	ID       uint64 `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(300);not null"`
	ImageURL string `json:"image_url" gorm:"type:varchar(500);default:'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Intel-Core-i7-9700K.jpg'"`
	Status   string `json:"status" gorm:"type:varchar(30);check:status IN ('enabled', 'deleted');not null"`
	Quantity uint64 `json:"quantity" gorm:"default:0;not null"`
	Material string `json:"material" gorm:"type:varchar(50);not null"`
	Height   uint64 `json:"height" gorm:"not null"`
	Width    uint64 `json:"width" gorm:"not null"`
	Depth    uint64 `json:"depth" gorm:"not null"`
	Barcode  uint64 `json:"barcode" gorm:"unique;check:barcode > 0;not null"`
}

type ItemsData struct {
	Items      []Item
	Filter     string
	SearchText string
	Status     string
}
