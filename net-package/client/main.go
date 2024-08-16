package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if response, err := RequestAndResponse(conn, "request data 1"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Response: ", response)
	}

	if response, err := RequestAndResponse(conn, "request data 2"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Response: ", response)
	}

	if response, err := RequestAndResponse(conn, "request data 3"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Response: ", response)
	}

	if response, err := RequestAndResponse(conn, "close"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Response: ", response)
	}
}

func RequestAndResponse(conn net.Conn, requestData string) (string, error) {
	requestBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	_, err = conn.Write(requestBytes)
	if err != nil {
		return "", err
	}

	responseData := make([]byte, 1024)
	n, err := conn.Read(responseData)
	if err != nil {
		return "", err
	}

	return string(responseData[:n]), nil
}
