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
	"time"

	pb "github.com/axel-arroyo/sd-lab-3/gen/proto"
	"google.golang.org/grpc"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServer
}

const (
	portFulcrum   = ":50050"
	fulcrumNumber = 1
)

var (
	ipFulcrum    = [3]string{"10.6.43.77", "10.6.43.78", "10.6.43.79"}
	vectorClocks = make(map[string]*pb.Vector)
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

// Registra la informacion de las llamadas a los servicios
func writeToLog(request string, planeta_afectado string, ciudad_afectada string, nuevo_numero_soldados int32, nuevo_nombre_ciudad string) {
	// Crear carpeta del planeta si no existe
	if !planetFolderExists(planeta_afectado) {
		os.Mkdir("fulcrum/planets/"+planeta_afectado, 0777)
	}
	// Crear archivo de log y planeta si no existe
	if !planetFileExists(planeta_afectado) {
		// Crear archivo del planeta
		file, err := os.Create("fulcrum/planets/" + planeta_afectado + "/" + planeta_afectado + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// Crear log del planeta
		file, err = os.Create("fulcrum/planets/" + planeta_afectado + "/" + planeta_afectado + "_log.txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// initialize vector clock
		vectorClocks[planeta_afectado] = &pb.Vector{X: 0, Y: 0, Z: 0}
	}
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
	filename, err := os.OpenFile("fulcrum/planets/"+req.NombrePlaneta+"/"+req.NombrePlaneta+".txt", os.O_RDWR, 0644)
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
	fmt.Println("VectorClockMerge")
	fmt.Println(req.X, req.Y, req.Z)
	planet := req.NombrePlaneta
	// revise if planet exists in map
	if _, ok := vectorClocks[planet]; !ok {
		vectorClocks[planet] = &pb.Vector{X: req.X, Y: req.Y, Z: req.Z}
		return &pb.Empty{}, nil
	}
	switch req.Ip {
	case ipFulcrum[0]:
		// merge fulcrum1
	case ipFulcrum[1]:
		// update Y value of vector clock to the max value of Y
		if vectorClocks[planet].Y < req.Y {
			vectorClocks[planet].Y = req.Y
		}
	case ipFulcrum[2]:
		// update Z value of vector clock to the max value of Z
		if vectorClocks[planet].Z < req.Z {
			vectorClocks[planet].Z = req.Z
		}
	}
	return &pb.Empty{}, nil
}

// Corre en el fulcrum1, envía los cambios totales a fulcrum2 y fulcrum3, además de actualizar su vector clock
func (s *FulcrumServer) Merge(stream pb.Fulcrum_MergeServer) error {
	// for each line received, update local files and return vectorClock
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// send local planets files to fulcrum2 and fulcrum3
			go MergeOtherFulcrums()
			// update vector clock in fulcrum 2 and 3
			stream.SendAndClose(&pb.VectorClocks{VectorClocks: vectorClocks})
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

func (s *FulcrumServer) MergeFulcrums(stream pb.Fulcrum_MergeFulcrumsServer) error {
	// remove local files
	os.RemoveAll("fulcrum/planets")
	// create folder planets
	os.Mkdir("fulcrum/planets", 0777)
	// for each line received, update local files
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// close stream
			fmt.Println("Merge finished")
			stream.SendAndClose(&pb.Empty{})
			return nil
		}
		if err != nil {
			return err
		}
		line := req.Line
		// format line
		command := strings.Split(line, " ")[0]
		planet_name := strings.Split(line, " ")[1]
		city_name := strings.Split(line, " ")[2]
		// update local files
		switch command {
		case "AddCity":
			var number int32
			if len(line) == 4 {
				num, _ := strconv.Atoi(strings.Split(line, " ")[3])
				number = int32(num)
			} else {
				number = 0
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

func MergeOtherFulcrums() {
	for _, ip := range ipFulcrum {
		if ip != ipFulcrum[0] {
			// connect to other fulcrums
			conn, err := grpc.Dial(ip, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewFulcrumClient(conn)
			// send planet files to other fulcrums
			for planet := range vectorClocks {
				planet_file, err := os.OpenFile("fulcrum/planets/"+planet+"/"+planet+".txt", os.O_RDWR, 0644)
				if err != nil {
					log.Fatalf("could not open file: %v", err)
				}
				defer planet_file.Close()
				stream, err := client.Merge(context.Background())
				// read file line by line
				scanner := bufio.NewScanner(planet_file)
				for scanner.Scan() {
					line := scanner.Text()
					// send line to other fulcrum
					if err := stream.Send(&pb.MergeRequest{Line: line}); err != nil {
						log.Fatal(err)
					}
				}
				planet_file.Close()
			}
		}
	}
}

// Corre en el fulcrum2 y 3, envían el vector clock al fulcrum1 + todos los cambios de cada planeta
func mergeRoutine() {
	// wait two minutes
	time.Sleep(time.Minute * 2)
	fmt.Println("Merge started")
	// send vectorClock to fulcrum1
	conn, err := grpc.Dial(ipFulcrum[0]+portFulcrum, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// get vectorClock from vectorClocks map for every planet in the map
	for planet, vectorClock := range vectorClocks {
		// send vectorClock to fulcrum1
		_, err := client.VectorClockMerge(ctx, &pb.VectorClock{
			X:             vectorClock.X,
			Y:             vectorClock.Y,
			Z:             vectorClock.Z,
			Ip:            ipFulcrum[fulcrumNumber],
			NombrePlaneta: planet})
		if err != nil {
			log.Fatalf("could not merge: %v", err)
		}
		// read planet log file and send all commands to fulcrum1
		filename, err := os.OpenFile("fulcrum/planets/"+planet+"/"+planet+"_log.txt", os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		stream, err := client.Merge(ctx)
		defer filename.Close()
		scanner := bufio.NewScanner(filename)
		// for each line in log file
		for scanner.Scan() {
			line := scanner.Text()
			// send line to fulcrum1
			if err := stream.Send(&pb.MergeRequest{Line: line}); err != nil {
				log.Fatal(err)
			}
		}
	}

	// restart all logs
	restartLog()
	// merge again
	mergeRoutine()
}

func main() {
	// initialize mergeRoutine in a goroutine on fulcrum2 and 3
	if fulcrumNumber > 0 {
		go mergeRoutine()
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
