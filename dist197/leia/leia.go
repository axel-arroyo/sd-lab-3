package main

import (
	"context"
	"fmt"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

const (
	portBroker = ":50051"
	ipBroker   = "10.6.43.80"
)

var lastVectorClock = make(map[string]*pb.GetNumberRebeldesResponse)

// compare which vector is the most recent one
func compareMostRecentVectorClock(lastReceived *pb.GetNumberRebeldesResponse, city string) {
	// revise if city exists in lastVectorClock
	if _, ok := lastVectorClock[city]; ok {
		// compare lastReceived with lastVectorClock[city]
		sumaLast := lastVectorClock[city].X + lastVectorClock[city].Y + lastVectorClock[city].Z
		sumaReceived := lastReceived.X + lastReceived.Y + lastReceived.Z
		// if received vector clock has a sum of its components larger than the local one, reassign vector clock to most recent
		if sumaReceived > sumaLast {
			lastVectorClock[city] = lastReceived
		}
	} else {
		// add city to lastVectorClock
		lastVectorClock[city] = lastReceived
	}
}

func GetNumberRebelds(planet string, city string) {
	// send request to broker
	conn, err := grpc.Dial(ipBroker+portBroker, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	clientBroker := pb.NewBrokerClient(conn)
	resp, err := clientBroker.GetNumberRebeldes(context.Background(), &pb.GetNumberRebeldesRequest{NombrePlaneta: planet, NombreCiudad: city})
	if err != nil {
		fmt.Println(err)
		return
	}
	// update local vector clock to consistency (value of map is X, Y, Z from vector clock and number as last index)
	lastReceived := resp
	compareMostRecentVectorClock(lastReceived, city)
	fmt.Printf("Numero de rebeldes: %d\n", lastVectorClock[city].NumeroRebeldes)
	return
}

func Menu() {
	fmt.Println("Esperando comando...")
	// read line
	var command string
	var planet string
	var city string
	fmt.Scanln(&command, &planet, &city)
	// while command is not empty
	for command != "" {
		if command == "GetNumberRebelds" {
			GetNumberRebelds(planet, city)
		} else {
			fmt.Println("Comando no reconocido")
		}
		fmt.Scanln(&command, &planet, &city)
	}
}

func main() {
	Menu()
}
