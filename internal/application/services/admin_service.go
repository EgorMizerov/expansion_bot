package services

import (
	"context"
	"io"
	"strconv"

	"github.com/EgorMizerov/expansion_bot/internal/application/command"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/domain/repository"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
)

type AdminService struct {
	driverRepository repository.DriverRepository
	carRepository    repository.CarRepository
	fleet            *fleet.Client
	jumpClient       *jump.JumpClient
}

func NewAdminService(driverRepository repository.DriverRepository, carRepository repository.CarRepository, fleet *fleet.Client, jumpClient *jump.JumpClient) *AdminService {
	return &AdminService{driverRepository: driverRepository, carRepository: carRepository, fleet: fleet, jumpClient: jumpClient}
}

func (self *AdminService) GetCards(ctx context.Context) (*command.GetCardsCommandResult, error) {
	cards, err := self.driverRepository.GetCards(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cards")
	}

	return &command.GetCardsCommandResult{
		Cards: cards,
		TinkoffCardsCount: lo.CountBy(cards, func(item entity.CardNumber) bool {
			return item.IsTinkoff()
		}),
		AnotherCardsCount: lo.CountBy(cards, func(item entity.CardNumber) bool {
			return !item.IsTinkoff()
		}),
	}, nil
}

func (self *AdminService) GetCardsXLSX(ctx context.Context) (io.Reader, error) {
	drivers, err := self.driverRepository.GetDrivers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a cards")
	}

	//payments, err := self.jumpClient.GetPayments(ctx)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to get a payments")
	//}

	file := excelize.NewFile()
	file.SetCellValue("Sheet1", "A1", "Telegram ID")
	file.SetCellValue("Sheet1", "B1", "ФИО")
	file.SetCellValue("Sheet1", "C1", "Банк Пп/В")
	file.SetCellValue("Sheet1", "D1", "Дата Пп/В")
	file.SetCellValue("Sheet1", "E1", "Сумма Пп/М")
	file.SetCellValue("Sheet1", "F1", "Банк П/В")
	file.SetCellValue("Sheet1", "G1", "Дата П/В")
	file.SetCellValue("Sheet1", "H1", "Сумма П/М")

	for index := range drivers {
		file.SetCellValue("Sheet1", "A"+strconv.Itoa(index+2), drivers[index].TelegramID)
		file.SetCellValue("Sheet1", "C"+strconv.Itoa(index+2), drivers[index].Fullname())
		file.SetCellValue("Sheet1", "D"+strconv.Itoa(index+2), drivers[index].Fullname())
		file.SetCellValue("Sheet1", "E"+strconv.Itoa(index+2), drivers[index].Fullname())
	}

	return file.WriteToBuffer()
}
