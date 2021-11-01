# Testable Services in Go

A simple go server 

## build tasks

- `go mod tidy`       - Removes any unnecessary dependencies from `go.mod`
- `go mod verify`     - Ensures that `go.sum` is fully up-to-date based on the dependencies in `go.mod`
- `go generate ./...` - Generates mock functions from interfaces
- `go test ./...`     - Unit test application
- `go build`          - Builds an executable of the server
