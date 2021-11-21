package db

import (
	"fmt"
	"github.com/MihailChapenko/chat/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

func Init(cfg *config.DB) {
	db, err := sqlx.Connect(cfg.Dialect, cfg.DataSource)
	if err != nil {
		fmt.Println("Database err: ", err)
		return
	}

	conn = db
}

// GetDB get database connection
func GetDB() *sqlx.DB {
	return conn
}
