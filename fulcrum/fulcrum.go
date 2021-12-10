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
	"time"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServer
}

const (
	portFulcrum = ":50050"
)

var (
	ipFulcrum   = [3]string{"10.6.43.77", "10.6.43.78", "10.6.43.79"}
	vectorClock = [3]int32{0, 0, 0}
)

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
}

func restartLog() {
	// Borrar archivo de log
	os.Remove("fulcrum/log.txt")
	// Crear archivo de log
	file, err := os.Create("fulcrum/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func planetFileExists(nombre_planeta string) bool {
	_, err := os.Stat("fulcrum/planets/" + nombre_planeta + ".txt")
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (s *FulcrumServer) AddCity(ctx context.Context, req *pb.AddCityRequest) (*pb.VectorClock, error) {
	// AddCity
	writeToLog("AddCity", req.NombrePlaneta, req.NombreCiudad, req.Soldados, "")
	addCityToFile(req.NombrePlaneta, req.NombreCiudad, req.Soldados)
	// Update vector clock
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.VectorClock, error) {
	writeToLog("DeleteCity", req.NombrePlaneta, req.NombreCiudad, 0, "")
	// Update vector clock
	if planetFileExists(req.NombrePlaneta) {
		deleteCity(req.NombrePlaneta, req.NombreCiudad)
	}
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.VectorClock, error) {
	writeToLog("UpdateName", req.NombrePlaneta, req.NombreCiudad, 0, req.NuevoNombre)
	// Update vector clock
	if planetFileExists(req.NombrePlaneta) {
		updateCiudad(req.NombrePlaneta, req.NombreCiudad, req.NuevoNombre)
	} else {
		addCityToFile(req.NombrePlaneta, req.NuevoNombre, 0)
	}
	// if fulcrum1: X: changes, fulcrum2 --> Y: changes, fulcrum3 --> Z: changes
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) UpdateNumber(ctx context.Context, req *pb.UpdateNumberRequest) (*pb.VectorClock, error) {
	writeToLog("UpdateNumber", req.NombrePlaneta, req.NombreCiudad, req.NuevoNumero, "")
	// Update vector clock
	if planetFileExists(req.NombrePlaneta) {
		updateSoldados(req.NombrePlaneta, req.NombreCiudad, req.NuevoNumero)
	} else {
		addCityToFile(req.NombrePlaneta, req.NombreCiudad, req.NuevoNumero)
	}
	// if fulcrum1: X: changes, fulcrum2 --> Y: changes, fulcrum3 --> Z: changes
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) GetNumberRebeldesFulcrum(ctx context.Context, req *pb.GetNumberRebeldesRequest) (*pb.GetNumberRebeldesResponse, error) {
	// read planet file
	filename, err := os.OpenFile("fulcrum/planets/"+req.NombrePlaneta+".txt", os.O_RDWR, 0644)
	if err != nil {
		return &pb.GetNumberRebeldesResponse{
			NumeroRebeldes: 0,
		}, nil
	}
	defer filename.Close()
	// read line with city name
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, req.NombreCiudad) {
			// get number of soldiers
			numero_soldados := strings.Split(line, " ")[2]
			numero_soldados_int, _ := strconv.Atoi(numero_soldados)
			return &pb.GetNumberRebeldesResponse{
				NumeroRebeldes: int32(numero_soldados_int),
			}, nil
		}
	}
	return &pb.GetNumberRebeldesResponse{
		NumeroRebeldes: 0,
	}, nil
}

func (s *FulcrumServer) VectorClockMerge(ctx context.Context, req *pb.VectorClock) (*pb.Empty, error) {
	// merge vector clocks
	switch req.Ip {
	case ipFulcrum[0]:
		// merge fulcrum1
	case ipFulcrum[1]:
		// merge fulcrum2
	case ipFulcrum[2]:
		// merge fulcrum3
	}
	return &pb.Empty{}, nil
}

func (s *FulcrumServer) Merge(stream pb.Fulcrum_MergeServer) error {
	return nil
}

func mergeRoutine() {
	// wait two minutes
	time.Sleep(time.Minute * 2)
	// send vectorClock to fulcrum1
	conn, err := grpc.Dial(ipFulcrum[0]+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.VectorClockMerge(ctx, &pb.VectorClock{X: vectorClock[0], Y: vectorClock[1], Z: vectorClock[2]})
	if err != nil {
		log.Fatalf("could not merge: %v", err)
	}
	// send all changes to fulcrum1
	// read log file
	filename, err := os.OpenFile("fulcrum/log.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := client.Merge(ctx)
	defer filename.Close()
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		line := scanner.Text()
		// send line to fulcrum1
		if err := stream.Send(&pb.MergeRequest{Line: line}); err != nil {
			log.Fatal(err)
		}
	}
	restartLog()
	// merge again
	mergeRoutine()
}

func main() {
	// Crear archivo log.txt
	restartLog()
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
