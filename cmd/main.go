package main

import (
	"fmt"
	"log"
	"net/http"

	ptth "github.com/mat-golang/server/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	fmt.Println("Listening on port :8080")
	return http.ListenAndServe(":8080", ptth.NewService())
}
