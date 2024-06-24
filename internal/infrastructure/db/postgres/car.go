package postgres

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres/dto"
	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type CarRepository struct {
	db *sqlx.DB
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (self *CarRepository) GetCarByID(ctx context.Context, id entity.CarID) (*entity.Car, error) {
	var car dto.CarDTO
	sql, args := sq.Select("id", "fleet_id", "brand", "model", "year", "color", "vin", "license_plate_number", "registration_certificate").
		From("cars").
		Where(sq.Eq{"id": uuid.UUID(id)}).
		PlaceholderFormat(sq.Dollar).MustSql()

	err := self.db.QueryRowxContext(ctx, sql, args...).StructScan(&car)
	return car.ToCar(), err
}

func (self *CarRepository) CreateCar(ctx context.Context, car *entity.Car) error {
	_, err := sq.Insert("cars").
		Columns("id", "fleet_id", "brand", "model", "year", "color", "vin", "license_plate_number", "registration_certificate").
		Values(uuid.UUID(car.ID), car.FleetID, car.Brand, car.Model, car.Year, car.Color, car.VIN, car.LicensePlateNumber, car.RegistrationCertificate).
		PlaceholderFormat(sq.Dollar).RunWith(self.db).ExecContext(ctx)
	return err
}
