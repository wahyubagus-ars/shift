#!/bin/bash

# Install wire tool (assuming you have go installed) and didn't have install google wire
go install github.com/google/wire/cmd/wire@latest

# Generate dependencies using wire before build image docker
wire ./cmd/app/provider

# Run docker-compose in detached mode (background)
docker-compose up -d
