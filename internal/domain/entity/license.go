package entity

import "time"

type DriverLicense struct {
	RegistrationCertificate string
	DrivingExperience       time.Time
	IssueDate               time.Time
	ExpiryDate              time.Time
	Country                 string
}

func NewDriverLicense(registrationCertificate string, drivingExperience time.Time, issueDate time.Time, expiryDate time.Time, country string) DriverLicense {
	return DriverLicense{RegistrationCertificate: registrationCertificate, DrivingExperience: drivingExperience, IssueDate: issueDate, ExpiryDate: expiryDate, Country: country}
}
