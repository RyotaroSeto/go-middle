package main

import (
	"sample/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	api.NewMux()
}
