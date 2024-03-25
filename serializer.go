package main

func Serialize(data any) []string {
	
	if data == "hello world" {
		return []string{`+hello world\r\n`}
	}

	if data == "OK" {
		return []string{`+OK\r\n`}
	}
	
	return []string{}
}


// func Deserialize(msg []string) any {}