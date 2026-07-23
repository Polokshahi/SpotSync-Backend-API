package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);default:driver"`

	Reservations []Reservation `gorm:"foreignKey:UserID"`
}