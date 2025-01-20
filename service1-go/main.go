package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World from Service 1 (Golang)")
	})

	fmt.Println("Service 1 (Golang) running on port 9001")
	http.ListenAndServe(":9001", nil)
}
