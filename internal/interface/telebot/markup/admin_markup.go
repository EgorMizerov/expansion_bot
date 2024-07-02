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
	DriverInfoFromCarToDriverInfoButton                = InlineButton{RXUnique: DriverInfoFromCarToDriverInfoRegexp.Endpoint()}
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

func DriverInfoMarkup(phoneNumber entity.PhoneNumber) *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{
				{Text: "Водительское удостоверение", Data: fmt.Sprintf("rx:di_s_dli:%s", phoneNumber)},
				{Text: "Автомобиль", Data: fmt.Sprintf("rx:di_s_ci:%s", phoneNumber)},
			},
		},
	}
}

func DriverCarInfoMarkup(phoneNumber entity.PhoneNumber) *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{{Text: "Вернуться назад", Data: fmt.Sprintf("rx:di_f_c_t_di:%s", phoneNumber)}},
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
