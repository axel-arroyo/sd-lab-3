package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

const (
	portFulcrum = ":50050"
	portBroker  = ":50054"
	ipBroker    = "10.6.43.80"
)

func getFulcrumIp() string {
	// connect to broker
	conn, err := grpc.Dial(ipBroker+portBroker, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to broker: %v", err)
	}
	defer conn.Close()
	clientBroker := pb.NewBrokerClient(conn)
	resp, err := clientBroker.GetFulcrum(context.Background(), &pb.GetFulcrumRequest{})
	if err != nil {
		log.Fatalf("Could not get Fulcrum IP: %v", err)
	}
	ipFulcrum := resp.IpFulcrum
	return ipFulcrum
}

func addCity(args string) {
	// get planet, city and optional number
	parsedLine := strings.Split(args, " ")
	planet := parsedLine[0]
	city := parsedLine[1]
	var number int = 0
	if len(parsedLine) == 3 {
		number, _ = strconv.Atoi(parsedLine[2])
	}
	// connect to broker and get Fulcrum IP
	ipFulcrum := getFulcrumIp()
	// connect to fulcrum
	conn, err := grpc.Dial(ipFulcrum+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to fulcrum: %v", err)
	}
	defer conn.Close()
	clientFulcrum := pb.NewFulcrumClient(conn)
	_, err = clientFulcrum.AddCity(context.Background(), &pb.AddCityRequest{NombrePlaneta: planet, NombreCiudad: city, Soldados: int32(number)})
	if err != nil {
		log.Fatalf("error contacting fulcrum: %v", err)
	}
	fmt.Println("City added")
}

func updateNumber(args string) {
	// get planet, city and optional number
	parsedLine := strings.Split(args, " ")
	planet := parsedLine[0]
	city := parsedLine[1]
	// read number as int32
	number, _ := strconv.Atoi(parsedLine[2])
	// get Fulcrum IP
	ipFulcrum := getFulcrumIp()
	// connect to fulcrum
	conn, err := grpc.Dial(ipFulcrum+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to fulcrum: %v", err)
	}
	defer conn.Close()
	clientFulcrum := pb.NewFulcrumClient(conn)
	_, err = clientFulcrum.UpdateNumber(context.Background(), &pb.UpdateNumberRequest{NombrePlaneta: planet, NombreCiudad: city, NuevoNumero: int32(number)})
	if err != nil {
		log.Fatalf("error contacting fulcrum: %v", err)
	}
	fmt.Println("Number updated")
}

func updateName(args string) {
	parsedLine := strings.Split(args, " ")
	planet := parsedLine[0]
	old_city_name := parsedLine[1]
	new_city_name := parsedLine[2]
	// get Fulcrum IP
	ipFulcrum := getFulcrumIp()
	// connect to fulcrum
	conn, err := grpc.Dial(ipFulcrum+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to fulcrum: %v", err)
	}
	defer conn.Close()
	clientFulcrum := pb.NewFulcrumClient(conn)
	_, err = clientFulcrum.UpdateName(context.Background(), &pb.UpdateNameRequest{NombrePlaneta: planet, NombreCiudad: old_city_name, NuevoNombre: new_city_name})
	if err != nil {
		log.Fatalf("error contacting fulcrum: %v", err)
	}
	fmt.Println("Name updated")
}

func deleteCity(args string) {
	parsedLine := strings.Split(args, " ")
	planet := parsedLine[0]
	city := parsedLine[1]
	// get Fulcrum IP
	ipFulcrum := getFulcrumIp()
	// connect to fulcrum
	conn, err := grpc.Dial(ipFulcrum+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to fulcrum: %v", err)
	}
	defer conn.Close()
	clientFulcrum := pb.NewFulcrumClient(conn)
	_, err = clientFulcrum.DeleteCity(context.Background(), &pb.DeleteCityRequest{NombrePlaneta: planet, NombreCiudad: city})
	if err != nil {
		log.Fatalf("error contacting fulcrum: %v", err)
	}
	fmt.Println("City deleted")
}

func Menu() {
	var line string
	fmt.Println("Esperando comando...")
	// read line
	fmt.Scanln(&line)
	// while line is not empty
	for line != "" {
		// command is first word
		command := line[:strings.Index(line, " ")]
		// args is the rest
		args := line[strings.Index(line, " ")+1:]
		switch command {
		case "AddCity":
			addCity(args)
		case "UpdateNumber":
			updateNumber(args)
		case "UpdateName":
			updateName(args)
		case "DeleteCity":
			deleteCity(args)
		default:
			fmt.Println("Comando no reconocido")
		}
		fmt.Println("Esperando comando...")
		fmt.Scanln(&line)
	}
}

func main() {
	Menu()
}
