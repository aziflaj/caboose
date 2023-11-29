package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"

	"github.com/aziflaj/caboose/sarge"
)

func main() {
	l, err := net.Listen("tcp", ":6900")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("Listening on port 6900")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			// io.WriteString(conn, "Hello from TCP server\n")
			buffer := make([]byte, 1024)
			var bufferSize int

			// buffer read loop
			for {
				n, err := conn.Read(buffer)
				bufferSize += n
				if errors.Is(err, io.EOF) {
					fmt.Println("eof reached")
					break
				} else if err != nil {
					panic(err)
				}

				break
			}

			data := buffer[:bufferSize]

			// fmt.Println(string(data))
			// parse RESP
			req, err := sarge.Deserialize(string(data))
			if err != nil {
				panic(err)
			}

			fmt.Println(req.([]string))
			reqArr := req.([]string)
			fmt.Println(reflect.TypeOf(req))

			// DO THE AI (if-else)
			var res string
			if reqArr[0] == "PING" {
				if len(reqArr) > 1 {
					res = sarge.SerializeArray(reqArr[1:])
				} else {
					res = sarge.SerializeBulkString("PONG")
				}
			}

			// respond with RESP
			io.Copy(conn, strings.NewReader(res))

			// ???
			// profit
			conn.Close()
		}(conn)
	}
}
