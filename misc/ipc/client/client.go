package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    socketPath := "/home/mugund10/tourofgo/misc/ipc/run"

    conn, err := net.Dial("unix", socketPath)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    message := "Hello from client"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error sending message:", err)
        os.Exit(1)
    }

    var buf [1024]byte
    n, err := conn.Read(buf[:])
    if err != nil {
        fmt.Println("Error reading response:", err)
        os.Exit(1)
    }

    response := string(buf[:n])
    fmt.Println("Received response:", response)
}
