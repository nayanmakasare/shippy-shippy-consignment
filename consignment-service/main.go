package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/nayanmakasare/shippy-shippy-consignment/consignment-service/proto/consignment"
	"github.com/nayanmakasare/shippy-shippy-consignment/shippy-service-vessel/proto/vessel"
	"log"
	"os"
)

const (
	DB_HOST  = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/shippy?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	port = ":50051"
	)

func main()  {

	srv := micro.NewService(micro.Name("shippy.service.consignment"))

	srv.Init()

	uri := os.Getenv("DB_HOST")

	if uri == "" {
		uri = DB_HOST
	}

	client, err := CreateClient(uri)

	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignment")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vessel.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}




















