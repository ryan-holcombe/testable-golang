# Testable Services in Go
[![Go](https://github.com/ryan-holcombe/testable-golang/actions/workflows/go.yml/badge.svg)](https://github.com/ryan-holcombe/testable-golang/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/ryan-holcombe/testable-golang/branch/main/graph/badge.svg?token=083O6ONW1P)](https://codecov.io/gh/ryan-holcombe/testable-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryan-holcombe/testable-golang)](https://goreportcard.com/report/github.com/ryan-holcombe/testable-golang)

https://medium.com/@rholcombe30/building-go-microservices-focused-on-testability-d6164751275d

A simple go server

## build tasks

- `go mod tidy`       - Removes any unnecessary dependencies from `go.mod`
- `go mod verify`     - Ensures that `go.sum` is fully up-to-date based on the dependencies in `go.mod`
- `go generate ./...` - Generates mock functions from interfaces
- `go test ./...`     - Unit test application
- `go build`          - Builds an executable of the server
