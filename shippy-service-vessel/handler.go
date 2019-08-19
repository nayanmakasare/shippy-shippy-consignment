package main


import (
	pb "github.com/nayanmakasare/shippy-shippy-consignment/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

// Our gRPC service handler
type service struct {
	collection *mongo.Collection
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.collection}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()

	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}