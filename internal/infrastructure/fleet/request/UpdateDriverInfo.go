package request

type UpdateDriverBody struct {
	DriverID      string               `json:"driver_id"`
	Accounts      UpdateDriverAccounts `json:"accounts"`
	DriverProfile UpdateDriverProfile  `json:"driver_profile"`
}

type UpdateDriverAccounts struct {
	BalanceLimit string `json:"balance_limit"`
}

type UpdateDriverProfile struct {
	Address                  string                               `json:"address"`
	AutomaticReceiptCreation bool                                 `json:"automatic_receipt_creation"`
	BalanceDenyOnlyCard      bool                                 `json:"balance_deny_onlycard"`
	BankAccounts             []any                                `json:"bank_accounts"`
	CarID                    string                               `json:"car_id"`
	Deaf                     bool                                 `json:"deaf"`
	EmergencyPersonContacts  []any                                `json:"emergency_person_contacts"`
	FirstName                string                               `json:"first_name"`
	LastName                 string                               `json:"last_name"`
	MiddleName               string                               `json:"middle_name"`
	HiddenFields             []any                                `json:"hidden_fields"`
	HireDate                 string                               `json:"hire_date"`
	Identifications          []any                                `json:"identifications"`
	WorkRuleID               string                               `json:"work_rule_id"`
	WorkStatus               string                               `json:"work_status"`
	Providers                []string                             `json:"providers"`
	Phones                   []string                             `json:"phones"`
	PaymentServiceID         string                               `json:"payment_service_id"`
	LicenseExperience        UpdateDriverProfileLicenseExperience `json:"license_experience"`
	DriverLicense            UpdateDriverProfileDriverLicense     `json:"driver_license"`
}

type UpdateDriverProfileLicenseExperience struct {
	TotalSince string `json:"total_since"`
}

type UpdateDriverProfileDriverLicense struct {
	BirthDate      string `json:"birth_date"`
	Country        string `json:"country"`
	ExpirationDate string `json:"expiration_date"`
	IssueDate      string `json:"issue_date"`
	Number         string `json:"number"`
}
