package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func HandleClient(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	fmt.Fprintln(conn,
		"       _nnnn_\n"+
			"      dGGGGMMb\n"+
			"     @p~qp~~qMb\n"+
			"     M|@||@) M|\n"+
			"     @,----.JM|\n"+
			"    JS^\\__/  qKL\n"+
			"   dZP        qKRb\n"+
			"  dZP          qKKb\n"+
			" fZP            SMMb\n"+
			" HZM            MMMM\n"+
			" FqM            MMMM\n"+
			"__| \".        |\\dS\"qML\n"+
			"|    `.       | `' \\Zq\n"+
			"_)      \\.___.,|     .'\n"+
			"\\____   )MMMMMP|   .'\n"+
			"     `-'       `--'",
	)

	var name string

	// We get a non empty name
	for {
		fmt.Fprint(conn, "[ENTER YOUR NAME]: ")

		if !scanner.Scan() {
			return
		}

		name = strings.TrimSpace(scanner.Text())

		if name == "" {
			fmt.Fprintln(conn, "ERROR: name cannot be empty")
			continue
		}

		if errMessage := addClient(name, conn); errMessage != "" {
			fmt.Fprintln(conn, errMessage)

			if errMessage == "ERROR: server is full" {
				return
			}

			continue
		}

		break
	}

	defer removeClient(conn)

	log.Printf("%s has connected", name)

	readMessages(conn, scanner, name)

	log.Printf("%s disconnected", name)
}
