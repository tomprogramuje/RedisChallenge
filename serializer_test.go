package main

import "testing"

func TestSerialize(t *testing.T) {
	cases := []struct {
		Description string
		Message     []string
		Want        string
	}{
		{"`+OK\r\n` gets converted to 'OK'", []string{`+OK\r\n`}, "OK"},
		{"`+hello world\r\n` gets converted to 'hello world'", []string{`+hello world\r\n`}, "hello world"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Serialize(test.Message)
			if got != test.Want {
				t.Errorf("failed conversion, got %s want %s", got, test.Want)
			}
		})
	}
}
