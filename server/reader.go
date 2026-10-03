package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func readMessages(conn net.Conn, scanner *bufio.Scanner, name string) {
	for {
		// Show prompt before waiting for message
		fmt.Fprint(conn, prompt(name))

		if !scanner.Scan() {
			break
		}

		message := strings.TrimSpace(scanner.Text())

		// Do not broadcast empty messages
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")

		formattedMessage := fmt.Sprintf(
			"[%s][%s]:%s",
			timestamp,
			name,
			message,
		)

		broadcast(formattedMessage, conn)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}
}
