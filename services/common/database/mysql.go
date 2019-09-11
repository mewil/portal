package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseType                 = "mysql"
	databaseMaxConnectionRetries = 10
	databaseConnectionRetryDelay = 10 * time.Second
)

type DB interface {
	Close() error
	Exec(query string, args ...interface{}) (Result, error)
	Ping() error
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
}

type Result interface{}

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

type Row interface {
	Scan(dest ...interface{}) error
}

func NewDatabase(user, password, host, dbName, port string) (DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=1m", user, password, host, port, dbName)
	for retries := 0; ; retries++ {
		if db, err := getDB(sourceName); err == nil {
			return db, nil
		} else if retries == databaseMaxConnectionRetries {
			return nil, fmt.Errorf("failed to connect to database after %d retries", databaseMaxConnectionRetries)
		}
		time.Sleep(databaseConnectionRetryDelay)
	}
}

func getDB(sourceName string) (DB, error) {
	db, err := sql.Open("mysql", sourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return dbAdapter{db}, nil
}

type dbAdapter struct {
	*sql.DB
}

func (db dbAdapter) Close() error {
	return db.DB.Close()
}

func (db dbAdapter) Exec(query string, args ...interface{}) (Result, error) {
	return db.DB.Exec(query, args...)
}

func (db dbAdapter) Ping() error {
	return db.DB.Ping()
}

func (db dbAdapter) Query(query string, args ...interface{}) (Rows, error) {
	return db.DB.Query(query, args...)
}

func (db dbAdapter) QueryRow(query string, args ...interface{}) Row {
	return db.DB.QueryRow(query, args...)
}
