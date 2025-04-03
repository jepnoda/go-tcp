package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()

	if error != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	listener, error := net.Listen("tcp", ":"+port)
	if error != nil {
		fmt.Println("Error Starting TCP Server: ", error)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("TCP Server is listening on port " + port)
	for {
		conn, error := listener.Accept()
		if error != nil {
			fmt.Println("Error accepting connection: ", error)
			continue
		}
		go handleTcpConnection(conn)
	}
}

func handleTcpConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New Client Connected: ", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Recieved: ", text)
		conn.Write([]byte("Echo: " + text + "\n"))
	}

	if error := scanner.Err(); error != nil {
		fmt.Println("Connection Error: ", error)
	}

	fmt.Println("Client Disconnected ", conn.RemoteAddr())
}
