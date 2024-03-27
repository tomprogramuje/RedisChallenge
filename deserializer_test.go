package main

import (
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	cases := []struct{
		Description string
		Message [1]string
		Want any
	}{
		{"`$-1\r\n` gets converted to nil", [1]string{`$-1\r\n`}, nil},
		{"`+OK\r\n` gets converted to 'OK'", [1]string{`+OK\r\n`}, "OK"},
		{"`:28\r\n` gets converted to 28", [1]string{`:28\r\n`}, 28},
		{"`$4\r\n3.14\r\n` gets converted to 3.14", [1]string{`$4\r\n3.14\r\n`}, 3.14},
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
