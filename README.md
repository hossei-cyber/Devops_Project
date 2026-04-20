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

- Go 1.23 or higher
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

- Individual local runs default to `localhost:8080`
- In Kubernetes, the services use ports `8081`, `8082`, and `8083`

### Test the API

```bash
curl http://localhost:8080/products
```

## Build Commands

### Build all services

```bash
make build-all
```

### Build a specific service

```bash
make build service=auth-service
make build service=product-service
make build service=checkout-service
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

Run the services in Docker containers.

### Build Docker images

```bash
docker build --build-arg SERVICE=auth-service -t hosseicyber/webshop-auth:latest .
docker build --build-arg SERVICE=product-service -t hosseicyber/webshop-product:latest .
docker build --build-arg SERVICE=checkout-service -t hosseicyber/webshop-checkout:latest .
```

### Run Docker containers

```bash
docker run -p 8081:8080 hosseicyber/webshop-auth:latest
docker run -p 8082:8080 hosseicyber/webshop-product:latest
docker run -p 8083:8080 hosseicyber/webshop-checkout:latest
```

### Docker Hub

1. Login to Docker Hub

   ```bash
   docker login
   ```

2. Tag the images

   ```bash
   docker tag hosseicyber/webshop-auth:latest hosseicyber/webshop-auth:<version>
   docker tag hosseicyber/webshop-product:latest hosseicyber/webshop-product:<version>
   docker tag hosseicyber/webshop-checkout:latest hosseicyber/webshop-checkout:<version>
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

## Kubernetes Deployment

The application is containerized and ready for Kubernetes deployment with separate environments.

### Prerequisites

- Kubernetes cluster (minikube for local development)
- kubectl configured
- Docker images built and pushed

### Quick Start with Minikube

1. Start minikube cluster

   ```bash
   minikube start
   ```

2. Deploy to production environment

   ```bash
   kubectl apply -k k8s/overlays/production
   ```

3. Check deployment status

   ```bash
   kubectl get all -n webshop-prod
   ```

4. Access services via port forwarding

   ```bash
   kubectl port-forward -n webshop-prod svc/prod-auth-service 8081:8081
   kubectl port-forward -n webshop-prod svc/prod-product-service 8082:8082
   kubectl port-forward -n webshop-prod svc/prod-checkout-service 8083:8083
   ```

### Environment Configuration

#### Production Environment
- **Namespace**: `webshop-prod`
- **Replicas**: auth `1`, product `2`, checkout `3`
- **Resources**: Higher (128Mi RAM, 500m CPU)
- **Access**: ClusterIP services inside the cluster

### Service Endpoints in Kubernetes

- **Auth Service**: `prod-auth-service:8081`
- **Product Service**: `prod-product-service:8082`
- **Checkout Service**: `prod-checkout-service:8083`

### Deployment Commands

```bash
# Deploy production environment
kubectl apply -k k8s/overlays/production

# Delete production environment
kubectl delete -k k8s/overlays/production
```

### Testing Kubernetes Services

```bash
# Port forward for direct access
kubectl port-forward -n webshop-prod svc/prod-product-service 8082:8082

# Test the API
curl http://localhost:8082/products
```

## ArgoCD GitOps Deployment

ArgoCD provides GitOps-based continuous deployment for the webshop application. It automatically syncs your Kubernetes manifests from the Git repository to your cluster.

### Prerequisites

- Kubernetes cluster (minikube for local development)
- kubectl configured and connected to your cluster
- Git repository with Kubernetes manifests (this repo)

### Step 1: Install ArgoCD

1. Create ArgoCD namespace and install ArgoCD:

   ```bash
   kubectl create namespace argocd
   kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
   ```

2. Wait for ArgoCD components to be ready:

   ```bash
   kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
   ```

### Step 2: Access ArgoCD UI

1. **Option A: Port Forward (Recommended for local development)**

   ```bash
   kubectl port-forward svc/argocd-server -n argocd 8080:443
   ```

   Access ArgoCD at: `https://localhost:8080` (accept the self-signed certificate)

2. **Option B: Expose via NodePort (Alternative)**

   ```bash
   kubectl patch svc argocd-server -n argocd -p '{"spec":{"type":"NodePort"}}'
   minikube service argocd-server -n argocd --url
   ```

### Step 3: Get ArgoCD Admin Password

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

**Login credentials:**
- Username: `admin`
- Password: (output from the command above)

### Step 4: Build and Push Docker Images

Before deploying with ArgoCD, ensure your Docker images are built and available:

1. **Build all service images:**

   ```bash
   docker build --build-arg SERVICE=auth-service -t hosseicyber/webshop-auth:latest .
   docker build --build-arg SERVICE=product-service -t hosseicyber/webshop-product:latest .
   docker build --build-arg SERVICE=checkout-service -t hosseicyber/webshop-checkout:latest .
   ```

2. **Login to Docker Hub:**

   ```bash
   docker login
   ```

3. **Push images to Docker Hub:**

   ```bash
   docker push hosseicyber/webshop-auth:latest
   docker push hosseicyber/webshop-product:latest
   docker push hosseicyber/webshop-checkout:latest
   ```

4. **For local Minikube development:**

   ```bash
   minikube image load hosseicyber/webshop-auth:latest
   minikube image load hosseicyber/webshop-product:latest
   minikube image load hosseicyber/webshop-checkout:latest
   ```

### Step 5: Deploy Applications with ArgoCD

1. Apply the ArgoCD application manifests:

   ```bash
   kubectl apply -f argocd/webshop-application.yaml
   ```

2. Verify applications are created:

   ```bash
   kubectl get applications -n argocd
   ```

### Step 6: Sync and Monitor Applications

1. **Check application status:**

   ```bash
   kubectl get applications -n argocd
   ```

2. **View application details:**

   ```bash
   kubectl describe application webshop-prod -n argocd
   ```

### Application Configuration

The ArgoCD applications are configured as follows:

#### Production Environment
- **Application Name**: `webshop-prod`
- **Source**: `k8s/overlays/production`
- **Destination**: `webshop-prod` namespace
- **Sync Policy**: Automated with prune and self-heal

### GitOps Workflow

1. **Make changes** to your Kubernetes manifests in the `k8s/` directory
2. **Commit and push** changes to Git repository
3. **ArgoCD automatically detects** changes and syncs them to the cluster
4. **Monitor deployments** via ArgoCD UI
