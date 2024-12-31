package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"type:varchar(120);not null"`
	LastName  string `gorm:"type:varchar(120);not null"`
	Email     string `gorm:"type:varchar(120);not null;unique_index"`
	TenantId  int
}
