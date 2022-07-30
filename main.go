package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pranav"
	password = "pranav"
	dbname   = "go_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// --- Inserting a record
	// sqlQuery := `
	//  INSERT INTO users(first_name, last_name, age, email)
	//  VALUES($1, $2, $3, $4);
	//  `
	//
	// _, err = db.Exec(sqlQuery, "aaron", "stone", 28, "aaronsavior@mail.com")
	// if err != nil {
	// 	panic(err)
	// }

	// --- Inserting a record and returns the ID
	// sqlQuery1 := `
	//  INSERT INTO users (first_name, last_name, age, email)
	//  VALUES ($1, $2, $3, $4)
	//  RETURNING id;
	//  `
	//
	// id := 0
	// err = db.QueryRow(sqlQuery1, "ruby", "margaret", 15, "ruby@io.com").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("new record id:", id)

	// --- Updating a record
	// sqlQuery2 := `UPDATE users SET age = $2 WHERE id = $1`
	// _, err = db.Exec(sqlQuery2, 3, 18)
	// if err != nil {
	//   panic(err)
	// }

	// -- Deleting a record
	sqlQuery3 := `delete from users where id = $1`
	res, err := db.Exec(sqlQuery3, 4)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Rows affected: ", count)
}
