package sarge_test

import (
	"testing"

	"github.com/aziflaj/caboose/sarge"
)

func TestParseSimpleString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present SimpleString", "+howdy\r\n", "howdy"},
		{"Empty SimpleString", "+\r\n", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleString, _ := sarge.ParseSimpleString(tc.input)
			if simpleString != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, simpleString)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present Error",
			"-ERR what did you break this time?\r\n",
			"ERR what did you break this time?",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err, _ := sarge.ParseError(tc.input)
			if err != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, err)
			}
		})
	}
}

func TestParseInteger(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Positive Integer", ":42\r\n", 42},
		{"Explicitly Positive Integer", ":+42\r\n", 42},
		{"Negative Integer", ":-38\r\n", -38},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			integer, _ := sarge.ParseInteger(tc.input)
			if integer != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, integer)
			}
		})
	}
}

func TestParseBulkString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Present BulkString", "$5\r\nhowdy\r\n", "howdy"},
		{"Shorter BulkString", "$3\r\nhey how are you\r\n", "hey"},
		{"Empty BulkString", "$0\r\n\r\n", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bulkString, _ := sarge.ParseBulkString(tc.input)
			if bulkString != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, bulkString)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{"Empty Array", "*0\r\n", []string{}},
		{"Integer Array", "*4\r\n:-1\r\n:+2\r\n:30\r\n:7\r\n", []string{"-1", "2", "30", "7"}},
		{
			"Present Array",
			"*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$1\r\n!\r\n",
			[]string{"hello", "world", "!"},
		},
		{
			"Echo command",
			"*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n",
			[]string{"ECHO", "hello"},
		},
		{
			"Set command",
			"*3\r\n$3\r\nSET\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			[]string{"SET", "hello", "world"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			array, _ := sarge.ParseArray(tc.input)
			if len(array) != len(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, array)
			}

			for i := range array {
				if array[i] != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected, array)
				}
			}
		})
	}
}
