package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServer
}

const (
	portFulcrum   = ":50050"
	fulcrumNumber = 0
)

var (
	ipFulcrum    = [3]string{"10.6.43.77", "10.6.43.78", "10.6.43.79"}
	vectorClocks = make(map[string]*pb.Vector)
	mutex        = &sync.Mutex{}
	canReceive   = false
)

func printVectorClocks() {
	fmt.Println("VECTOR CLOCKS:")
	for planeta, vector := range vectorClocks {
		fmt.Println(planeta, vector.X, vector.Y, vector.Z)
	}
}

func planetFolderExists(nombre_planeta string) bool {
	// Revise if folder exists inside planets folder
	_, err := os.Stat("fulcrum/planets/" + nombre_planeta)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func planetFileExists(nombre_planeta string) bool {
	_, err := os.Stat("fulcrum/planets/" + nombre_planeta + "/" + nombre_planeta + ".txt")
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func addOneToVectorClock(nombre_planeta string) {
	// Add one to vector clock
	switch fulcrumNumber {
	case 0:
		vectorClocks[nombre_planeta].X++
	case 1:
		vectorClocks[nombre_planeta].Y++
	case 2:
		vectorClocks[nombre_planeta].Z++
	}
}

func createPlanet(nombre_planeta string) {
	// create folder for planet
	if !planetFolderExists(nombre_planeta) {
		fmt.Println("Creating folder for planet: ", nombre_planeta)
		os.Mkdir("fulcrum/planets/"+nombre_planeta, 0777)
	}
	// Create files for planet
	if !planetFileExists(nombre_planeta) {
		fmt.Println("Creating files for planet: ", nombre_planeta)
		// create planet file
		file, err := os.Create("fulcrum/planets/" + nombre_planeta + "/" + nombre_planeta + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// create planet log file
		file, err = os.Create("fulcrum/planets/" + nombre_planeta + "/" + nombre_planeta + "_log.txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// initialize vector clock if this doesnt exists
		if _, ok := vectorClocks[nombre_planeta]; !ok {
			vectorClocks[nombre_planeta] = &pb.Vector{X: 0, Y: 0, Z: 0}
		}
	}
}

// Registra la informacion de las llamadas a los servicios
func writeToLog(request string, planeta_afectado string, ciudad_afectada string, nuevo_numero_soldados int32, nuevo_nombre_ciudad string) {
	// Abrir archivo de log
	file, ferr := os.OpenFile("fulcrum/planets/"+planeta_afectado+"/"+planeta_afectado+"_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if ferr != nil {
		log.Fatal(ferr)
	}
	defer file.Close()
	// update vector clock
	addOneToVectorClock(planeta_afectado)
	switch request {
	case "AddCity":
		_, err := file.WriteString(request + " " + planeta_afectado + " " + ciudad_afectada + " " + strconv.Itoa(int(nuevo_numero_soldados)) + "\n")
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
	filename := "fulcrum/planets/" + nombre_planeta + "/" + nombre_planeta + ".txt"
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
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+"/"+nombre_planeta+".txt", os.O_RDWR, 0644)
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
	// bool to check if city was found
	found := false
	// Actualizar soldados de la ciudad nombre_ciudad en el archivo planets/nombre_planeta.txt
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+"/"+nombre_planeta+".txt", os.O_RDWR, 0644)
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
			// update found
			found = true
			// change number of soldiers
			old_numero_soldados := strings.Split(line, " ")[2]
			line = strings.Replace(line, old_numero_soldados, strconv.Itoa(int(numero_soldados)), 1)
		}
		lines = append(lines, line)
	}
	if !found {
		// add city to file
		addCityToFile(nombre_planeta, nombre_ciudad, numero_soldados)
	} else {
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
}

// UpdateName
func updateCiudad(nombre_planeta string, nombre_ciudad string, nombre_reemplazo string) {
	// bool to check if city was found
	found := false
	filename, err := os.OpenFile("fulcrum/planets/"+nombre_planeta+"/"+nombre_planeta+".txt", os.O_RDWR, 0644)
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
			// update found
			found = true
			// change name of city
			line = strings.Replace(line, nombre_ciudad, nombre_reemplazo, 1)
		}
		lines = append(lines, line)
	}
	if !found {
		// add city to file
		addCityToFile(nombre_planeta, nombre_reemplazo, 0)
	} else {
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
}

func restartLog() {
	// remove all planets, reading vectorClocks map
	for planeta := range vectorClocks {
		// remove planet log file
		os.Remove("fulcrum/planets/" + planeta + "/" + planeta + "_log.txt")
		// Crear archivo de log
		file, err := os.Create("fulcrum/planets/" + planeta + "/" + planeta + "_log.txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}

func (s *FulcrumServer) AddCity(ctx context.Context, req *pb.AddCityRequest) (*pb.VectorClock, error) {
	createPlanet(req.NombrePlaneta)
	writeToLog("AddCity", req.NombrePlaneta, req.NombreCiudad, req.Soldados, "")
	addCityToFile(req.NombrePlaneta, req.NombreCiudad, req.Soldados)
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.VectorClock, error) {
	createPlanet(req.NombrePlaneta)
	writeToLog("DeleteCity", req.NombrePlaneta, req.NombreCiudad, 0, "")
	if planetFileExists(req.NombrePlaneta) {
		deleteCity(req.NombrePlaneta, req.NombreCiudad)
	}
	return &pb.VectorClock{}, nil
}

// Falta ver que funcione cuando no existe una ciudad/planeta
// Pasa lo mismo que con updateNumber
func (s *FulcrumServer) UpdateName(ctx context.Context, req *pb.UpdateNameRequest) (*pb.VectorClock, error) {
	createPlanet(req.NombrePlaneta)
	writeToLog("UpdateName", req.NombrePlaneta, req.NombreCiudad, 0, req.NuevoNombre)
	updateCiudad(req.NombrePlaneta, req.NombreCiudad, req.NuevoNombre)
	// if fulcrum1: X: changes, fulcrum2 --> Y: changes, fulcrum3 --> Z: changes
	return &pb.VectorClock{}, nil
}

// Falta ver que funcione cuando no existe una ciudad/planeta
func (s *FulcrumServer) UpdateNumber(ctx context.Context, req *pb.UpdateNumberRequest) (*pb.VectorClock, error) {
	createPlanet(req.NombrePlaneta)
	writeToLog("UpdateNumber", req.NombrePlaneta, req.NombreCiudad, req.NuevoNumero, "")
	updateSoldados(req.NombrePlaneta, req.NombreCiudad, req.NuevoNumero)
	// Con createPlanet siempre existe el file, hay que cambiar esto
	// if fulcrum1: X: changes, fulcrum2 --> Y: changes, fulcrum3 --> Z: changes
	return &pb.VectorClock{}, nil
}

func (s *FulcrumServer) GetNumberRebeldesFulcrum(ctx context.Context, req *pb.GetNumberRebeldesRequest) (*pb.GetNumberRebeldesResponse, error) {
	// read planet file
	filename, err := os.OpenFile("fulcrum/planets/"+req.NombrePlaneta+"/"+req.NombrePlaneta+".txt", os.O_RDWR, 0644)
	// if planet doesnt exist, vector clock is iniatilized as (0,0,0)
	if err != nil {
		return &pb.GetNumberRebeldesResponse{
			NumeroRebeldes: 0,
			X:              0,
			Y:              0,
			Z:              0,
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
				X:              vectorClocks[req.NombrePlaneta].X,
				Y:              vectorClocks[req.NombrePlaneta].Y,
				Z:              vectorClocks[req.NombrePlaneta].Z,
			}, nil
		}
	}
	// city does not exists, but vector clock can be iniatilized with certain values
	return &pb.GetNumberRebeldesResponse{
		NumeroRebeldes: 0,
		X:              vectorClocks[req.NombrePlaneta].X,
		Y:              vectorClocks[req.NombrePlaneta].Y,
		Z:              vectorClocks[req.NombrePlaneta].Z,
	}, nil
}

func (s *FulcrumServer) IsAvailable(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

// Corre en el fulcrum1, envía los cambios totales a fulcrum2 y fulcrum3, además de actualizar su vector clock
func (s *FulcrumServer) BidirectionalMerge(stream pb.Fulcrum_BidirectionalMergeServer) error {
	// for each line received, update local files
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// send local files to connected fulcrum
			for planet := range vectorClocks {
				// read planet file and send all lines to fulcrum2 and fulcrum3
				fmt.Println("Opening file fulcrum/planets/" + planet + "/" + planet + ".txt")
				filename, err := os.OpenFile("fulcrum/planets/"+planet+"/"+planet+".txt", os.O_RDWR, 0644)
				if err != nil {
					log.Fatalf("could not open file: %v", err)
				}
				defer filename.Close()
				scanner := bufio.NewScanner(filename)
				// for each line in planet file
				for scanner.Scan() {
					line := scanner.Text()
					fmt.Println("Sending line " + line)
					// send line to connected fulcrum
					if err := stream.Send(&pb.MergeResponse{Line: line}); err != nil {
						log.Fatal(err)
					}
				}
				// close file
				filename.Close()
			}
			return nil
		}
		if err != nil {
			return err
		}
		line := req.Line
		fmt.Println("received", line)
		// format line
		command := strings.Split(line, " ")[0]
		planet_name := strings.Split(line, " ")[1]
		city_name := strings.Split(line, " ")[2]
		// create planet file if it doesn't exist
		createPlanet(planet_name)
		// update local files
		switch command {
		case "AddCity":
			var number int32 = 0
			if len(line) == 4 {
				num, _ := strconv.Atoi(strings.Split(line, " ")[3])
				number = int32(num)
			}
			addCityToFile(planet_name, city_name, number)
		case "DeleteCity":
			deleteCity(planet_name, city_name)
		case "UpdateName":
			new_name := strings.Split(line, " ")[3]
			updateCiudad(planet_name, city_name, new_name)
		case "UpdateNumber":
			new_number, _ := strconv.Atoi(strings.Split(line, " ")[3])
			nuevoNumero := int32(new_number)
			updateSoldados(planet_name, city_name, nuevoNumero)
		}
	}
}

func (s *FulcrumServer) ClockMerge(ctx context.Context, req *pb.VectorClocks) (*pb.VectorClocks, error) {
	// merge vector clocks
	receivedVectorClocks := req.VectorClocks
	for planet, vectorClock := range receivedVectorClocks {
		// if vector clock is not in local vector clock, add it
		if _, ok := vectorClocks[planet]; !ok {
			vectorClocks[planet] = vectorClock
		} else {
			// if vector clock is in local vector clock, merge it
			newVectorClock := vectorClocks[planet]
			if vectorClock.X > newVectorClock.X {
				newVectorClock.X = vectorClock.X
			}
			if vectorClock.Y > newVectorClock.Y {
				newVectorClock.Y = vectorClock.Y
			}
			if vectorClock.Z > newVectorClock.Z {
				newVectorClock.Z = vectorClock.Z
			}
			vectorClocks[planet] = newVectorClock
		}
	}
	return &pb.VectorClocks{VectorClocks: vectorClocks}, nil
}

func MergeRoutine() {
	// wait two minutes
	time.Sleep(time.Second * 20)
	conn, err := grpc.Dial(ipFulcrum[0]+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stream, err := client.BidirectionalMerge(ctx)
	for planet, _ := range vectorClocks {
		// read planet log file and send all commands to fulcrum1
		filename, err := os.OpenFile("fulcrum/planets/"+planet+"/"+planet+"_log.txt", os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		defer filename.Close()
		scanner := bufio.NewScanner(filename)
		// for each line in log file
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Sending line " + line)
			// send line to fulcrum1
			if err := stream.Send(&pb.MergeRequest{Line: line}); err != nil {
				log.Fatal(err)
			}
		}
		// close file
		filename.Close()
	}
	// send close
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}
	// delete files
	os.RemoveAll("fulcrum/planets")
	os.Mkdir("fulcrum/planets", 0777)
	// receive files from fulcrum1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		line := req.Line
		planet := strings.Split(line, " ")[0]
		fmt.Println("received", line)
		createPlanet(planet)
		// write to planet file the line
		planetFile, err := os.OpenFile("fulcrum/planets/"+planet+"/"+planet+".txt", os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		defer planetFile.Close()
		planetFile.WriteString(line + "\n")
	}
	// merge vectorClocks
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	resp, err := client.ClockMerge(context.Background(), &pb.VectorClocks{VectorClocks: vectorClocks})
	if err != nil {
		log.Fatal(err)
	}
	vectorClocks = resp.VectorClocks
}

func main() {
	// initialize mergeRoutine in a goroutine on fulcrum2 and 3
	if fulcrumNumber > 0 {
		go MergeRoutine()
	}

	// automaticly remove local files and create folder planets when the server starts
	os.RemoveAll("fulcrum/planets")
	os.Mkdir("fulcrum/planets", 0777)

	// writeToLog("AddCity", "Tatooine", "Mos_Eisley", 1000, "")
	// addCityToFile("Tatooine", "Mos_Eisley", 1000)
	// writeToLog("AddCity", "Test", "Mos_Eisley", 1000, "")
	// addCityToFile("Test", "Mos_Eisley", 1000)
	// writeToLog("UpdateCity", "Tatooine", "Mos_Eisley", 0, "Rancagua")
	// updateCiudad("Tatooine", "Mos_Eisley", "Rancagua")
	// for planet, vectorClock := range vectorClocks {
	// 	fmt.Println(planet, vectorClock.X, vectorClock.Y, vectorClock.Z)
	// }

	// Escuchar en el puerto 50050
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFulcrumServer(grpcServer, &FulcrumServer{})
	log.Printf("Fulcrum server listening at 50050")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
