package main

import (
	"net"
	"testing"
)

func init() {
	go func() {
		establishConnection()
	}()
}

func TestServer(t *testing.T) {
	_, err := net.Dial("tcp", ":5678")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	
}
