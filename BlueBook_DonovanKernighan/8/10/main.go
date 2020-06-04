package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Chat starting")
	listener, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client chan<- string // Clients inbox msg

var (
	clients  = make(map[client]bool) // Current set of connected clients
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// The broadcaster thread manages the clients and messages traffic
func broadcaster() {
	// good illustration of how *select* works
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			sendMsg(msg)
		case cli := <-entering:
			// New client
			clients[cli] = true
		case cli := <-leaving:
			// Client leaving
			delete(clients, cli)
			close(cli)
		}
	}

}

func sendMsg(msg string) {
	for cli := range clients {
		cli <- msg
	}
}

func handleConn(conn net.Conn) {
	// wouldnt better ch := client{}
	ch := make(chan string) // outgoing client messages

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// Scan from stdout
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// receives messages broadcast to the client’s outgoing message channel and writes them to the client’s network connection
	// The client writer’s loop terminates when the broadcaster closesthe channel after receiving a leaving notification.
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

// TODO: Exercise 8.13: Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes.
