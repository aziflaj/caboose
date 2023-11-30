package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/aziflaj/caboose/vic"
)

func main() {
	l, err := net.Listen("tcp", ":6900")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// global store
	store := vic.NewKVStore()
	log.Println("Listening on port 6900")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go requestHandler(conn, store)
	}
}

func requestHandler(conn net.Conn, store *vic.KVStore) {
	data, err := readFromConn(conn)
	if err != nil {
		panic(err)
	}

	res := vic.HandleRequest(store, data)

	log.Println(res)

	io.Copy(conn, strings.NewReader(res))
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
