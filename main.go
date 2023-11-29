package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/aziflaj/caboose/megahash"
	"github.com/aziflaj/caboose/omalley"
	"github.com/aziflaj/caboose/sarge"
)

func main() {
	l, err := net.Listen("tcp", ":6900")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// global store
	store := megahash.NewMegahashTable()
	log.Println("Listening on port 6900")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go requestHandler(conn, store)
	}
}

func requestHandler(conn net.Conn, store *megahash.MegahashTable) {
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
	// log.Println(reqArr)

	// Step 2: Very Intricate AI (deeply nested if-else)
	command := reqArr[0]
	args := reqArr[1:]
	// log.Println(command, args)
	res := omalley.Execute(store, command, args)

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
