package telebot

import (
	"fmt"

	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	tele "github.com/EgorMizerov/telebot"
	"github.com/pkg/errors"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/markup"
)

type GuestHandler struct {
	guestRepository repository.GuestRepository
}

func NewGuestHandler(bot *Bot, guestRepository repository.GuestRepository) *GuestHandler {
	guest := &GuestHandler{
		guestRepository: guestRepository,
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
	guest := entity.NewGuest(
		entity.TelegramID(ctx.Message().Contact.UserID),
		entity.PhoneNumber(ctx.Message().Contact.PhoneNumber),
	)
	err := self.guestRepository.CreateGuest(ctx, guest)
	if err != nil && !errors.Is(err, repository.ErrGuestAlreadyExists) {
		return Error(ctx, errors.Wrap(err, "failed to create guest"))
	}

	return ctx.Send("Мы получили ваш контакт. Теперь заполните анкету!", markup.SignUpMarkup())
}
