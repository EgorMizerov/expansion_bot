package dto

import "time"

type DriverDTO struct {
	ID                             string    `db:"id"`
	TelegramID                     int       `db:"telegram_id"`
	FirstName                      string    `db:"first_name"`
	LastName                       string    `db:"last_name"`
	MiddleName                     string    `db:"middle_name"`
	Address                        string    `db:"address"`
	PhoneNumber                    string    `db:"phone_number"`
	CardNumber                     string    `db:"card_number"`
	ReferralKey                    *int      `db:"referral_key"`
	AcceptCash                     bool      `db:"accept_cash"`
	WorkRuleID                     string    `db:"work_rule_id"`
	WorkRuleUpdatedAt              time.Time `db:"work_rule_updated_at"`
	IsSelfEmployed                 bool      `db:"is_self_employed"`
	CarID                          string    `db:"car_id"`
	CreatedAt                      time.Time `db:"created_at"`
	LicenseDriverID                string    `db:"driver_id"`
	LicenseRegistrationCertificate string    `db:"registration_certificate"`
	LicenseDrivingExperience       time.Time `db:"driving_experience"`
	LicenseIssueDate               time.Time `db:"issue_date"`
	LicenseExpiryDate              time.Time `db:"expiry_date"`
	LicenseCountry                 string    `db:"country"`
}
