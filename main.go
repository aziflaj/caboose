package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/aziflaj/caboose/sarge"
)

func main() {
	l, err := net.Listen("tcp", ":6900")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	log.Println("Listening on port 6900")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go requestHandler(conn)
	}
}

func requestHandler(conn net.Conn) {
	data, err := readFromConn(conn)
	if err != nil {
		panic(err)
	}

	// Step 1: parse RESP
	req, err := sarge.Deserialize(string(data))
	if err != nil {
		panic(err)
	}

	reqArr := req.([]string)

	// Step 2: DO THE AI (extremely nested if-else)
	var res string
	if reqArr[0] == "PING" {
		if len(reqArr) > 1 {
			res = sarge.SerializeArray(reqArr[1:])
		} else {
			res = sarge.SerializeBulkString("PONG")
		}
	} else if reqArr[0] == "ECHO" {
		res = sarge.SerializeArray(reqArr[1:])
	} else {
		// format given command
		command := strings.Join(reqArr, " ")
		res = sarge.SerializeError("Unknown command: " + command)
	}

	// Step 3: respond with RESP
	io.Copy(conn, strings.NewReader(res))

	// Step 4: ???
	// Step 5: Profit
	conn.Close()
}

func readFromConn(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}
