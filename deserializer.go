package main

import (
	"errors"
	"strconv"
	"strings"
)

func Deserialize(msg [1]string) any {

	if msg[0] == bulkNull || msg[0] == arrayNull {
		return nil
	}

	withoutSuffix, found := strings.CutSuffix(msg[0], terminator)
	if !found {
		return nil
	}

	dataType := string(withoutSuffix[0])

	switch dataType {

	case errorType:
		err, _ := strings.CutPrefix(withoutSuffix, errorType)

		return errors.New(err)

	case stringType:
		string, _ := strings.CutPrefix(withoutSuffix, stringType)

		return string

	case bulkStringType:
		withoutPrefix, _ := strings.CutPrefix(withoutSuffix, bulkStringType)
		_, bulkString, _ := strings.Cut(withoutPrefix, terminator)

		float, err := strconv.ParseFloat(bulkString, 64)
		if err != nil {
			return bulkString
		}

		return float

	case sliceType:
		withoutPrefix, _ := strings.CutPrefix(withoutSuffix, sliceType)
		pre, data, _ := strings.Cut(withoutPrefix, terminator)
		numOfElem, _ := strconv.Atoi(pre)

		if string(data[0]) == IntType {
			sliceOfInt := make([]int, 0)
			for i := 0; i < numOfElem; i++ {
				withoutIntPrefix, _ := strings.CutPrefix(data, IntType)
				elem, rest, _ := strings.Cut(withoutIntPrefix, terminator)
				num, err := strconv.Atoi(elem)
				if err != nil {
					return err
				}

				sliceOfInt = append(sliceOfInt, num)
				data = rest
			}

			return sliceOfInt
		} else {
			sliceOfStrings := make([]string, 0)
			for i := 0; i < numOfElem; i++ {
				_, withoutLenght, _ := strings.Cut(data, terminator)
				cleanElem, rest, _ := strings.Cut(withoutLenght, terminator)
				sliceOfStrings = append(sliceOfStrings, cleanElem)
				data = rest
			}

			return sliceOfStrings
		}

	case IntType:
		withoutPrefix, _ := strings.CutPrefix(withoutSuffix, IntType)

		number, _ := strconv.Atoi(withoutPrefix)

		return number

	default:

		return nil
	}
}
