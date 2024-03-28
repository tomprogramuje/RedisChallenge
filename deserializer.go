package main

import (
	"strconv"
	"strings"
)

func Deserialize(msg [1]string) any {

	withoutSuffix, _ := strings.CutSuffix(msg[0], terminator)
	dataType := string(withoutSuffix[0])

	switch dataType {

	case stringType:
		string, _ := strings.CutPrefix(withoutSuffix, stringType)

		return string

	case bulkStringType:
		withoutPrefix, _ := strings.CutPrefix(withoutSuffix, bulkStringType)

		if withoutPrefix == "-1" {
			return nil
		}

		_, bulkString, _ := strings.Cut(withoutPrefix, terminator)

		float, err := strconv.ParseFloat(bulkString, 64)
		if err != nil {
			return bulkString
		}

		return float

	case IntType:
		withoutPrefix, _ := strings.CutPrefix(withoutSuffix, IntType)

		number, _ := strconv.Atoi(withoutPrefix)

		return number

	default:

		return nil
	}
}
