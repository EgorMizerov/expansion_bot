package dto

import (
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/google/uuid"
)

type DriverLicenseDTO struct {
	ID                      uuid.UUID `db:"id"`
	RegistrationCertificate string    `db:"registration_certificate"`
	DrivingExperience       time.Time `db:"driving_experience"`
	IssueDate               time.Time `db:"issue_date"`
	ExpiryDate              time.Time `db:"expiry_date"`
	Country                 string    `db:"country"`
}

func ToDriverLicenseDTO(license entity.DriverLicense) DriverLicenseDTO {
	return DriverLicenseDTO{
		ID:                      uuid.UUID(license.ID),
		RegistrationCertificate: license.RegistrationCertificate,
		DrivingExperience:       license.DrivingExperience,
		IssueDate:               license.IssueDate,
		ExpiryDate:              license.ExpiryDate,
		Country:                 license.Country,
	}
}
