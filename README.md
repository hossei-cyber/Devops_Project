# Go Webshop API - DevOps Assignment

A simple REST API webshop built in Go for the DevOps lecture. This project demonstrates modular Go architecture, JWT authentication, and basic CRUD operations.

## Features

- Product catalog with self-care items
- JWT-based authentication 
- RESTful API endpoints
- Modular Go project structure

## Project Structure

```
├── cmd/                    # Main application
│   └── main.go            # Entry point
├── internal/              # Private packages
│   ├── handlers/          # HTTP handlers
│   └── models/            # Data models
├── pkg/                   # Public packages
│   ├── auth/              # JWT utilities
│   └── utils/             # Helper functions
├── go.mod                 # Go modules
└── README.md              # Documentation
```

## API Endpoints

### Authentication
- `POST /auth/login` - User login (username: user, password: pass)
- `POST /auth/logout` - User logout

### Products
- `GET /products` - List all products
- `GET /products/{id}` - Get product details

### Orders
- `POST /checkout/placeorder` - Place order (requires authentication)

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation

1. **Clone repository**
   ```bash
   git clone https://github.com/hossei-cyber/Devops_Project.git
   cd Devops_Project
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run cmd/main.go
   ```

4. **Test the API**
   ```bash
   curl http://localhost:8080/products
   ```

### Build Commands

```bash
go build cmd/main.go    # Build executable
go fmt ./...           # Format code
go mod tidy           # Clean dependencies
```

## Testing Authentication

1. **Login to get token**
   ```bash
   curl -X POST -d "username=user&password=pass" http://localhost:8080/auth/login
   ```

2. **Use token for orders**
   ```bash
   curl -X POST -H "Authorization: Bearer YOUR_TOKEN_HERE" http://localhost:8080/checkout/placeorder
   ```