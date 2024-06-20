package types

type Contractor struct {
	Account       Account       `json:"account"`
	CarID         string        `json:"car_id,omitempty"`
	OrderProvider OrderProvider `json:"order_provider"`
	Person        Person        `json:"person"`
	Profile       Profile       `json:"profile"`
}

type Account struct {
	BalanceLimit                   string `json:"balance_limit"`
	BlockOrdersOnBalanceBelowLimit bool   `json:"block_orders_on_balance_below_limit"`
	PaymentServiceID               string `json:"payment_service_id"`
	WorkRuleID                     string `json:"work_rule_id"`
}

type OrderProvider struct {
	Partner  bool `json:"partner"`
	Platform bool `json:"platform"`
}

type Person struct {
	ContactInfo             ContactInfo             `json:"contact_info"`
	DriverLicense           DriverLicense           `json:"driver_license"`
	DriverLicenseExperience DriverLicenseExperience `json:"driver_license_experience"`
	FullName                FullName                `json:"full_name"`
	IsSelfemployed          bool                    `json:"is_selfemployed"`
	TaxIdentificationNumber string                  `json:"tax_identification_number,omitempty"`
}

type ContactInfo struct {
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone"`
}

type DriverLicense struct {
	BirthDate  string `json:"birth_date,omitempty"`
	Country    string `json:"country"`
	ExpiryDate string `json:"expiry_date"`
	IssueDate  string `json:"issue_date"`
	Number     string `json:"number"`
}

type DriverLicenseExperience struct {
	TotalSinceDate string `json:"total_since_date"`
}

type FullName struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
}

type Profile struct {
	Comment    string `json:"comment,omitempty"`
	Feedback   string `json:"feedback,omitempty"`
	FireDate   string `json:"fire_date,omitempty"`
	HireDate   string `json:"hire_date"`
	WorkStatus string `json:"work_status"`
}
