package markup

import (
	. "github.com/EgorMizerov/telebot"
)

var (
	AdminMenuButton                                = InlineButton{Text: "Меню администратора", Unique: AdminMenuData}
	CreateDriverButton                             = InlineButton{Text: "Регистрация пользователя", WebApp: &WebApp{URL: "https://my.jump.taxi/autoregistration/list/69433/car/edit"}}
	CreateDriverAbortButton                        = InlineButton{Text: "Прервать регистрацию", Unique: Admin_CreateDriver_AbortData}
	CreateDriverConfirmButton                      = InlineButton{Text: "Подтвердить регистрацию", Unique: Admin_CreateDriver_ConfirmData}
	CreateDriverInputIDButton                      = InlineButton{Text: "Идентификатор пользователя", Unique: Admin_CreateDriver_InputIDData}
	CreateDriverInputFullnameButton                = InlineButton{Text: "ФИО пользователя", Unique: Admin_CreateDriver_InputFullnameData}
	CreateDriverInputPhoneNumberButton             = InlineButton{Text: "Номер телефона", Unique: Admin_CreateDriver_InputPhoneNumberData}
	CreateDriverInputAddressButton                 = InlineButton{Text: "Адрес проживания", Unique: Admin_CreateDriver_InputAddressData}
	CreateDriverInputIsSelfEmployedButton          = InlineButton{Text: "Статус самозанятасти", Unique: Admin_CreateDriver_InputIsSelfEmployedData}
	CreateDriverInputWorkRuleButton                = InlineButton{Text: "Тариф", Unique: Admin_CreateDriver_InputTariffData}
	CreateDriverInputCardNumberButton              = InlineButton{Text: "Номер карты", Unique: Admin_CreateDriver_InputCardNumberData}
	CreateDriverInputDrivingExperienceButton       = InlineButton{Text: "Водительский стаж", Unique: Admin_CreateDriver_InputDrivingExperienceData}
	CreateDriverInputRegistrationCertificateButton = InlineButton{Text: "Серия и номер ВУ", Unique: Admin_CreateDriver_InputRegistrationCerteficateData}
	CreateDriverInputLicenseCountryButton          = InlineButton{Text: "Страна выдачи ВУ", Unique: Admin_CreateDriver_InputLicenseCountry}
	CreateDriverInputLicenseIssueDateButton        = InlineButton{Text: "Дата выдачи ВУ", Unique: Admin_CreateDriver_InputLicenseIssueDate}
	CreateDriverInputLicenseExpiryDateButton       = InlineButton{Text: "Дата окончания действия ВУ", Unique: Admin_CreateDriver_InputLicenseExpiryDate}
	CreateDriverInputCarBrandButton                = InlineButton{Text: "Марка автомобиля", Unique: Admin_CreateDriver_InputCarBrand}
	CreateDriverInputCarModelButton                = InlineButton{Text: "Модель автомобиля", Unique: Admin_CreateDriver_InputCarModel}
	CreateDriverInputCarColorButton                = InlineButton{Text: "Цвет автомобиля", Unique: Admin_CreateDriver_InputCarColor}
	CreateDriverInputCarYearButton                 = InlineButton{Text: "Год выпуска автомобиля", Unique: Admin_CreateDriver_InputCarYear}
	CreateDriverInputCarVINButton                  = InlineButton{Text: "VIN автомобиля", Unique: Admin_CreateDriver_InputCarVIN}
	CreateDriverInputCarLicensePlateNumberButton   = InlineButton{Text: "Гос. Номер автомобиля", Unique: Admin_CreateDriver_InputLicensePlateNumber}
	CreateDriverInputReferralKeyButton             = InlineButton{Text: "Реферальный ключ", Unique: Admin_CreateDriver_InputRefferalKey}

	CardsInfoXLSXButton = InlineButton{Text: "Сформировать отчёт", Unique: Admin_CreateDriver_InputRefferalKey}
)

func MenuMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{
			{CreateDriverButton},
		},
	}
}

func CreateDriverMarkup(nextTxt string, nextData string) *ReplyMarkup {
	return &ReplyMarkup{
		InlineKeyboard: [][]InlineButton{
			{CreateDriverInputIDButton},
			{CreateDriverInputFullnameButton, CreateDriverInputPhoneNumberButton},
			{CreateDriverInputAddressButton, CreateDriverInputIsSelfEmployedButton},
			{CreateDriverInputWorkRuleButton, CreateDriverInputCardNumberButton},
			{CreateDriverInputDrivingExperienceButton, CreateDriverInputRegistrationCertificateButton},
			{CreateDriverInputLicenseCountryButton, CreateDriverInputLicenseIssueDateButton},
			{CreateDriverInputLicenseExpiryDateButton, CreateDriverInputCarBrandButton},
			{CreateDriverInputCarModelButton, CreateDriverInputCarColorButton},
			{CreateDriverInputCarYearButton, CreateDriverInputCarVINButton},
			{CreateDriverInputCarLicensePlateNumberButton, CreateDriverInputReferralKeyButton},
			{CreateDriverAbortButton},
			{{Text: nextTxt, Unique: nextData}},
		},
		ResizeKeyboard: true,
	}
}

func CardsInfoMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		ResizeKeyboard: true,
		InlineKeyboard: [][]InlineButton{},
	}
}
