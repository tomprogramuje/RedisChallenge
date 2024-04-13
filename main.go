package main

import "log"

func main() {
	storage := NewFredis()
	server := NewServer(":5678", storage)
	log.Println("listening on 5678...")
	server.Run()
}