# Webshop

A simple REST API webshop built in Go for the DevOps lecture. It demonstrates a modular Go architecture, JWT authentication, and basic CRUD operations.

## Features

- Product catalog with self-care items
- JWT-based authentication
- RESTful API endpoints
- Modular Go project structure

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

3. Run the application
   ```bash
   go run cmd/main.go
   ```

4. Test the API
   ```bash
   curl http://localhost:8080/products
   ```

### Build Commands

```bash
go build cmd/main.go  # Build executable
go fmt ./...          # Format code
go mod tidy           # Clean dependencies
```

## Authentication Demo

1. Login to get a token
   ```bash
   curl -X POST -d "username=user&password=pass" http://localhost:8080/auth/login
   ```

2. Use the token for orders
   ```bash
   curl -X POST -H "Authorization: Bearer YOUR_TOKEN_HERE" http://localhost:8080/checkout/placeorder
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

Run the application in a Docker container:

1. Build the Docker image
   ```bash
   docker build -t webshop .
   ```

2. Run the Docker container
   ```bash
   docker run -p 8080:8080 webshop
   ```

3. Login to Docker Hub
   ```bash
   docker login
   ```

4. Push the image to Docker Hub
   ```bash
   docker tag webshop:latest hosseicyber/webshop:<version>
   docker push hosseicyber/webshop:<version>
   ```

5. Pull the image from Docker Hub
   ```bash
   docker pull hosseicyber/webshop:<version>
   ```

Note: replace `<version>` with the actual version tag (currently `1.0.0`).






