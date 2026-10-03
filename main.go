package main

import (
	"log"
	"net"

	functions "netcat/server"
)

func main() {
	logFile, err := functions.Logger()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	port := functions.Validation()
	if port == "" {
		return
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("Listening on the port :" + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go functions.HandleClient(conn)
	}
}
