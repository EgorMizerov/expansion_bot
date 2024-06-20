package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/application/command"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type AdminService interface {
	CreateDriver(ctx context.Context, driver *entity.Driver, car *entity.Car) error
	GetCards(ctx context.Context) (*command.GetCardsCommandResult, error)
}
