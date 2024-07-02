package template

import (
	"strings"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type DriverRegistrationData struct {
	FullName          string
	PhoneNumber       string
	Address           string
	WorkRule          entity.WorkRule
	DrivingExperience time.Time
	LicenseNumber     string
	LicenseCountry    string
	LicenseIssueDate  time.Time
	LicenseExpiryDate time.Time
	IsSelfEmployed    bool
	CardNumber        CardNumber
	CarBrand          string
	CarModel          string
	CarColor          string
	CarYear           int
	CarVIN            string
	CarNumber         string
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

func (self DriverRegistrationData) FirstName() string {
	return strings.Fields(self.FullName)[1]
}

func (self DriverRegistrationData) LastName() string {
	return strings.Fields(self.FullName)[0]
}

func (self DriverRegistrationData) MiddleName() string {
	return strings.Fields(self.FullName)[2]
}

type CardsInfoData struct {
	TinkoffCardsCount int
	AnotherCardsCount int
}

type RegistrationApplicationsData struct {
	Items []RegistrationApplication
}

type RegistrationApplication struct {
	ID       int
	Status   string
	Date     time.Time
	Fullname string
	Link     string
}

func (self RegistrationApplication) FormattedTime(t time.Time) string {
	defaultTime := time.Time{}
	if t == defaultTime {
		return ""
	}
	return t.Format(time.DateOnly)
}

type DriversListData struct {
	Items []DriversListItem
}

type DriversListItem struct {
	PhoneNumber string
	Fullname    string
}

type DriverInfoData struct {
	ID             string
	Fullname       string
	PhoneNumber    string
	City           string
	IsSelfEmployed bool
	WorkRule       *string
}

func (self DriverInfoData) Nullable(arg *string) string {
	if arg == nil {
		return "-"
	}
	return *arg
}

type DriversCarInfoData struct {
	Brand  string
	Model  string
	Color  string
	Year   int
	VIN    string
	Number string
}

type DriversLicenseInfoData struct {
	LicenseNumber            string
	LicenseIssueDate         time.Time
	LicenseExpiryDate        time.Time
	LicenseDrivingExperience int
	LicenseCountry           string
}

func (self DriversLicenseInfoData) FormattedTime(t time.Time) string {
	return t.Format(time.DateOnly)
}
