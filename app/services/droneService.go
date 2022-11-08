package services

import (
	"housinganywhere/dns/app/domains"
	"housinganywhere/dns/app/models"
)

type DroneService interface {
	GetLocation(request *models.LocationRequest) (models.LocationResponse, error)
}
type droneService struct {
	SectorID float64
}

func NewDroneService(sectorId float64) DroneService {
	return &droneService{SectorID: sectorId}
}

func (s *droneService) GetLocation(request *models.LocationRequest) (models.LocationResponse, error) {

	location, err := models.CreateLocation(request)
	if err != nil {
		return models.LocationResponse{}, err
	}

	drone := domains.CreateDrone(s.SectorID)
	val, err := drone.GetLocation(location)
	if err != nil {
		return models.LocationResponse{}, err
	}

	return models.CreateLocationResponse(val), nil

}
