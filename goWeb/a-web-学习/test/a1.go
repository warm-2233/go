package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server start....")
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("end..")

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cnm")
	fmt.Println(r.URL)
	fmt.Println(r.Header.Get("User-Agent"))
	fmt.Println(r.Method)
	fmt.Println(r.Host)
}
