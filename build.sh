#!/bin/bash
set -e

echo "Building jc-project"
echo "vetting code"
go vet ./...
echo "go vet PASSED"
echo "running unit tests"
go test -race ./...
echo "go test PASSED"
echo "building..."
go build -o jc-project cmd/main.go
echo "build FINISHED"
echo "./jc-project to run test harness"
echo "INFO: use go run -race cmd/main.go to run test harness and check for race conditions"
