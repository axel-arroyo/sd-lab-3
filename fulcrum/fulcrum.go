package main

import (
	"bufio"
	"context"
	"errors"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServer
}

// Registra la informacion de las llamadas a los servicios
func writeToLog(request string, planeta_afectado string, ciudad_afectada string, nuevo_numero_soldados int32, nuevo_nombre_ciudad string) {
	file, ferr := os.OpenFile("fulcrum/log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if ferr != nil {
		log.Fatal(ferr)
	}
	defer file.Close()
	switch request {
	case "AddCity":
		_, err := file.WriteString(request + " " + planeta_afectado + " " + ciudad_afectada + "\n")
		if err != nil {
			log.Fatal(err)
		}
		break
	case "DeleteCity":
		_, err := file.WriteString(request + " " + planeta_afectado + " " + ciudad_afectada + "\n")
		if err != nil {
			log.Fatal(err)
		}
		break
	case "UpdateNumber":
		_, err := file.WriteString(request + " " + planeta_afectado + " " + ciudad_afectada + " " + strconv.Itoa(int(nuevo_numero_soldados)) + "\n")
		if err != nil {
			log.Fatal(err)
		}
		break
	case "UpdateName":
		_, err := file.WriteString(request + " " + planeta_afectado + " " + ciudad_afectada + " " + nuevo_nombre_ciudad + "\n")
		if err != nil {
			log.Fatal(err)
		}
		break
	}
	file.Close()
}

// AddCity
func addCityToFile(nombre_planeta string, nombre_ciudad string, numero_soldados int32) {
	filename := "fulcrum/planets/" + nombre_planeta + ".txt"
	// Chequear si el archivo no existe para crearlo
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		textFile, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		textFile.Close()
	}
	textFile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = textFile.WriteString(nombre_planeta + " " + nombre_ciudad + " " + strconv.Itoa(int(numero_soldados)) + "\n")
	if err != nil {
		log.Fatal(err)
	}
	textFile.Close()
	writeToLog("AddCity", nombre_planeta, nombre_ciudad, numero_soldados, "")
}

// DeleteCity
func deleteCity(nombre_planeta string, nombre_ciudad string) {
	// delete city line from planet file
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+".txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()
	// read file
	scanner := bufio.NewScanner(filename)
	var lines []string
	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), nombre_ciudad) {
			lines = append(lines, scanner.Text())
		}
	}
	// rewrite file without deleted line
	filename.Truncate(0)
	filename.Seek(0, 0)
	for _, line := range lines {
		_, err = filename.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	filename.Close()
	writeToLog("DeleteCity", nombre_planeta, nombre_ciudad, 0, "")
}

// UpdateNumber
func updateSoldados(nombre_planeta string, nombre_ciudad string, numero_soldados int32) {
	// Actualizar soldados de la ciudad nombre_ciudad en el archivo planets/nombre_planeta.txt
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+".txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()
	// read file
	scanner := bufio.NewScanner(filename)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, nombre_ciudad) {
			// change number of soldiers
			old_numero_soldados := strings.Split(line, " ")[2]
			line = strings.Replace(line, old_numero_soldados, strconv.Itoa(int(numero_soldados)), 1)
		}
		lines = append(lines, line)
	}
	// rewrite file with new number of soldiers
	filename.Truncate(0)
	filename.Seek(0, 0)
	for _, line := range lines {
		_, err = filename.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	filename.Close()
	writeToLog("UpdateNumber", nombre_planeta, nombre_ciudad, numero_soldados, "")
}

// UpdateName
func updateCiudad(nombre_planeta string, nombre_ciudad string, nombre_reemplazo string) {
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+".txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()
	// read file
	scanner := bufio.NewScanner(filename)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, nombre_ciudad) {
			// change city name
			// old_ciudad := strings.Split(line, " ")[1]
			line = strings.Replace(line, nombre_ciudad, nombre_reemplazo, 1)
		}
		lines = append(lines, line)
	}
	// rewrite file with new city name
	filename.Truncate(0)
	filename.Seek(0, 0)
	for _, line := range lines {
		_, err = filename.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	filename.Close()
	writeToLog("UpdateName", nombre_planeta, nombre_ciudad, 0, nombre_reemplazo)
}

func (s *FulcrumServer) AddCity(ctx context.Context, req *pb.AddCityRequest) (*pb.AddCityResponse, error) {
	log.Printf("AddCity: %v", req)
	return nil, nil
}

func (s *FulcrumServer) DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.DeleteCityResponse, error) {
	log.Printf("DeleteCity: %v", req)
	return nil, nil
}

func (s *FulcrumServer) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.UpdateNameResponse, error) {
	log.Printf("UpdateName: %v", req)
	return nil, nil
}

func (s *FulcrumServer) UpdateNumber(ctx context.Context, req *pb.UpdateNumberRequest) (*pb.UpdateNumberResponse, error) {
	log.Printf("UpdateNumber: %v", req)
	return nil, nil
}

func main() {
	// Crear archivo log.txt
	filename, ferr := os.Create("fulcrum/log.txt")
	if ferr != nil {
		log.Fatal(ferr)
	}
	filename.Close()
	// addCityToFile("Tatooine", "Mos_Eisley", 1000)
	// updateSoldados("Tatooine", "Mos_Eisley", 2000)
	// updateCiudad("Tatooine", "Mos_Eisley", "Mos_Eisley_2")
	// addCityToFile("Tatooine", "Mos_Espa", 2000)
	// addCityToFile("Tatooine", "Mos_Pelgo", 3000)
	// addCityToFile("Coruscant", "Coruscant", 1000)
	// addCityToFile("Kamino", "Tipoca_City", 1000)
	// addCityToFile("Kamino", "Derem_City", 0)
	// deleteCity("Tatooine", "Mos_Eisley_2")
	// Escuchar en el puerto 50050
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFulcrumServer(grpcServer, &FulcrumServer{})
	log.Printf("Server listening at 50050")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
