package main

import (
	"net"
	"testing"
)

func TestPortIsAvailable(t *testing.T) {
	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		t.Fatalf("port 8989 is not empty: %v", err)
	}
	defer listener.Close()
}
