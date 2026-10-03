package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)
// Continuously send messages to server
func sender(scanner *bufio.Scanner, conn net.Conn) {
	for scanner.Scan() {
		message := scanner.Text()

		fmt.Fprintln(conn, message)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}
