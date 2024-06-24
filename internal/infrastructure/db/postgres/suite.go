package postgres

import (
	"context"
	"time"

	"github.com/EgorMizerov/expansion_bot/migrations"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type SQLSuite struct {
	suite.Suite

	ctx context.Context
	db  *sqlx.DB
}

func (self *SQLSuite) Setup() {
	self.ctx = context.Background()

	pgContainer, err := postgres.RunContainer(self.ctx,
		testcontainers.WithImage("postgres:latest"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if !self.NoError(err) {
		self.T().FailNow()
	}

	self.T().Cleanup(func() {
		if err := pgContainer.Terminate(self.ctx); err != nil {
			self.T().Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connectionStr, err := pgContainer.ConnectionString(self.ctx)
	if !self.NoError(err) {
		self.T().FailNow()
	}

	db, err := sqlx.Connect("pgx", connectionStr)
	if !self.NoError(err) {
		self.T().FailNow()
	}

	self.db = db

	err = migrations.Migrate(db, "./../../../../migrations")
	if !self.NoError(err) {
		self.T().FailNow()
	}
}
