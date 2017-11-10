package models

import (
	"database/sql"

	_ "github.com/lib/pq" //PostgreSQL Driver
)

// ConnectToDb - Initializes Connection to the postgres DB
func ConnectToDb() *sql.DB {
	connStr := `host=qdjjtnkv.db.elephantsql.com 
				user=birlcvav 
				password=19QZsLOLiARafPCZOatrZobWlRYF5XGL 
				dbname=birlcvav`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
