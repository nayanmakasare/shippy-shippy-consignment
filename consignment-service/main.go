package main

import (
	"context"
	"fmt"
	"github.com/micro/protobuf/protoc-gen-go/micro"
	pb "github.com/nayanmakasare/shippy-shippy-consignment/proto/consignment"
	"log"
)

const DB_HOST  = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"

func main()  {
	srv := micro.NewService(micro.Name("shippy.service.consignment"))
	srv.Init()

	db, err := CreateClient(DB_HOST)
	defer db.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	consignmentCollection := db.Database("shippy").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}

	h := &Handler{repository}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}




















