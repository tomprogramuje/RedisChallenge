package main

func Deserialize(msg [1]string) any {
	
	if msg == [1]string{`+OK\r\n`} {
		return "OK"
	}
	return nil
}