package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "Hello, world!\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	for {
		fmt.Println(status)
	}
	// if err != nil {
	// 	fmt.Println("Error: %s", err)
	// }
	// buff := make([]byte, 0, 256)

	// for {
	// 	_, err := conn.Read(buff)
	// 	if err != nil {
	// 		fmt.Println("Error: %s", err)
	// 		break
	// 	}
	// 	fmt.Println(buff)
	// }
}
