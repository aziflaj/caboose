package sarge_test

import (
	"testing"

	"github.com/aziflaj/caboose/sarge"
)

func TestSerializeSimpleString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present SimpleString", "howdy", "+howdy\r\n"},
		{"Empty SimpleString", "", "+\r\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleString := sarge.SerializeSimpleString(tc.input)
			if simpleString != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, simpleString)
			}
		})
	}
}

func TestSerializeError(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present Error",
			"ERR what did you break this time?",
			"-ERR what did you break this time?\r\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := sarge.SerializeError(tc.input)
			if err != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, err)
			}
		})
	}
}

func TestSerializeInteger(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"Positive Integer", 42, ":42\r\n"},
		{"Negative Integer", -38, ":-38\r\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i := sarge.SerializeInteger(tc.input)
			if i != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, i)
			}
		})
	}
}

func TestSerializeBulkString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present BulkString", "To be or not to be", "$18\r\nTo be or not to be\r\n"},
		{"Empty BulkString", "", "$0\r\n\r\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bulkString := sarge.SerializeBulkString(tc.input)
			if bulkString != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, bulkString)
			}
		})
	}
}

func TestSerializeArray(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Empty Array", []string{}, "*0\r\n"},
		{
			"Present Array",
			[]string{"hello", "world", "!"},
			"*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$1\r\n!\r\n",
		},
		{
			"Echo command",
			[]string{"ECHO", "hello"},
			"*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n",
		},
		{
			"Set command",
			[]string{"SET", "hello", "world"},
			"*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			array := sarge.SerializeArray(tc.input)
			if array != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, array)
			}
		})
	}

	// TODO: Fix this test
	// t.Run("Integer Array", func(t *testing.T) {
	// 	input := []int{-1, 2, 30, 7}
	// 	expected := "*4\r\n:-1\r\n:+2\r\n:30\r\n:7\r\n"

	// 	array := sarge.SerializeArray[int](input)
	// 	if array != expected {
	// 		t.Errorf("Expected %v, got %v", expected, array)
	// 	}
	// })
}
