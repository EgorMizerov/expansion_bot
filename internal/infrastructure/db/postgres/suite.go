package postgres

import (
	"context"
	"regexp"
	"strconv"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	sqlMock sqlmock.Sqlmock
	db      *sqlx.DB
	ctx     context.Context
}

func (self *Suite) Setup() {
	db, mock, err := sqlmock.New()
	if !self.NoError(err) {
		self.T().FailNow()
	}

	self.db = sqlx.NewDb(db, "pgx")
	self.sqlMock = mock
	self.ctx = context.Background()
}

func (self *Suite) prepareQuery(query string) string {
	var count int
	return regexp.QuoteMeta(
		regexp.MustCompile(`:\w+`).ReplaceAllStringFunc(query, func(string) string {
			count++
			return "$" + strconv.Itoa(count)
		}))
}
