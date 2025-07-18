# 🌐 Portfolio Web App

A simple portfolio web application built with **Go**, styled with **HTML/CSS**, containerized using **Docker**, and deployed to **Kubernetes** using **Helm**.

---

## 📁 Project Structure

portfolio/
├── app/
│ ├── main.go # Entry point of the Go web app
│ ├── Dockerfile # Docker image definition
│ ├── data.json # JSON file containing portfolio data
│ ├── static/ # Static files (CSS, images)
│ └── templates/ # HTML templates
├── k8s/
│ ├── helm/portfolio-chart/ # Helm chart for Kubernetes deployment
│ ├── deploy.sh # Simple deployment script
│ └── README.md # Kubernetes deployment guide
├── .github/workflows/
│ └── ci-cd.yaml # GitHub Actions CI/CD pipeline
└── README.md # Project documentation

---

## 🚀 Features

- 🖥️ Web interface to view and submit portfolios
- 🐳 Containerized with Docker
- ☸️ Deployable to Kubernetes using Helm
- 🔁 CI/CD pipeline with GitHub Actions

---

## 🧪 Run Locally

### ✅ Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/)

---

### 💻 Running the App Locally (Without Docker)

```bash
cd app
go run main.go
# Then open http://localhost:8080
```

### 🐳 Running with Docker

```bash
# Build the Docker image
docker build -t portfolio-app ./app

# Run the container
docker run -p 8080:8080 portfolio-app
```

---

## ☸️ Kubernetes Deployment

### Quick Deploy

```bash
cd k8s
./deploy.sh
```

### Access the Application

```bash
kubectl port-forward svc/portfolio-portfolio-chart 8080:8080 -n portfolio
# Open http://localhost:8080
```

For more details, see [k8s/README.md](k8s/README.md).

---

## 🔄 CI/CD Pipeline

The GitHub Actions pipeline automatically:
- Runs tests
- Builds Docker image
- Pushes to Docker Hub

---

## 🛠️ Development

### Building the Application

```bash
cd app
go build -o portfolio main.go
```

### Running Tests

```bash
cd app
go test -v
```

### Building Docker Image

```bash
docker build -t snirrdev/portfolio:latest ./app
```

### Pushing to Registry

```bash
docker push snirrdev/portfolio:latest
```
