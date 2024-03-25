package main

import (
	"fmt"
	"strings"
)

const (
	terminator     = `\r\n`
	stringType     = `+`
	bulkStringType = `$`
	sliceType      = `*`
	errorType      = `-`
	intType        = `:`
	bulkNull       = `$-1\r\n`
	arrayNull      = `*-1\r\n`
)

func Serialize(data any) []string {
	switch d := data.(type) {
	case nil:
		return []string{`$-1\r\n`}
	case string:

		return []string{stringType + d + terminator}

	case []string:

		if len(d) == 0 {
			return []string{`*-1\r\n`}
		}

		var msgBuilder strings.Builder

		for _, text := range d {
			textLength := fmt.Sprint(len(text))
			msgBuilder.WriteString(bulkStringType)
			msgBuilder.WriteString(textLength)
			msgBuilder.WriteString(terminator)
			msgBuilder.WriteString(text)
			msgBuilder.WriteString(terminator)
		}
		msg := msgBuilder.String()

		return []string{sliceType + fmt.Sprint(len(d)) + terminator + msg}

	default:

		return []string{}
	}
}

// func Deserialize(msg []string) any {}
