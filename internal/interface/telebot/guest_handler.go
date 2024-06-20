package telebot

import (
	"fmt"

	tele "github.com/EgorMizerov/telebot"

	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/markup"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type GuestHandler struct{}

func NewGuestHandler(bot *Bot) *GuestHandler {
	guest := &GuestHandler{}

	bot.HandleStart(entity.GuestRole, guest.Start)

	return guest
}

func (self *GuestHandler) Start(ctx tele.Context) error {
	defer ctx.Delete()

	message := fmt.Sprintf("Добро пожаловать, %s!\nДля регистрации вам необходимо отправить ваш контакт.", ctx.Sender().FirstName)

	return ctx.EditOrSend(message, markup.SignUpMarkup(), tele.ModeMarkdown)
}
