package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sample/api"

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
	con, err := connectDB()
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewMux(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
