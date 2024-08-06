package services

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	"github.com/pkg/errors"
)

type GuestService struct {
	guestRepository repository.GuestRepository
}

func NewGuestService(guestRepository repository.GuestRepository) *GuestService {
	return &GuestService{guestRepository: guestRepository}
}

func (self *GuestService) CreateGuest(ctx context.Context, telegramID entity.TelegramID, phone entity.PhoneNumber) error {
	guest := entity.NewGuest(telegramID, phone)

	err := self.guestRepository.CreateGuest(ctx, guest)
	if err != nil {
		if errors.Is(err, repository.ErrGuestAlreadyExists) {
			return interfaces.ErrGuestAlreadyExists
		}
		return errors.Wrap(err, "failed to create guest")
	}

	return nil
}
