package client

import (
	"log"
	"net/http"

	"github.com/oanaOM/weatherapp/server"
)

func main() {
	http.Handle("/", &server.MyHandler{Greeting: "Hello"})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalf("Ups we got this error %v", err)
	}
}
