package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	fleet2 "github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2"
	"github.com/EgorMizerov/expansion_bot/migrations"
	tele "github.com/EgorMizerov/telebot"
	"github.com/caarlos0/env/v11"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	"github.com/EgorMizerov/expansion_bot/config"
	"github.com/EgorMizerov/expansion_bot/internal/application/services"
	"github.com/EgorMizerov/expansion_bot/internal/database"
	postgres2 "github.com/EgorMizerov/expansion_bot/internal/infrastructure/db/postgres"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump"
	redis2 "github.com/EgorMizerov/expansion_bot/internal/infrastructure/redis"
	"github.com/EgorMizerov/expansion_bot/internal/interface/rest"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot"
	"github.com/EgorMizerov/expansion_bot/logger"
)

func main() {
	godotenv.Load()
	// Config
	config, err := env.ParseAs[config.Config]()
	if err != nil {
		panic("faield to parse env: " + err.Error())
	}

	// Logger
	logger := logger.Logger(logger.LogType(config.LogType), logger.LogLevel(config.LogLevel))

	// Bot
	teleBot, err := tele.NewBot(tele.Settings{
		Token:     config.BotToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	})
	if err != nil {
		logger.Error("failed to init bot", slog.Any("error", err))
		return
	}

	// Database
	db, err := database.ConnectToDatabase(database.ConnectConfig{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Database: config.PostgresDatabase,
	})
	if err != nil {
		logger.Error("failed to connect to database", slog.Any("error", err))
		return
	}

	if err = migrations.Migrate(db, "./migrations"); err != nil {
		logger.Error("failed to up migrations", slog.Any("error", err))
		return
	}

	// Dependencies
	fleetClient := fleet.NewClient(config.FleetHost, config.FleetParkID, config.FleetClientID, config.FleetAPIKey)
	fleet2Client := fleet2.NewFleetClient(config.FleetHost, config.FleetParkID, config.FleetClientID, config.FleetAPIKey)
	jumpClient := jump.NewJumpClient(config.JumpTaxiHost, config.JumpTaxiClientKey)
	redisClient := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)})

	if redisClient.Ping().Err() != nil {
		panic(err)
	}
	stateMachine := redis2.NewStateMachine(redisClient)

	adminId, _ := strconv.Atoi(config.AdminID)
	bot := telebot.NewBot(teleBot, stateMachine, int64(adminId))

	driverRepository := postgres2.NewDriverRepository(db)
	carRepository := postgres2.NewCarRepository(db)
	registrationApplicationRepository := postgres2.NewRegistrationApplicationRepository(db)

	adminService := services.NewAdminService(driverRepository, carRepository, fleetClient, jumpClient)
	registrationApplicationService := services.NewRegistrationApplicationService(fleet2Client, jumpClient, registrationApplicationRepository, driverRepository, carRepository)
	telebot.NewAdminHandler(bot, stateMachine, adminService)
	telebot.NewGuestHandler(bot)

	go func() {
		err := http.ListenAndServe("localhost:8080", rest.NewJumpWebhook(registrationApplicationService))
		if err != nil {
			panic(err)
		}
	}()

	logger.Info("start polling...")
	bot.Start()
}
