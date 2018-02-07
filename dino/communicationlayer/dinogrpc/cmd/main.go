package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"github.com/skyfish81/udemy-modern-golang/dino/communicationlayer/dinogrpc"
	"github.com/skyfish81/udemy-modern-golang/dino/databaselayer"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// export GOPATH="/Users/ryan/gohome:/Users/ryan/dev/lab/golang-tutorial"
// GRPC_GO_LOG_SEVERITY_LEVEL=INFO go run main.go -op s
// GRPC_GO_LOG_SEVERITY_LEVEL=INFO go run main.go -op c
func main() {
	op := flag.String("op", "s", "s for server, and c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		runGRPCServer()
	case "c":
		runGRPCClient()
	}
}

func runGRPCServer() {
	fmt.Println("runGRPCServer...")
	grpclog.Println("Starting GRPC Server")
	lis, err := net.Listen("tcp", ":8282")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		fmt.Printf("failed to listen: %v", err)
	}
	grpclog.Println("Listening on 127.0.0.1;8282")
	// var opts []grpc.ServerOption
	// grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	dinoServer, err := dinogrpc.NewDinoGrpcServer(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	dinogrpc.RegisterDinoServiceServer(grpcServer, dinoServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func runGRPCClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := dinogrpc.NewDinoServiceClient(conn)
	input := ""
	fmt.Println("All animals? (y/n)")
	fmt.Scanln(&input)
	if strings.ToLower(input) == "y" {
		fmt.Println("Received input:", input)
		animals, err := client.GetAllAnimals(context.Background(), &dinogrpc.Request{})
		if err != nil {
			log.Fatal(err)
		}

		for {
			animal, err := animals.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				grpclog.Fatal(err)
			}
			grpclog.Println(animal)
		}
		return
	}
	fmt.Println("Nickname?")
	fmt.Scanln(&input)
	a, err := client.GetAnimal(context.Background(), &dinogrpc.Request{Nickname: input})
	if err != nil {
		log.Fatal(err)
	}
	grpclog.Println(*a)

}
