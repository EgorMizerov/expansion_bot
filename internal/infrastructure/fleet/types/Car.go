package types

import (
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type CreateCarRequest struct {
	Car *entity.Car
}

func (self CreateCarRequest) ToBody() *Car {
	return &Car{
		VehicleSpecifications: VehicleSpecifications{
			Model: self.Car.Model,
			Brand: self.Car.Brand,
			Color: self.Car.Color,
			Year:  int64(self.Car.Year),
			Vin:   self.Car.VIN,
		},
		VehicleLicenses: VehicleLicenses{
			LicencePlateNumber:      self.Car.LicensePlateNumber,
			RegistrationCertificate: self.Car.RegistrationCertificate,
		},
		ParkProfile: ParkProfile{
			Status:   "working",
			Callsign: self.Car.LicensePlateNumber,
		},
	}
}

type Car struct {
	VehicleSpecifications VehicleSpecifications `json:"vehicle_specifications"`
	VehicleLicenses       VehicleLicenses       `json:"vehicle_licenses"`
	ParkProfile           ParkProfile           `json:"park_profile"`
}

type ParkProfile struct {
	Status   string `json:"status"`
	Callsign string `json:"callsign"`
}

type VehicleLicenses struct {
	LicencePlateNumber      string `json:"licence_plate_number"`
	RegistrationCertificate string `json:"registration_certificate"`
}

type VehicleSpecifications struct {
	Model string `json:"model"`
	Brand string `json:"brand"`
	Color string `json:"color"`
	Year  int64  `json:"year"`
	Vin   string `json:"vin"`
}

type VehicleID struct {
	VehicleID string `json:"vehicle_id"`
}
