package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message received:", string(msg))
		newMsg := strings.ToUpper(msg)
		conn.Write([]byte(newMsg + "\n"))
	}
}
