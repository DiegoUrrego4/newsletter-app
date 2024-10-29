package main

import "github.com/DiegoUrrego4/newsletter-app/cmd/server"

func main() {
	s := server.NewServer()
	s.Run()
}
