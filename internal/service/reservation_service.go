package service

import (
	"errors"

	"gorm.io/gorm"

	"spotsync/internal/models"
	"spotsync/internal/repository"
)

type ReservationService struct {
	repo repository.ReservationRepository
	db   *gorm.DB
}

func NewReservationService(
	repo repository.ReservationRepository,
	db *gorm.DB,
) *ReservationService {

	return &ReservationService{
		repo: repo,
		db:   db,
	}
}


// Create Reservation
func (s *ReservationService) CreateReservation(
	zoneID uint,
	userID uint,
	licensePlate string,
) (*models.Reservation, error) {

	var reservation *models.Reservation

	err := s.db.Transaction(func(tx *gorm.DB) error {

		zone, err := s.repo.LockZone(tx, zoneID)

		if err != nil {
			return err
		}


		count, err := s.repo.GetActiveCount(tx, zoneID)

		if err != nil {
			return err
		}


		// Capacity Check
		if count >= int64(zone.TotalCapacity) {
			return errors.New("parking zone is full")
		}


		reservation = &models.Reservation{
			UserID:       userID,
			ZoneID:       zoneID,
			LicensePlate: licensePlate,
			Status:       "active",
		}


		if err := s.repo.CreateReservation(tx, reservation); err != nil {
			return err
		}


		return nil
	})


	if err != nil {
		return nil, err
	}


	return reservation, nil
}



// Get Logged User Reservations
func (s *ReservationService) GetMyReservations(
	userID uint,
) ([]models.Reservation, error) {

	return s.repo.GetUserReservations(userID)
}



// Cancel Reservation
func (s *ReservationService) CancelReservation(
	reservationID uint,
	userID uint,
) error {


reservation, err := s.repo.GetReservationByID(reservationID)

	if err != nil {
		return err
	}


	// Ownership check
	if reservation.UserID != userID {
		return errors.New("unauthorized")
	}


	reservation.Status = "cancelled"


	return s.repo.UpdateReservation(reservation)
}



// Get All Reservations (Admin)
func (s *ReservationService) GetAllReservations(
) ([]models.Reservation, error) {

	return s.repo.GetAllReservations()
}