package main

import (
	"fmt"
	"net/http"
	"time"
)

var x = -3

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Now serving localhost:80")
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a web page! %s %d", time.Now(), calculate(x))
	x++
}

func calculate(x int) int {
	result := x + 3
	return result
}
