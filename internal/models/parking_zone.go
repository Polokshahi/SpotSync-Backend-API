package models

import "gorm.io/gorm"

type ParkingZone struct {
	gorm.Model

	Name          string  `gorm:"size:100;not null"`
	Type          string  `gorm:"type:varchar(30);not null"`
	TotalCapacity int     `gorm:"not null"`
	PricePerHour  float64 `gorm:"type:decimal(10,2);not null"`

	Reservations []Reservation `gorm:"foreignKey:ZoneID"`
}