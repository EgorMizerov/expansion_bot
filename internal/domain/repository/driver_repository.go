package repository

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type DriverRepository interface {
	CreateDriver(ctx context.Context, driver *entity.Driver) error
	GetDrivers(ctx context.Context) ([]*entity.Driver, error)
	GetCards(ctx context.Context) ([]entity.CardNumber, error)
}
