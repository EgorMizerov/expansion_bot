package entity

import (
	"fmt"
	"time"
)

type RegistrationApplicationID int

type RegistrationApplication struct {
	ID     RegistrationApplicationID
	Status string
	Date   time.Time

	PhoneNumber *string
	LastName    *string
	FirstName   *string
	MiddleName  *string
	City        *string

	LicenseNumber         *string
	LicenseTotalSinceDate *time.Time
	LicenseIssueDate      *time.Time
	LicenseExpiryDate     *time.Time
	LicenseCountry        *string

	CarBrand   *string
	CarModel   *string
	CarYear    *int
	CarColor   *string
	CarVIN     *string
	CarNumber  *string
	CarLicense *string

	WorkRule *WorkRule
}

func NewRegistrationApplication(id int, date time.Time) *RegistrationApplication {
	return &RegistrationApplication{
		ID:     RegistrationApplicationID(id),
		Status: "draft",
		Date:   date,
	}
}

func (self *RegistrationApplication) Fullname() string {
	if self.LastName == nil || self.FirstName == nil {
		return "имя не указано"
	}
	if self.MiddleName != nil {
		return fmt.Sprintf("%s %s %s", *self.LastName, *self.FirstName, *self.MiddleName)
	}
	return fmt.Sprintf("%s %s", *self.LastName, *self.FirstName)
}

func (self *RegistrationApplication) Link() string {
	return fmt.Sprintf("https://my.jump.taxi/autoregistration/list/%d", self.ID)
}

func (self *RegistrationApplication) SetPhone(phone string) {
	self.PhoneNumber = &phone
}

func (self *RegistrationApplication) SetPersonInfo(lastName string, firstName string, middleName *string, city string) {
	self.LastName = &lastName
	self.FirstName = &firstName
	self.MiddleName = middleName
	self.City = &city
}

func (self *RegistrationApplication) SetDriverLicense(country string, number string, issueDate time.Time, expiryDate time.Time, totalSinceDate time.Time) {
	self.LicenseNumber = &number
	self.LicenseTotalSinceDate = &totalSinceDate
	self.LicenseIssueDate = &issueDate
	self.LicenseExpiryDate = &expiryDate
	self.LicenseCountry = &country
}

func (self *RegistrationApplication) SetCar(year int, number string, license string, vin string, brand string, model string, color string) {
	self.CarBrand = &brand
	self.CarModel = &model
	self.CarYear = &year
	self.CarColor = &color
	self.CarVIN = &vin
	self.CarNumber = &number
	self.CarLicense = &license
}

func (self *RegistrationApplication) SetStatus(status string) {
	self.Status = status
}

func (self *RegistrationApplication) SetWorkRule(rule WorkRule) {
	self.WorkRule = &rule
}
