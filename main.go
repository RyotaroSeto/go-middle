package main

import (
	"database/sql"
	"fmt"
	"sample/api"
	"sample/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "docker"
	dbHost := "mysql"
	dbPort := "3306"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 1000
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)

	api.NewMux()
}
