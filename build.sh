#!/bin/bash
set -e

echo "Building jc-project"
echo "vetting code"
go vet ./...
echo "go vet PASSED"
echo "running unit tests"
go test -race ./...
echo "go test PASSED"
echo
echo "INFO: use go run cmd/main.go to run test harness"
echo "done"
