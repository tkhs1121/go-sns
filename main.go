package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, requset *http.Request) {
	fmt.Fprintln(writer, "Hello world!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
