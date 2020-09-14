package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		} else {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("[.] ", conn.RemoteAddr().String(), " connected.")
	conn.Write([]byte("Hello, World!"))
}
