package main

import (
	"net"
	"testing"
	"time"
)

func init() {
	storage := NewFredis()
	tcp := NewServer(":5678", storage)
	go func() {
		tcp.Run()
	}()
}

func TestServer(t *testing.T) {
	cases := []struct {
		Description string
		Command     []string
		Want        string
	}{
		{"returns PONG after sending PING command", []string{"PING"}, "PONG"},
		{"returns message back after sending ECHO command", []string{"ECHO", "Hello World!"}, "Hello World!"},
		{"returns OK after SET command", []string{"SET", "Name", "John"}, "OK"},
		{"returns 'John' after GET command", []string{"GET", "Name"}, "John"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":5678")
			if err != nil {
				t.Error("could not connect to server: ", err)
			}
			defer conn.Close()

			cmd := Serialize(test.Command)

			if _, err := conn.Write([]byte(cmd)); err != nil {
				t.Error("could not write to TCP server")
			}
			time.Sleep(time.Millisecond * 10)

			response, err := readMessage(conn)
			if err != nil {
				t.Error(err)
			}

			got := Deserialize(response)
			want := test.Want

			if got != test.Want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}
