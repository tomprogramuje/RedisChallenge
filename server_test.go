package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go func() {
		if err := establishConnection(); err != nil {
			t.Errorf("error starting server: %v", err)
		}
	}()
	time.Sleep(time.Millisecond * 100)

	t.Run("returns PONG after sending PING command", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":5678")
		if err != nil {
			t.Error("could not connect to server: ", err)
		}
		defer conn.Close()

		cmd := Serialize("PING")
		if _, err := conn.Write([]byte(cmd)); err != nil {
			t.Error("could not write to TCP server")
		}
		time.Sleep(time.Millisecond * 50)
		response, _ := readMessage(conn)
		got := Deserialize(response)
		want := "PONG"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("returns message back after sending ECHO command", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":5678")
		if err != nil {
			t.Error("could not connect to server: ", err)
		}
		defer conn.Close()

		cmd := Serialize([]string{"ECHO", "Hello World!"})
		if _, err := conn.Write([]byte(cmd)); err != nil {
			t.Error("could not write to TCP server")
		}
		time.Sleep(time.Millisecond * 50)
		response, err := readMessage(conn)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("response", response)
		got := Deserialize(response)
		fmt.Println("got", got)
		want := "Hello World!"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
