# Hospital Management System

A modern, microservices-based hospital management system built with Go. This system helps hospitals manage patients, appointments, medical records, and billing through a set of interconnected services.

## Features

- 🏥 Complete hospital management solution
- 🔐 Secure authentication and authorization
- 👥 Patient and doctor management
- 📅 Appointment scheduling
- 📋 Medical records management
- 💰 Billing and invoicing
- 📨 Notification system

## System Architecture

The system is built using a microservices architecture with the following components:

- **User Service** (Port 8001): Handles user authentication and management
- **Patient Service** (Port 8002): Manages patient information
- **Appointment Service** (Port 8003): Handles scheduling and appointments
- **Medical Records Service** (Port 8004): Manages patient medical records
- **Billing Service** (Port 8005): Handles invoicing and payments
- **Notification Service** (Port 8006): Manages system notifications

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL 15+
- RabbitMQ
- Redis
- Make

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/hospital-system.git
cd hospital-system
```

2. Install development tools:
```bash
make install-tools
```

3. Set up environment variables:
```bash
# Copy example env files for each service
cp services/user-service/.env.example services/user-service/.env
cp services/patient-service/.env.example services/patient-service/.env
# ... repeat for other services
```

4. Start the services:

### Development mode (with hot reload):
```bash
make run-all
```

### Using Docker:
```bash
make docker-up
```

## Development

### Project Structure
```
hospital-system/
├── docker-compose.yml
├── Makefile
├── scripts/
│   ├── run-all.sh
│   └── watch.sh
└── services/
    ├── user-service/
    ├── patient-service/
    └── ...
```

### Available Make Commands

- `make run-all`: Start all services with hot reload
- `make build-all`: Build all services
- `make test-all`: Run tests for all services
- `make docker-up`: Start all services using Docker
- `make docker-down`: Stop all Docker services
- `make run-service service=user port=8001`: Run a specific service

### Adding a New Service

1. Create a new service:
```bash
make new-service  # Enter service name when prompted
```

2. Implement the service endpoints following the project structure
3. Add the service to `docker-compose.yml`
4. Update the `scripts/run-all.sh` file

## API Documentation

### User Service (8001)
```
POST /api/v1/users/register - Register new user
POST /api/v1/users/login - User login
GET /api/v1/users/{id} - Get user details
```

### Patient Service (8002)
```
POST /api/v1/patients - Create patient record
GET /api/v1/patients/{id} - Get patient details
PUT /api/v1/patients/{id} - Update patient information
```

### Appointment Service (8003)
```
POST /api/v1/appointments - Create appointment
GET /api/v1/appointments/{id} - Get appointment details
PUT /api/v1/appointments/{id} - Update appointment
```

For complete API documentation, see the [API Documentation](./docs/api.md) file.

## Configuration

Each service has its own configuration file in `.env`:

```env
SERVER_PORT=8001
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hospital
```

## Testing

Run tests for all services:
```bash
make test-all
```

Run tests for a specific service:
```bash
cd services/user-service
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Coding Standards

- Follow Go best practices and idioms
- Use meaningful variable and function names
- Write tests for new features
- Update documentation when needed

## Deployment

### Using Docker

```bash
docker-compose up -d
```

### Manual Deployment

1. Build the services:
```bash
make build-all
```

2. Set up the databases and message broker
3. Configure environment variables
4. Start each service

## Monitoring

The system includes:
- Health check endpoints for each service
- Prometheus metrics
- Grafana dashboards
- Distributed tracing with Jaeger

## Security

- JWT-based authentication
- Role-based access control
- Data encryption at rest
- Secure communication between services
- Regular security updates

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Gin Web Framework
- PostgreSQL
- RabbitMQ
- And all other open source projects used

## Contact

Your Name - your.email@example.com
Project Link: https://github.com/yourusername/hospital-system


Would you like me to:
1. Add more specific details to any section?
2. Create additional documentation files like API docs or deployment guides?
3. Provide examples of configuration files or environment variables?
