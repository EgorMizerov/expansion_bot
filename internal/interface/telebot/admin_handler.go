package telebot

import (
	"context"
	"fmt"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/common"
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
	bot.Handle(markup.AdminUsersRegistrationApplicationsButton.Text, admin.RegistrationApplications)
	bot.Handle(markup.AdminUsersButton.Text, admin.DriversList)
	bot.Handle(`rx:\+7\d{10}`, admin.GetDriverByPhone(false))
	bot.Handle(&markup.DriverInfoShowCarInfoButton, admin.GetDriversCarInfo)
	bot.Handle(&markup.DriverInfoShowCarInfoBackButton, admin.GetDriverByPhone(true))

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
				Date:     item.Date,
				Fullname: item.Fullname(),
				Link:     item.Link(),
			}
		}),
	}

	return ctx.EditOrSend(
		template.ParseTemplate(template.RegistrationApplicationsTemplate, registrationApplicationsData),
	)
}

func (self *AdminHandler) DriversList(ctx tele.Context) error {
	car := entity.NewCar("test_fleet_id", "test_brand", "test_model", 2012, "test_color", "test_vin", "test_number", "test_license")
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "rus")
	driver1 := entity.NewDriver(111, "Мизеров", "Егор", "Мизеров", common.Point("Викторович"), "test_city", "+79956908933", car.ID, license)
	driver2 := entity.NewDriver(222, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "+79222112253", car.ID, license)

	var drivers = []*entity.Driver{driver1, driver2}
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
		err = nil
		driver = &entity.Driver{ID: entity.DriverID(uuid.New()), FirstName: "Егор", LastName: "Мизеров", PhoneNuber: "+79956908933"}
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
