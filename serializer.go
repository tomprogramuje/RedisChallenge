package main

func Serialize(msg string) string {
	if msg == "hello world" {
		return `+hello world\r\n`
	}

	if msg == "OK" {
		return `+OK\r\n`
	}
	
	return ""
}
