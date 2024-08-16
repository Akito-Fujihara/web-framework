package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error: ", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		time.Sleep(2 * time.Second)
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Data received: ", string(buf[:n]))

		if string(buf[:n]) == `"close"` {
			conn.Close()
			return
		}

		responseData := "response data"
		responseDataBytes, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		_, err = conn.Write(responseDataBytes)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
}
