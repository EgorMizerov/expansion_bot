package markup

import (
	. "github.com/EgorMizerov/telebot"
)

var (
	AdminUsersButton                         = ReplyButton{Text: "Список пользователей"}
	AdminUsersRegistrationApplicationsButton = ReplyButton{Text: "Заявки на регистрацию"}
	DriverInfoShowDriverLicenseInfoButton    = InlineButton{Text: "Водительские права", Unique: DriverInfoShowDriverLicenseInfoData}
	DriverInfoShowCarInfoButton              = InlineButton{Text: "Автомобиль", Unique: DriverInfoShowCarInfoData}
	DriverInfoShowCarInfoBackButton          = InlineButton{Text: "Вернуться назад", Unique: DriverInfoShowCarInfoBackData}
)

func AdminMenuMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]ReplyButton{
			{AdminUsersButton},
			{AdminUsersRegistrationApplicationsButton},
		},
	}
}

func DriverInfoMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{DriverInfoShowDriverLicenseInfoButton, DriverInfoShowCarInfoButton},
		},
	}
}

func DriverCarInfoMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{DriverInfoShowCarInfoBackButton},
		},
	}
}

func CardsInfoMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{},
	}
}
