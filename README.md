# Go HTTP Server

This is a simple HTTP server written in Go.

## Usage

After running the server using `go run main.go`, you can access the server at `http://localhost:8080/`.

You can also use `curl` to send a request to the server:

- Get the root path:

  ```bash
  curl http://localhost:8080/
  ```

- POST a user:

  ```bash
  curl -X POST -d '{"name":"Ali"}' localhost:8080/users
  ```
