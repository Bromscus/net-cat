package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Start typing: ")
	for scanner.Scan() {
		if scanner.Text() == "ping" {
			time.Sleep(500 * (time.Millisecond))
			fmt.Println("pong")
		}
		if strings.ToLower(scanner.Text()) == "exit" {
			os.Exit(0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
