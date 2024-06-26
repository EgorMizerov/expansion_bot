package dto

import "github.com/EgorMizerov/expansion_bot/internal/domain/entity"

type GuestDTO struct {
	PhoneNumber string `db:"phone_number"`
	TelegramID  int    `db:"telegram_id"`
}

func ToGuestDTO(guest *entity.Guest) *GuestDTO {
	return &GuestDTO{
		PhoneNumber: guest.PhoneNumber.String(),
		TelegramID:  int(guest.TelegramID),
	}
}

func (self *GuestDTO) ToEntity() *entity.Guest {
	return &entity.Guest{
		TelegramID:  entity.TelegramID(self.TelegramID),
		PhoneNumber: entity.PhoneNumber(self.PhoneNumber),
	}
}
