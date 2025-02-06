package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	go main()              // Запускаем сервер в отдельной горутине
	time.Sleep(1 * time.Second) // Даем серверу время запуститься

	// Пробуем подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	conn.Close()
}

func TestMessageExchange(t *testing.T) {

	client1, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Client1 connection error: %v", err)
	}
	defer client1.Close()

	client2, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Client2 connection error: %v", err)
	}
	defer client2.Close()


	reader1 := bufio.NewScanner(client1)
	if reader1.Scan() {
		fmt.Println("Client1 received:", reader1.Text())
	}

	// Читаем приветственное сообщение от сервера для client2
	reader2 := bufio.NewScanner(client2)
	if reader2.Scan() {
		fmt.Println("Client2 received:", reader2.Text())
	}

	// Отправляем сообщение от client1
	message := "Hello from Client1"
	fmt.Fprintln(client1, message)

	// Проверяем, получил ли client2 сообщение
	received := false
	timeout := time.After(2 * time.Second)

	for {
		select {
		case <-timeout:
			if !received {
				t.Fatal("Client2 did not receive the expected message")
			}
			return
		default:
			if reader2.Scan() {
				text := reader2.Text()
				fmt.Println("Client2 received:", text)
				if strings.Contains(text, message) {
					received = true
					return
				}
			}
		}
	}
}


func TestClientDisconnect(t *testing.T) {
	// Подключаем клиента
	client, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	reader := bufio.NewScanner(client)
	if reader.Scan() {
		fmt.Println("Received:", reader.Text())
	}

	// Закрываем соединение
	client.Close()

	// Даем серверу обработать отключение
	time.Sleep(1 * time.Second)
}
