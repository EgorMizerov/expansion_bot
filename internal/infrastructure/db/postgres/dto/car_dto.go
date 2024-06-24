package dto

import (
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/google/uuid"
)

type CarDTO struct {
	ID                      uuid.UUID `db:"id"`
	FleetID                 string    `db:"fleet_id"`
	Brand                   string    `db:"brand"`
	Model                   string    `db:"model"`
	Year                    int       `db:"year"`
	Color                   string    `db:"color"`
	VIN                     string    `db:"vin"`
	LicensePlateNumber      string    `db:"license_plate_number"`
	RegistrationCertificate string    `db:"registration_certificate"`
}

func (self CarDTO) ToCar() *entity.Car {
	return &entity.Car{
		ID:                      entity.CarID(self.ID),
		FleetID:                 entity.CarFleetID(self.FleetID),
		Brand:                   self.Brand,
		Model:                   self.Model,
		Year:                    self.Year,
		Color:                   self.Color,
		VIN:                     self.VIN,
		LicensePlateNumber:      self.LicensePlateNumber,
		RegistrationCertificate: self.RegistrationCertificate,
	}
}
