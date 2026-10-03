package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handler() {
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	serverScanner := bufio.NewScanner(conn)

	if !serverScanner.Scan() {
		return
	}
	fmt.Println(serverScanner.Text())

	if !scanner.Scan() {
		return
	}

	fmt.Fprintln(conn, scanner.Text())

	// Read welcome/error response
	if !serverScanner.Scan() {
		return
	}

	response := serverScanner.Text()

	if strings.HasPrefix(response, "ERROR:") {
		fmt.Println(response)
		return
	}
	fmt.Println(response)

	go receiver(serverScanner)
	sender(scanner, conn)
}
