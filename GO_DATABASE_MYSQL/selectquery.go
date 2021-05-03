package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/tuts")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

}

// var tag Tag
// // Execute the query
// err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
// if err != nil {
//     panic(err.Error()) // proper error handling instead of panic in your app
// }

// log.Println(tag.ID)
// log.Println(tag.Name)

// with the exec we can insert query
// sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"
//     res, err := db.Exec(sql)

//     if err != nil {
//         panic(err.Error())
//     }

//     lastId, err := res.LastInsertId()

//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("The last inserted row id: %d\n", lastId)
