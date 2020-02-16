package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func Init() {
	var err error
	Client, _ = sql.Open("mysql", "root:vishal1132@(172.17.0.2:3306)/urlshortener?parseTime=true")
	err = Client.Ping()
	if err != nil {
		//sql connection unsuccessful and this should be fatal users cannot continue without this
		log.Println("error in pinging mysql database with error ", err)
	} else {
		//continue the application
		log.Println("Mysql Connection established successfully and ping to database successful")
	}
}
