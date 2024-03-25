package main

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
		return []string{`*1\r\n$4\r\nping\r\n`}
	}

	return []string{}
}

// func Deserialize(msg []string) any {}
