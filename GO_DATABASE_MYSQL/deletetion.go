package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {

	db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "DELETE FROM cities WHERE id IN (2, 4, 6)"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The statement affected %d rows\n", affectedRows)
}
