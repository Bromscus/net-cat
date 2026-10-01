package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type ClientInfo struct {
	sync.Mutex
	nextID int
	names  []string
}

var clientInfo ClientInfo

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

		clientInfo.Lock()
		clientInfo.nextID++
		id := clientInfo.nextID
		clientInfo.Unlock()

		go handleClient(conn, id)
	}
}

func handleClient(conn net.Conn, id int) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Printf("Client %d name read error: %v", id, err)
		}
		return
	}

	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		fmt.Fprintln(conn, "ERR: name cannot be empty")
		return
	}

	clientInfo.Lock()
	for _, existingName := range clientInfo.names {
		if strings.EqualFold(existingName, name) {
			clientInfo.Unlock()
			fmt.Fprintln(conn, "ERR: name is already in use")
			return
		}
	}
	clientInfo.names = append(clientInfo.names, name)
	clientInfo.Unlock()
	defer func() {
		clientInfo.Lock()
		for index, existingName := range clientInfo.names {
			if existingName == name {
				clientInfo.names = append(clientInfo.names[:index], clientInfo.names[index+1:]...)
				break
			}
		}
		clientInfo.Unlock()
	}()

	fmt.Fprintf(conn, "OK: Welcome %s (ID: %d)\n", name, id)
	fmt.Printf("Client %d connected as %s\n", id, name)

	for scanner.Scan() {
		message := scanner.Text()

		fmt.Printf("Received from %s (ID: %d): %s\n", name, id, message)

		if strings.EqualFold(strings.TrimSpace(message), "ping") {
			fmt.Fprintln(conn, "pong")
		} else {
			fmt.Fprintln(conn, "Server received:", message)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}

	fmt.Printf("Client %s (ID: %d) disconnected\n", name, id)
}
