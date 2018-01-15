package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {

	var source string

	source = database_config.username +
		":" +
		database_config.password +
		"@tcp(" +
		database_config.host +
		":" +
		database_config.port +
		")/" +
		database_config.db_name

	db, _ := sql.Open("mysql", source)

	if db.Ping() != nil {

		return nil, db.Ping()
	}

	return db, nil
}
