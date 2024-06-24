package postgres

import (
	"testing"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/suite"
)

type RegistrationApplicationRepositorySQLTests struct {
	SQLSuite

	repository *RegistrationApplicationRepository
}

func (self *RegistrationApplicationRepositorySQLTests) SetupTest() {
	self.Setup()
	self.repository = NewRegistrationApplicationRepository(self.db)
}

func (self *RegistrationApplicationRepositorySQLTests) TestSaveRegistrationApplication() {
	application := entity.NewRegistrationApplication(111, time.Now())

	err := self.repository.SaveRegistrationApplication(self.ctx, application)

	self.NoError(err)
}

func (self *RegistrationApplicationRepositorySQLTests) TestGetRegistrationApplication() {
	applicationId := entity.RegistrationApplicationID(111)
	t := time.Date(2024, 06, 23, 0, 0, 0, 0, time.UTC)
	application := entity.NewRegistrationApplication(int(applicationId), t)
	err := self.repository.SaveRegistrationApplication(self.ctx, application)
	self.NoError(err)

	result, err := self.repository.GetRegistrationApplication(self.ctx, applicationId)

	self.NoError(err)
	self.Equal(application, result)
}

func TestRegistrationApplicationRepository(t *testing.T) {
	suite.Run(t, new(RegistrationApplicationRepositorySQLTests))
}
