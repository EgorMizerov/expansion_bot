package entity

import (
	"time"

	"github.com/google/uuid"
)

type DriverLicenseID uuid.UUID

type DriverLicense struct {
	ID                      DriverLicenseID
	RegistrationCertificate string
	DrivingExperience       time.Time
	IssueDate               time.Time
	ExpiryDate              time.Time
	Country                 string
}

func NewDriverLicense(registrationCertificate string, drivingExperience time.Time, issueDate time.Time, expiryDate time.Time, country string) DriverLicense {
	return DriverLicense{ID: DriverLicenseID(uuid.New()), RegistrationCertificate: registrationCertificate, DrivingExperience: drivingExperience, IssueDate: issueDate, ExpiryDate: expiryDate, Country: country}
}
