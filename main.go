package main

import (
	"database/sql"
	"fmt"
	"log"
	"sample/api"
	"sample/handlers"
	"sample/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "docker"
	dbHost     = "mysql"
	dbPort     = "3306"
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := handlers.NewMyAppController(ser)
	api.NewMux(con)
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
