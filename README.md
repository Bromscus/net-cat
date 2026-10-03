# NetCat

## Description

The project main idea is to build a lightweight TCP chat application in Go. It allows multiple users to connect to the same server, choose a username, send messages in real time, and receive join and leave notifications from the other participants.

The server manages connected clients, preserves chat history for new users, and broadcasts incoming messages to all clients except the sender. The application is designed to be minimal, reliable and easy to run from the terminal.

## Authors

- ahmsalem
-asami

## Usage: how to run

### Run directly with Go

1. From the root project directory, start the server:

```bash
go run .
```

You can also specify a custom port:

```bash
go run . 8989
```

2. Open a second terminal and connect to the server using a TCP client such as Netcat:

```bash
nc localhost 8989
```

3. Enter a username when prompted.

4. Start sending messages to the chat room.

5. Open additional terminals to simulate multiple users in the same chat.

## Implementation details: algorithm

1. Start a TCP listener on the chosen port.
2. Accept incoming connections and handle each client in its own goroutine.
3. Validate the username to ensure it is not empty, already in use, or the server is full.
4. Store each connected client in the shared chat state.
5. Send previous chat history to new clients when they join.
6. Read incoming text from each user and format it with a timestamp and username.
7. Broadcast each message to all other clients.
8. Refresh the prompt for the remaining users after each message.
9. Detect client disconnects and notify the rest of the chat.
10. Continue handling incoming connections until the server stops.

## Project Structure

```txt
net-cat/
├── go.mod
├── main.go
├── README.md
├── client/
│   ├── handler.go
│   ├── receiver.go
│   └── sender.go
└── server/
    ├── adder.go
    ├── broadcast.go
    ├── handler.go
    ├── prompt.go
    ├── reader.go
    ├── remover.go
    └── validation.go
```
