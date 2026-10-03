package server

import (
	"log"
	"net"
	"os"
	"sync"
)

type Client struct {
	name string
	conn net.Conn
}

type Chat struct {
	sync.Mutex
	clients []Client
	history []string
}

const maxClients = 10

var chat Chat

func Validation() string {
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		log.Println("[USAGE]: ./TCPChat $port")
		return ""
	}
	return port
}
