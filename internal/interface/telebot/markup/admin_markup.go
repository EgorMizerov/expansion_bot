package markup

import (
	"fmt"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	. "github.com/EgorMizerov/telebot"
)

var (
	AdminUsersButton                         = ReplyButton{Text: "Список пользователей"}
	AdminUsersRegistrationApplicationsButton = ReplyButton{Text: "Заявки на регистрацию"}
	DriverInfoShowDriverLicenseInfoButton    = InlineButton{Text: "Водительские права", Unique: DriverInfoShowDriverLicenseInfoData}
	DriverInfoShowCarInfoButton              = InlineButton{Text: "Автомобиль", Unique: DriverInfoShowCarInfoData}
	DriverInfoShowCarInfoBackButton          = InlineButton{Text: "Вернуться назад", Unique: DriverInfoShowCarInfoBackData}

	SetFixWorkRuleForApplicationButton                 = InlineButton{RXUnique: SetFixWorkRuleForApplicationRegexp.Endpoint()}
	SetFixSelfEmployedWorkRuleForApplicationButton     = InlineButton{RXUnique: SetFixSelfEmployedWorkRuleForApplicationRegexp.Endpoint()}
	SetPercentSelfEmployedWorkRuleForApplicationButton = InlineButton{RXUnique: SetPercentSelfEmployedWorkRuleForApplicationRegexp.Endpoint()}
	SetPercentWorkRuleForApplicationButton             = InlineButton{RXUnique: SetPercentWorkRuleForApplicationRegexp.Endpoint()}
	SetPerDayWorkRuleForApplicationButton              = InlineButton{RXUnique: SetPerDayWorkRuleForApplicationRegexp.Endpoint()}
	ConfirmRegistrationApplicationButton               = InlineButton{RXUnique: ConfirmRegistrationApplicationRegexp.Endpoint()}
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

func ChooseWorkRuleMarkup(registrationApplicationID entity.RegistrationApplicationID) *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{{Text: entity.FixSelfEmployedWorkRule.Name, Data: fmt.Sprintf("rx:fsewr:%d", registrationApplicationID)}, {Text: entity.PercentSelfEmployedWorkRule.Name, Data: fmt.Sprintf("rx:psewr:%d", registrationApplicationID)}},
			{{Text: entity.FixWorkRule.Name, Data: fmt.Sprintf("rx:fwr:%d", registrationApplicationID)}, {Text: entity.PercentWorkRule.Name, Data: fmt.Sprintf("rx:pwr:%d", registrationApplicationID)}},
			{{Text: entity.PerDayWorkRule.Name, Data: fmt.Sprintf("rx:pdwr:%d", registrationApplicationID)}},
		},
	}
}

func ConfirmRegistrationApplicationMarkup(registrationApplicationID entity.RegistrationApplicationID) *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{{Text: "Подтвердить", Data: fmt.Sprintf("rx:cf_ra:%d", registrationApplicationID)}},
		},
	}
}

func CardsInfoMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{},
	}
}
