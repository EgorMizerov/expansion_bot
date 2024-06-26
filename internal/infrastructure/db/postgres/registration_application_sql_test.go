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

func (self *RegistrationApplicationRepositorySQLTests) TestGetRegistrationApplications() {
	applicationId := entity.RegistrationApplicationID(111)
	t := time.Date(2024, 06, 23, 0, 0, 0, 0, time.UTC)
	application := entity.NewRegistrationApplication(int(applicationId), t)
	err := self.repository.SaveRegistrationApplication(self.ctx, application)
	self.NoError(err)

	result, err := self.repository.GetRegistrationApplications(self.ctx)

	self.NoError(err)
	self.Contains(result, application)
}

func (self *RegistrationApplicationRepositorySQLTests) TestGetRegistrationApplicationsIfNoRows() {
	result, err := self.repository.GetRegistrationApplications(self.ctx)

	self.NoError(err)
	self.Nil(result)
}

func (self *RegistrationApplicationRepositorySQLTests) TestGetRegistrationApplicationsWhereNotRegistered() {
	registeredApplication := entity.NewRegistrationApplication(int(entity.RegistrationApplicationID(111)), time.Date(2023, 06, 23, 0, 0, 0, 0, time.UTC))
	registeredApplication.SetStatus("registered")
	newApplication := entity.NewRegistrationApplication(int(entity.RegistrationApplicationID(2222)), time.Date(2024, 06, 23, 0, 0, 0, 0, time.UTC))
	newApplication.SetStatus("new")
	err := self.repository.SaveRegistrationApplication(self.ctx, registeredApplication)
	self.NoError(err)
	err = self.repository.SaveRegistrationApplication(self.ctx, newApplication)
	self.NoError(err)

	result, err := self.repository.GetRegistrationApplications(self.ctx)

	self.NoError(err)
	self.Equal(result, []*entity.RegistrationApplication{newApplication})
}

func (self *RegistrationApplicationRepositorySQLTests) TestGetRegistrationApplicationsOrderBy() {
	oldApplication := entity.NewRegistrationApplication(int(entity.RegistrationApplicationID(111)), time.Date(2023, 06, 23, 0, 0, 0, 0, time.UTC))
	newApplication := entity.NewRegistrationApplication(int(entity.RegistrationApplicationID(2222)), time.Date(2024, 06, 23, 0, 0, 0, 0, time.UTC))
	err := self.repository.SaveRegistrationApplication(self.ctx, oldApplication)
	self.NoError(err)
	err = self.repository.SaveRegistrationApplication(self.ctx, newApplication)
	self.NoError(err)

	result, err := self.repository.GetRegistrationApplications(self.ctx)

	self.NoError(err)
	self.Equal(result, []*entity.RegistrationApplication{newApplication, oldApplication})
}

func TestRegistrationApplicationRepository(t *testing.T) {
	suite.Run(t, new(RegistrationApplicationRepositorySQLTests))
}
