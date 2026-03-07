# Webshop

A simple REST API webshop built in Go for the DevOps lecture. It demonstrates a modular Go architecture, JWT authentication, a microservice-based structure, and basic API operations.

## Features

- Product catalog with self-care items
- JWT-based authentication
- RESTful API endpoints
- Modular Go project structure
- Split into three microservices:
  - `auth-service`
  - `product-service`
  - `checkout-service`

## Architecture

The application is structured as three Go microservices:

- `auth-service` — handles authentication endpoints
- `product-service` — provides product catalog endpoints
- `checkout-service` — handles order placement

Shared code is organized in reusable packages, including:
- `internal/...` for service-specific handlers
- `pkg/...` for shared helpers, authentication, and models

## API Endpoints

### Authentication

- `POST /auth/login` — User login (username: `user`, password: `pass`)
- `POST /auth/logout` — User logout

### Products

- `GET /products` — List all products
- `GET /products/{id}` — Get product details

### Orders

- `POST /checkout/placeorder` — Place order (requires authentication)

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- Docker (optional, for containerized execution)

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/hossei-cyber/Devops_Project.git
   cd Devops_Project
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

## Run locally

### Start individual services

#### Auth service
```bash
go run ./auth-service/cmd/main.go
```

#### Product service
```bash
go run ./product-service/cmd/main.go
```

#### Checkout service
```bash
go run ./checkout-service/cmd/main.go
```

### Default service ports

- `auth-service` → `localhost:8081`
- `product-service` → `localhost:8082`
- `checkout-service` → `localhost:8083`

### Test the API

```bash
curl http://localhost:8082/products
```

## Build Commands

### Build all services
```bash
make build-all
```

### Build a specific service
```bash
make build SERVICE=auth-service
make build SERVICE=product-service
make build SERVICE=checkout-service
```

### Additional Go commands
```bash
go fmt ./...
go mod tidy
go test ./...
```

## Authentication Demo

1. Login to get a token
   ```bash
   curl -X POST -d "username=user&password=pass" http://localhost:8081/auth/login
   ```

2. Use the token for orders
   ```bash
   curl -X POST -H "Authorization: Bearer YOUR_TOKEN_HERE" http://localhost:8083/checkout/placeorder
   ```

## Version Control Standards

### Branching Strategy

- Use feature branches for new features and bug fixes.
- Naming conventions:
  - Features: `feature/feature-name`
  - Bug fixes: `fix/bug-description`
  - Refactor: `refactor/description`
  - Documentation: use the `documentation` branch for docs-only changes
- Merge into `main` after code review.

### Commit Messages

- Use clear, descriptive messages.
- Feature commits: `feat: add new feature`
- Bug fixes: `fix: resolve bug description`
- Documentation: `docs: update documentation for feature`
- Refactoring: `refactor: improve code structure`

## Dockerization

Run the services in Docker containers.

### Build Docker images

```bash
docker build --build-arg SERVICE=auth-service -t webshop-auth .
docker build --build-arg SERVICE=product-service -t webshop-product .
docker build --build-arg SERVICE=checkout-service -t webshop-checkout .
```

### Run Docker containers

```bash
docker run -p 8081:8081 webshop-auth
docker run -p 8082:8082 webshop-product
docker run -p 8083:8083 webshop-checkout
```

### Docker Hub

1. Login to Docker Hub
   ```bash
   docker login
   ```

2. Tag the images
   ```bash
   docker tag webshop-auth:latest hosseicyber/webshop-auth:<version>
   docker tag webshop-product:latest hosseicyber/webshop-product:<version>
   docker tag webshop-checkout:latest hosseicyber/webshop-checkout:<version>
   ```

3. Push the images
   ```bash
   docker push hosseicyber/webshop-auth:<version>
   docker push hosseicyber/webshop-product:<version>
   docker push hosseicyber/webshop-checkout:<version>
   ```

4. Pull the images
   ```bash
   docker pull hosseicyber/webshop-auth:<version>
   docker pull hosseicyber/webshop-product:<version>
   docker pull hosseicyber/webshop-checkout:<version>
   ```

Note: replace `<version>` with the actual version tag.

## Project Goal in the DevOps Lecture

This repository is developed incrementally throughout the DevOps lecture. The webshop serves as the base project for applying topics such as:
- repository structuring
- containerization
- CI/CD
- Kubernetes
- and later DevOps practices in the course