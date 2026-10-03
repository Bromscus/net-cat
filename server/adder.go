package server

import (
	"fmt"
	"net"
	"strings"
)

func addClient(name string, conn net.Conn) string {
	chat.Lock()
	defer chat.Unlock()

	// Check duplicate names
	for _, client := range chat.clients {
		if strings.EqualFold(client.name, name) {
			return "ERROR: name is already in use"
		}
	}

	// Check maximum clients
	if len(chat.clients) >= maxClients {
		return "ERROR: server is full"
	}

	fmt.Fprintf(conn, "Welcome %s\n", name)

	// Send previous messages
	for _, message := range chat.history {
		fmt.Fprintln(conn, message)
	}

	// Add client
	chat.clients = append(chat.clients, Client{
		name: name, conn: conn,
	})

	joinMessage := fmt.Sprintf("%s HAS JOINED THE CHAT", name)

	// Tell everyone else
	for _, client := range chat.clients {
		if client.conn != conn {
			fmt.Fprintln(client.conn)
			fmt.Fprintln(client.conn, joinMessage)
			fmt.Fprint(client.conn, prompt(client.name))
		}
	}

	return ""
}
