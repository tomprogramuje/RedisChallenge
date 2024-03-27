package main

import (
	"strconv"
	"strings"
)

func Deserialize(msg [1]string) any {

	dataType := string(msg[0][0])

	switch dataType {
		
	case "+":
		withoutPrefix, _ := strings.CutPrefix(msg[0], `+`)
		string, _ := strings.CutSuffix(withoutPrefix, `\r\n`)

		return string

	case "$":
		withoutPrefix, _ := strings.CutPrefix(msg[0], `$`)
		bulkString, _ := strings.CutSuffix(withoutPrefix, `\r\n`)

		if bulkString == "-1" {
			return nil
		}

		return bulkString

	case ":":
		withoutPrefix, _ := strings.CutPrefix(msg[0], `:`)
		withoutSuffix, _ := strings.CutSuffix(withoutPrefix, `\r\n`)

		number, _ := strconv.Atoi(withoutSuffix)

		return number

	default:

		return nil
	}
}
