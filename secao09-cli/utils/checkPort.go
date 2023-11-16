package utils

import (
	"log"
	"net"
	"strings"
	"time"
)

func CheckPort(host string, ports string) {
	portList := strings.Split(ports, ",")
	for _, port := range portList {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			log.Printf("Porta %s está fechada\n", port)
		}
		if conn != nil {
			log.Printf("Porta %s está aberta\n", port)
			defer conn.Close()
		}
	}
}
