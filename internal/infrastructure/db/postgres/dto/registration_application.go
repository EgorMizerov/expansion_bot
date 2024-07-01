package dto

import (
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type RegistrationApplication struct {
	ID     int       `db:"id"`
	Status string    `db:"status"`
	Date   time.Time `db:"date"`

	PhoneNumber *string `db:"phone_number"`
	LastName    *string `db:"last_name"`
	FirstName   *string `db:"first_name"`
	MiddleName  *string `db:"middle_name"`
	City        *string `db:"city"`

	LicenseNumber         *string    `db:"license_number"`
	LicenseTotalSinceDate *time.Time `db:"license_total_since_date"`
	LicenseIssueDate      *time.Time `db:"license_issue_date"`
	LicenseExpiryDate     *time.Time `db:"license_expiry_date"`
	LicenseCountry        *string    `db:"license_country"`

	CarBrand   *string `db:"car_brand"`
	CarModel   *string `db:"car_model"`
	CarYear    *int    `db:"car_year"`
	CarColor   *string `db:"car_color"`
	CarVIN     *string `db:"car_vin"`
	CarNumber  *string `db:"car_number"`
	CarLicense *string `db:"car_license"`
	WorkRuleID *string `db:"work_rule_id"`
}

func ToRegistrationApplication(application *entity.RegistrationApplication) *RegistrationApplication {
	if application == nil {
		return nil
	}
	var workRuleID *string
	if application.WorkRule != nil {
		workRuleID = common.Point(application.WorkRule.ID)
	}

	return &RegistrationApplication{
		ID:                    int(application.ID),
		Status:                application.Status,
		Date:                  application.Date,
		PhoneNumber:           application.PhoneNumber,
		LastName:              application.LastName,
		FirstName:             application.FirstName,
		MiddleName:            application.MiddleName,
		City:                  application.City,
		LicenseNumber:         application.LicenseNumber,
		LicenseTotalSinceDate: application.LicenseTotalSinceDate,
		LicenseIssueDate:      application.LicenseIssueDate,
		LicenseExpiryDate:     application.LicenseExpiryDate,
		LicenseCountry:        application.LicenseCountry,
		CarBrand:              application.CarBrand,
		CarModel:              application.CarModel,
		CarYear:               application.CarYear,
		CarColor:              application.CarColor,
		CarVIN:                application.CarVIN,
		CarNumber:             application.CarNumber,
		CarLicense:            application.CarLicense,
		WorkRuleID:            workRuleID,
	}
}

func (self RegistrationApplication) ToRegistrationApplication() *entity.RegistrationApplication {
	var workRule *entity.WorkRule
	if self.WorkRuleID != nil {
		workRule = common.Point(entity.WorkRuleFromID(*self.WorkRuleID))
	}
	return &entity.RegistrationApplication{
		ID:                    entity.RegistrationApplicationID(self.ID),
		Status:                self.Status,
		Date:                  self.Date,
		PhoneNumber:           self.PhoneNumber,
		LastName:              self.LastName,
		FirstName:             self.FirstName,
		MiddleName:            self.MiddleName,
		City:                  self.City,
		LicenseNumber:         self.LicenseNumber,
		LicenseTotalSinceDate: self.LicenseTotalSinceDate,
		LicenseIssueDate:      self.LicenseIssueDate,
		LicenseExpiryDate:     self.LicenseExpiryDate,
		LicenseCountry:        self.LicenseCountry,
		CarBrand:              self.CarBrand,
		CarModel:              self.CarModel,
		CarYear:               self.CarYear,
		CarColor:              self.CarColor,
		CarVIN:                self.CarVIN,
		CarNumber:             self.CarNumber,
		CarLicense:            self.CarLicense,
		WorkRule:              workRule,
	}
}
