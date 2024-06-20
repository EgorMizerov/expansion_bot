package entity

import "time"

type Application struct {
	Action    ApplicationAction `json:"action"`
	Item      ApplicationItem   `json:"item"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type ApplicationItem struct {
	ID                     int                      `json:"id"`
	Phone                  *string                  `json:"phone,omitempty"`
	Direction              *ApplicationDirection    `json:"direction,omitempty"`
	Profession             *ApplicationProfession   `json:"profession,omitempty"`
	PersonInfo             *ApplicationPersonInfo   `json:"person_info,omitempty"`
	IsCarDriver            *bool                    `json:"is_car_driver,omitempty"`
	Date                   *time.Time               `json:"date,omitempty"`
	DriverLicense          *DriverLicense           `json:"driver_license,omitempty"`
	Car                    *Car                     `json:"car,omitempty"`
	Status                 *ApplicationStatus       `json:"status,omitempty"`
	ApplicationIntegration []ApplicationIntegration `json:"application_integration,omitempty"`
}

type ApplicationPersonInfo struct {
	LastName          string  `json:"last_name"`
	FirstName         string  `json:"first_name"`
	MiddleName        *string `json:"middle_name,omitempty"`
	BirthDate         *string `json:"birth_date,omitempty"`
	City              *string `json:"city,omitempty"`
	DriverLicenseTaxi *string `json:"driver_license_taxi,omitempty"`
	Comment           *string `json:"comment,omitempty"`
	//TODO: Files
}

type ApplicationStatus struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type ApplicationProfession struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ApplicationDirection struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type ApplicationDriverLicense struct {
	Country             ApplicationCountry `json:"country"`
	Number              string             `json:"number"`
	IssueDate           string             `json:"issue_date"`            // YYYY-MM-DD
	ExpiryDate          *string            `json:"expiry_date,omitempty"` // YYYY-MM-DD
	ExpiryDateUnlimited *bool              `json:"expiry_date_unlimited,omitempty"`
	TotalSinceDate      string             `json:"total_since_date"` // YYYY-MM-DD
	//TODO: Files
}

type ApplicationCountry struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

type ApplicationCar struct {
	Model   IDName  `json:"model"`
	Color   IDName  `json:"color"`
	Year    int     `json:"year"`
	Number  string  `json:"number"`
	License string  `json:"license"`
	VIN     *string `json:"vin,omitempty"`
	NoVIN   *bool   `json:"no_vin,omitempty"`
	// TODO: files
}

type ApplicationIntegration struct {
	IntegratorID int    `json:"integrator_id"`
	Name         string `json:"name"`
}

type ApplicationApplicationIntegration struct {
	ID            int `json:"id"`
	IntegrationID int `json:"integration_id"`
	//Manager       *IDName `json:"manager,omitempty"`
	Name    string            `json:"name"`
	Status  ApplicationStatus `json:"status"`
	Comment *string           `json:"comment,omitempty"`
	Group   IDName            `json:"group"`
}

type IDName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ApplicationAction string

func (self ApplicationAction) Is(action ApplicationAction) bool {
	return self == action
}

const (
	NewApplicationAction    ApplicationAction = "new_application"
	ChangeApplicationAction ApplicationAction = "change_application"
	NewIntegrationAction    ApplicationAction = "new_integration"
	ChangeIntegrationAction ApplicationAction = "change_integration"
)
