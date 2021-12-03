package main

import (
	"context"
	"log"
	"net"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type BrokerServer struct {
	pb.UnimplementedBrokerServer
}

// func (s *BrokerServer) AddCity(ctx context.Context, req *pb.AddCityRequest) (*pb.AddCityBrokerResponse, error) {
// 	log.Printf("AddCity: %v", req)
// 	return nil, nil
// }

// func (s *BrokerServer) DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.DeleteCityBrokerResponse, error) {
// 	log.Printf("DeleteCity: %v", req)
// 	return nil, nil
// }

// func (s *BrokerServer) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.UpdateNameBrokerResponse, error) {
// 	log.Printf("UpdateName: %v", req)
// 	return nil, nil
// }

// func (s *BrokerServer) UpdateNumber(ctx context.Context, req *pb.UpdateNumberRequest) (*pb.UpdateNumberBrokerResponse, error) {
// 	log.Printf("UpdateNumber: %v", req)
// 	return nil, nil
// }

func (s *BrokerServer) GetNumberRebeldes(ctx context.Context, req *pb.GetNumberRebeldesRequest) (*pb.GetNumberRebeldesResponse, error) {
	// Leia llama esta wea
	// Consulta al Fulcrum
	// Response a Leia
	log.Printf("NumberRebeldes: %v", req)
	return nil, nil
}

func (s *BrokerServer) GetFulcrum(ctx context.Context, req *pb.GetFulcrumRequest) (*pb.GetFulcrumResponse, error) {
	// Informantes llaman para saber a que fulcrum se refiere
	log.Printf("GetFulcrum: %v", req)
	return nil, nil
}

func main() {
	// Escuchar en el puerto 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServer(grpcServer, &BrokerServer{})
	log.Printf("Server listening at 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
