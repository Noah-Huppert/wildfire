package main

import (
	"fmt"
	"net"
	"flag"
)

func main() {
	listenErrors := 0
	readErrors := 0
	connections := 0
	goodConnections := 0

	host := flag.String("host", "", "server host address")
	port := flag.Int("port", -1, "server port")
	flag.Parse()

	receiveBuff := make([]byte, 2048)
	addr := net.UDPAddr{
		IP: net.ParseIP(*host),
		Port: *port,
	}

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error listening on %v:%v (error: %v)\n", *host, *port, err)
		listenErrors++
	}

	for {
		_, _, err := server.ReadFromUDP(receiveBuff)
		connections++
		if err != nil {
			fmt.Printf("Failed to receive: %v", err)
			readErrors++
		} else {
			goodConnections++
		}


		fmt.Printf("connections: %v -> good connections: %v | listen errors: %v | readErrors: %v\n", connections, goodConnections, listenErrors, readErrors)
	}
}
