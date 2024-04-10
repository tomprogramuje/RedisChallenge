package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func Deserialize(msg string) any {

	if msg == bulkNull || msg == arrayNull {
		return nil
	}

	dataType, msgWithoutSuffix := getDataType(msg)

	switch dataType {

	case errorType:

		return deserializeErrType(msgWithoutSuffix)

	case stringType:
		
		return deserializeStringType(msgWithoutSuffix)

	case bulkStringType:
		
		return deserializeBulkStringType(msgWithoutSuffix)

	case sliceType:
		
		return deserializeSliceType(msgWithoutSuffix)

	case IntType:
		
		return deserializeIntType(msgWithoutSuffix)

	default:

		return nil
	}
}

func getDataType(msg string) (dataType, msgWithoutSuffix string) {
	msgWithoutSuffix, found := strings.CutSuffix(msg, terminator)
	if !found {
		return "", ""
	}
	dataType = string(msgWithoutSuffix[0])

	return dataType, msgWithoutSuffix
}

func deserializeErrType(msg string) error {
	err, _ := strings.CutPrefix(msg, errorType)

	return errors.New(err)
}
	

func deserializeStringType(msg string) string {
	string, _ := strings.CutPrefix(msg, stringType)

	return string
}
	

func deserializeBulkStringType(msg string) any {
	withoutPrefix, _ := strings.CutPrefix(msg, bulkStringType)
	_, bulkString, _ := strings.Cut(withoutPrefix, terminator)

	float, err := strconv.ParseFloat(bulkString, 64)
	if err != nil {
		return bulkString
	}

	return float
}

func deserializeSliceType(msg string) any {
	withoutPrefix, _ := strings.CutPrefix(msg, sliceType)
		pre, data, _ := strings.Cut(withoutPrefix, terminator)
		numOfElem, _ := strconv.Atoi(pre)

		if string(data[0]) == IntType {
			
			sliceOfInt, err := deserializeIntSlices(numOfElem, data)
			if err != nil {
				log.Println(err)
			}

			return sliceOfInt
			
		} else {

			return deserializeStringSlices(numOfElem, data)
		}
}

func deserializeIntSlices(numOfElem int, data string) ([]int, error) {
	sliceOfInt := make([]int, 0)
	for i := 0; i < numOfElem; i++ {
		withoutIntPrefix, _ := strings.CutPrefix(data, IntType)
		elem, rest, _ := strings.Cut(withoutIntPrefix, terminator)
		num, err := strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}

		sliceOfInt = append(sliceOfInt, num)
		data = rest
	}

	return sliceOfInt, nil
}

func deserializeStringSlices(numOfElem int, data string) []string {
	sliceOfString := make([]string, 0)
	for i := 0; i < numOfElem; i++ {
		_, withoutLenght, _ := strings.Cut(data, terminator)
		cleanElem, rest, _ := strings.Cut(withoutLenght, terminator)
		sliceOfString = append(sliceOfString, cleanElem)
		data = rest
	}

	return sliceOfString
}


func deserializeIntType(msg string) int {
	withoutPrefix, _ := strings.CutPrefix(msg, IntType)
	number, _ := strconv.Atoi(withoutPrefix)

	return number
}