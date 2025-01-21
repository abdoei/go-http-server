# Go HTTP Server

This is a simple HTTP server written in Go.

## Usage

After running the server using `go run main.go`, from the browser, you can access the root at `http://localhost:8080/` and the users at `http://localhost:8080/users/{ID-number}` after replacing `{ID-number}` with the user ID that you POSTed.

You can also use `curl` to send a request to the server:

- Get the root path:

  ```bash
  curl http://localhost:8080/
  ```

- POST a user:

  ```bash
  curl -X POST -d '{"name":"Ali"}' localhost:8080/users
  ```
