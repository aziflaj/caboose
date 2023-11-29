package main

import (
	"fmt"

	"github.com/aziflaj/caboose/sarge"
)

func main() {
	var input string = "$5\r\nhowdy\r\n"

	respType, err := sarge.ParseRESPType(input)
	if err != nil {
		panic(err)
	}

	switch respType {
	case sarge.SimpleString:
		fmt.Println(sarge.ParseSimpleString(input))
	case sarge.Error:
		fmt.Println(sarge.ParseError(input))
	case sarge.Integer:
		fmt.Println(sarge.ParseInteger(input))
	case sarge.BulkString:
		fmt.Println(sarge.ParseBulkString(input))
	case sarge.Array:
		fmt.Println(sarge.ParseArray(input))
	default:
		panic("WTF")
	}
}
