package main

import (
	"fmt"
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

func main() {
	// client.NewUDPClient() but does not support query. typicall is used for insert data
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := queryDB(c, "dino", `select * from weightmeasures where "animal_type" = 'Velociraptor'`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Begin to loop through result")

	for _, v := range res {
		log.Println("messages: ", v.Messages)
		for _, s := range v.Series {
			log.Println("series name:", s.Name)
			log.Println("series columns:", s.Columns)
			log.Println("series values size:", len(s.Values))
			for _, row := range s.Values {
				log.Println(row)
			}
			// log.Println("series values:", s.Values)
		}
	}

}

func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	response, err := c.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error()
	}

	return response.Results, nil
}
