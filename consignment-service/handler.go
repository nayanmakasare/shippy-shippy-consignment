package main

import (
	"context"
	pb "github.com/nayanmakasare/shippy-shippy-consignment/proto/consignment"
)

type Handler struct {
	Repository
}

func (h *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response)  error {
	return Create(req)
}

func (h *Handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error  {
	consignments, err := GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return err
}


