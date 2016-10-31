package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "gcardbmanage:L3xH-4G2@tcp(10.15.206.203:3306)/gocar")
	checkErr(err)

	rows, err := db.Query("SELECT pbid,name FROM gcar_brand_parent")
	checkErr(err)

	for rows.Next() {
		var pbid int
		var name string
		//var department string
		//var created string
		err = rows.Scan(&pbid, &name)
		checkErr(err)
		fmt.Println(pbid)
		fmt.Println(name)
	}

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
