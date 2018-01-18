package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	gorm.Model
	// ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	AnimalType string `gorm:type:TEXT"`
	Nickname   string `gorm:type:TEXT"`
	Zone       int    `gorm:type:INTEGER"`
	Age        int
}

// https://godoc.org/github.com/lib/pq
func main() {
	//connect to the database
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})
	db.Table("dinos").DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{}) // will add any missing fields, will add 's' to the struct name
	db.Table("dinos").CreateTable(&animal{})

	a := animal{
		AnimalType: "Tyrannosaurus rex",
		Nickname:   "rex",
		Zone:       1,
		Age:        11,
	}
	db.Debug().Create(&a) //vs create()
	db.Table("dinos").Create(&a)

	a = animal{
		AnimalType: "Velociraptor",
		Nickname:   "rapto",
		Zone:       2,
		Age:        15,
	}
	db.Save(&a)

	//updates
	// db.Table("animals").Where("nickname = ? and zone = ?", "rapto", 2).Update("age", 16)

	//queries
	animals := []animal{}
	db.Find(&animals, "age > ?", 12) //first
	fmt.Println(animals)
	dinoAnimals := []animal{}
	db.Debug().Table("dinos").Find(&dinoAnimals, "age > ?", 10) //first
	fmt.Println(dinoAnimals)

	// err = insertSampleData(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	/*
		//general query with arguments
		rows, err := db.Query("select * from animals where age > $1", 5) // both $ and ? are supported
		handlerows(rows, err)

		//query a single row
		row := db.QueryRow("select * from animals where age > ?", 5)
		a := animal{}
		err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a)

		//insert a row
		/* 	result, err := db.Exec("insert into animals (animal_type, nickname, zone, age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
		   	if err != nil {
		   		log.Fatal(err)
		   	}

		   	fmt.Println(result.LastInsertId())
		   	fmt.Println(result.RowsAffected())
	*/
	//update a row
	/* result, err := db.Exec("Update animals set age = $1 where id = $2", 16, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId()) //not supported here
	fmt.Println(result.RowsAffected()) */

	/*
		var id int
		db.QueryRow("Update animals set age = $1 where id = $2 returning id", 16, 2).Scan(&id)
		fmt.Println("Affected object id:", id)
	//*/

	//prepare queries to be used multiple times
	// fmt.Println("Prepared Statements...")
	// stmt, err := db.Prepare("select * from animals where age > $1")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// //let's try with age > 5
	// rows, err = stmt.Query(5)
	// handlerows(rows, err)

	// //let's try with age > 10
	// rows, err = stmt.Query(10)
	// handlerows(rows, err)

	// testTransaction(db)

}

// func handlerows(rows *sql.Rows, err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	animals := []animal{}
// 	for rows.Next() {
// 		a := animal{}
// 		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		animals = append(animals, a)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(animals)
// }

// func testTransaction(db *sql.DB) {
// 	fmt.Println("Transactions...")
// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tx.Rollback()
// 	stmt, err := tx.Prepare("select * from animals where age > $1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(15)
// 	handlerows(rows, err)
// 	handlerows(stmt.Query(17))
// 	results, err := tx.Exec("update animals set age = $1 where id = $2", 18, 2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(results.RowsAffected())
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func createDatabase(db *sql.DB) error {
// 	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS
// 			animals(id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			animal_type TEXT,
// 			nickname TEXT,
// 			zone INTEGER,
// 			age INTEGER)`)
// 	return err
// }

// func insertSampleData(db *sql.DB) error {
// 	_, err := db.Exec(`insert into animals (animal_type, nickname, zone, age)
// 	values ('Tyrannosaurus rex', 'rex', 1, 10), ('Velociraptor', 'rapto',2, 15)`)
// 	return err
// }
