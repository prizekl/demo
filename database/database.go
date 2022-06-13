package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	if ping := db.Ping(); ping != nil {
		return nil, err
	}

	fmt.Println("Connected to database server")
	return db, nil
}
