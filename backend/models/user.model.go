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
