package request

import (
	"strings"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type CreateDriverProfileRequest struct {
	Driver     *entity.Driver
	WorkRuleID string
	HireDate   time.Time
}

func (self CreateDriverProfileRequest) ToBody() DriverProfile {
	return DriverProfile{
		Account: struct {
			BalanceLimit string "json:\"balance_limit\""
			WorkRuleID   string "json:\"work_rule_id\""
		}{
			BalanceLimit: "0",
			WorkRuleID:   self.WorkRuleID,
		},
		OrderProvider: struct {
			Partner  bool "json:\"partner\""
			Platform bool "json:\"platform\""
		}{
			Partner:  true,
			Platform: true,
		},
		Person: struct {
			ContactInfo struct {
				Address string "json:\"address,omitempty\""
				Phone   string "json:\"phone\""
			} "json:\"contact_info\""
			DriverLicense struct {
				Country    string "json:\"country\""
				ExpiryDate string "json:\"expiry_date\""
				IssueDate  string "json:\"issue_date\""
				Number     string "json:\"number\""
			} "json:\"driver_license\""
			DriverLicenseExperience struct {
				TotalSinceDate string "json:\"total_since_date\""
			} "json:\"driver_license_experience\""
			FullName struct {
				FirstName  string "json:\"first_name\""
				LastName   string "json:\"last_name\""
				MiddleName string "json:\"middle_name\""
			} "json:\"full_name\""
		}{
			ContactInfo: struct {
				Address string `json:"address,omitempty"`
				Phone   string "json:\"phone\""
			}{
				Phone: string(self.Driver.PhoneNuber),
			},
			DriverLicense: struct {
				Country    string "json:\"country\""
				ExpiryDate string "json:\"expiry_date\""
				IssueDate  string "json:\"issue_date\""
				Number     string "json:\"number\""
			}{
				Country:    self.Driver.DriverLicense.Country,
				ExpiryDate: self.Driver.DriverLicense.ExpiryDate.Format("2006-01-02"),
				IssueDate:  self.Driver.DriverLicense.IssueDate.Format("2006-01-02"),
				Number:     strings.ReplaceAll(self.Driver.DriverLicense.RegistrationCertificate, " ", ""),
			},
			DriverLicenseExperience: struct {
				TotalSinceDate string "json:\"total_since_date\""
			}{
				TotalSinceDate: self.Driver.DriverLicense.DrivingExperience.Format("2006-01-02"),
			},
			FullName: struct {
				FirstName  string "json:\"first_name\""
				LastName   string "json:\"last_name\""
				MiddleName string "json:\"middle_name\""
			}{
				FirstName:  self.Driver.FirstName,
				LastName:   self.Driver.LastName,
				MiddleName: self.Driver.MiddleName,
			},
		},
		Profile: struct {
			HireDate string "json:\"hire_date\""
		}{
			HireDate: self.HireDate.Format("2006-01-02"),
		},
	}
}

type DriverProfile struct {
	Account struct {
		BalanceLimit string `json:"balance_limit"`
		WorkRuleID   string `json:"work_rule_id"`
	} `json:"account"`
	OrderProvider struct {
		Partner  bool `json:"partner"`
		Platform bool `json:"platform"`
	} `json:"order_provider"`
	Person struct {
		ContactInfo struct {
			Address string `json:"address,omitempty"`
			Phone   string `json:"phone"`
		} `json:"contact_info"`
		DriverLicense struct {
			Country    string `json:"country"`
			ExpiryDate string `json:"expiry_date"`
			IssueDate  string `json:"issue_date"`
			Number     string `json:"number"`
		} `json:"driver_license"`
		DriverLicenseExperience struct {
			TotalSinceDate string `json:"total_since_date"`
		} `json:"driver_license_experience"`
		FullName struct {
			FirstName  string `json:"first_name"`
			LastName   string `json:"last_name"`
			MiddleName string `json:"middle_name"`
		} `json:"full_name"`
	} `json:"person"`
	Profile struct {
		HireDate string `json:"hire_date"`
	} `json:"profile"`
}
