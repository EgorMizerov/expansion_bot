package entity

type Guest struct {
	TelegramID  TelegramID
	PhoneNumber PhoneNumber
}

func NewGuest(telegramID TelegramID, phoneNumber PhoneNumber) *Guest {
	return &Guest{TelegramID: telegramID, PhoneNumber: phoneNumber}
}
