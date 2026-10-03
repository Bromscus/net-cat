package client

import (
	"bufio"
	"fmt"
	"log"
)

// Continuously receive messages from server
func receiver(serverScanner *bufio.Scanner) {
	for serverScanner.Scan() {
		fmt.Println(serverScanner.Text())
	}
	if err := serverScanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}
