package database

import (
	"database/sql"
	"fmt"
	"schedule/config"
)

var Db *sql.DB

func Init() {
	conSte := "root:" + config.DatabasePasswd + "@tcp(127.0.0.1:3306)/nft"
	var err error
	Db, err = sql.Open("mysql", conSte)
	if err != nil {
		fmt.Print(err)
	}
}
