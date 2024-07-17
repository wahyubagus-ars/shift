#!/bin/bash

echo "Running build.sh"

if [ "$ENV" = "local" ]; then
  air
elif [ "$ENV" = "master" ]; then
  ./main
else
  echo "Unknown environment: $ENV"
  exit 1
fi
