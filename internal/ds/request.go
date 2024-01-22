package ds

import (
	"time"

	"gorm.io/gorm"
)

/*
	Request Status Code:
	draft - 1
	deleted - 2
	formed - 3
	completed - 4
	rejected - 5
*/

type Request struct {
	gorm.Model
	ID             uint64    `json:"id" gorm:"primary_key"`
	Status         string    `json:"status" gorm:"type:varchar(30);check:status in ('draft','deleted','formed','completed','rejected');not null"`
	CreationDate   time.Time `json:"creationDate" gorm:"type:timestamp;not null"`
	FormationDate  time.Time `json:"formationDate" gorm:"type:timestamp"`
	CompletionDate time.Time `json:"completionDate" gorm:"type:timestamp"`
	CreatorID      uint64    `json:"creatorID"`
	AdminID        uint64    `json:"adminID"`
	User           User      `json:"user" gorm:"foreignkey:CreatorID"`
}

type RequestInfo struct {
	Request
	Email string `json:"email"`
}
