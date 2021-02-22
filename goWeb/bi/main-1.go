package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "wello world", r.URL.Path)
}

func main() {
	fmt.Println("==============================================")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
