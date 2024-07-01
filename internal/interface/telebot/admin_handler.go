package telebot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/middleware"
	tele "github.com/EgorMizerov/telebot"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
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
	fsm                            FSM
	adminService                   interfaces.AdminService
	driverService                  interfaces.DriverService
	registrationApplicationService interfaces.RegistrationApplicationService
}

func NewAdminHandler(bot *Bot, stateMachine FSM, driverService interfaces.DriverService, adminService interfaces.AdminService, registrationApplicationService interfaces.RegistrationApplicationService) *AdminHandler {
	admin := &AdminHandler{
		fsm:                            stateMachine,
		adminService:                   adminService,
		driverService:                  driverService,
		registrationApplicationService: registrationApplicationService,
	}

	bot.HandleStart(entity.AdminRole, admin.Menu)
	bot.Handle(markup.AdminUsersButton.Text, admin.DriversList, middleware.AdminAuth())
	bot.Handle(markup.DriverPhoneRegexp.Endpoint(), admin.GetDriverByPhone(false), middleware.AdminAuth())
	{ // Registration Applications
		bot.Handle(markup.AdminUsersRegistrationApplicationsButton.Text, admin.RegistrationApplications, middleware.AdminAuth())
		bot.Handle(markup.RegistrationApplicationIDRegexp.Endpoint(), admin.EditRegistrationApplications, middleware.AdminAuth())
		bot.Handle(markup.ConfirmRegistrationApplicationRegexp.Endpoint(), admin.ConfirmRegistrationApplication, middleware.AdminAuth())

		{ // Set work rule for Registration Applications
			bot.Handle(markup.SetFixWorkRuleForApplicationRegexp.Endpoint(), admin.SetWorkRuleForApplication(entity.FixWorkRule), middleware.AdminAuth())
			bot.Handle(markup.SetFixSelfEmployedWorkRuleForApplicationRegexp.Endpoint(), admin.SetWorkRuleForApplication(entity.FixSelfEmployedWorkRule), middleware.AdminAuth())
			bot.Handle(markup.SetPercentWorkRuleForApplicationRegexp.Endpoint(), admin.SetWorkRuleForApplication(entity.PercentWorkRule), middleware.AdminAuth())
			bot.Handle(markup.SetPercentSelfEmployedWorkRuleForApplicationRegexp.Endpoint(), admin.SetWorkRuleForApplication(entity.PercentSelfEmployedWorkRule), middleware.AdminAuth())
			bot.Handle(markup.SetPerDayWorkRuleForApplicationRegexp.Endpoint(), admin.SetWorkRuleForApplication(entity.PerDayWorkRule), middleware.AdminAuth())
		}
	}

	bot.Handle(&markup.DriverInfoShowCarInfoButton, admin.GetDriversCarInfo, middleware.AdminAuth())
	bot.Handle(&markup.DriverInfoShowCarInfoBackButton, admin.GetDriverByPhone(true), middleware.AdminAuth())

	return admin
}

func (self *AdminHandler) Menu(ctx tele.Context) error {
	return ctx.EditOrSend("Меню администратора", markup.AdminMenuMarkup())
}

func (self *AdminHandler) RegistrationApplications(ctx tele.Context) error {
	registrationApplications, err := self.registrationApplicationService.GetRegistrationApplications(ctx)
	if err != nil {
		return Error(ctx, err)
	}

	var registrationApplicationsData = template.RegistrationApplicationsData{
		Items: lo.Map(registrationApplications, func(item *entity.RegistrationApplication, index int) template.RegistrationApplication {
			return template.RegistrationApplication{
				ID:       int(item.ID),
				Status:   item.Status,
				Date:     item.Date,
				Fullname: item.Fullname(),
				Link:     item.Link(),
			}
		}),
	}

	return ctx.EditOrSend(
		template.ParseTemplate(template.RegistrationApplicationsTemplate, registrationApplicationsData), tele.ModeDefault,
	)
}

func (self *AdminHandler) EditRegistrationApplications(ctx tele.Context) error {
	id, _ := strconv.Atoi(ctx.Text())
	registrationApplication, err := self.registrationApplicationService.GetRegistrationApplication(ctx, entity.RegistrationApplicationID(id))
	if err != nil {
		return Error(ctx, errors.Wrap(err, "failed to get registration application"))
	}
	if registrationApplication.Status != "registered" {
		return ctx.Send("Редактировать можно только те заявки, что находятся в статусе \"registered\". Для этого их нужно принять в jump!")
	}

	return ctx.Send("Выберите тариф", markup.ChooseWorkRuleMarkup(registrationApplication.ID))
}

func (self *AdminHandler) ConfirmRegistrationApplication(ctx tele.Context) error {
	defer ctx.Delete()

	registrationApplicationID := entity.RegistrationApplicationID(markup.Regexp(ctx.Callback().Data).GetNumber())
	registrationApplication, err := self.registrationApplicationService.GetRegistrationApplication(ctx, registrationApplicationID)
	if err != nil {
		return Error(ctx, errors.Wrap(err, "failed to get registration application"))
	}
	registrationApplication.SetStatus("closed")

	err = self.registrationApplicationService.ConfirmRegistrationApplication(ctx, registrationApplication)
	if err != nil {
		return Error(ctx, errors.Wrap(err, "failed to confirm registration application"))
	}

	return ctx.RespondAlert("Регистрация пользователя подтверждена!")
}

