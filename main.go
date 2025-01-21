package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var request_number int = 0
var users_cache map[int]User = make(map[int]User)
var users_cache_mutex sync.RWMutex

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handle_root)
	mux.HandleFunc("POST /users", create_user)
	mux.HandleFunc("GET /users/{id}", get_user)

	fmt.Println("Server is listening on port :8080")
	http.ListenAndServe(":8080", mux)
}

func handle_root(
	response_writer http.ResponseWriter,
	request *http.Request,
) {
	request_number += 1
	fmt.Fprintf(response_writer, "Hello, World!\nrequest No.%d\n", request_number)
}

func create_user(
	response_writer http.ResponseWriter,
	request *http.Request,
) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(response_writer, "Name is required", http.StatusBadRequest)
		return
	}

	response_writer.WriteHeader(http.StatusCreated)

	
	users_cache_mutex.Lock()
	var user_id int
	user_id = request_number
	request_number++
	users_cache_mutex.Unlock()
	
	users_cache[user_id] = user
	fmt.Printf("User created: %s\t|\tkey: %d\n", user.Name, user_id)
}

func get_user(
	response_writer http.ResponseWriter,
	request *http.Request,
) {
	var user_id int
	// TODO: test case of /users/001 shall not equal to /users/1
	fmt.Sscanf(request.URL.Path, "/users/%d", &user_id)

	users_cache_mutex.RLock()
	user, ok := users_cache[user_id]
	users_cache_mutex.RUnlock()

	if !ok {
		http.Error(response_writer, fmt.Sprintf("User with id:%d not found", user_id), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(response_writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response_writer.WriteHeader(http.StatusOK)
	response_writer.Header().Set("Content-Type", "application/json")
	response_writer.Write(j)
}