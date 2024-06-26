package entity

type State int

const (
	CreateDriver_ReceiveUserID State = iota + 1
	CreateDriver_ReceiveFullanme
	CreateDriver_ReceivePhoneNumber
	CreateDriver_ReceiveAddress
	CreateDriver_ReceiveIsSelfEmployed
	CreateDriver_ReceiveTariff
	CreateDriver_ReceiveCardNumber
	CreateDriver_ReceiveDrivingExperience
	CreateDriver_ReceiveRegistrationCertificate
	CreateDriver_ReceiveLicenseCountry
	CreateDriver_ReceiveLicenseIssueDate
	CreateDriver_ReceiveLicenseExpiryDate
	CreateDriver_ReceiveCarBrand
	CreateDriver_ReceiveCarModel
	CreateDriver_ReceiveCarColor
	CreateDriver_ReceiveCarYear
	CreateDriver_ReceiveCarVIN
	CreateDriver_ReceiveLicensePlateNumber
	CreateDriver_ReceiveReferralKey

	ChangePhoneNumber_GetNumber
)
