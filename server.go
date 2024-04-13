package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Fredis map[string]any

func NewFredis() Fredis {
	return Fredis{}
}

type TCPServer struct {
	addr    string
	storage Fredis
	server  net.Listener
}

type Server interface {
	Run() error
	Close() error
}

func NewServer(addr string, storage Fredis) *TCPServer {
	server := new(TCPServer)

	server.addr = addr
	server.storage = storage

	return server
}

func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}

	defer t.server.Close()

	for {
		conn, err := t.server.Accept()
		if err != nil {
			log.Println("could not accept connection:", err)
			break
		}

		go t.handleConnection(conn)
	}

	return nil
}

func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader, err := readMessage(conn)
	if err != nil {
		log.Println(err)
	}

	data := Deserialize(reader).([]string)

	switch data[0] {

	case "PING":
		msg := Serialize("PONG")
		conn.Write([]byte(msg))

	case "ECHO":
		msg := Serialize(data[1])
		conn.Write([]byte(msg))

	case "SET":
		t.storage[data[1]] = data[2]
		msg := Serialize("OK")
		conn.Write([]byte(msg))

	case "GET":
		msg := Serialize(t.storage["Name"])
		conn.Write([]byte(msg))
	}
}

func readMessage(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	var message strings.Builder

	err := conn.SetReadDeadline(time.Now().Add(time.Millisecond * 50))
	if err != nil {
		return "", err
	}

	for scanner.Scan() {
		line := scanner.Text()
		message.WriteString(line)
		message.WriteString("\r\n")
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
		return "", err
	}

	return message.String(), nil
}
