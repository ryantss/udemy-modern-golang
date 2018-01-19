package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type animal struct {
	// ID         int    `bson:"id"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//
	animalcollection := session.DB("Dino").C("animals")
	// animals := []interface{}{
	// 	animal{
	// 		AnimalType: "Tyrannosaurus rex",
	// 		Nickname:   "rex",
	// 		Zone:       1,
	// 		Age:        11,
	// 	},
	// 	animal{
	// 		AnimalType: "Velociraptor",
	// 		Nickname:   "rapto",
	// 		Zone:       2,
	// 		Age:        15,
	// 	},
	// 	animal{
	// 		AnimalType: "Velociraptor",
	// 		Nickname:   "rapto",
	// 		Zone:       2,
	// 		Age:        9,
	// 	},
	// }

	// //create document
	// err = animalcollection.Insert(animals...)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//update document
	// err = animalcollection.Update(bson.M{"nickname": "rapto"}, bson.M{"$set": bson.M{"age": 18}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//delete document
	// err = animalcollection.Remove(bson.M{"nickname": "rapto"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//query
	query := bson.M{
		"age": bson.M{
			"$gt": 5},
		"zone": bson.M{
			"$in": []int{1, 2},
		},
	}
	results := []animal{}
	animalcollection.Find(query).All(&results) //.one
	fmt.Println(results)

	// animalcollection.Find(bson.M{}).All(&results) //.one
	// fmt.Println("return all", results)

	//one result
	result := animal{}
	animalcollection.Find(query).One(&result)
	fmt.Println(result)
}
