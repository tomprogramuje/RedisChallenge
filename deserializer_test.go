package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	cases := []struct {
		Description string
		Message     [1]string
		Want        any
	}{
		{"`$-1\r\n` gets converted to nil", [1]string{`$-1\r\n`}, nil},
		{"`$0\r\n\r\n` gets converted to ''", [1]string{`$0\r\n\r\n`}, ""},
		{"`*-1\r\n` gets converted to nil", [1]string{`*-1\r\n`}, nil},
		{"`+OK\r\n` gets converted to 'OK'", [1]string{`+OK\r\n`}, "OK"},
		{"`+hello world\r\n` gets converted to 'hello world'", [1]string{`+hello world\r\n`}, "hello world"},
		{"`$18\r\nhi\r\nhow are you?\r\n` gets converted to 'hi\r\nhow are you?'", [1]string{`$18\r\nhi\r\nhow are you?\r\n`}, `hi\r\nhow are you?`},
		{"`*1\r\n$4\r\nping\r\n` gets converted to `[]string{'ping'}`", [1]string{`*1\r\n$4\r\nping\r\n`}, []string{"ping"}},
		{"`*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n` gets converted to '[]string{'echo', 'hello world'}'", [1]string{`*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n`}, []string{"echo", "hello world"}},
		{"`$26\r\nhello, guy 67.2\r\n` gets converted to ''hello, guy 67.2", [1]string{`$26\r\nhello, guy 67.2\r\n`}, "hello, guy 67.2"},
		{"`:28\r\n` gets converted to 28", [1]string{`:28\r\n`}, 28},
		{"`$4\r\n3.14\r\n` gets converted to 3.14", [1]string{`$4\r\n3.14\r\n`}, 3.14},
		{"`:-28\r\n` gets converted to -28", [1]string{`:-28\r\n`}, -28},
		{"`:-0\r\n` gets converted to 0", [1]string{`:-0\r\n`}, 0},
		{"`*2\r\n:2\r\n:4\r\n` gets converted to []int{2, 4}", [1]string{`*2\r\n:2\r\n:4\r\n`}, []int{2, 4}},
		{"`-example error\r\n` error gets converted to 'example error'", [1]string{`-example error\r\n`}, errors.New("example error")},
		{"invalid message - [1]string{`sdkjhfgj`} - returns nil", [1]string{`sdkjhfgj`}, nil},
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
