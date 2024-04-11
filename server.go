package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
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
			reader, err := readMessage(c)
			if err != nil {
				log.Println(err)
			}
			if strings.HasPrefix(reader, "*") {
				data := Deserialize(reader).([]string)
				if data[0] == "ECHO" {
					msg := Serialize(data[1])
					c.Write([]byte(msg))
				}
			} else {
				data := Deserialize(reader).(string)
				fmt.Println(data)
				if data == "PING" {
					msg := Serialize("PONG")
					c.Write([]byte(msg))
				}
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
		if line == "\r\n" {
			break
		}
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
		return "", err
	}

	return message.String(), nil
}
