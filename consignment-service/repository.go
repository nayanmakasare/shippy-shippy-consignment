package main

import (
	pb "github.com/nayanmakasare/shippy-shippy-consignment/proto/consignment"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([] *pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func(mongoRepo *MongoRepository) Create(consignment *pb.Consignment) error{
	 _, err :=  mongoRepo.collection.InsertOne(context.Background(), consignment)
	 if err != nil{
	 	return err
	 }
	 return err
}

func(mongoRepo *MongoRepository) GetAll()([]*pb.Consignment, error){
	cur, err := mongoRepo.collection.Find(context.Background(),nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()){
		var consignment *pb.Consignment
		if err := cur.Decode(consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments,err;
}



























