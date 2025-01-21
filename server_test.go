package main

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

func startServer() *exec.Cmd {
	cmd := exec.Command("go", "run", os.Getenv("PWD"), "/main.go")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second) // Give the server time to start
	return cmd
}

func stopServer(cmd *exec.Cmd) {
	err := cmd.Process.Kill()
	if err != nil {
		panic(err)
	}
}

func TestRoot(t *testing.T) {
	pid := startServer()
	defer stopServer(pid)
	response, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", response.Status)
	}
}

func TestCreateUser(t *testing.T) {
	pid := startServer()
	defer stopServer(pid)
	response, err := http.Post(
		"http://localhost:8080/users", 
		"application/json", 
		bytes.NewReader([]byte(`{"name":"Ali"}`)),
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status Created; got %v", response.Status)
	}
}
