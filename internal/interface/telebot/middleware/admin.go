package middleware

import (
	"strconv"

	"github.com/EgorMizerov/expansion_bot/config"
	"github.com/EgorMizerov/telebot"
)

func AdminAuth() telebot.MiddlewareFunc {
	return func(handlerFunc telebot.HandlerFunc) telebot.HandlerFunc {
		return func(context telebot.Context) error {
			if strconv.Itoa(int(context.Sender().ID)) == config.Config.AdminID {
				return handlerFunc(context)
			}
			return nil
		}
	}
}
