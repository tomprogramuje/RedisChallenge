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
	IntType     = `:`
	bulkNull       = `$-1\r\n`
	arrayNull      = `*-1\r\n`
)

func Serialize(data any) [1]string {
	switch d := data.(type) {
	case nil:
		return [1]string{`$-1\r\n`}
	case string:

		return [1]string{stringType + d + terminator}

	case []string:

		if len(d) == 0 {
			return [1]string{`*-1\r\n`}
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

		return [1]string{sliceType + fmt.Sprint(len(d)) + terminator + msg}

	case int:
		
		return [1]string{IntType + fmt.Sprint(d) + terminator}

	default:

		return [1]string{}
	}
}

// func Deserialize(msg []string) any {}
