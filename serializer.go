package main

import "fmt"

const terminator = `\r\n`
const stringType = `+`
const bulkStringType = `$`
const sliceType = `*`

func Serialize(data any) []string {
	switch d := data.(type) {
	case string:

		return []string{stringType + d + terminator}

	case []string:

		msg := ""

		for _, text := range d {
			textLength := fmt.Sprint(len(text))
			msgPart := bulkStringType + textLength + terminator + text + terminator
			msg += msgPart
		}

		return []string{sliceType + fmt.Sprint(len(d)) + terminator + msg}

	default:

		return []string{}
	}
}

// func Deserialize(msg []string) any {}
