package main

import (
	"bufio"
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

		response, _ := bufio.NewReader(conn).ReadString('\n')
		got := Deserialize(response)
		want := "PONG"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
