package main

// These tests need to be ran sequentially

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
	// Arrange
	pid := startServer()
	defer stopServer(pid)

	// Act
	response, err := http.Get("http://localhost:8080/")

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", response.Status)
	}
}

func TestCreateUser(t *testing.T) {
	// Arrange
	pid := startServer()
	defer stopServer(pid)

	// Act
	response, err := http.Post(
		"http://localhost:8080/users",
		"application/json",
		bytes.NewReader([]byte(`{"name":"Ali"}`)),
	)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status Created; got %v", response.Status)
	}
}

func TestGetUser(t *testing.T) {
	// Arrange
	pid := startServer()
	defer stopServer(pid)

	//Act
	// Create a user
	response, err := http.Post(
		"http://localhost:8080/users",
		"application/json",
		bytes.NewReader([]byte(`{"name":"Ali"}`)),
	)
	// Assert for create user
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status Created; got %v", response.Status)
	}

	// Get the user
	response, err = http.Get("http://localhost:8080/users/0/")

	// Assert for get user
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", response.Status)
	}
}

func TestDeleteUser(t *testing.T) {
	// Arrange
	pid := startServer()
	stopServer(pid)

	// Act
	// Create a user
	response, err := http.Post(
		"http://localhost:8080/users",
		"application/json",
		bytes.NewReader([]byte(`{"name":"Ali"}`)),
	)
	// Assert for create user
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status Created; got %v", response.Status)
	}

	// Delete the user
	request, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/0/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err = http.DefaultClient.Do(request)

	// Assert for delete user
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status No Content; got %v", response.Status)
	}
}
