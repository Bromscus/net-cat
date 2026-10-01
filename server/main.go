package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// server
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read from the client
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(buf[:n]))

	fmt.Fprint(conn, "Echo - "+string(buf[:n]))
}
