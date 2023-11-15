package main

import (
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type TopURL struct {
	GlobalRanking string // posicao 0
	DomainRanking string // posicao 1
	Address       string // posicao 2
	Country       string // posicao 3
}

func checkCert(server string, ranking string) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 120 * time.Second}, "tcp", "www."+server+":443", nil)
	if err != nil {
		log.Printf("Erro ao conectar com site %s ranking %s: %s", server, ranking, err)
		return
	}
	defer conn.Close()
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	currentTime := time.Now()
	diff := expiry.Sub(currentTime)
	log.Printf("Tempo restante para o server %s expirar: %1.f dias. Ranking: %s", server, math.Round(diff.Hours()/24), ranking)
}

func CreateURLList(urlList *os.File) []TopURL {
	csvReader := csv.NewReader(urlList)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var urls []TopURL

	for _, line := range data {
		if line[3] == "br" {
			urls = append(urls, TopURL{
				GlobalRanking: line[0],
				DomainRanking: line[1],
				Address:       line[2],
				Country:       line[3],
			})
		}
	}
	return urls
}

func DownloadMillionDomains() error {
	if _, err := os.Stat("majestic_million.csv"); errors.Is(err, os.ErrNotExist) {
		log.Println("Downloading million domains file...")
		url := "https://downloads.majestic.com/majestic_million.csv"
		out, err := os.Create("majestic_million.csv")
		if err != nil {
			return err
		}
		defer out.Close()
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Bad status: %s", res.Status)
		}

		_, err = io.Copy(out, res.Body)
		if err != nil {
			return err
		}
		log.Println("Download finished")
	} else {
		log.Println("majestic_million.csv file already exists")
	}
	return nil
}

func main() {
	logFile, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo result.log", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	var wg sync.WaitGroup

	err = DownloadMillionDomains()
	if err != nil {
		log.Fatalln("An error ocurred on file download: ", err)
	}

	csvFile, err := os.OpenFile("majestic_million.csv", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalln("An error ocurred on file open: ", err)
	}
	defer csvFile.Close()

	urls := CreateURLList(csvFile)

	for _, url := range urls {
		wg.Add(1)
		go func(url TopURL) {
			checkCert(url.Address, url.DomainRanking)
			defer wg.Done()
		}(url)
	}

	wg.Wait()
}
