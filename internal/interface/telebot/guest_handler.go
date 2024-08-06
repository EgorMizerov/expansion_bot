package telebot

import (
	"fmt"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/markup"
	tele "github.com/EgorMizerov/telebot"
	"github.com/pkg/errors"
)

type GuestHandler struct {
	guestService interfaces.GuestService
}

func NewGuestHandler(bot *Bot, guestRepository interfaces.GuestService) *GuestHandler {
	guest := &GuestHandler{
		guestService: guestRepository,
	}

	bot.HandleStart(entity.GuestRole, guest.Start)
	bot.Handle(tele.OnContact, guest.OnContact)

	return guest
}

func (self *GuestHandler) Start(ctx tele.Context) error {
	return ctx.Send(
		fmt.Sprintf(`Добро пожаловать, %s! Для того, чтобы начать пользоваться ботом вам необходимо отправить свой номер телефона и заполнить анкету.`, ctx.Sender().FirstName),
		markup.SendContactMarkup(),
	)
}

func (self *GuestHandler) OnContact(ctx tele.Context) error {
	err := self.guestService.CreateGuest(
		ctx,
		entity.TelegramID(ctx.Sender().ID),
		entity.PhoneNumber(ctx.Message().Contact.PhoneNumber),
	)
	if errors.Is(err, interfaces.ErrGuestAlreadyExists) {
		return ctx.Send("Мы получили ваш контакт. Теперь заполните анкету!", markup.SignUpMarkup())
	}
	if err != nil {
		return Error(ctx, err)
	}

	return ctx.Send("Мы получили ваш контакт. Теперь заполните анкету!", markup.SignUpMarkup())
}
