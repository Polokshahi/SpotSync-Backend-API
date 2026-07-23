package repository

import (
	"spotsync/internal/models"

	"gorm.io/gorm"
)

type ZoneRepository interface {
	Create(zone *models.ParkingZone) error
	GetAll() ([]models.ParkingZone, error)
	GetByID(id uint) (*models.ParkingZone, error)
	Update(zone *models.ParkingZone) error
	Delete(id uint) error
}

type zoneRepo struct {
	db *gorm.DB
}

func NewZoneRepository(db *gorm.DB) ZoneRepository {
	return &zoneRepo{
		db: db,
	}
}

func (r *zoneRepo) Create(zone *models.ParkingZone) error {
	return r.db.Create(zone).Error
}

func (r *zoneRepo) GetAll() ([]models.ParkingZone, error) {

	var zones []models.ParkingZone

	err := r.db.Find(&zones).Error

	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (r *zoneRepo) GetByID(id uint) (*models.ParkingZone, error) {

	var zone models.ParkingZone

	err := r.db.First(&zone, id).Error

	if err != nil {
		return nil, err
	}

	return &zone, nil
}

func (r *zoneRepo) Update(zone *models.ParkingZone) error {
	return r.db.Save(zone).Error
}

func (r *zoneRepo) Delete(id uint) error {
	return r.db.Delete(&models.ParkingZone{}, id).Error
}