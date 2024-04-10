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
			conn.Close()
		}(conn)
	}
}
