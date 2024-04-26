// Launch microservice server- main.go
package main

import (
	"log"

	"github.com/anakole1/github-action/microservice"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
