package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	cases := []struct {
		Description string
		Data        any
		Want        string
	}{
		{"nil gets converted to `$-1\r\n`", nil, "$-1\r\n"},
		{"'' gets converted to `$0\r\n\r\n`", "", "$0\r\n\r\n"},
		{"[]string{} gets converted to `*-1\r\n`", []string{}, "*-1\r\n"},
		{"'OK' gets converted to `+OK\r\n`", "OK", "+OK\r\n"},
		{"'hello world' gets converted to `+hello world\r\n`", "hello world", "+hello world\r\n"},
		{"'hi\r\nhow are you?' gets converted to `$16\r\nhi\r\nhow are you?\r\n`", "hi\r\nhow are you?", "$16\r\nhi\r\nhow are you?\r\n"},
		{"`[]string{'ping'}` gets converted to `*1\r\n$4\r\nping\r\n`", []string{"ping"}, "*1\r\n$4\r\nping\r\n"},
		{"'[]string{'echo', 'hello world'}' gets converted to `*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n`", []string{"echo", "hello world"}, "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n"},
		{"[]int{2, 4} gets converted to *2\r\n:2\r\n:4\r\n`", []int{2, 4}, "*2\r\n:2\r\n:4\r\n"},
		{"[]float64{3.14, 2.19} gets converted to `*2\r\n$4\r\n3.14\r\n$4\r\n2.19\r\n`", []float64{3.14, 2.19}, "*2\r\n$4\r\n3.14\r\n$4\r\n2.19\r\n"},
		{"28 gets converted to `:28\r\n`", 28, ":28\r\n"},
		{"-28 gets converted to `:-28\r\n`", -28, ":-28\r\n"},
		{"-0 gets converted to `:0\r\n`", -0, ":0\r\n"},
		{"3.14 gets converted to `$4\r\n3.14\r\n`", 3.14, "$4\r\n3.14\r\n"},
		{"'example error' error gets converted to `-example error\r\n`", errors.New("example error"), "-example error\r\n"},
		{"invalid data - map[int]string{1: 'apple'} - returns [1]string{}", map[int]string{1: "apple"}, "invalid data"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Serialize(test.Data)
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("failed conversion, got %s want %s", got, test.Want)
			}
		})
	}
}
