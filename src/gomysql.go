package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"	
	"database/sql"
)

func main(){
	fmt.Println("go Mysql tutorial")
	db, err := sql.Open("mysql","root@Rahas!@20!8@tcp(192.168.24.70:3306)/ipd_management")

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(db)

	defer db.Close()

	fmt.Println("okee")
}