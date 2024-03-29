package main

import (
	"errors"
	"net"
)

func establishConnection() (err error) {
	server, err := net.Listen("tcp", ":5678")
	if err != nil {
		return
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			return errors.New("could not accept connection")
		}
		if conn == nil {
			return errors.New("could not create connection")
		}
		conn.Close()
	}
}
