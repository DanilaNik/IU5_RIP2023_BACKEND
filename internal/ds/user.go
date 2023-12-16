package ds

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
)

type User struct {
	//gorm.Model
	ID       uint64    `json:"id" gorm:"primary_key"`
	UserName string    `json:"userName" gorm:"type:varchar(100);not null"`
	Login    string    `json:"login" gorm:"type:text;not null"`
	Password string    `json:"password" gorm:"type:varchar(100);not null"`
	Email    string    `json:"email" gorm:"unique;type:varchar(100);not null"`
	Role     role.Role `json:"role" gorm:"type:varchar(20);check:role IN ('Admin', 'Moderator', 'User');not null"`
	ImageURL string    `json:"image_url" gorm:"type:varchar(500);default:'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/default-avatar.png'"`
}
