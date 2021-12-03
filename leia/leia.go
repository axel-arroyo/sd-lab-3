package main

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

const (
	portBroker = ":50051"
	ipBroker   = "10.6.43.80"
)

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
	number := resp.NumeroRebeldes
	fmt.Printf("Numero de rebeldes: %d\n", number)
	return
}

func Menu() {
	var line string
	fmt.Println("Esperando comando...")
	// read line
	fmt.Scanln(&line)
	// while line is not empty
	for line != "" {
		// parse line
		tokens := strings.Split(line, " ")
		command := tokens[0]
		planet := tokens[1]
		city := tokens[2]
		if command == "GetNumberRebelds" {
			GetNumberRebelds(planet, city)
		} else {
			fmt.Println("Comando não reconhecido")
		}
		fmt.Scanln(&line)
	}
}

func main() {
	Menu()
}
