#!/bin/bash

# Run Go application in the background
go run cmd/main.go &

# Run Envoy
envoy -c api-gateway/api-gateway.yaml
