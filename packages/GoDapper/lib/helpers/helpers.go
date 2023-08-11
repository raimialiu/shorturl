package helpers

import (
	"database/sql"
	"fmt"
)

func ConnectToDb(dialect, connectionString string) (db *sql.DB) {
	db, err := sql.Open(dialect, connectionString)
	if err != nil {
		fmt.Println(db)
		return db
	}
	panic(err)
}
