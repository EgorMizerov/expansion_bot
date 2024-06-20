package markup

import . "github.com/EgorMizerov/telebot"

func SignUpMarkup() *ReplyMarkup {
	markup := &ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		ReplyKeyboard: [][]ReplyButton{
			{{Text: "Отправить контакт", WebApp: &WebApp{
				URL: "https://my.jump.taxi/autoregistration/form/c26193",
			}}},
		},
	}

	return markup
}
