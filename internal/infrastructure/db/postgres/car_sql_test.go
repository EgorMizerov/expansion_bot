package postgres

import (
	"testing"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/suite"
)

type CarRepositoryTests struct {
	SQLSuite

	repository *CarRepository
}

func (self *CarRepositoryTests) SetupTest() {
	self.Setup()
	self.repository = NewCarRepository(self.db)
}

func (self *CarRepositoryTests) TestCreateCar() {
	car := entity.NewCar("test_fleet_id", "test_brand", "test_model", 2012, "test_color", "test_vin", "test_number", "test_license")

	err := self.repository.CreateCar(self.ctx, car)

	self.NoError(err)
}

func (self *CarRepositoryTests) TestGetCarByID() {
	car := entity.NewCar("test_fleet_id", "test_brand", "test_model", 2012, "test_color", "test_vin", "test_number", "test_license")
	err := self.repository.CreateCar(self.ctx, car)
	self.NoError(err)

	result, err := self.repository.GetCarByID(self.ctx, car.ID)

	self.NoError(err)
	self.Equal(car, result)
}

func TestCarRepositoryTests(t *testing.T) {
	suite.Run(t, new(CarRepositoryTests))
}
