package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	terminator                = "\r\n"
	stringType                = "+"
	bulkStringType            = "$"
	sliceType                 = "*"
	errorType                 = "-"
	IntType                   = ":"
	bulkNull                  = "$-1\r\n"
	arrayNull                 = "*-1\r\n"
	simpleStringByteThreshold = 20
)

func Serialize(data any) string {
	switch d := data.(type) {

	case nil:

		return "$-1\r\n"

	case error:

		return errorType + d.Error() + terminator

	case string:

		return serializeStringType(d)

	case []string:

		return serializeStringSlices(d)

	case int:

		return IntType + fmt.Sprint(d) + terminator

	case []int:

		return serializeIntSlices(d)

	case float64:

		return serializeFloatType(d)

	case []float64:

		return serializeFloatSlices(d)

	default:

		return "invalid data"
	}
}

func serializeStringType(d string) string {
	if d == "" {
		return "$0\r\n\r\n"
	}

	if strings.Contains(d, "\r\n") || len([]byte(d)) > simpleStringByteThreshold {

		return bulkStringType + fmt.Sprint(len([]byte(d))) + terminator + d + terminator
	}

	return stringType + d + terminator
}

func serializeStringSlices(d []string) string {
	if len(d) == 0 {
		return "*-1\r\n"
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

	return sliceType + fmt.Sprint(len(d)) + terminator + msg
}

func serializeIntSlices(d []int) string {
	if len(d) == 0 {
		return "*-1\r\n"
	}

	var msgBuilder strings.Builder

	for _, num := range d {
		msgBuilder.WriteString(IntType)
		msgBuilder.WriteString(fmt.Sprint(num))
		msgBuilder.WriteString(terminator)
	}
	msg := msgBuilder.String()

	return sliceType + fmt.Sprint(len(d)) + terminator + msg
}

func serializeFloatType(d float64) string {
	dToString := strconv.FormatFloat(d, 'f', 2, 64)

	return bulkStringType + fmt.Sprint(len(dToString)) + terminator + dToString + terminator
}

func serializeFloatSlices(d []float64) string {
	if len(d) == 0 {
		return "*-1\r\n"
	}

	convSlice := []string{}
	for _, num := range d {
		dToString := strconv.FormatFloat(num, 'f', 2, 64)
		convSlice = append(convSlice, dToString)
	}

	var msgBuilder strings.Builder

	for _, flNum := range convSlice {
		textLength := fmt.Sprint(len(flNum))
		msgBuilder.WriteString(bulkStringType)
		msgBuilder.WriteString(textLength)
		msgBuilder.WriteString(terminator)
		msgBuilder.WriteString(flNum)
		msgBuilder.WriteString(terminator)
	}
	msg := msgBuilder.String()

	return sliceType + fmt.Sprint(len(convSlice)) + terminator + msg
}