package repository

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type DriverRepository interface {
	CreateDriver(ctx context.Context, driver *entity.Driver) error
	GetDriverByID(ctx context.Context, driverID entity.DriverID) (*entity.Driver, error)
	GetDriverByPhoneNumber(ctx context.Context, phoneNumber entity.PhoneNumber) (*entity.Driver, error)
	GetDrivers(ctx context.Context) ([]*entity.Driver, error)
	UpdateDriver(crx context.Context, driver *entity.Driver) error
	GetCards(ctx context.Context) ([]entity.CardNumber, error)
}
