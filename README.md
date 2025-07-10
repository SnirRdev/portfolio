# 🌐 Portfolio Web App

A simple portfolio web application built with **Go**, styled with **HTML/CSS**, containerized using **Docker**, and deployed to **Kubernetes** using **Helm**. CI/CD is handled via **GitHub Actions**, and optionally managed via **ArgoCD**.

---

## 📁 Project Structure

portfolio/
├── app/
│ ├── main.go # Entry point of the Go web app
│ ├── Dockerfile # Docker image definition
│ ├── data.json # JSON file containing portfolio data
│ ├── static/ # Static files (CSS, images)
│ └── templates/ # HTML templates
├── helm/
│ └── portfolio-chart/ # Helm chart for Kubernetes deployment
│ ├── templates/
│ └── values.yaml
├── k8s/
│ └── argo-app.yaml # Argo CD application definition (GitOps)
├── .github/workflows/
│ └── ci-cd.yaml # GitHub Actions CI/CD pipeline
└── README.md # Project documentation

---

## 🚀 Features

- 🖥️ Web interface to view and submit portfolios
- 🐳 Containerized with Docker
- ☸️ Deployable to Kubernetes using Helm
- 🔁 CI/CD pipeline with GitHub Actions
- 📦 GitOps ready with ArgoCD

---

## 🧪 Run Locally

### ✅ Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/)
- (Optional) [Argo CD](https://argo-cd.readthedocs.io/en/stable/)

---

### 💻 Running the App Locally (Without Docker)

```bash
cd app
go run main.go
Then open http://localhost:8080

🐳 Docker
Build the Docker image:
bash
Copy code
docker build -t portfolio-app ./app
Run the container:
bash
Copy code
docker run -p 8080:8080 portfolio-app
