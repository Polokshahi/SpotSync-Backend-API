package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	ZoneID uint `gorm:"not null"`

	LicensePlate string `gorm:"size:15;not null"`
	Status        string `gorm:"type:varchar(20);default:active"`

	User User `gorm:"foreignKey:UserID"`
	Zone ParkingZone `gorm:"foreignKey:ZoneID"`
}