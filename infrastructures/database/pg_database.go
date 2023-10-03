package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
	"touchedFlowed/features/utils"
)

var Instance *pgx.Conn

func connect() {
	if Instance == nil {
		db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
		if err != nil {
			panic(err)
		}
		Instance = db
	}
}

type pgDatabase struct {
	db *pgx.Conn
}

func (d pgDatabase) Query(query string, args ...interface{}) (utils.Rows, error) {
	return d.db.Query(context.Background(), query, args...)
}

func NewPgDatabase() utils.Database {
	connect()
	return &pgDatabase{db: Instance}
}
