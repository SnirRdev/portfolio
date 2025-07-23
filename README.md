# 🌐 Portfolio Web Application

A modern portfolio web application built with **Python** (FastAPI), containerized with **Docker**, and deployed to **Kubernetes** using **Helm** and **ArgoCD** for GitOps workflows.

---

## 📁 Project Structure

```
portfolio/
├── app/
│   ├── main.py            # FastAPI application entry point
│   ├── db.py              # Database connection and models
│   └── requirements.txt   # Python dependencies
├── k8s/
│   ├── argocd/            # ArgoCD application configurations
│   │   └── application.yaml
│   └── helm/
│       └── portfolio-chart/  # Helm chart for Kubernetes deployment
│           ├── charts/       # Chart dependencies
│           ├── templates/    # Kubernetes manifests
│           ├── Chart.yaml    # Chart metadata
│           └── values.yaml   # Default configuration values
└── Dockerfile              # Container image definition
```

---

## 🚀 Features

- 🐍 **FastAPI Backend**: Modern, fast web framework for building APIs with Python
- 🐳 **Containerized**: Easy deployment with Docker
- ☸️ **Kubernetes Native**: Deployable to any Kubernetes cluster
- 🔄 **GitOps Workflow**: Managed with ArgoCD for continuous delivery
- 📊 **PostgreSQL Database**: Persistent data storage
- 🔒 **Production-Ready**: Includes health checks, metrics, and logging

---

## 🛠️ Prerequisites

- [Docker](https://www.docker.com/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/)
- [minikube](https://minikube.sigs.k8s.io/docs/start/) or access to a Kubernetes cluster
- [ArgoCD](https://argoproj.github.io/argo-cd/) (for GitOps deployment)

---

## 🚀 Quick Start

### Local Development

1. **Build and run with Docker:**
   ```bash
   docker build -t portfolio-app .
   docker run -p 8000:8000 portfolio-app
   ```

2. **Access the API:**
   ```
   http://localhost:8000
   ```

### Running Locally (Development)

1. **Set up a Python virtual environment:**
   ```bash
   python -m venv venv
   source venv/bin/activate  # On Windows: .\venv\Scripts\activate
   ```

2. **Install dependencies:**
   ```bash
   pip install -r app/requirements.txt
   ```

3. **Run the FastAPI application:**
   ```bash
   cd app
   uvicorn main:app --reload
   ```

---

## ☸️ Kubernetes Deployment

### Prerequisites
- A running Kubernetes cluster
- `kubectl` configured to communicate with your cluster
- Helm installed

### Install with Helm

1. **Add the PostgreSQL repository (if not already added):**
   ```bash
   helm repo add bitnami https://charts.bitnami.com/bitnami
   ```

2. **Install the chart:**
   ```bash
   helm install portfolio ./k8s/helm/portfolio-chart
   ```

### Access the Application

1. **Port-forward the service:**
   ```bash
   kubectl port-forward svc/portfolio 8000:8000
   ```

2. **Access the application:**
   ```
   http://localhost:8000
   ```

### ArgoCD Deployment (GitOps)

If you're using ArgoCD, you can deploy the application using the provided Application manifest:

```bash
kubectl apply -f k8s/argocd/application.yaml
```

---

## 🛠️ Development

### Environment Variables

The application can be configured using the following environment variables:

- `DATABASE_URL`: PostgreSQL connection string (default: `postgresql://postgres:postgres@postgres:5432/portfolio`)
- `ENVIRONMENT`: Runtime environment (default: `development`)
- `PORT`: Port to run the application on (default: `8000`)

### Running Tests

```bash
# Install test dependencies
pip install -r requirements-dev.txt

# Run tests
pytest
```

### Linting and Formatting

```bash
# Run black formatter
black .

# Run flake8 linter
flake8
```

### Building for Production

```bash
# Build Docker image
docker build -t your-registry/portfolio:latest .

# Push to container registry
docker push your-registry/portfolio:latest
```

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
