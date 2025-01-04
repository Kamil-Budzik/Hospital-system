.PHONY: run-all build-all test-all docker-up docker-down

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest

# Run all services with hot reload
run-all:
	./scripts/run-all.sh

# Build all services
build-all:
	@for service in services/*; do \
		if [ -d $$service ]; then \
			echo "Building $$service..."; \
			(cd $$service && go build -o bin/app); \
		fi \
	done

# Run tests for all services
test-all:
	@for service in services/*; do \
		if [ -d $$service ]; then \
			echo "Testing $$service..."; \
			(cd $$service && go test ./...); \
		fi \
	done

# Start all services with Docker Compose
docker-up:
	docker-compose up --build

# Stop all Docker services
docker-down:
	docker-compose down

# Run a specific service with hot reload
run-service:
	./scripts/watch.sh services/$(service) $(port)

# Create a new service
new-service:
	@read -p "Enter service name: " name; \
	mkdir -p services/$$name-service; \
	cp -r templates/service/* services/$$name-service/; \
	cd services/$$name-service && go mod init hospital-system/services/$$name-service
