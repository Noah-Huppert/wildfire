package main

import (
	"fmt"
	"net"
	"bufio"
	"flag"
	"strings"
)

func udpSend(addr string, data string) {
	buffer := make([]byte, 2048)

	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("Failed to connect to %v(error: %v)", addr, err)
	}

	fmt.Fprintf(conn, "ping")
	 _, err = bufio.NewReader(conn).Read(buffer)
	 if err != nil {
		 fmt.Printf("Failed to read respond (error: %v)", err)
	 } else {
		 fmt.Printf("buffer: %v", buffer)
	 }
}

func main() {
	var known = flag.String("k", "", "one or more known server ips seperated by spaces") 
	flag.Parse()
	known = strings.Split(known, " ")

	for _, ip := range known {
		udpSend(ip, "hello")
	}
}
