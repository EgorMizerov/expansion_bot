package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type DriverID uuid.UUID
type JumpID int
type TelegramID int
type FleetID string

type Driver struct {
	// ID's
	ID         DriverID    // Внутренний ID
	TelegramID *TelegramID // Привязку к телеграмму осуществляет администратор
	JumpID     JumpID
	FleetID    FleetID

	// Person Info
	FirstName      string
	MiddleName     *string
	LastName       string
	City           string
	IsSelfEmployed bool
	PhoneNuber     PhoneNumber
	DriverLicense  DriverLicense

	// Account Info
	WorkRule          *WorkRule
	WorkRuleUpdatedAt *time.Time
	AcceptCash        bool
	CarID             CarID

	CreatedAt time.Time
}

func NewDriver(jumpID JumpID, fleetID FleetID, firstName string, lastName string, middleName *string, city string, phoneNumber PhoneNumber, carID CarID, driverLicense DriverLicense) *Driver {
	return &Driver{
		ID:            DriverID(uuid.New()),
		TelegramID:    nil,
		JumpID:        jumpID,
		FleetID:       fleetID,
		FirstName:     firstName,
		MiddleName:    middleName,
		LastName:      lastName,
		City:          city,
		PhoneNuber:    phoneNumber,
		DriverLicense: driverLicense,
		AcceptCash:    false,
		CarID:         carID,
		CreatedAt:     time.Unix(time.Now().Unix(), 0),
	}
}

func (self *Driver) Fullname() string {
	if self.MiddleName != nil {
		return fmt.Sprintf("%s %s %s", self.LastName, self.FirstName, *self.MiddleName)
	}
	return fmt.Sprintf("%s %s", self.LastName, self.FirstName)
}

type PhoneNumber string

func (self PhoneNumber) String() string {
	return string(self)
}

type CardNumber string

func (self CardNumber) IsTinkoff() bool {
	bin := strings.TrimSpace(string(self))[0:6]
	if bin == "553691" || bin == "220070" {
		return true
	}
	return false
}
