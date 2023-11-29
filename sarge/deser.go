package sarge

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type RESPType int

const (
	SimpleString = iota
	Error
	Integer
	BulkString
	Array
)

func ParseRESPType(input string) (RESPType, error) {
	switch input[0] {
	case '+':
		return SimpleString, nil
	case '-':
		return Error, nil
	case ':':
		return Integer, nil
	case '$':
		return BulkString, nil
	case '*':
		return Array, nil
	default:
		// TODO: don't return Error type
		return Error, errors.New("Invalid RESP type " + input)
	}
}

func ParseBulkString(input string) string {
	reader := readerForString(input[1:])

	// read the string length
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	// convert the string length to an int
	length, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}

	// read the string
	buf := make([]byte, length)
	_, err = reader.Read(buf)
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func ParseSimpleString(input string) string {
	reader := readerForString(input[1:])

	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	return string(line)
}

func ParseError(input string) string {
	return "Not Implemented"
}

func ParseInteger(input string) int {
	return 0
}

func ParseArray(input string) []string {
	return []string{}
}

func readerForString(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}
