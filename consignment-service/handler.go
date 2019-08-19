package main

import (
	"context"
	"github.com/nayanmakasare/shippy-shippy-consignment/shippy-service-vessel/proto/vessel"
	pb "github.com/nayanmakasare/shippy-shippy-consignment/consignment-service/proto/consignment"
	"log"
)

type handler struct {
	repository
	vesselClient vessel.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error  {

	vesselSpec := &vessel.Specification{Capacity: int32(len(req.Containers))}

	vesselResponse,err :=  s.vesselClient.FindAvailable(ctx, vesselSpec)

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	if err = s.repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error  {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}


























