package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	cases := []struct {
		Description string
		Message     string
		Want        any
	}{
		{"`$-1\r\n` gets converted to nil", "$-1\r\n", nil},
		{"`$0\r\n\r\n` gets converted to ''", "$0\r\n\r\n", ""},
		{"`*-1\r\n` gets converted to nil", "*-1\r\n", nil},
		{"`+OK\r\n` gets converted to 'OK'", "+OK\r\n", "OK"},
		{"`+hello world\r\n` gets converted to 'hello world'", "+hello world\r\n", "hello world"},
		{"`$18\r\nhi\r\nhow are you?\r\n` gets converted to 'hi\r\nhow are you?'", "$16\r\nhi\r\nhow are you?\r\n", "hi\r\nhow are you?"},
		{"`*1\r\n$4\r\nping\r\n` gets converted to `[]string{'ping'}`", "*1\r\n$4\r\nping\r\n", []string{"ping"}},
		{"`*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n` gets converted to '[]string{'echo', 'hello world'}'", "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n", []string{"echo", "hello world"}},
		{"`$26\r\nhello, guy 67.2\r\n` gets converted to ''hello, guy 67.2", "$26\r\nhello, guy 67.2\r\n", "hello, guy 67.2"},
		{"`:28\r\n` gets converted to 28", ":28\r\n", 28},
		{"`$4\r\n3.14\r\n` gets converted to 3.14", "$4\r\n3.14\r\n", 3.14},
		{"`:-28\r\n` gets converted to -28", ":-28\r\n", -28},
		{"`:-0\r\n` gets converted to 0", ":-0\r\n", 0},
		{"`*2\r\n:2\r\n:4\r\n` gets converted to []int{2, 4}", "*2\r\n:2\r\n:4\r\n", []int{2, 4}},
		{"`-example error\r\n` error gets converted to 'example error'", "-example error\r\n", errors.New("example error")},
		{"invalid message - `sdkjhfgj` - returns nil", "sdkjhfgj", nil},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Deserialize(test.Message)
			if !reflect.DeepEqual(got, test.Want) {
				t.Error("failed conversion,", got, "doesn't equal", test.Want)
			}
		})
	}
}
