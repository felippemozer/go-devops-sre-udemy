package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ServerName string
	ServerURL  string
	ExecTime   float64
	StatusCode int
	FailDate   string
}

func AppendDowntimeList(DowntimeFile *os.File, DowntimeList []Server) {
	csvWriter := csv.NewWriter(DowntimeFile)
	for _, server := range DowntimeList {
		line := []string{server.ServerName, server.ServerURL, server.FailDate, fmt.Sprintf("%.2f", server.ExecTime), fmt.Sprintf("%d", server.StatusCode)}
		csvWriter.Write(line)
	}
	csvWriter.Flush()
}

func CreateServerList(ServerFile *os.File) []Server {
	csvReader := csv.NewReader(ServerFile)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var servers []Server

	for i, line := range data {
		if i > 0 {
			server := Server{
				ServerName: line[0],
				ServerURL:  line[1],
			}
			servers = append(servers, server)
		}
	}

	return servers
}

func ProcessServer(ServerList []Server) []Server {
	var DownServers []Server
	for _, server := range ServerList {
		now := time.Now()
		get, err := http.Get(server.ServerURL)

		if err != nil {
			fmt.Println("Servidor", server.ServerName, "está fora")
			fmt.Println(err)
			fmt.Println()
			server.StatusCode = 0
			server.FailDate = now.Format("02/01/2006 15:04:05")
			DownServers = append(DownServers, server)
			continue
		}

		server.ExecTime = time.Since(now).Seconds()
		server.StatusCode = get.StatusCode
		if server.StatusCode != 200 {
			server.FailDate = now.Format("02/01/2006 15:04:05")
			DownServers = append(DownServers, server)
		}

		fmt.Println("Execução do site", server.ServerName)
		fmt.Println("Status:", server.StatusCode)
		fmt.Println("Tempo decorrido:", server.ExecTime, "segundos")
		fmt.Println()
	}

	return DownServers

}

func OpenFiles(ServerFileName string, DowntimeFileName string) (*os.File, *os.File) {
	ServerFile, err := os.OpenFile(ServerFileName, os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	DowntimeFile, err := os.OpenFile(DowntimeFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return ServerFile, DowntimeFile

}

func main() {
	ServerFile, DowntimeFile := OpenFiles(os.Args[1], os.Args[2])
	defer ServerFile.Close()
	defer DowntimeFile.Close()

	ServerList := CreateServerList(ServerFile)
	DownServerList := ProcessServer(ServerList)
	AppendDowntimeList(DowntimeFile, DownServerList)
}
