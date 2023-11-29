package sarge

import "fmt"

func SerializeSimpleString(input string) string {
	return fmt.Sprintf("+%s\r\n", input)
}

func SerializeError(input string) string {
	return fmt.Sprintf("-%s\r\n", input)
}

func SerializeInteger(input int) string {
	return fmt.Sprintf(":%d\r\n", input)
}

func SerializeBulkString(input string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
}

func SerializeArray(input []string) string {
	return fmt.Sprintf("*%d\r\n%s", len(input), serializeArrayElements(input))
}

func serializeArrayElements(input []string) string {
	var result string
	for _, element := range input {
		result += SerializeBulkString(element)
	}
	return result
}
