package postgres

import (
	"testing"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/suite"
)

type GuestRepositorySQLTests struct {
	SQLSuite

	repository *GuestRepository
}

func (self *GuestRepositorySQLTests) SetupTest() {
	self.Setup()
	self.repository = NewGuestRepository(self.db)
}

func (self *GuestRepositorySQLTests) TestGetGuestByPhoneNumber() {
	guest := self.createGuest()

	result, err := self.repository.GetGuestByPhoneNumber(self.ctx, guest.PhoneNumber)

	self.NoError(err)
	self.Equal(guest, result)
}

func (self *GuestRepositorySQLTests) TestGetGuestByPhoneNumberIfNotFound() {
	_, err := self.repository.GetGuestByPhoneNumber(self.ctx, entity.PhoneNumber("test_number"))

	self.ErrorIs(err, repository.ErrGuestNotFound)
}

func (self *GuestRepositorySQLTests) TestCreateGuest() {
	guest := &entity.Guest{
		TelegramID:  10,
		PhoneNumber: "+79957685466",
	}

	err := self.repository.CreateGuest(self.ctx, guest)

	self.NoError(err)
}

func (self *GuestRepositorySQLTests) TestCreateGuestAlreadyExists() {
	guest := &entity.Guest{
		TelegramID:  10,
		PhoneNumber: "+79957685466",
	}

	err := self.repository.CreateGuest(self.ctx, guest)
	self.NoError(err)

	err = self.repository.CreateGuest(self.ctx, guest)
	self.ErrorIs(err, repository.ErrGuestAlreadyExists)
}

func (self *GuestRepositorySQLTests) createGuest() *entity.Guest {
	guest := &entity.Guest{
		TelegramID:  10,
		PhoneNumber: "+79957685466",
	}

	err := self.repository.CreateGuest(self.ctx, guest)
	if !self.NoError(err) {
		self.T().FailNow()
	}

	return guest
}

func TestGuestRepositorySQLTests(t *testing.T) {
	suite.Run(t, new(GuestRepositorySQLTests))
}
