package server

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	Greeting string
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %v", mh.Greeting)))
}

// func SumTwo(a, b int) int {
// 	return a + b
// }
