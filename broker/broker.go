package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type BrokerServer struct {
	pb.UnimplementedBrokerServer
}

const (
	portFulcrum = ":50050"
)

var (
	ipFulcrum = [3]string{"10.6.43.77", "10.6.43.78", "10.6.43.79"}
)

func (s *BrokerServer) GetNumberRebeldes(ctx context.Context, req *pb.GetNumberRebeldesRequest) (*pb.GetNumberRebeldesResponse, error) {
	// Leia llama esta wea
	// Consulta al Fulcrum
	// Response a Leia
	log.Printf("NumberRebeldes: %v", req)
	return nil, nil
}

func (s *BrokerServer) GetFulcrum(ctx context.Context, req *pb.GetFulcrumRequest) (*pb.GetFulcrumResponse, error) {
	// Informantes llaman para saber a que fulcrum se refiere
	// pick a random ip from ipFulcrum
	index := rand.Intn(3)
	ip := ipFulcrum[index]
	// return ip
	return &pb.GetFulcrumResponse{IpFulcrum: ip}, nil
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
