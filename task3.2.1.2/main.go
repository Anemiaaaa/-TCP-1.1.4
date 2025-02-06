package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return
	}
	method, path := parts[0], parts[1]

	if method == "GET" && strings.TrimSpace(path) == "/" {
		response := "HTTP/1.1 200 OK\n" +
			"Content-Type: text/html\n\n" +
			"<!DOCTYPE html>\n<html>\n<head>\n<title>Webserver</title>\n</head>\n<body>\nhello world\n</body>\n</html>"
		conn.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\n" +
			"404 Not Found"
		conn.Write([]byte(response))
	}
}

func main() {
	listener,_ := net.Listen("tcp", ":8080")
	defer listener.Close()
	fmt.Println("Server is running on port 8080...")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go handleConnection(conn)
		}
	}()

	// получаем данные с сервера
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))
}
