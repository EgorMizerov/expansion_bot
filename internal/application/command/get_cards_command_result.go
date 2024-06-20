package command

import "github.com/EgorMizerov/expansion_bot/internal/domain/entity"

type GetCardsCommandResult struct {
	Cards             []entity.CardNumber
	TinkoffCardsCount int
	AnotherCardsCount int
}
