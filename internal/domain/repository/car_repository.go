package repository

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type CarRepository interface {
	CreateCar(ctx context.Context, car *entity.Car) error
}
