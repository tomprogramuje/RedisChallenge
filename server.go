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

func establishConnection() (err error) {
	fredis := make(map[string]any)
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
			reader, err := readMessage(c)
			if err != nil {
				log.Println(err)
			}

			data := Deserialize(reader).([]string)

			switch data[0] {
				
			case "PING":
				msg := Serialize("PONG")
				c.Write([]byte(msg))

			case "ECHO":
				msg := Serialize(data[1])
					c.Write([]byte(msg))

			case "SET":
				fredis[data[1]] = data[2]
				msg := Serialize("OK")
				c.Write([]byte(msg))
			}
		}(conn)
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
