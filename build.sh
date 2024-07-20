#!/bin/sh

echo "Running build.sh"

if [ "$ENV" = "local" ]; then
  go install github.com/air-verse/air@latest
  air
elif [ "$ENV" = "master" ]; then
  /app/binary
else
  echo "Unknown environment: $ENV"
  exit 1
fi
