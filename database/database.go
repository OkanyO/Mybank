package database

import (
	"nuel/go-bank/helpers"

	"database/sql"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
	Accounts []Account
}

type Account struct {
	ID      int
	Name    string
	Balance int
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func dbCall(query string) *sql.Rows {
	db := connectDB()

	call, err := db.Query(query)

	helpers.HandleErr(err)
	return call
}

var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("host=localhost, port=5432, user=nuel, dbname=bankapp, password=123, sslmode=disable")
	helpers.HandleErr(err)
	// Set up connection pool

	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
