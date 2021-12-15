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
	// get the number of rebelds from one random fulcrum
	index := rand.Intn(3)
	conn, err := grpc.Dial(ipFulcrum[index]+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)
	r, err := c.GetNumberRebeldesFulcrum(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received: %d", r.NumeroRebeldes)
	// send updated vector clock to leia
	return &pb.GetNumberRebeldesResponse{NumeroRebeldes: r.NumeroRebeldes, X: r.X, Y: r.Y, Z: r.Z, Ip: ipFulcrum[index]}, nil
}

func (s *BrokerServer) GetFulcrum(ctx context.Context, req *pb.GetFulcrumRequest) (*pb.GetFulcrumResponse, error) {
	// pick a random available fulcrum
	var ip string
	for {
		index := rand.Intn(3)
		ip := ipFulcrum[index]
		// dial to ip
		conn, err := grpc.Dial(ip+portFulcrum, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewFulcrumClient(conn)
		_, err = c.IsAvailable(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		break
	}
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
