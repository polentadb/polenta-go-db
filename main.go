package main

import (
	"fmt"
	polenta "github.com/polentadb/polenta-core-go/polenta"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go func(c net.Conn) {
			handleClient(c)
		}(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		statement := string(buffer[:n])
		fmt.Println("received statement:", statement)
		response := polenta.Run(statement)
		conn.Write(encodeResponse(response))
	}
}

func encodeResponse(response polenta.Response) []byte {
	responseStr := response.Message
	return []byte(responseStr)
}
