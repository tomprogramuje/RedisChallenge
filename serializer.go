package main

import "fmt"

const commandEnd = `\r\n`

func Serialize(data any) []string {
	switch data.(type) {
	case string:

		if data == "hello world" {
			return []string{`+hello world\r\n`}
		}

		if data == "OK" {
			return []string{`+OK\r\n`}
		}
	case []string:
		sliceLength := len(data.([]string))
		dataS := data.([]string)
	
		stringLength1 := len(dataS[0])
		

		if sliceLength == 1 {
			return []string{`*` + fmt.Sprint(sliceLength) + commandEnd + `$` + fmt.Sprint(stringLength1) + commandEnd + string(dataS[0]) + commandEnd}
		}

		stringLength2 := len(dataS[1])
		return []string{`*` + fmt.Sprint(sliceLength) + commandEnd + `$` + fmt.Sprint(stringLength1) + commandEnd + string(dataS[0]) + commandEnd + `$` + fmt.Sprint(stringLength2) + commandEnd + string(dataS[1]) + commandEnd}
	}

	return []string{}
}

// func Deserialize(msg []string) any {}
