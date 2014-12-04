package main

import (
	"github.com/benschw/satisapi-go/satisapi"
	"log"
)

func main() {
	s := &satisapi.Server{}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
