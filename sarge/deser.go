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

func Deserialize(input string) (interface{}, error) {
	respType, err := assessRESPType(input)
	if err != nil {
		return nil, err
	}

	switch respType {
	case SimpleString:
		return ParseSimpleString(input)
	case Error:
		return ParseError(input)
	case Integer:
		return ParseInteger(input)
	case BulkString:
		return ParseBulkString(input)
	case Array:
		return ParseArray(input)
	default:
		return nil, errors.New("WTF")
	}
}

func ParseBulkString(input string) (string, error) {
	reader := readerForString(input)

	// read the string length
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	strLine := string(line)

	str, err := readBulkString(reader, strLine[1:])
	if err != nil {
		return "", err
	}

	return str, nil
}

func ParseSimpleString(input string) (string, error) {
	reader := readerForString(input[1:])

	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line), nil
}

func ParseError(input string) (string, error) {
	reader := readerForString(input[1:])
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line), nil
}

func ParseInteger(input string) (int, error) {
	reader := readerForString(input[1:])

	line, _, err := reader.ReadLine()
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ParseArray(input string) ([]string, error) {
	reader := readerForString(input[1:])

	// read the array length
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}

	length, err := strconv.Atoi(string(line))
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, length)

	for i := 0; i < length; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		strLine := string(line)

		// find the type of the line
		respType, err := assessRESPType(strLine)
		if err != nil {
			return nil, err
		}

		// parse the line based on its type
		switch respType {
		case SimpleString:
			// TODO: parse next line and append
		case Error:
			resp, err := ParseError(strLine)
			if err != nil {
				return nil, err
			}
			result = append(result, resp)
		case Integer:
			resp, err := ParseInteger(strLine)
			if err != nil {
				return nil, err
			}
			result = append(result, strconv.Itoa(resp))
		case BulkString:
			str, err := readBulkString(reader, strLine[1:]) // [1:] to skip the leading '$'
			if err != nil {
				return nil, err
			}

			result = append(result, str)
		case Array:
			fallthrough // we're not handling nested arrays
		default:
			return nil, errors.New("WTF")
		}
	}

	return result, nil
}

func readerForString(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}

func assessRESPType(input string) (RESPType, error) {
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

func readBulkString(reader *bufio.Reader, strLine string) (string, error) {
	// read the string length
	length, err := strconv.Atoi(strLine)
	if err != nil {
		return "", err
	}

	// read the actual string
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line[:length]), nil
}
