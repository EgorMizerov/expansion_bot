package telebot

import (
	"context"

	tele "github.com/EgorMizerov/telebot"
	"github.com/pkg/errors"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	hcontext "github.com/EgorMizerov/expansion_bot/internal/interface/telebot/context"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/markup"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/template"
)

type FSM interface {
	SetState(userId int64, state entity.State)
	GetState(userId int64) (entity.State, error)
	Clear(userID int64)
	SaveRegistrationData(ctx context.Context, data template.DriverRegistrationData) error
	GetRegistrationData(ctx context.Context) (template.DriverRegistrationData, error)
}

type AdminHandler struct {
	fsm          FSM
	adminService interfaces.AdminService
}

func NewAdminHandler(bot *Bot, stateMachine FSM, adminService interfaces.AdminService) *AdminHandler {
	admin := &AdminHandler{
		fsm:          stateMachine,
		adminService: adminService,
	}

	bot.HandleStart(entity.AdminRole, admin.Menu)
	bot.Handle(&markup.AdminMenuButton, admin.Menu)
	//bot.Handle(&markup.CreateDriverButton, admin.CreateDriver)
	//bot.Handle(&markup.CreateDriverConfirmButton, admin.CreateDriverConfirm)
	//bot.Handle(&markup.CreateDriverAbortButton, admin.CreateDriverAbort)
	//bot.Handle(&markup.CreateDriverInputIDButton, admin.CreateDriverInputUserID)
	//bot.Handle(&markup.CreateDriverInputFullnameButton, admin.CreateDriverInputFullname)
	//bot.Handle(&markup.CreateDriverInputPhoneNumberButton, admin.CreateDriverInputPhoneNumber)
	//bot.Handle(&markup.CreateDriverInputAddressButton, admin.CreateDriverInputAddress)
	//bot.Handle(&markup.CreateDriverInputIsSelfEmployedButton, admin.CreateDriverInputIsSelfEmployed)
	//bot.Handle(&markup.CreateDriverInputWorkRuleButton, admin.CreateDriverInputTariff)
	//bot.Handle(&markup.CreateDriverInputCardNumberButton, admin.CreateDriverInputCardNumber)
	//bot.Handle(&markup.CreateDriverInputDrivingExperienceButton, admin.CreateDriverInputDrivingExperience)
	//bot.Handle(&markup.CreateDriverInputRegistrationCertificateButton, admin.CreateDriverInputRegistrationCertificate)
	//bot.Handle(&markup.CreateDriverInputLicenseCountryButton, admin.CreateDriverInputLicenseCountry)
	//bot.Handle(&markup.CreateDriverInputLicenseIssueDateButton, admin.CreateDriverInputLicenseIssueDate)
	//bot.Handle(&markup.CreateDriverInputLicenseExpiryDateButton, admin.CreateDriverInputLicenseExpiryDate)
	//bot.Handle(&markup.CreateDriverInputCarBrandButton, admin.CreateDriverInputCarBrand)
	//bot.Handle(&markup.CreateDriverInputCarModelButton, admin.CreateDriverInputCarModel)
	//bot.Handle(&markup.CreateDriverInputCarColorButton, admin.CreateDriverInputCarColor)
	//bot.Handle(&markup.CreateDriverInputCarYearButton, admin.CreateDriverInputCarYear)
	//bot.Handle(&markup.CreateDriverInputCarVINButton, admin.CreateDriverInputCarVIN)
	//bot.Handle(&markup.CreateDriverInputCarLicensePlateNumberButton, admin.CreateDriverInputLicensePlateNumber)
	//bot.Handle(&markup.CreateDriverInputReferralKeyButton, admin.CreateDriverInputReferralKey)

	//bot.HandleState(entity.CreateDriver_ReceiveUserID, admin.CreateDriverReceiveUserID)
	//bot.HandleState(entity.CreateDriver_ReceiveFullanme, admin.CreateDriverReceiveFullname)
	//bot.HandleState(entity.CreateDriver_ReceivePhoneNumber, admin.CreateDriverReceivePhoneNumber)
	//bot.HandleState(entity.CreateDriver_ReceiveAddress, admin.CreateDriverReceiveAddress)
	//bot.HandleState(entity.CreateDriver_ReceiveIsSelfEmployed, admin.CreateDriverReceiveIsSelfEmployed)
	//bot.HandleState(entity.CreateDriver_ReceiveTariff, admin.CreateDriverReceiveTariff)
	//bot.HandleState(entity.CreateDriver_ReceiveCardNumber, admin.CreateDriverReceiveCardNumber)
	//bot.HandleState(entity.CreateDriver_ReceiveDrivingExperience, admin.CreateDriverReceiveDrivingExperience)
	//bot.HandleState(entity.CreateDriver_ReceiveRegistrationCertificate, admin.CreateDriverReceiveRegistrationCertificate)
	//bot.HandleState(entity.CreateDriver_ReceiveLicenseCountry, admin.CreateDriverReceiveLicenseCountry)
	//bot.HandleState(entity.CreateDriver_ReceiveLicenseIssueDate, admin.CreateDriverReceiveLicenseIssueDate)
	//bot.HandleState(entity.CreateDriver_ReceiveLicenseExpiryDate, admin.CreateDriverReceiveLicenseExpiryDate)
	//bot.HandleState(entity.CreateDriver_ReceiveCarBrand, admin.CreateDriverReceiveCarBrand)
	//bot.HandleState(entity.CreateDriver_ReceiveCarModel, admin.CreateDriverReceiveCarModel)
	//bot.HandleState(entity.CreateDriver_ReceiveCarColor, admin.CreateDriverReceiveCarColor)
	//bot.HandleState(entity.CreateDriver_ReceiveCarYear, admin.CreateDriverReceiveCarYear)
	//bot.HandleState(entity.CreateDriver_ReceiveCarVIN, admin.CreateDriverReceiveCarVIN)
	//bot.HandleState(entity.CreateDriver_ReceiveLicensePlateNumber, admin.CreateDriverReceiveLicensePlateNumber)
	//bot.HandleState(entity.CreateDriver_ReceiveReferralKey, admin.CreateDriverReceiveReferralKey)

	return admin
}

