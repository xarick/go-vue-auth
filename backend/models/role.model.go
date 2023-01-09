package models

type Role struct {
	ID     uint   `gorm:"autoIncrement; primarykey"`
	Name   string `gorm:"type:varchar(255); unique; not null"`
	Status bool   `gorm:"type:bool; default:true"`
}
