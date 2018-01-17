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

	//query a single row
	row := db.QueryRow("select * from Dino.animals where age > ?", 10)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	//insert a row
	/* 	result, err := db.Exec("insert into animals (animal_type, nickname, zone, age) values ('Carnotaurus', 'Carno', 3, 22)")
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	*/

	//update a row
	result, err := db.Exec("Update Dino.animals set age = ? where id=?", 16, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
