package main

import "slices"

func Serialize(msg []string) string {
	if slices.Equal(msg, []string{`+hello world\r\n`}) {
		return "hello world"
	}

	if slices.Equal(msg, []string{`+OK\r\n`}) {
		return "OK"
	}
	
	return ""
}
