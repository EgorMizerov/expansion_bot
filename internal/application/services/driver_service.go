package services

import (
	"context"
	"database/sql"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	"github.com/pkg/errors"
)

type DriverService struct {
	driverRepository repository.DriverRepository
	guestRepository  repository.GuestRepository
}

func NewDriverService(driverRepository repository.DriverRepository) *DriverService {
	return &DriverService{driverRepository: driverRepository}
}

func (self *DriverService) GetDriverByPhoneNumber(ctx context.Context, phone entity.PhoneNumber) (*entity.Driver, error) {
	driver, err := self.driverRepository.GetDriverByPhoneNumber(ctx, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, interfaces.ErrDriverNotFound
		}
		return nil, err
	}

	return driver, nil
}

func (self *DriverService) GetDrivers(ctx context.Context) ([]*entity.Driver, error) {
	return self.driverRepository.GetDrivers(ctx)
}
