package postgres

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres/dto"
	"github.com/google/uuid"
	"github.com/samber/lo"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	createDriverQuery = `INSERT INTO drivers 
		(id, telegram_id, fleet_id, jump_id, first_name, last_name, middle_name, city, phone_number, created_at, accept_cash, is_self_employed, car_id, driver_license_id, work_rule_id, work_rule_updated_at)
		VALUES (:driver_id, :telegram_id, :fleet_id, :jump_id, :first_name, :last_name, :middle_name, :city, :phone_number, :created_at, :accept_cash, :is_self_employed, :car_id, :driver_license_id, :work_rule_id, :work_rule_updated_at)`

	createDriverLicenseQuery = `INSERT INTO driver_license
		(id, registration_certificate, driving_experience, issue_date, expiry_date, country)
		VALUES (:license_id, :registration_certificate, :driving_experience, :issue_date, :expiry_date, :country)`

	updateDriverQuery = `UPDATE drivers SET
		telegram_id=:telegram_id, phone_number=:phone_number, accept_cash=:accept_cash, work_rule_id=:work_rule_id, work_rule_updated_at=:work_rule_updated_at, is_self_employed=:is_self_employed
		WHERE id=:driver_id`

	getDriversQuery = `SELECT d.id AS driver_id, d.telegram_id, d.fleet_id, d.jump_id, d.first_name, d.last_name, d.middle_name, d.city, d.phone_number, d.created_at, d.accept_cash, d.is_self_employed, d.car_id, d.driver_license_id, dl.id AS license_id, dl.registration_certificate, dl.driving_experience, dl.issue_date, dl.expiry_date, dl.country 
		FROM drivers AS d
		JOIN driver_license AS dl ON d.driver_license_id = dl.id`

	getDriverByIDQuery          = `SELECT id AS driver_id, telegram_id, fleet_id, jump_id, first_name, last_name, middle_name, city, phone_number, created_at, accept_cash, is_self_employed, car_id, driver_license_id FROM drivers WHERE id=:driver_id`
	getDriverByPhoneNumberQuery = `SELECT * FROM drivers WHERE phone_number=:phone_number`
	getDriverLicenseByIDQuery   = `SELECT * FROM driver_license WHERE id=:license_id`
)

type DriverRepository struct {
	db DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: DB{db}}
}

func (self *DriverRepository) CreateDriver(ctx context.Context, driver *entity.Driver) error {
	_, err := self.db.NamedExecContext(ctx, createDriverLicenseQuery, dto.ToDriverLicenseDTO(driver.DriverLicense))
	if err != nil {
		return errors.Wrap(err, "failed to insert to driver_license table")
	}
	_, err = self.db.NamedExecContext(ctx, createDriverQuery, dto.ToDriverDTO(driver))
	if err != nil {
		return errors.Wrap(err, "failed to insert to drivers table")
	}
	return nil
}

func (self *DriverRepository) GetDriverByID(ctx context.Context, driverId entity.DriverID) (*entity.Driver, error) {
	var driver dto.DriverDTO
	var license dto.DriverLicenseDTO
	if err := self.db.NamedQueryRowContext(ctx, getDriverByIDQuery, map[string]interface{}{"id": uuid.UUID(driverId)}).
		StructScan(&driver); err != nil {
		return nil, err
	}
	if err := self.db.NamedQueryRowContext(ctx, getDriverLicenseByIDQuery, map[string]interface{}{"id": driver.DriverLicenseID}).
		StructScan(&license); err != nil {
		return nil, err
	}
	return driver.ToDriver(license.ToDriverLicense()), nil
}

func (self *DriverRepository) GetDriverByPhoneNumber(ctx context.Context, phone entity.PhoneNumber) (*entity.Driver, error) {
	var driver dto.DriverDTO
	var license dto.DriverLicenseDTO
	if err := self.db.NamedQueryRowContext(ctx, getDriverByPhoneNumberQuery, map[string]interface{}{"phone_number": string(phone)}).
		StructScan(&driver); err != nil {
		return nil, err
	}
	if err := self.db.NamedQueryRowContext(ctx, getDriverLicenseByIDQuery, map[string]interface{}{"id": driver.DriverLicenseID}).
		StructScan(&license); err != nil {
		return nil, err
	}
	return driver.ToDriver(license.ToDriverLicense()), nil
}

func (self *DriverRepository) GetDrivers(ctx context.Context) ([]*entity.Driver, error) {
	rows, err := self.db.QueryxContext(ctx, getDriversQuery)
	if err != nil {
		return nil, err
	}
	drivers, err := StructListScan[dto.DriverWithLicenseDTO](rows)
	if err != nil {
		return nil, err
	}

	return lo.Map(drivers, func(item dto.DriverWithLicenseDTO, _ int) *entity.Driver {
		return item.ToEntity()
	}), nil
}

func (self *DriverRepository) UpdateDriver(ctx context.Context, driver *entity.Driver) error {
	_, err := self.db.NamedExecContext(ctx, updateDriverQuery, dto.ToDriverDTO(driver))
	return err
}

func (self *DriverRepository) GetCards(ctx context.Context) ([]entity.CardNumber, error) {
	sql, _ := sq.Select("card_number").From("drivers").MustSql()
	rows, err := self.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	var cards []entity.CardNumber
	for rows.Next() {
		var card entity.CardNumber
		err = rows.Scan(&card)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan rows")
		}
		cards = append(cards, card)
	}

	return cards, nil
}
