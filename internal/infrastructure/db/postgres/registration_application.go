package postgres

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres/dto"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

const (
	saveRegistrationApplicationQuery = `INSERT INTO registration_application (id, status, date, phone_number, last_name, first_name, middle_name, city, license_number, license_total_since_date, license_issue_date, license_expiry_date, license_country, car_brand, car_model, car_year, car_color, car_vin, car_number, car_license, work_rule_id)
		VALUES (:id, :status, :date, :phone_number, :last_name, :first_name, :middle_name, :city, :license_number, :license_total_since_date, :license_issue_date, :license_expiry_date, :license_country, :car_brand, :car_model, :car_year, :car_color, :car_vin, :car_number, :car_license, :work_rule_id)
		ON CONFLICT (id) DO UPDATE
		SET status=:status, date=:date, phone_number=:phone_number, last_name=:last_name, first_name=:first_name, middle_name=:middle_name, city=:city, license_number=:license_number, license_total_since_date=:license_total_since_date, license_issue_date=:license_issue_date, license_expiry_date=:license_expiry_date, license_country=:license_country, car_brand=:car_brand, car_model=:car_model, car_year=:car_year, car_color=:car_color, car_vin=:car_vin, car_number=:car_number, car_license=:car_license, work_rule_id=:work_rule_id`

	getRegistrationApplicationQuery = `SELECT
    	       id, status, date, phone_number, last_name, first_name, middle_name, city, license_number, license_total_since_date, license_issue_date, license_expiry_date, license_country, car_brand, car_model, car_year, car_color, car_vin, car_number, car_license, work_rule_id
		  FROM registration_application
		 WHERE id=:id`

	getRegistrationApplicationsQuery = `SELECT id, status, date, phone_number, last_name, first_name, middle_name, city, license_number, license_total_since_date, license_issue_date, license_expiry_date, license_country, car_brand, car_model, car_year, car_color, car_vin, car_number, car_license, work_rule_id
		  FROM registration_application
         WHERE status != 'closed'
      ORDER BY date DESC`
)

type RegistrationApplicationRepository struct {
	db *sqlx.DB
}

func NewRegistrationApplicationRepository(db *sqlx.DB) *RegistrationApplicationRepository {
	return &RegistrationApplicationRepository{db: db}
}

func (self *RegistrationApplicationRepository) GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error) {
	var application dto.RegistrationApplication
	sql, args, _ := sqlx.Named(getRegistrationApplicationQuery, map[string]interface{}{"id": applicationID})
	sql = sqlx.Rebind(sqlx.BindType("postgres"), sql)
	err := self.db.QueryRowxContext(ctx, sql, args...).StructScan(&application)
	return application.ToRegistrationApplication(), err
}

func (self *RegistrationApplicationRepository) GetRegistrationApplications(ctx context.Context) ([]*entity.RegistrationApplication, error) {
	rows, err := self.db.QueryxContext(ctx, getRegistrationApplicationsQuery)
	if err != nil {
		return nil, err
	}
	registrationApplications, err := StructListScan[dto.RegistrationApplication](rows)
	if err != nil {
		return nil, err
	}

	return lo.Map(registrationApplications, func(item dto.RegistrationApplication, _ int) *entity.RegistrationApplication {
		return item.ToRegistrationApplication()
	}), nil
}

func (self *RegistrationApplicationRepository) SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	_, err := self.db.NamedExecContext(ctx, saveRegistrationApplicationQuery, dto.ToRegistrationApplication(application))
	return err
}
