package markup

import . "github.com/EgorMizerov/telebot"

var SignUpButton = InlineButton{Text: "Регистрация", WebApp: &WebApp{URL: "https://my.jump.taxi/autoregistration/form/c26193"}}

func SignUpMarkup() *ReplyMarkup {
	markup := &ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		ReplyKeyboard: [][]ReplyButton{
			{{Text: "Заполнить анкету", WebApp: &WebApp{
				URL: "https://my.jump.taxi/autoregistration/form/c26193",
			}}},
		},
	}

	return markup
}

func SendContactMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]ReplyButton{
			{{Text: "Поделиться контактам", Contact: true}},
		},
	}
}
