package domains

import (
	"housinganywhere/dns/app/models"
)

type Drone struct {
	SectorID float64 `json:"sid"`
}

func CreateDrone(sectorId float64) Drone {
	return Drone{SectorID: sectorId}
}

func (d *Drone) GetLocation(location models.Location) (float64, error) {
	//calculate location based on drone
	calculation := location.X*d.SectorID + location.Y*d.SectorID + location.Z*d.SectorID + location.Vel
	return calculation, nil
}
