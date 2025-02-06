// Код сервера
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving  = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {
	fmt.Println("Starting server on port 8000")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server is running. Waiting for connections...")
	go broadcaster()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}


// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}