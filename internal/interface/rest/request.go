package rest

import "time"

type Request struct {
	Action    RequestAction `json:"action"`
	Item      RequestItem   `json:"item"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type RequestItem struct {
	ID                     int            `json:"id"`
	Phone                  *string        `json:"phone,omitempty"`
	Direction              *Direction     `json:"direction,omitempty"`
	Profession             *Profession    `json:"profession,omitempty"`
	PersonInfo             *PersonInfo    `json:"person_info,omitempty"`
	IsCarDriver            *bool          `json:"is_car_driver,omitempty"`
	Date                   *time.Time     `json:"date,omitempty"`
	DriverLicense          *DriverLicense `json:"driver_license,omitempty"`
	Car                    *Car           `json:"car,omitempty"`
	Status                 *Status        `json:"status,omitempty"`
	ApplicationIntegration []Integration  `json:"application_integration,omitempty"`
}

type PersonInfo struct {
	LastName          string  `json:"last_name"`
	FirstName         string  `json:"first_name"`
	MiddleName        *string `json:"middle_name,omitempty"`
	BirthDate         *string `json:"birth_date,omitempty"`
	City              *string `json:"city,omitempty"`
	DriverLicenseTaxi *string `json:"driver_license_taxi,omitempty"`
	Comment           *string `json:"comment,omitempty"`
	//TODO: Files
}

type Status struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type Profession struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Direction struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type DriverLicense struct {
	Country             Country `json:"country"`
	Number              string  `json:"number"`
	IssueDate           string  `json:"issue_date"`            // YYYY-MM-DD
	ExpiryDate          *string `json:"expiry_date,omitempty"` // YYYY-MM-DD
	ExpiryDateUnlimited *bool   `json:"expiry_date_unlimited,omitempty"`
	TotalSinceDate      string  `json:"total_since_date"` // YYYY-MM-DD
	//TODO: Files
}

type Country struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

type Car struct {
	Model   IDName  `json:"model"`
	Color   IDName  `json:"color"`
	Year    int     `json:"year"`
	Number  string  `json:"number"`
	License string  `json:"license"`
	VIN     *string `json:"vin,omitempty"`
	NoVIN   *bool   `json:"no_vin,omitempty"`
	// TODO: files
}

type Integration struct {
	IntegratorID int    `json:"integrator_id"`
	Name         string `json:"name"`
}

type ApplicationIntegration struct {
	ID            int `json:"id"`
	IntegrationID int `json:"integration_id"`
	//Manager       *IDName `json:"manager,omitempty"`
	Name    string  `json:"name"`
	Status  Status  `json:"status"`
	Comment *string `json:"comment,omitempty"`
	Group   IDName  `json:"group"`
}

type IDName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RequestAction string

func (self RequestAction) Is(action RequestAction) bool {
	return self == action
}

const (
	NewApplicationAction    RequestAction = "new_application"
	ChangeApplicationAction RequestAction = "change_application"
	NewIntegrationAction    RequestAction = "new_integration"
	ChangeIntegrationAction RequestAction = "change_integration"
)
