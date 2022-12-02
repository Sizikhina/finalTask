package main

import (
	"log"
	"net/http"
	"zlata/service"
)

func main() {
	err := http.ListenAndServe("localhost:8080", service.NewRouter())
	if err != nil {
		log.Fatalln(err)
	}
}
