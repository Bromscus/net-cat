package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()

		fmt.Println("Received:", message)

		if strings.EqualFold(strings.TrimSpace(message), "ping") {
			fmt.Fprintln(conn, "pong")
		} else {
			fmt.Fprintln(conn, "Server received:", message)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}

	fmt.Println("Client disconnected")
}