func (self *AdminHandler) SetWorkRuleForApplication(rule entity.WorkRule) func(ctx tele.Context) error {
	return func(ctx tele.Context) error {
		registrationApplicationID := entity.RegistrationApplicationID(markup.Regexp(ctx.Callback().Data).GetNumber())
		registrationApplication, err := self.registrationApplicationService.GetRegistrationApplication(ctx, registrationApplicationID)
		if err != nil {
			return Error(ctx, errors.Wrap(err, "failed to get registration application"))
		}
		registrationApplication.SetWorkRule(rule)
		err = self.registrationApplicationService.SaveRegistrationApplication(ctx, registrationApplication)
		if err != nil {
			return Error(ctx, errors.Wrap(err, "failed to save registration application"))
		}
		return ctx.Edit(template.ParseTemplate(template.RegistrationApplicationTemplate, template.DriverRegistrationData{
			FullName:          registrationApplication.Fullname(),
			PhoneNumber:       *registrationApplication.PhoneNumber,
			Address:           *registrationApplication.City,
			DrivingExperience: *registrationApplication.LicenseTotalSinceDate,
			LicenseNumber:     *registrationApplication.LicenseNumber,
			LicenseCountry:    *registrationApplication.LicenseCountry,
			LicenseIssueDate:  *registrationApplication.LicenseIssueDate,
			LicenseExpiryDate: *registrationApplication.LicenseExpiryDate,
			WorkRule:          *registrationApplication.WorkRule,
			CarBrand:          *registrationApplication.CarBrand,
			CarModel:          *registrationApplication.CarModel,
			CarColor:          *registrationApplication.CarColor,
			CarYear:           *registrationApplication.CarYear,
			CarVIN:            *registrationApplication.CarVIN,
			CarNumber:         *registrationApplication.CarNumber,
		}), markup.ConfirmRegistrationApplicationMarkup(registrationApplicationID))
	}
}

func (self *AdminHandler) DriversList(ctx tele.Context) error {
	drivers, err := self.driverService.GetDrivers(ctx)
	if err != nil {
		return Error(ctx, errors.Wrap(err, "failed to get drivers"))
	}
	var driversListData = template.DriversListData{
		Items: lo.Map(drivers, func(item *entity.Driver, index int) template.DriversListItem {
			return template.DriversListItem{
				PhoneNumber: item.PhoneNuber.String(),
				Fullname:    item.Fullname(),
			}
		}),
	}

	return ctx.Send(
		template.ParseTemplate(template.DriversListTemplate, driversListData),
	)
}

func (self *AdminHandler) GetDriverByPhone(edit bool) func(ctx tele.Context) error {
	return func(ctx tele.Context) error {
		driver, err := self.driverService.GetDriverByPhoneNumber(ctx, entity.PhoneNumber(ctx.Text()))
		if err != nil {
			if errors.Is(err, interfaces.ErrDriverNotFound) {
				return ctx.Send(fmt.Sprintf("Водителя с номером %s не суещствует!", ctx.Text()))
			}
			return Error(ctx, err)
		}

		message := template.ParseTemplate(template.DriverInfoTemplate, template.DriverInfoData{
			ID:             uuid.UUID(driver.ID).String(),
			Fullname:       driver.Fullname(),
			PhoneNumber:    driver.PhoneNuber.String(),
			City:           driver.City,
			IsSelfEmployed: driver.IsSelfEmployed,
			WorkRule:       driver.WorkRule.StringPointer(),
		})

		if edit {
			return ctx.Edit(message, markup.DriverInfoMarkup())
		} else {
			return ctx.Send(message, markup.DriverInfoMarkup())
		}
	}
}

func (self *AdminHandler) GetDriversCarInfo(ctx tele.Context) error {
	car := entity.NewCar("some_id", "Lada", "Vesta", 2012, "Белый", "somevin", "П123НС159", "9923742455")
	return ctx.Edit(
		template.ParseTemplate(template.DriversCarInfoTemplate, template.DriversCarInfoData{
			Brand:  car.Brand,
			Model:  car.Model,
			Color:  car.Color,
			Year:   car.Year,
			VIN:    car.VIN,
			Number: car.LicensePlateNumber,
		}),
		markup.DriverCarInfoMarkup(),
	)
}

func (self *AdminHandler) CardsInfo(ctx tele.Context) error {
	cardsResult, err := self.adminService.GetCards(context.TODO())
	if err != nil {
		return Error(ctx, errors.Wrap(err, "failed to get cards info"))
	}

	return ctx.EditOrSend(
		template.ParseTemplate(template.GetCardsInfo, template.CardsInfoData{
			TinkoffCardsCount: cardsResult.TinkoffCardsCount,
			AnotherCardsCount: cardsResult.AnotherCardsCount,
		}),
	)
}
