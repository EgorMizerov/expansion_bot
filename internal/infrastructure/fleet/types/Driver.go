package types

type Driver struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	LicenseNumber string `json:"license_number"`
	WorkRuleID    string `json:"work_rule_id"`
}
