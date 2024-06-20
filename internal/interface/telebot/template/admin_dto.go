package template

import (
	"strconv"
	"strings"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"

	tele "github.com/EgorMizerov/telebot"
)

type DriverRegistrationData struct {
	UserID                  int
	ReferralKey             int64
	FullName                string
	PhoneNumber             string
	Address                 string
	DrivingExperience       time.Time
	RegistrationCertificate string
	LicenseCountry          string
	LicenseIssueDate        time.Time
	LicenseExpiryDate       time.Time
	IsSelfEmployed          bool
	CardNumber              CardNumber
	WorkRule                entity.WorkRule
	CarBrand                string
	CarModel                string
	CarColor                string
	CarYear                 time.Time
	CarVIN                  string
	LicensePlateNumber      string

	InputMessage string
	Message      *tele.Message
}

type CardNumber string

func (self CardNumber) String() string {
	runes := []rune(self)
	var builder strings.Builder
	for i := 1; i-1 < len(runes); i++ {
		builder.WriteRune(runes[i-1])
		if i%4 == 0 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func (self DriverRegistrationData) FormattedTime(t time.Time) string {
	defaultTime := time.Time{}
	if t == defaultTime {
		return ""
	}
	return t.Format(time.DateOnly)
}

func (self DriverRegistrationData) FormattedYear() string {
	defaultTime := time.Time{}
	if self.CarYear == defaultTime {
		return ""
	}
	return strconv.Itoa(self.CarYear.Year())
}

func (self DriverRegistrationData) FirstName() string {
	return strings.Fields(self.FullName)[1]
}

func (self DriverRegistrationData) LastName() string {
	return strings.Fields(self.FullName)[0]
}

func (self DriverRegistrationData) MiddleName() string {
	return strings.Fields(self.FullName)[2]
}

func (self DriverRegistrationData) ToDriver() *entity.Driver {
	var refferalKey *int64
	if self.ReferralKey != 0 {
		refferalKey = &self.ReferralKey
	}

	driverLicense := entity.NewDriverLicense(
		self.RegistrationCertificate,
		self.DrivingExperience,
		self.LicenseIssueDate,
		self.LicenseExpiryDate,
		self.LicenseCountry,
	)

	return entity.NewDriver(
		int64(self.UserID),
		self.FirstName(),
		self.LastName(),
		self.MiddleName(),
		entity.PhoneNumber(self.PhoneNumber),
		refferalKey,
		driverLicense,
		self.WorkRule,
		self.IsSelfEmployed,
		entity.CardNumber(self.CardNumber),
	)
}

func (self DriverRegistrationData) ToCar() *entity.Car {
	return entity.NewCar(self.CarBrand, self.CarModel, self.CarYear.Year(), self.CarColor, self.CarVIN, self.LicensePlateNumber, self.RegistrationCertificate)
}

type CardsInfoData struct {
	TinkoffCardsCount int
	AnotherCardsCount int
}
