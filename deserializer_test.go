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
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Deserialize(test.Message)
			if !reflect.DeepEqual(got, test.Want) {
				t.Error("failed conversion, got", got, "doesn't equal want", test.Want)
			}
		})
	}
}
