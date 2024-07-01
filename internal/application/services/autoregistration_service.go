package services

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	fleet2 "github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/request"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/types"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type RegistrationApplicationService struct {
	fleetClient                       *fleet2.FleetClient
	jumpClient                        *jump.JumpClient
	registrationApplicationRepository repository.RegistrationApplicationRepository
	driverRepository                  repository.DriverRepository
	carRepository                     repository.CarRepository
	guestRepository                   repository.GuestRepository
}

func NewRegistrationApplicationService(
	fleetClient *fleet2.FleetClient,
	jumpClient *jump.JumpClient,
	registrationApplicationRepository repository.RegistrationApplicationRepository,
	driverRepository repository.DriverRepository,
	carRepository repository.CarRepository,
	guestRepository repository.GuestRepository,
) *RegistrationApplicationService {
	return &RegistrationApplicationService{
		fleetClient:                       fleetClient,
		jumpClient:                        jumpClient,
		registrationApplicationRepository: registrationApplicationRepository,
		driverRepository:                  driverRepository,
		carRepository:                     carRepository,
		guestRepository:                   guestRepository,
	}
}

func (self *RegistrationApplicationService) GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error) {
	return self.registrationApplicationRepository.GetRegistrationApplication(ctx, applicationID)
}

func (self *RegistrationApplicationService) GetRegistrationApplications(ctx context.Context) ([]*entity.RegistrationApplication, error) {
	return self.registrationApplicationRepository.GetRegistrationApplications(ctx)
}

func (self *RegistrationApplicationService) SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	return self.registrationApplicationRepository.SaveRegistrationApplication(ctx, application)
}

func (self *RegistrationApplicationService) ConfirmRegistrationApplication(ctx context.Context, registrationApplication *entity.RegistrationApplication) error {
	jumpDriver, err := self.jumpClient.GetDriverByPhoneNumber(ctx, *registrationApplication.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "failed to get driver from jump")
	}
	fleetDriver, err := self.getDriverFromFleetByPhoneNumber(ctx, *registrationApplication.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "failed to get driver from fleet")
	}
	guest, err := self.guestRepository.GetGuestByPhoneNumber(ctx, entity.PhoneNumber(*registrationApplication.PhoneNumber))
	if err != nil {
		return errors.Wrap(err, "failed to get guest by phone number")
	}

	car := entity.NewCar(
		entity.CarFleetID(fleetDriver.Car.ID),
		*registrationApplication.CarBrand,
		*registrationApplication.CarModel,
		*registrationApplication.CarYear,
		*registrationApplication.CarColor,
		*registrationApplication.CarVIN,
		*registrationApplication.CarNumber,
		*registrationApplication.CarLicense,
	)
	err = self.carRepository.CreateCar(ctx, car)
	if err != nil {
		return errors.Wrap(err, "failed to create car")
	}

	isSelfEmployed := func(rule entity.WorkRule) bool {
		return rule == entity.FixSelfEmployedWorkRule || rule == entity.PercentSelfEmployedWorkRule
	}(*registrationApplication.WorkRule)

	driver := entity.NewDriver(
		entity.JumpID(jumpDriver.ID),
		entity.FleetID(fleetDriver.DriverProfile.ID),
		guest.TelegramID,
		*registrationApplication.FirstName,
		*registrationApplication.LastName,
		registrationApplication.MiddleName,
		*registrationApplication.City,
		entity.PhoneNumber(*registrationApplication.PhoneNumber),
		car.ID,
		entity.NewDriverLicense(
			*registrationApplication.LicenseNumber,
			*registrationApplication.LicenseTotalSinceDate,
			*registrationApplication.LicenseIssueDate,
			*registrationApplication.LicenseExpiryDate,
			*registrationApplication.LicenseCountry,
		),
		isSelfEmployed,
		*registrationApplication.WorkRule,
	)
	err = self.driverRepository.CreateDriver(ctx, driver)
	if err != nil {
		return errors.Wrap(err, "failed to create driver")
	}

	registrationApplication.SetStatus("closed")
	err = self.registrationApplicationRepository.SaveRegistrationApplication(ctx, registrationApplication)
	if err != nil {
		return errors.Wrap(err, "failed to save registration application")
	}

	return nil
}

func (self *RegistrationApplicationService) getDriverFromFleetByPhoneNumber(ctx context.Context, phoneNumber string) (*types.GetDriversItem, error) {
	body := request.GetDriversRequest{
		OrderBy:        request.GetDriversRequestOrderByCreatedAt,
		OrderDirection: request.OrderByDesc,
		Limit:          10,
		Offset:         0,
	}
	drivers, err := self.fleetClient.GetDrivers(ctx, body)
	if err != nil {
		return nil, err
	}

	driver, ok := lo.Find(drivers.Items, func(item types.GetDriversItem) bool {
		return lo.Contains(item.DriverProfile.Phones, phoneNumber)
	})
	if !ok {
		return nil, errors.New("failed to find driver profile")
	}
	return &driver, nil
}
