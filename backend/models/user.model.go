package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"autoIncrement; primarykey"`
	Name      string `gorm:"type:varchar(255); not null"`
	Email     string `gorm:"type:varchar(255); uniqueIndex; not null"`
	Password  string `gorm:"type:varchar(255); not null"`
	RoleId    int    `gorm:"default:2"`
	Status    bool   `gorm:"type:bool; default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	RoleId int    `json:"role_id,omitempty"`
	Status bool   `json:"status,omitempty"`
}

type SignUpUser struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
