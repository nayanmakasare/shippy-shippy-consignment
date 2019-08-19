package main


import (
	"context"
	"fmt"
	"log"

	"os"

	pb "github.com/nayanmakasare/shippy-shippy-consignment/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/shippy?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"

)

func createDummyData(repo Repository) {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateClient(host)
	defer session.Disconnect(context.Background())

	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}
	collection := session.Database("shippy").Collection("vessels")
	repo := &VesselRepository{collection}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{collection})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
