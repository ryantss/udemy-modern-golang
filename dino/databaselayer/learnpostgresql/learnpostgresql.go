package main

import (
	"database/sql"
	"fmt"
	"log"
)

import _ "github.com/lib/pq"

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

// https://godoc.org/github.com/lib/pq
func main() {
	//connect to the database
	// 	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"

	db, err := sql.Open("postgres", "user=ryan dbname=dino sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//general query with arguments
	rows, err := db.Query("select * from animals where age > $1", 5) // $ insterad of ?
	handlerows(rows, err)

	//query a single row
	row := db.QueryRow("select * from animals where age > $1", 5)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	//insert a row
	// result, err := db.Exec("insert into animals (animal_type, nickname, zone, age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(result.LastInsertId()) //not supported
	// fmt.Println(result.RowsAffected())

	//update a row
	/* 	result, err := db.Exec("Update animals set age = $1 where id = $2", 16, 2)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Println(result.LastInsertId()) //not supported here
		   fmt.Println(result.RowsAffected())
	*/
	/*
		var id int
		db.QueryRow("Update animals set age = $1 where id = $2 returning id", 16, 2).Scan(&id)
		fmt.Println("Affected object id:", id)
	*/

	//prepare queries to be used multiple times
	fmt.Println("Prepared Statements...")
	stmt, err := db.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//let's try with age > 5
	rows, err = stmt.Query(5)
	handlerows(rows, err)

	//let's try with age > 10
	rows, err = stmt.Query(10)
	handlerows(rows, err)

	testTransaction(db)

}

func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(animals)
}

func testTransaction(db *sql.DB) {
	fmt.Println("Transactions...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(15)
	handlerows(rows, err)
	handlerows(stmt.Query(17))
	results, err := tx.Exec("update animals set age = $1 where id = $2", 18, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.RowsAffected())
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
