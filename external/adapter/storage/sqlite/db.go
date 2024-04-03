package sqlite

import (
	"database/sql"

	"github.com/khanhtranrk/justice-alpha/justice_chain/adapter/config"

	_ "github.com/mattn/go-sqlite3"
)

func New(config *config.DB) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", "./litego.db")

  if err != nil {
    return nil, err
  }

  return db, nil
}
