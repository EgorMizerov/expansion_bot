package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/pkg/errors"
)

var (
	// ErrDriverNotFound is returned when a driver is not found.
	ErrDriverNotFound = errors.New("the driver was not found")
)

//go:generate mockery --name DriverService
type DriverService interface {
	GetDriverByPhoneNumber(ctx context.Context, phone entity.PhoneNumber) (*entity.Driver, error)
	GetDrivers(ctx context.Context) ([]*entity.Driver, error)
}
