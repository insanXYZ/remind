package main

import (
	"log"
	"remind-daemon/server"
)

func main() {
	s := server.NewServer()
	err := s.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

}
