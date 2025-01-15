package main

import "remind/daemon/server"

func main() {

	s := server.NewServer()
	err := s.Run()
	if err != nil {
		panic(err.Error())
	}

}
