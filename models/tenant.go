package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null"`
	Slug   string `gorm:"type:varchar(255);not null;unique_index"`
	Status bool   `gorm:"type:bool;default:true"`
}