func (self *AdminHandler) Menu(ctx tele.Context) error {
	return ctx.EditOrSend("Меню администратора", markup.MenuMarkup())
}

//func (self *AdminHandler) CreateDriver(ctx tele.Context) error {
//	var driverRegistrationData = template.DriverRegistrationData{Message: ctx.Message()}
//	err := self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	return ctx.EditOrSend(
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести идентификатор", markup.Admin_CreateDriver_InputIDData),
//	)
//}
//
//func (self *AdminHandler) CreateDriverInputUserID(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveUserID)
//	return ctx.EditOrSend("Введите идентификатор пользователя")
//}
//
//func (self *AdminHandler) CreateDriverReceiveUserID(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	userID, err := strconv.Atoi(ctx.Message().Text)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to convert message text to int"))
//	}
//	driverRegistrationData.UserID = userID
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести ФИО", markup.Admin_CreateDriver_InputFullnameData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputFullname(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveFullanme)
//	return ctx.EditOrSend("Введите ФИО пользователя")
//}
//
//func (self *AdminHandler) CreateDriverReceiveFullname(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	fullnameParts := strings.Split(ctx.Message().Text, " ")
//	if len(fullnameParts) < 3 {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести ФИО", markup.Admin_CreateDriver_InputFullnameData),
//		)
//		return err
//	}
//	driverRegistrationData.FullName = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести номер телефона", markup.Admin_CreateDriver_InputPhoneNumberData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputPhoneNumber(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceivePhoneNumber)
//	return ctx.EditOrSend("Введите номер телефона пользователя")
//}
//
//func (self *AdminHandler) CreateDriverReceivePhoneNumber(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	isPhone, err := regexp.MatchString(`^\+7\d{10}`, ctx.Message().Text)
//	if err != nil || !isPhone {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести номер телефона", markup.Admin_CreateDriver_InputPhoneNumberData),
//		)
//		return err
//	}
//	driverRegistrationData.PhoneNumber = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести адрес проживания", markup.Admin_CreateDriver_InputAddressData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputAddress(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveAddress)
//	return ctx.EditOrSend("Введите адрес проживания пользователя")
//}
//
//func (self *AdminHandler) CreateDriverReceiveAddress(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	driverRegistrationData.Address = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести статус самозанятасти", markup.Admin_CreateDriver_InputIsSelfEmployedData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputIsSelfEmployed(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveIsSelfEmployed)
//	return ctx.EditOrSend("Введите статус самозанятасти пользователя (`Да` или `Нет`)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveIsSelfEmployed(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	if strings.ToLower(ctx.Message().Text) == "да" {
//		driverRegistrationData.IsSelfEmployed = true
//	} else {
//		driverRegistrationData.IsSelfEmployed = false
//	}
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Выбрать тариф", markup.Admin_CreateDriver_InputTariffData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputTariff(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveTariff)
//	return ctx.EditOrSend("Выберите тариф (`процент`, `фикс/заказ`, `фикс/день`)", tele.ModeMarkdown)
//}
//
//func (self *AdminHandler) CreateDriverReceiveTariff(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	if strings.ToLower(ctx.Message().Text) == "процент" && driverRegistrationData.IsSelfEmployed {
//		driverRegistrationData.WorkRule = entity.PercentSelfEmployedWorkRule
//	} else if strings.ToLower(ctx.Message().Text) == "процент" && !driverRegistrationData.IsSelfEmployed {
//		driverRegistrationData.WorkRule = entity.PercentWorkRule
//	} else if strings.ToLower(ctx.Message().Text) == "фикс/заказ" && !driverRegistrationData.IsSelfEmployed {
//		driverRegistrationData.WorkRule = entity.FixWorkRule
//	} else if strings.ToLower(ctx.Message().Text) == "фикс/заказ" && driverRegistrationData.IsSelfEmployed {
//		driverRegistrationData.WorkRule = entity.FixSelfEmployedWorkRule
//	} else if strings.ToLower(ctx.Message().Text) == "фикс/день" && driverRegistrationData.IsSelfEmployed {
//		driverRegistrationData.WorkRule = entity.PerDayWorkRule
//	} else {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Выбрать тариф", markup.Admin_CreateDriver_InputTariffData),
//		)
//		return err
//	}
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести номер карты", markup.Admin_CreateDriver_InputCardNumberData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCardNumber(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCardNumber)
//	return ctx.EditOrSend("Введите номер карты пользователя")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCardNumber(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	cardNumber := strings.ReplaceAll(ctx.Message().Text, " ", "")
//	isCardNumber, err := regexp.MatchString(`^\d{16}`, cardNumber)
//	if err != nil || !isCardNumber {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести номер карты", markup.Admin_CreateDriver_InputCardNumberData),
//		)
//		return err
//	}
//	driverRegistrationData.CardNumber = template.CardNumber(cardNumber)
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести водительский стаж", markup.Admin_CreateDriver_InputDrivingExperienceData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputDrivingExperience(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveDrivingExperience)
//	return ctx.EditOrSend("Введите водительский стаж с пользователя (yyyy-mm-dd)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveDrivingExperience(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	drivingExperience, err := time.Parse(time.DateOnly, ctx.Message().Text)
//	if err != nil {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести водительский стаж", markup.Admin_CreateDriver_InputDrivingExperienceData),
//		)
//		return err
//	}
//	driverRegistrationData.DrivingExperience = drivingExperience
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести серию и номер ВУ", markup.Admin_CreateDriver_InputRegistrationCerteficateData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputRegistrationCertificate(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveRegistrationCertificate)
//	return ctx.EditOrSend("Введите серию и номер водительского удостоверения (xx xx xxxxxx)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveRegistrationCertificate(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	ok, err := regexp.MatchString(`\d{2} \d{2} \d{6}`, ctx.Message().Text)
//	if err != nil || !ok {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести серию и номер ВУ", markup.Admin_CreateDriver_InputRegistrationCerteficateData),
//		)
//		return err
//	}
//	driverRegistrationData.RegistrationCertificate = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести страну выдачи ВУ", markup.Admin_CreateDriver_InputLicenseCountry),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputLicenseCountry(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveLicenseCountry)
//	return ctx.EditOrSend("Введите стару выдачи водительского удостоверения")
//}
//
//func (self *AdminHandler) CreateDriverReceiveLicenseCountry(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	// TODO: only rus
//	driverRegistrationData.LicenseCountry = "rus"
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести дату выдачи ВУ", markup.Admin_CreateDriver_InputLicenseIssueDate),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputLicenseIssueDate(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveLicenseIssueDate)
//	return ctx.EditOrSend("Введите дату выдачи водительского удостоверения (yyyy-mm-dd)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveLicenseIssueDate(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	licenseIssueDate, err := time.Parse(time.DateOnly, ctx.Message().Text)
//	if err != nil {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести дату выдачи ВУ", markup.Admin_CreateDriver_InputLicenseIssueDate),
//		)
//		return err
//	}
//	driverRegistrationData.LicenseIssueDate = licenseIssueDate
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести дату окончания действия ВУ", markup.Admin_CreateDriver_InputLicenseExpiryDate),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputLicenseExpiryDate(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveLicenseExpiryDate)
//	return ctx.EditOrSend("Введите дату окончания действия водительского удостоверения (yyyy-mm-dd)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveLicenseExpiryDate(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	licenseExpiryDate, err := time.Parse(time.DateOnly, ctx.Message().Text)
//	if err != nil {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести дату окончания действия ВУ", markup.Admin_CreateDriver_InputLicenseExpiryDate),
//		)
//		return err
//	}
//	driverRegistrationData.LicenseExpiryDate = licenseExpiryDate
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести марку автомобиля", markup.Admin_CreateDriver_InputCarBrand),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCarBrand(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCarBrand)
//	return ctx.EditOrSend("Введите марку автомобиля")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCarBrand(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	driverRegistrationData.CarBrand = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести модель автомобиля", markup.Admin_CreateDriver_InputCarModel),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCarModel(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCarModel)
//	return ctx.EditOrSend("Введите модель автомобиля")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCarModel(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	driverRegistrationData.CarModel = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Выбрать цвет автомобиля", markup.Admin_CreateDriver_InputCarColor),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCarColor(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCarColor)
//	return ctx.EditOrSend("Выберите цвет автомобиля (`Белый`, `Желтый`, `Бежевый`, `Черный`, `Голубой`, `Серый`, `Красный`, `Оранжевый`, `Синий`, `Зеленый`, `Коричневый`, `Фиолетовый`, `Розовый`)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCarColor(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	colors := strings.Split("Белый, Желтый, Бежевый, Черный, Голубой, Серый, Красный, Оранжевый, Синий, Зеленый, Коричневый, Фиолетовый, Розовый", ", ")
//	if !slices.Contains(colors, ctx.Message().Text) {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Выбрать цвет автомобиля", markup.Admin_CreateDriver_InputCarColor),
//		)
//		return err
//	}
//	driverRegistrationData.CarColor = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести год выпуска автомобиля", markup.Admin_CreateDriver_InputCarYear),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCarYear(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCarYear)
//	return ctx.EditOrSend("Введите год выпуска автомобиля (yyyy)")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCarYear(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	carYear, err := time.Parse("2006", ctx.Message().Text)
//	if err != nil {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести год выпуска автомобиля", markup.Admin_CreateDriver_InputLicenseIssueDate),
//		)
//		return err
//	}
//	driverRegistrationData.CarYear = carYear
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести VIN номер автомобиля", markup.Admin_CreateDriver_InputCarVIN),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputCarVIN(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveCarVIN)
//	return ctx.EditOrSend("Введите VIN номер автомобиля")
//}
//
//func (self *AdminHandler) CreateDriverReceiveCarVIN(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	driverRegistrationData.CarVIN = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Ввести гос. Номер автомобиля", markup.Admin_CreateDriver_InputLicensePlateNumber),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputLicensePlateNumber(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveLicensePlateNumber)
//	return ctx.EditOrSend("Введите гос. Номер автомобиля")
//}
//
//func (self *AdminHandler) CreateDriverReceiveLicensePlateNumber(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	driverRegistrationData.LicensePlateNumber = ctx.Message().Text
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Завершить регистрацию", markup.Admin_CreateDriver_ConfirmData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverInputReferralKey(ctx tele.Context) error {
//	self.fsm.SetState(ctx.Sender().ID, entity.CreateDriver_ReceiveReferralKey)
//	return ctx.EditOrSend("Введите реферальный ключ")
//}
//
//func (self *AdminHandler) CreateDriverReceiveReferralKey(ctx tele.Context) error {
//	defer ctx.Delete()
//
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	referralKey, err := strconv.Atoi(ctx.Message().Text)
//	if err != nil {
//		_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//			template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//			markup.CreateDriverMarkup("Ввести реферальный ключ", markup.Admin_CreateDriver_InputRefferalKey),
//		)
//		return err
//	}
//	driverRegistrationData.ReferralKey = int64(referralKey)
//
//	err = self.fsm.SaveRegistrationData(context.TODO(), driverRegistrationData)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to save driver registration data"))
//	}
//
//	_, err = ctx.Bot().Edit(driverRegistrationData.Message,
//		template.ParseTemplate(template.DriverRegistrationTemplate, driverRegistrationData),
//		markup.CreateDriverMarkup("Завершить регистрацию", markup.Admin_CreateDriver_ConfirmData),
//	)
//	return err
//}
//
//func (self *AdminHandler) CreateDriverConfirm(ctx tele.Context) error {
//	driverRegistrationData, err := self.fsm.GetRegistrationData(context.TODO())
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to get driver registration data"))
//	}
//
//	err = self.adminService.CreateDriver(
//		context.TODO(),
//		driverRegistrationData.ToDriver(),
//		driverRegistrationData.ToCar(),
//	)
//	if err != nil {
//		return handleError(ctx, errors.Wrap(err, "failed to create driver"))
//	}
//
//	return self.Menu(ctx)
//}
//
//func (self *AdminHandler) CreateDriverAbort(ctx tele.Context) error {
//	self.fsm.Clear(ctx.Sender().ID)
//	return self.Menu(ctx)
//}

func (self *AdminHandler) CardsInfo(ctx tele.Context) error {
	cardsResult, err := self.adminService.GetCards(context.TODO())
	if err != nil {
		return handleError(ctx, errors.Wrap(err, "failed to get cards info"))
	}

	return ctx.EditOrSend(
		template.ParseTemplate(template.GetCardsInfo, template.CardsInfoData{
			TinkoffCardsCount: cardsResult.TinkoffCardsCount,
			AnotherCardsCount: cardsResult.AnotherCardsCount,
		}),
	)
}

func handleError(ctx tele.Context, err error) error {
	hcontext.Logger(ctx).Error(err.Error())
	return nil
}
