package main

import (
	"fmt"
	"net/http"
)

var num_requests int = 0

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle_root)

	fmt.Println("Server is listening on port :8080")
	http.ListenAndServe(":8080", mux)
}

func handle_root(
	response_writer http.ResponseWriter, 
	request *http.Request,
	) {
		num_requests += 1
		fmt.Fprintf(response_writer, "Hello, World!\nNumber of requests: %d\n", num_requests)
}
