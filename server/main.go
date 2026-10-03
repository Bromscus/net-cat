package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	name string
	conn net.Conn
}

type Chat struct {
	sync.Mutex
	clients []Client
	history []string
}

const maxClients = 10

var chat Chat

func main() {
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: go run . $port")
		return
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server started on port", port)

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

	name, ok := readName(conn, scanner)
	if !ok {
		return
	}

	if errMessage := addClient(name, conn); errMessage != "" {
		fmt.Fprintln(conn, errMessage)
		return
	}

	defer removeClient(conn)

	fmt.Printf("%s has connected\n", name)

	readMessages(conn, scanner, name)

	fmt.Printf("%s disconnected\n", name)
}

func readName(conn net.Conn, scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println("Error reading name:", err)
		}

		return "", false
	}

	name := strings.TrimSpace(scanner.Text())

	if name == "" {
		fmt.Fprintln(conn, "ERROR: name cannot be empty")
		return "", false
	}

	return name, true
}

func addClient(name string, conn net.Conn) string {
	chat.Lock()
	defer chat.Unlock()

	for _, client := range chat.clients {
		if strings.EqualFold(client.name, name) {
			return "ERROR: name is already in use"
		}
	}

	if len(chat.clients) >= maxClients {
		return "ERROR: server is full"
	}

	fmt.Fprintf(conn, "Welcome %s\n", name)

	for _, message := range chat.history {
		fmt.Fprintln(conn, message)
	}

	chat.clients = append(chat.clients, Client{name: name, conn: conn})

	joinMessage := fmt.Sprintf("%s HAS JOINED THE CHAT", name)

	for _, client := range chat.clients {
		if client.conn != conn {
			fmt.Fprintln(client.conn, joinMessage)
		}
	}

	return ""
}

func removeClient(conn net.Conn) {
	chat.Lock()
	defer chat.Unlock()

	for i, client := range chat.clients {
		if client.conn == conn {
			name := client.name

			chat.clients = append(chat.clients[:i], chat.clients[i+1:]...)

			message := fmt.Sprintf("%s HAS LEFT THE CHAT", name)

			for _, client := range chat.clients {
				fmt.Fprintln(client.conn, message)
			}

			fmt.Println(message)

			return
		}
	}
}

func readMessages(conn net.Conn, scanner *bufio.Scanner, name string) {
	for scanner.Scan() {
		message := strings.TrimSpace(scanner.Text())

		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")

		formattedMessage := fmt.Sprintf("[%s][%s]:%s", timestamp, name, message)

		broadcast(formattedMessage, conn)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}
}

func broadcast(message string, sender net.Conn) {
	chat.Lock()
	defer chat.Unlock()

	chat.history = append(chat.history, message)

	fmt.Println(message)

	for _, client := range chat.clients {
		if client.conn != sender {
			fmt.Fprintln(client.conn, message)
		}
	}
}
