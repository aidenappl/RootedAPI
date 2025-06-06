package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aidenappl/rootedapi/env"
	"github.com/go-sql-driver/mysql"
)

const (
	DefaultListLimit = 50
	MaximumListLimit = 100

	ErNoReferencedRow     = 1215
	ErDupEntry            = 1062
	ErDupEntryWithKeyName = 1586
)

func PingDB() error {
	db, err := sql.Open("mysql", env.RootedDB)
	if err != nil {
		return err
	}

	ping := db.Ping()
	db.Close()

	return ping
}

var DB = func() *sql.DB {
	db, err := sql.Open("postgres", env.RootedDB)
	if err != nil {
		panic(fmt.Errorf("error opening database: %w", err))
	}

	return db
}()

type Queryable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func ExtractDBErrorCode(err error) uint16 {
	var sqlErr *mysql.MySQLError
	if errors.As(err, &sqlErr) {
		return sqlErr.Number
	} else {
		return 0
	}
}
