package dinogrpc

import (
	"fmt"
	"udemy-modern-golang/dino/communicationlayer/dinogrpc"
	"udemy-modern-golang/dino/databaselayer"

	context "golang.org/x/net/context"
)

type DinoGrpcServer struct {
	dbHandler databaselayer.DinoHandler
}

func NewDinoGrpServer(dbtype uint8, connstring string) (*DinoGrpcServer, error) {
	handler, err := databaselayer.GetDatabaseHandler(dbtype, connstring) //databaselayer.MONGODB, "mongodb://127.0.0.1"
	if err != nil {
		return nil, fmt.Errorf("Could not create a database handler object, error %v", err)
	}
	return &DinoGrpcServer{
		dbHandler: handler,
	}, nil
}

func (server *DinoGrpcServer) GetAnimal(ctx context.Context, r *Request) (*dinogrpc.Animal, error) {
	animal, err := server.dbHandler.GetDynoByNickname(r.GetNickname())
	return convertToDinoGRPCAnimal(animal), err
}

func (server *DinoGrpcServer) GetAllAnimal(*dinogrpc.Request, dinogrpc.DinoService_GetAllAnimalServer) error {
	animals, err := server.dbHandler.GetAvailableDynos()
	if err != nil {
		return err
	}

	for _, animal := range animals {
		grpcAnimal := convertToDinoGRPCAnimal(animal)
		if err := streamSend(grpcAnimal); err != nil {
			return err
		}
	}

	return nil
}

func convertToDinoGRPCAnimal(animal databaselayer.Animal) dinogrpc.Animal {
	return &dinogrpc.Animal{
		Id:       int32(animal.ID),
		Nickname: animal.Nickname,
		Zone:     animal.Zone,
		Age:      animal.Age,
	}
}
