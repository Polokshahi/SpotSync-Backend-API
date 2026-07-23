package repository

import (
	"spotsync/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepository interface {

	// Transaction Methods
	LockZone(
		tx *gorm.DB,
		zoneID uint,
	) (*models.ParkingZone, error)


	GetActiveCount(
		tx *gorm.DB,
		zoneID uint,
	) (int64, error)


	CreateReservation(
		tx *gorm.DB,
		reservation *models.Reservation,
	) error


	// Normal Methods
	GetUserReservations(
		userID uint,
	) ([]models.Reservation, error)


	GetReservationByID(
		id uint,
	) (*models.Reservation, error)


	UpdateReservation(
		reservation *models.Reservation,
	) error


	GetAllReservations(
	) ([]models.Reservation, error)
}



type reservationRepo struct {
	db *gorm.DB
}



func NewReservationRepository(
	db *gorm.DB,
) ReservationRepository {

	return &reservationRepo{
		db: db,
	}
}



// =================================
// Transaction Methods
// =================================


func (r *reservationRepo) LockZone(
	tx *gorm.DB,
	zoneID uint,
) (*models.ParkingZone, error) {


	var zone models.ParkingZone


	err := tx.
		Clauses(
			clause.Locking{
				Strength: "UPDATE",
			},
		).
		First(&zone, zoneID).
		Error


	if err != nil {
		return nil, err
	}


	return &zone, nil
}



func (r *reservationRepo) GetActiveCount(
	tx *gorm.DB,
	zoneID uint,
) (int64, error) {


	var count int64


	err := tx.
		Model(&models.Reservation{}).
		Where(
			"zone_id = ? AND status = ?",
			zoneID,
			"active",
		).
		Count(&count).
		Error


	return count, err
}



func (r *reservationRepo) CreateReservation(
	tx *gorm.DB,
	reservation *models.Reservation,
) error {

	return tx.Create(reservation).Error
}



// =================================
// Normal Methods
// =================================


func (r *reservationRepo) GetUserReservations(
	userID uint,
) ([]models.Reservation, error) {


	var reservations []models.Reservation


	err := r.db.
		Preload("Zone").
		Where(
			"user_id = ?",
			userID,
		).
		Find(&reservations).
		Error


	if err != nil {
		return nil, err
	}


	return reservations, nil
}



func (r *reservationRepo) GetAllReservations(
) ([]models.Reservation, error) {


	var reservations []models.Reservation


	err := r.db.
		Preload("User").
		Preload("Zone").
		Find(&reservations).
		Error


	if err != nil {
		return nil, err
	}


	return reservations, nil
}



func (r *reservationRepo) GetReservationByID(
	id uint,
) (*models.Reservation, error) {


	var reservation models.Reservation


	err := r.db.
		First(
			&reservation,
			id,
		).
		Error


	if err != nil {
		return nil, err
	}


	return &reservation, nil
}



func (r *reservationRepo) UpdateReservation(
	reservation *models.Reservation,
) error {


	return r.db.Save(reservation).Error
}