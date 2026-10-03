package main

import (
	"fmt"
	"log"
	"net"

	functions "netcat/server"
)

func main() {
	port := functions.Validation()
	if port == "" {
		return
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Listening on the port :" + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go functions.HandleClient(conn)
	}
}
