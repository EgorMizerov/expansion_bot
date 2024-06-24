package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

var testError = errors.New("test_error")

type DriverRepositoryTests struct {
	suite.Suite

	repository *DriverRepository
	sqlMock    sqlmock.Sqlmock
	ctx        context.Context
}

func (self *DriverRepositoryTests) SetupTest() {
	db, mock, err := sqlmock.New()
	if !self.NoError(err) {
		self.T().FailNow()
	}

	self.repository = NewDriverRepository(sqlx.NewDb(db, "pgx"))
	self.sqlMock = mock
	self.ctx = context.Background()
}

func (self *DriverRepositoryTests) TearDownTest() {
	self.NoError(self.sqlMock.ExpectationsWereMet())
}

func (self *DriverRepositoryTests) TestCreateDriver() {
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "rus")
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()), license)
	driverDTO := dto.ToDriverDTO(driver)
	licenseDTO := dto.ToDriverLicenseDTO(license)

	self.sqlMock.ExpectExec("INSERT INTO driver_license").
		WithArgs(licenseDTO.ID, licenseDTO.RegistrationCertificate, license.DrivingExperience, license.IssueDate, license.ExpiryDate, license.Country).
		WillReturnResult(sqlmock.NewResult(0, 1))
	self.sqlMock.ExpectExec("INSERT INTO drivers").
		WithArgs(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.CreatedAt, driverDTO.AcceptCash, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := self.repository.CreateDriver(self.ctx, driver)

	self.NoError(err)
}

func (self *DriverRepositoryTests) TestCreateDriverFailedToInsertDriverLicense() {
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country")
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()), license)
	licenseDTO := dto.ToDriverLicenseDTO(license)

	self.sqlMock.ExpectExec("INSERT INTO driver_license").
		WithArgs(licenseDTO.ID, licenseDTO.RegistrationCertificate, license.DrivingExperience, license.IssueDate, license.ExpiryDate, license.Country).
		WillReturnError(testError)

	err := self.repository.CreateDriver(self.ctx, driver)

	self.ErrorIs(err, testError)
}

func (self *DriverRepositoryTests) TestCreateDriverFailedToInsertDriver() {
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country")
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()), license)
	licenseDTO := dto.ToDriverLicenseDTO(license)
	driverDTO := dto.ToDriverDTO(driver)

	self.sqlMock.ExpectExec("INSERT INTO driver_license").
		WithArgs(licenseDTO.ID, licenseDTO.RegistrationCertificate, license.DrivingExperience, license.IssueDate, license.ExpiryDate, license.Country).
		WillReturnResult(sqlmock.NewResult(0, 1))
	self.sqlMock.ExpectExec("INSERT INTO drivers").
		WithArgs(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.CreatedAt, driverDTO.AcceptCash, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID).
		WillReturnError(testError)

	err := self.repository.CreateDriver(self.ctx, driver)

	self.ErrorIs(err, testError)
}

func TestDriverRepositoryTests(t *testing.T) {
	suite.Run(t, new(DriverRepositoryTests))
}
