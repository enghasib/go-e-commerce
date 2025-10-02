package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString() string {
	// user => ecommerce
	// password => postgres
	// host => postgres
	// port => 5432
	// db name => ecommerce

	return "user=postgres password=postgres host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := getConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("DB err:", err)
		return nil, err
	}

	return dbCon, nil

}
