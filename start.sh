#!/bin/bash

echo "Running tests..."
go test ./...
if [ $? -ne 0 ]; then
  echo "Tests failed. Exiting."
  exit 1
fi

echo "Starting application..."
go run main.go