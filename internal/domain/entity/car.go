package entity

import "github.com/google/uuid"

type CarID uuid.UUID
type CarFleetID string

type Car struct {
	ID                      CarID
	FleetID                 CarFleetID
	Brand                   string
	Model                   string
	Year                    int
	Color                   string
	VIN                     string
	LicensePlateNumber      string
	RegistrationCertificate string
}

func NewCar(fleetID CarFleetID, brand string, model string, year int, color string, VIN string, licensePlateNumber string, registrationCertificate string) *Car {
	return &Car{
		ID:                      CarID(uuid.New()),
		FleetID:                 fleetID,
		Brand:                   brand,
		Model:                   model,
		Year:                    year,
		Color:                   color,
		VIN:                     VIN,
		LicensePlateNumber:      licensePlateNumber,
		RegistrationCertificate: registrationCertificate,
	}
}

func (self *Car) SetID(id CarID) {
	self.ID = id
}
