package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// client
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(connection, "Hello!")
	buffer := make([]byte, 1024)
	n, err := connection.Read(buffer)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(buffer[:n]))
	connection.Close()
}
