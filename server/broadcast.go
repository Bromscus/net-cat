package server

import (
	"fmt"
	"net"
)

func broadcast(message string, sender net.Conn) {
	chat.Lock()
	defer chat.Unlock()

	// Save message in history
	chat.history = append(chat.history, message)

	fmt.Println(message)

	// Send to every client except sender
	for _, client := range chat.clients {
		if client.conn != sender {
			// Move away from their current prompt
			fmt.Fprintln(client.conn)

			// Show incoming message
			fmt.Fprintln(client.conn, message)

			// Give them their prompt again
			fmt.Fprint(client.conn, prompt(client.name))
		}
	}
}
