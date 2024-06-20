package telebot

import (
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"

	tele "github.com/EgorMizerov/telebot"
)

type Bot struct {
	*tele.Bot
	stateHandlers map[entity.State]func(ctx tele.Context) error
	startHandlers map[entity.Role]func(ctx tele.Context) error
}

func NewBot(teleBot *tele.Bot, fsm FSM, adminID int64) *Bot {
	bot := &Bot{
		Bot:           teleBot,
		stateHandlers: make(map[entity.State]func(ctx tele.Context) error),
		startHandlers: make(map[entity.Role]func(ctx tele.Context) error),
	}

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		state, err := fsm.GetState(ctx.Sender().ID)
		if err != nil {
			return nil
		}

		handler, ok := bot.stateHandlers[state]
		if ok {
			return handler(ctx)
		}
		return nil
	})

	bot.Handle("/start", func(ctx tele.Context) error {
		if ctx.Sender().ID == adminID {
			return bot.startHandlers[entity.AdminRole](ctx)
		} else {
			return bot.startHandlers[entity.GuestRole](ctx)
		}
		return nil
	})

	return bot
}

func (self *Bot) HandleState(state entity.State, fn func(ctx tele.Context) error) {
	self.stateHandlers[state] = fn
}

func (self *Bot) HandleStart(role entity.Role, fn func(ctx tele.Context) error) {
	self.startHandlers[role] = fn
}
