package server

import (
	"fmt"
	"net"
)

func removeClient(conn net.Conn) {
	chat.Lock()
	defer chat.Unlock()

	for i, client := range chat.clients {
		if client.conn == conn {
			name := client.name

			// Remove client
			chat.clients = append(
				chat.clients[:i], chat.clients[i+1:]...)

			message := fmt.Sprintf("%s HAS LEFT THE CHAT.", name)

			// Tell remaining clients
			for _, client := range chat.clients {
				fmt.Fprintln(client.conn)
				fmt.Fprintln(client.conn, message)
				fmt.Fprint(client.conn, prompt(client.name))
			}

			fmt.Println(message)
			return
		}
	}
}
