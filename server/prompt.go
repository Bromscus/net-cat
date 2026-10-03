package server

import (
	"fmt"
	"time"
)

func prompt(name string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	return fmt.Sprintf("[%s][%s]:", timestamp, name)
}
