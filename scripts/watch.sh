# scripts/watch.sh
#!/bin/bash

SERVICE_DIR=$1
PORT=$2

cd $SERVICE_DIR

# Install air if not present
which air || go install github.com/cosmtrek/air@latest

# Run with air
air -c .air.toml
