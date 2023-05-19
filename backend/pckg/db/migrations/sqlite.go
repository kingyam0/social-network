package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DB global variable to hold db connection
var m *migrate.Migrate

// run migrations
func RunMigration(sourceUrl, databaseUrl string) *migrate.Migrate {
	m, err := migrate.New(sourceUrl, databaseUrl)
	if err != nil {
		fmt.Print(err.Error())
	}

	m.Up()

	return m
}

// remove Migrations
func RemoveMigrations(m *migrate.Migrate) {
	m.Down()
}

// ConnectDB opens a conncection to the database
func ConnectDB(filename string) *sql.DB {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		fmt.Println("panic +++++++++++++++ LINE 16 ConnectDB.go")
		panic(err.Error())
	}

	return db
}

// CloseDB will close database when needed
// func CloseDB() {
// 	DB.Close()
// }
