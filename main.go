package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		// do something here
	}
	users := make([]User, 0)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(user.Id, user.Name)
		users = append(users, user)
	}
	rows.Close()
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
	defer db.Close()
}

type User struct {
	Id   int
	Name string
}
