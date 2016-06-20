package main

import (
	"fmt"
	"net"
	//"bufio"
	"flag"
	"strings"
)

var connectErrors = 0
var connections = 0
var goodConnections = 0

func udpSend(addr string, data string) {
	conn, err := net.Dial("udp", addr)
	connections++
	if err != nil {
		fmt.Printf("Failed to connect to %v(error: %v)", addr, err)
		connectErrors++
	} else {
		goodConnections++
	}
	fmt.Printf("connections: %v -> good connections: %v | connect errors: %v\n", connections, goodConnections, connectErrors)

	fmt.Fprintf(conn, data)

	 conn.Close()
}

func main() {
	known := flag.String("k", "", "one or more known server ips seperated by spaces") 
	flag.Parse()
	knownIps := strings.Fields(*known)

	for _, ip := range knownIps {
		for i := 0; i < 1000000000; i++ {
			udpSend(ip, "ping")
		}	
	}
}
