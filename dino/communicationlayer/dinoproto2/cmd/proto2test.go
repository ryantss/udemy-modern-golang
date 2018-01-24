package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"udemy-modern-golang/dino/communicationlayer/dinoproto2"

	"github.com/golang/protobuf/proto"
)

/*
1- We will serialize some data via proto2
2- We will send this data via TCP to a different service
3- We will deserialize the data via proto2, and print out the extracted values

A- A TCP client needs to be written to send the data
B- A TCP server to receive the data
*/

func main() {
	op := flag.String("op", "s", "s for server, c for client, and ") //proto2test -op s => will run as a server, proto2test -op c as client
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}
}

func RunProto2Server() {
	l, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		log.Println("Accepted...")
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}
			a := &dinoproto2.Animal{}
			proto.Unmarshal(data, a)
			fmt.Println(a)
		}(c)
	}
}

func RunProto2Client() {
	a := &dinoproto2.Animal{
		Id:         proto.Int(1),
		AnimalType: proto.String("Raptor"),
		Nickname:   proto.String("rapto"),
		Zone:       proto.Int(3),
		Age:        proto.Int(20),
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}
