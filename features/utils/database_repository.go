package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryTx(tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error)
	ExecTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error)
	QueryRowTx(tx *sql.Tx, query string, args ...interface{}) *sql.Row
}

type database struct {
	db *sql.DB
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d database) QueryTx(tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.Query(query, args...)
}

func (d database) ExecTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return tx.Exec(query, args...)
}

func (d database) QueryRowTx(tx *sql.Tx, query string, args ...interface{}) *sql.Row {
	return tx.QueryRow(query, args...)
}

func NewDatabase() Database {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	return &database{
		db: db,
	}
}
