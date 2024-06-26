package repository

import (
	"context"
	"errors"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

var (
	// ErrGuestAlreadyExists is returned when trying to create a guest that already exists.
	ErrGuestAlreadyExists = errors.New("the guest already exists")
	// ErrGuestNotFound is returned when a guest is not found.
	ErrGuestNotFound = errors.New("the guest was not found")
)

type GuestRepository interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	GetGuestByPhoneNumber(ctx context.Context, number entity.PhoneNumber) (*entity.Guest, error)
}
