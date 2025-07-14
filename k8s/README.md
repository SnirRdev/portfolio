# 🚀 Simple Kubernetes Deployment

This directory contains a simple Kubernetes setup for the portfolio application.

## 📁 Files

```
k8s/
├── helm/portfolio-chart/     # Helm chart
├── deploy.sh                 # Simple deployment script
├── deploy-argocd.sh         # ArgoCD deployment script
├── argo-app.yaml            # ArgoCD application
└── README.md                # This file
```

## 🎯 Quick Deploy

### Option 1: Helm (Simple)
```bash
cd k8s
./deploy.sh
```

### Option 2: ArgoCD (GitOps)
```bash
cd k8s
./deploy-argocd.sh
```

## 🌐 Access the Application

```bash
kubectl port-forward svc/portfolio-portfolio-chart 8080:8080 -n portfolio
# Open http://localhost:8080
```

## 🔍 Check Status

```bash
# Check pods
kubectl get pods -n portfolio

# Check services
kubectl get svc -n portfolio
```

## 🧹 Cleanup

```bash
# Remove Helm deployment
helm uninstall portfolio -n portfolio
kubectl delete namespace portfolio

# Remove ArgoCD application
kubectl delete -f argo-app.yaml
```

## 📋 Prerequisites

- Kubernetes cluster (minikube, Docker Desktop, etc.)
- kubectl
- helm (for Option 1) 