package postgres

import (
	"testing"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/suite"
)

type DriverRepositorySQLTests struct {
	SQLSuite

	repository *DriverRepository
}

func (self *DriverRepositorySQLTests) SetupTest() {
	self.Setup()
	self.repository = NewDriverRepository(self.db)
}

func (self *DriverRepositorySQLTests) TestCreateDriver() {
	car := entity.NewCar("test_fleet_id", "test_brand", "test_model", 2012, "test_color", "test_vin", "test_number", "test_license")
	license := entity.NewDriverLicense("test_certificate", time.Unix(1, 0), time.Unix(2, 0), time.Unix(3, 0), "rus")
	driver := entity.NewDriver(111, "test_fleet_id", "test_first_name", "test_last_name", common.Point("test_middle_name"), "test_city", "test_phone", car.ID, license)
	self.createCar(car)

	err := self.repository.CreateDriver(self.ctx, driver)

	self.NoError(err)
}

func (self *DriverRepositorySQLTests) createCar(car *entity.Car) {
	err := NewCarRepository(self.db).CreateCar(self.ctx, car)
	if !self.NoError(err) {
		self.T().FailNow()
	}
}

func TestDriverRepositorySQLTests(t *testing.T) {
	suite.Run(t, new(DriverRepositorySQLTests))
}
