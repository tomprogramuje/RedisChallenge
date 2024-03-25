package main

import (
	"slices"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		Description string
		Data        any
		Want        []string
	}{
		{"'OK' gets converted to `+OK\r\n`", "OK", []string{`+OK\r\n`}},
		{"'hello world' gets converted to `+hello world\r\n`", "hello world", []string{`+hello world\r\n`}},
		{"`[]string{'ping'}` gets converted to '*1\r\n$4\r\nping\r\n'", []string{"ping"}, []string{`*1\r\n$4\r\nping\r\n`}},
		{"'[]string{'echo', 'hello world'}' gets converted to `*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n`", []string{"echo", "hello world"}, []string{`*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n`}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Serialize(test.Data)
			if !slices.Equal(got, test.Want) {
				t.Errorf("failed conversion, got %s want %s", got, test.Want)
			}
		})
	}
}
