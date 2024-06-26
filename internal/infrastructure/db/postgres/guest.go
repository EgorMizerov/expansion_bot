package postgres

import (
	"context"
	"database/sql"
	"strings"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres/dto"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	createGuestQuery           = `INSERT INTO guests (phone_number, telegram_id) VALUES (:phone_number, :telegram_id)`
	getGuestByPhoneNumberQuery = `SELECT phone_number, telegram_id FROM guests WHERE phone_number=:phone_number`
)

type GuestRepository struct {
	db DB
}

func NewGuestRepository(db *sqlx.DB) *GuestRepository {
	return &GuestRepository{db: DB{db}}
}

func (self *GuestRepository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	_, err := self.db.NamedExecContext(ctx, createGuestQuery, dto.ToGuestDTO(guest))
	return self.error(err)
}

func (self *GuestRepository) GetGuestByPhoneNumber(ctx context.Context, number entity.PhoneNumber) (*entity.Guest, error) {
	var guest dto.GuestDTO
	err := self.db.NamedQueryRowContext(ctx, getGuestByPhoneNumberQuery, map[string]interface{}{"phone_number": number.String()}).
		StructScan(&guest)
	return guest.ToEntity(), self.error(err)
}

func (self *GuestRepository) error(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "SQLSTATE 23505") {
		return repository.ErrGuestAlreadyExists
	}
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrGuestNotFound
	}
	return err
}
