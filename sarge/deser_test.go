package sarge_test

import (
	"testing"

	"github.com/aziflaj/caboose/sarge"
)

func TestParseRESPType(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected sarge.RESPType
	}{
		{"SimpleString", "+", sarge.SimpleString},
		{"Error", "-", sarge.Error},
		{"Integer", ":", sarge.Integer},
		{"BulkString", "$", sarge.BulkString},
		{"Array", "*", sarge.Array},
		{"Invalid", "!", sarge.WTF},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			respType, _ := sarge.ParseRESPType(tc.input)
			if respType != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, respType)
			}
		})
	}
}

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
			simpleString := sarge.ParseSimpleString(tc.input)
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
			err := sarge.ParseError(tc.input)
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
			integer := sarge.ParseInteger(tc.input)
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
			bulkString := sarge.ParseBulkString(tc.input)
			if bulkString != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, bulkString)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	t.Skip()
}
