package main

import "testing"

func TestSerialize(t *testing.T) {
	cases := []struct {
		Description string
		Message     string
		Want        string
	}{
		{"'OK' gets converted to `+OK\r\n`", "OK", `+OK\r\n`},
		{"'hello world' gets converted to `+hello world\r\n`", "hello world", `+hello world\r\n`},
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
