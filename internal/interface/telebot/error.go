package telebot

import (
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/context"
	tele "github.com/EgorMizerov/telebot"
)

func CallbackError(ctx tele.Context, err error) error {
	context.Logger(ctx).Error(err.Error())
	return ctx.RespondAlert("Упс! Что-то пошло не так. Попробуйте повторить действие или обратитесь в службу поддержки!")
}

func Error(ctx tele.Context, err error) error {
	context.Logger(ctx).Error(err.Error())
	return ctx.Send("Упс! Что-то пошло не так. Попробуйте повторить действие или обратитесь в службу поддержки!")
}
