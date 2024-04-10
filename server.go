package main

import (
	"log"
	"net"
)

func establishConnection() (err error) {
	l, err := net.Listen("tcp", ":5678")
	if err != nil {
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			msg := Serialize("PONG")
			c.Write([]byte(msg[0]))
		}(conn)
	}
}
