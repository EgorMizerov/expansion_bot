package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/application/command"
)

type AdminService interface {
	GetCards(ctx context.Context) (*command.GetCardsCommandResult, error)
}
