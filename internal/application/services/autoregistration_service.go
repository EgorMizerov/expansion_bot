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
}

func NewRegistrationApplicationService(fleetClient *fleet2.FleetClient, jumpClient *jump.JumpClient, registrationApplicationRepository repository.RegistrationApplicationRepository, driverRepository repository.DriverRepository, carRepository repository.CarRepository) *RegistrationApplicationService {
	return &RegistrationApplicationService{fleetClient: fleetClient, jumpClient: jumpClient, registrationApplicationRepository: registrationApplicationRepository, driverRepository: driverRepository, carRepository: carRepository}
}

func (self *RegistrationApplicationService) GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error) {
	return self.registrationApplicationRepository.GetRegistrationApplication(ctx, applicationID)
}

func (self *RegistrationApplicationService) SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	return self.registrationApplicationRepository.SaveRegistrationApplication(ctx, application)
}

func (self *RegistrationApplicationService) ConfirmRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	jumpDriver, err := self.jumpClient.GetDriverByPhoneNumber(ctx, *application.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "failed to get driver from jump")
	}
	fleetDriver, err := self.getDriverFromFleetByPhoneNumber(ctx, *application.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "failed to get driver from fleet")
	}

	car := entity.NewCar(entity.CarFleetID(fleetDriver.Car.ID), *application.CarBrand, *application.CarModel, *application.CarYear, *application.CarColor, *application.CarVIN, *application.CarNumber, *application.CarLicense)
	err = self.carRepository.CreateCar(ctx, car)
	if err != nil {
		return errors.Wrap(err, "failed to create car")
	}

	driver := entity.NewDriver(entity.JumpID(jumpDriver.ID), entity.FleetID(fleetDriver.DriverProfile.ID), *application.FirstName, *application.LastName, application.MiddleName, *application.City, entity.PhoneNumber(*application.PhoneNumber), car.ID,
		entity.NewDriverLicense(*application.LicenseNumber, *application.LicenseTotalSinceDate, *application.LicenseIssueDate, *application.LicenseExpiryDate, *application.LicenseCountry))
	err = self.driverRepository.CreateDriver(ctx, driver)
	if err != nil {
		return errors.Wrap(err, "failed to create driver")
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
