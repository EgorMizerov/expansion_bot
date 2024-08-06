package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/pkg/errors"
)

var (
	// ErrGuestAlreadyExists is returned when a guest already exists.
	ErrGuestAlreadyExists = errors.New("the guest already exists")
)

type GuestService interface {
	CreateGuest(ctx context.Context, telegramID entity.TelegramID, phone entity.PhoneNumber) error
}
