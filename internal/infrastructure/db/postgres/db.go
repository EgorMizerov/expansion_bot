package postgres

import (
	"context"
	"database/sql"
	"unsafe"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

type DB struct {
	*sqlx.DB
}

func (self *DB) NamedQueryRowContext(ctx context.Context, query string, arg interface{}) *sqlx.Row {
	rows, err := self.NamedQueryContext(ctx, query, arg)
	if err != nil {
		return self.rowWithError(err)
	}
	return self.rowFromRows(rows)
}

type myRow struct {
	err    error
	unsafe bool
	rows   *sql.Rows
	Mapper *reflectx.Mapper
}

func (self *DB) rowFromRows(rows *sqlx.Rows) *sqlx.Row {
	rowDst := new(sqlx.Row)
	rowSrc := &myRow{
		err:    rows.Err(),
		unsafe: false,
		rows:   rows.Rows,
		Mapper: self.DB.Mapper,
	}

	copy(
		(*(*[unsafe.Sizeof(sqlx.Row{})]byte)(unsafe.Pointer(rowDst)))[:],
		(*(*[unsafe.Sizeof(myRow{})]byte)(unsafe.Pointer(rowSrc)))[:],
	)

	return rowDst
}

func (self *DB) rowWithError(err error) *sqlx.Row {
	rowDst := new(sqlx.Row)
	rowSrc := &myRow{err: err}

	copy(
		(*(*[unsafe.Sizeof(sqlx.Row{})]byte)(unsafe.Pointer(rowDst)))[:],
		(*(*[unsafe.Sizeof(myRow{})]byte)(unsafe.Pointer(rowSrc)))[:],
	)

	return rowDst
}

func StructListScan[T any](rows *sqlx.Rows) ([]T, error) {
	var items []T
	for rows.Next() {
		var item T
		err := rows.StructScan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
