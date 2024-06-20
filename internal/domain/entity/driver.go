package entity

import (
	"fmt"
	"strings"
	"time"
)

type DriverID string

type Driver struct {
	ID                DriverID
	TelegramID        int64
	FirstName         string
	MiddleName        string
	LastName          string
	Address           string
	PhoneNuber        PhoneNumber
	CardNumber        CardNumber
	ReferralKey       *int64
	AcceptCash        bool
	DriverLicense     DriverLicense
	WorkRule          WorkRule
	WorkRuleUpdatedAt time.Time
	IsSelfEmployed    bool
	CarID             CarID
	CreatedAt         time.Time
}

func NewDriver(telegramID int64, firstName string, lastName string, middleName string, phoneNumber PhoneNumber, referralKey *int64, license DriverLicense, workRule WorkRule, isSelfEmployed bool, cardNumber CardNumber) *Driver {
	return &Driver{
		TelegramID:        telegramID,
		FirstName:         firstName,
		MiddleName:        middleName,
		LastName:          lastName,
		PhoneNuber:        phoneNumber,
		ReferralKey:       referralKey,
		AcceptCash:        false,
		DriverLicense:     license,
		CreatedAt:         time.Now(),
		WorkRule:          workRule,
		WorkRuleUpdatedAt: time.Now(),
		IsSelfEmployed:    isSelfEmployed,
		CardNumber:        cardNumber,
	}
}

func (self *Driver) SetDriverID(id DriverID) {
	self.ID = id
}

func (self *Driver) SetCarID(id CarID) {
	self.CarID = id
}

func (self *Driver) Fullname() string {
	return fmt.Sprintf("%s %s %s", self.LastName, self.FirstName, self.MiddleName)
}

type PhoneNumber string

type CardNumber string

func (self CardNumber) IsTinkoff() bool {
	bin := strings.TrimSpace(string(self))[0:6]
	if bin == "553691" || bin == "220070" {
		return true
	}
	return false
}
