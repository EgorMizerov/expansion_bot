package postgres

import (
	"context"
	"regexp"
	"strconv"
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

	self.sqlMock.ExpectExec(self.prepareQuery(createDriverLicenseQuery)).
		WithArgs(licenseDTO.ID, licenseDTO.RegistrationCertificate, license.DrivingExperience, license.IssueDate, license.ExpiryDate, license.Country).
		WillReturnResult(sqlmock.NewResult(0, 1))
	self.sqlMock.ExpectExec(self.prepareQuery(createDriverQuery)).
		WithArgs(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.CreatedAt, driverDTO.AcceptCash, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := self.repository.CreateDriver(self.ctx, driver)

	self.NoError(err)
}

func (self *DriverRepositoryTests) TestCreateDriverFailedToInsertDriverLicense() {
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country")
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()), license)
	licenseDTO := dto.ToDriverLicenseDTO(license)

	self.sqlMock.ExpectExec(self.prepareQuery(createDriverLicenseQuery)).
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

	self.sqlMock.ExpectExec(self.prepareQuery(createDriverLicenseQuery)).
		WithArgs(licenseDTO.ID, licenseDTO.RegistrationCertificate, license.DrivingExperience, license.IssueDate, license.ExpiryDate, license.Country).
		WillReturnResult(sqlmock.NewResult(0, 1))
	self.sqlMock.ExpectExec(self.prepareQuery(createDriverQuery)).
		WithArgs(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.CreatedAt, driverDTO.AcceptCash, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID).
		WillReturnError(testError)

	err := self.repository.CreateDriver(self.ctx, driver)

	self.ErrorIs(err, testError)
}

func (self *DriverRepositoryTests) TestGetDriverByID() {
	driver := entity.NewDriver(12, "test_fleet_id", "test_first_name", "test_last_name", nil, "test_city", "test_phone", entity.CarID(uuid.New()),
		entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country"))
	driverDTO := dto.ToDriverDTO(driver)
	licenseDTO := dto.ToDriverLicenseDTO(driver.DriverLicense)

	driverRow := sqlmock.NewRows([]string{"id", "telegram_id", "fleet_id", "jump_id", "first_name", "last_name", "middle_name", "city", "phone_number", "accept_cash", "work_rule_id", "work_rule_updated_at", "is_self_employed", "car_id", "driver_license_id", "created_at"}).
		AddRow(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.AcceptCash, driverDTO.WorkRuleID, driverDTO.WorkRuleUpdatedAt, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID, driverDTO.CreatedAt)
	licenseRow := sqlmock.NewRows([]string{"id", "registration_certificate", "driving_experience", "issue_date", "expiry_date", "country"}).
		AddRow(licenseDTO.ID, licenseDTO.RegistrationCertificate, licenseDTO.DrivingExperience, licenseDTO.IssueDate, licenseDTO.ExpiryDate, licenseDTO.Country)

	self.sqlMock.ExpectQuery(self.prepareQuery(getDriverByIDQuery)).
		WithArgs(driverDTO.ID).
		WillReturnRows(driverRow)
	self.sqlMock.ExpectQuery(self.prepareQuery(getDriverLicenseByIDQuery)).
		WithArgs(licenseDTO.ID).
		WillReturnRows(licenseRow)

	result, err := self.repository.GetDriverByID(self.ctx, driver.ID)

	self.NoError(err)
	self.Equal(driver, result)
}

func (self *DriverRepositoryTests) TestGetDriverByIDFailedIfDriverNamedQueryRowContextReturnsError() {
	driver := entity.NewDriver(12, "test_fleet_id", "test_first_name", "test_last_name", nil, "test_city", "test_phone", entity.CarID(uuid.New()),
		entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country"))
	driverDTO := dto.ToDriverDTO(driver)

	self.sqlMock.ExpectQuery(self.prepareQuery(getDriverByIDQuery)).
		WithArgs(driverDTO.ID).
		WillReturnError(testError)

	_, err := self.repository.GetDriverByID(self.ctx, driver.ID)

	self.ErrorIs(err, testError)
}

func (self *DriverRepositoryTests) TestGetDriverByIDFailedIfLicenseNamedQueryRowContextReturnsError() {
	driver := entity.NewDriver(12, "test_fleet_id", "test_first_name", "test_last_name", nil, "test_city", "test_phone", entity.CarID(uuid.New()),
		entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country"))
	driverDTO := dto.ToDriverDTO(driver)
	licenseDTO := dto.ToDriverLicenseDTO(driver.DriverLicense)

	driverRow := sqlmock.NewRows([]string{"id", "telegram_id", "fleet_id", "jump_id", "first_name", "last_name", "middle_name", "city", "phone_number", "accept_cash", "work_rule_id", "work_rule_updated_at", "is_self_employed", "car_id", "driver_license_id", "created_at"}).
		AddRow(driverDTO.ID, driverDTO.TelegramID, driverDTO.FleetID, driverDTO.JumpID, driverDTO.FirstName, driverDTO.LastName, driverDTO.MiddleName, driverDTO.City, driverDTO.PhoneNumber, driverDTO.AcceptCash, driverDTO.WorkRuleID, driverDTO.WorkRuleUpdatedAt, driverDTO.IsSelfEmployed, driverDTO.CarID, driverDTO.DriverLicenseID, driverDTO.CreatedAt)

	self.sqlMock.ExpectQuery(self.prepareQuery(getDriverByIDQuery)).
		WithArgs(driverDTO.ID).
		WillReturnRows(driverRow)
	self.sqlMock.ExpectQuery(self.prepareQuery(getDriverLicenseByIDQuery)).
		WithArgs(licenseDTO.ID).
		WillReturnError(testError)

	_, err := self.repository.GetDriverByID(self.ctx, driver.ID)

	self.ErrorIs(err, testError)
}

func (self *DriverRepositoryTests) TestUpdateDriver() {
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()),
		entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country"))
	driverDTO := dto.ToDriverDTO(driver)

	self.sqlMock.ExpectExec(self.prepareQuery(updateDriverQuery)).
		WithArgs(driverDTO.TelegramID, driverDTO.PhoneNumber, driverDTO.AcceptCash, driverDTO.WorkRuleID, driverDTO.WorkRuleUpdatedAt, driverDTO.IsSelfEmployed, driverDTO.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := self.repository.UpdateDriver(self.ctx, driver)

	self.NoError(err)
}

func (self *DriverRepositoryTests) TestUpdateDriverFailedIfExecQueryReturnsError() {
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", entity.CarID(uuid.New()),
		entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "test_country"))
	driverDTO := dto.ToDriverDTO(driver)

	self.sqlMock.ExpectExec(self.prepareQuery(updateDriverQuery)).
		WithArgs(driverDTO.TelegramID, driverDTO.PhoneNumber, driverDTO.AcceptCash, driverDTO.WorkRuleID, driverDTO.WorkRuleUpdatedAt, driverDTO.IsSelfEmployed, driverDTO.ID).
		WillReturnError(testError)

	err := self.repository.UpdateDriver(self.ctx, driver)

	self.ErrorIs(err, testError)
}

func (self *DriverRepositoryTests) prepareQuery(query string) string {
	var count int
	return regexp.QuoteMeta(
		regexp.MustCompile(`:\w+`).ReplaceAllStringFunc(query, func(string) string {
			count++
			return "$" + strconv.Itoa(count)
		}))
}

func TestDriverRepositoryTests(t *testing.T) {
	suite.Run(t, new(DriverRepositoryTests))
}
