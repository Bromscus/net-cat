package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// 1. Connect to the server's IP and Port (Triggers 3-Way Handshake)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close() // Ensure the connection closes at the end

	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(conn)

	fmt.Print("Enter your name: ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println("Error reading name:", err)
		}
		return
	}

	fmt.Fprintln(conn, scanner.Text())
	welcome, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving server response:", err)
		return
	}
	if strings.HasPrefix(welcome, "ERR:") {
		fmt.Print(welcome)
		return
	}
	fmt.Print(welcome)
	fmt.Println("Start typing:")

	for scanner.Scan() {
		message := scanner.Text()

		if strings.EqualFold(message, "exit") {
			return
		}

		fmt.Fprintln(conn, message)

		reply, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received reply: %s", reply)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}
}
