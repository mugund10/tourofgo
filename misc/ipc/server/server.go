package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	socketPath := "/home/mugund10/tourofgo/misc/ipc/run"

	// Remove any existing socket file
	if err := os.RemoveAll(socketPath); err != nil {
		panic(err)
	}

	// Create a Unix domain socket listener
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	message := string(buf[:n])
	fmt.Println("Received message:", message)

	response := "Message received: " + message
	conn.Write([]byte(response))
}
