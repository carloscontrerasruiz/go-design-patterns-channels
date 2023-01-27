package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	port = flag.Int("p", 3090, "Port")
	host = flag.String("h", "localhost", "Host")
)

func HandlerConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWriter(conn, message)

	clentName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clentName)

	messages <- fmt.Sprintf("Client conected %s\n", clentName)
	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)

	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s", clentName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s say goodbye", clentName)
}

func MessageWriter(conn net.Conn, messages <-chan string) {
	for _, message := range <-messages {
		fmt.Fprintln(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Println("Error fatal")
		log.Fatal(err)
	}

	go Broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go HandlerConnection(conn)

	}

}
