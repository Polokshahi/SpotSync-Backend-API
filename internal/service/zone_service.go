package service

import (
	"spotsync/internal/dto"
	"spotsync/internal/models"
	"spotsync/internal/repository"
)

type ZoneService struct {
	repo repository.ZoneRepository
}

func NewZoneService(repo repository.ZoneRepository) *ZoneService {
	return &ZoneService{
		repo: repo,
	}
}

func (s *ZoneService) CreateZone(req dto.CreateZoneRequest) (*models.ParkingZone, error) {

	zone := &models.ParkingZone{
		Name: req.Name,
		Type: req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour: req.PricePerHour,
	}

	err := s.repo.Create(zone)

	if err != nil {
		return nil, err
	}

	return zone, nil
}

func (s *ZoneService) GetAllZones() ([]models.ParkingZone, error) {
	return s.repo.GetAll()
}

func (s *ZoneService) GetZoneByID(id uint) (*models.ParkingZone, error) {
	return s.repo.GetByID(id)
}