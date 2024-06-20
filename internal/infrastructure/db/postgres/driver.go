package postgres

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DriverRepository struct {
	db *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (self *DriverRepository) CreateDriver(ctx context.Context, driver *entity.Driver) error {
	_, err := sq.Insert("drivers").
		Columns("id", "telegram_id", "first_name", "last_name", "middle_name", "address", "phone_number", "card_number", "referral_key", "created_at", "accept_cash", "is_self_employed", "work_rule_id", "work_rule_updated_at", "car_id").
		Values(driver.ID, driver.TelegramID, driver.FirstName, driver.LastName, driver.MiddleName, driver.Address, driver.PhoneNuber, driver.CardNumber, driver.ReferralKey, driver.CreatedAt, driver.AcceptCash, driver.IsSelfEmployed, driver.WorkRule.ID, driver.WorkRuleUpdatedAt, driver.CarID).
		PlaceholderFormat(sq.Dollar).RunWith(self.db).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to insert to user table")
	}

	_, err = sq.Insert("driver_license").
		Columns("driver_id", "registration_certificate", "driving_experience", "issue_date", "expiry_date", "country").
		Values(driver.ID, driver.DriverLicense.RegistrationCertificate, driver.DriverLicense.DrivingExperience, driver.DriverLicense.IssueDate, driver.DriverLicense.ExpiryDate, driver.DriverLicense.Country).
		PlaceholderFormat(sq.Dollar).RunWith(self.db).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to insert to driver_license table")
	}

	return nil
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

func (self *DriverRepository) GetDrivers(ctx context.Context) ([]*entity.Driver, error) {
	sql, _ := sq.Select("d.id", "d.telegram_id", "d.first_name", "d.last_name", "d.middle_name", "d.address", "d.phone_number", "d.card_number", "d.referral_key", "d.accept_cash", "d.work_rule_id", "d.work_rule_updated_at", "d.car_id", "d.created_at", "dl.driver_id", "dl.registration_certificate", "dl.driving_experience", "dl.issue_date", "dl.expiry_date", "dl.country").
		From("drivers AS d").Join("driver_license AS dl ON d.id = dl.driver_id").MustSql()

	rows, err := self.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute a query")
	}

	var drivers []*entity.Driver
	for rows.Next() {
		var driver entity.Driver
		err = rows.StructScan(&driver)
		if err != nil {
			return nil, errors.Wrap(err, "failed to struct scan")
		}
		drivers = append(drivers, &driver)
	}

	return drivers, nil
}
