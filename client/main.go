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
	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	serverScanner := bufio.NewScanner(conn)
	fmt.Println(
		"       _nnnn_\n" +
			"      dGGGGMMb\n" +
			"     @p~qp~~qMb\n" +
			"     M|@||@) M|\n" +
			"     @,----.JM|\n" +
			"    JS^\\__/  qKL\n" +
			"   dZP        qKRb\n" +
			"  dZP          qKKb\n" +
			" fZP            SMMb\n" +
			" HZM            MMMM\n" +
			" FqM            MMMM\n" +
			"__| \".        |\\dS\"qML\n" +
			"|    `.       | `' \\Zq\n" +
			"_)      \\.___.,|     .'\n" +
			"\\____   )MMMMMP|   .'\n" +
			"     `-'       `--'",
	)
	fmt.Print("[ENTER YOUR NAME]: ")

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
	fmt.Println("Start typing:")

	// Continuously receive messages from server
	go func() {
		for serverScanner.Scan() {
			fmt.Println(serverScanner.Text())
		}
	}()

	// Continuously send messages to server
	for scanner.Scan() {
		message := scanner.Text()

		if strings.EqualFold(message, "exit") {
			return
		}

		fmt.Fprintln(conn, message)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
	if err := serverScanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}
