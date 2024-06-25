package dto

import (
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DriverDTO struct {
	ID                uuid.UUID  `db:"id"`
	TelegramID        *int       `db:"telegram_id"`
	FleetID           string     `db:"fleet_id"`
	JumpID            int        `db:"jump_id"`
	FirstName         string     `db:"first_name"`
	LastName          string     `db:"last_name"`
	MiddleName        *string    `db:"middle_name"`
	City              string     `db:"city"`
	PhoneNumber       string     `db:"phone_number"`
	AcceptCash        bool       `db:"accept_cash"`
	WorkRuleID        *string    `db:"work_rule_id"`
	WorkRuleUpdatedAt *time.Time `db:"work_rule_updated_at"`
	IsSelfEmployed    *bool      `db:"is_self_employed"`
	CarID             uuid.UUID  `db:"car_id"`
	DriverLicenseID   uuid.UUID  `db:"driver_license_id"`
	CreatedAt         time.Time  `db:"created_at"`
}

func ToDriverDTO(driver *entity.Driver) DriverDTO {
	dto := DriverDTO{
		ID:                uuid.UUID(driver.ID),
		FleetID:           string(driver.FleetID),
		JumpID:            int(driver.JumpID),
		FirstName:         driver.FirstName,
		LastName:          driver.LastName,
		MiddleName:        driver.MiddleName,
		City:              driver.City,
		PhoneNumber:       string(driver.PhoneNuber),
		AcceptCash:        driver.AcceptCash,
		WorkRuleUpdatedAt: driver.WorkRuleUpdatedAt,
		IsSelfEmployed:    driver.IsSelfEmployed,
		CarID:             uuid.UUID(driver.CarID),
		DriverLicenseID:   uuid.UUID(driver.DriverLicense.ID),
		CreatedAt:         driver.CreatedAt,
	}
	if driver.TelegramID != nil {
		dto.TelegramID = common.Point(int(*driver.TelegramID))
	}
	if driver.WorkRule != nil {
		dto.WorkRuleID = &driver.WorkRule.ID
	}
	return dto
}

func (self DriverDTO) ToDriver(license *entity.DriverLicense) *entity.Driver {
	driver := &entity.Driver{
		ID:             entity.DriverID(self.ID),
		TelegramID:     nil,
		JumpID:         entity.JumpID(self.JumpID),
		FleetID:        entity.FleetID(self.FleetID),
		FirstName:      self.FirstName,
		MiddleName:     self.MiddleName,
		LastName:       self.LastName,
		City:           self.City,
		IsSelfEmployed: self.IsSelfEmployed,
		PhoneNuber:     entity.PhoneNumber(self.PhoneNumber),
		DriverLicense:  *license,
		AcceptCash:     self.AcceptCash,
		CarID:          entity.CarID(self.CarID),
		CreatedAt:      self.CreatedAt,
	}

	if !lo.IsNil(self.WorkRuleID) && !lo.IsNil(self.WorkRuleUpdatedAt) {
		driver.WorkRule = common.Point(entity.WorkRuleFromID(*self.WorkRuleID))
		driver.WorkRuleUpdatedAt = self.WorkRuleUpdatedAt
	}

	return driver
}
