package sarge

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type RESPType int

const (
	WTF RESPType = iota
	SimpleString
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
		return WTF, errors.New("Invalid RESP type " + input)
	}
}

func ParseBulkString(input string) string {
	reader := readerForString(input)

	// read the string length
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	strLine := string(line)

	// convert the string length to an int
	length, err := strconv.Atoi(strLine[1:])
	if err != nil {
		panic(err)
	}

	str, err := readStringFromReader(reader, length)
	if err != nil {
		panic(err)
	}

	return str
}

func readStringFromReader(reader *bufio.Reader, length int) (string, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line[:length]), nil
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
	reader := readerForString(input[1:])
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	return string(line)
}

func ParseInteger(input string) int {
	reader := readerForString(input[1:])

	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}

	return i
}

func ParseArray(input string) []string {
	reader := readerForString(input[1:])

	// read the array length
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	length, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}

	result := make([]string, 0, length)

	for i := 0; i < length; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		strLine := string(line)

		// find the type of the line
		respType, err := ParseRESPType(strLine)
		if err != nil {
			panic(err)
		}

		// parse the line based on its type
		switch respType {
		case SimpleString:
			// TODO: parse next line and append
		case Error:
			result = append(result, ParseError(strLine))
		case Integer:
			result = append(result, strconv.Itoa(ParseInteger(strLine)))
		case BulkString:
			length, err := strconv.Atoi(strLine[1:])
			if err != nil {
				panic(err)
			}

			str, err := readStringFromReader(reader, length)
			if err != nil {
				panic(err)
			}

			result = append(result, str)
		case Array:
			fallthrough // we're not handling nested arrays
		default:
			panic("WTF")
		}
	}

	return result
}

func readerForString(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}
