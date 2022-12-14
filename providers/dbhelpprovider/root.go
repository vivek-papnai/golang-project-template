package dbhelpprovider

import (
	"VivekPapnaiAtRS/template/providers"
	"github.com/jmoiron/sqlx"
)

type DBHelper struct {
	DB *sqlx.DB
}

func NewDBHelper(db *sqlx.DB) providers.DBHelpProvider {
	return &DBHelper{
		DB: db,
	}
}
