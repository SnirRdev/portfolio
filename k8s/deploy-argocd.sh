#!/bin/bash

# Simple ArgoCD deployment script
set -e

echo "🚀 Deploying portfolio application via ArgoCD..."

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl is not installed. Please install kubectl first."
    exit 1
fi

# Check if we can connect to the cluster
if ! kubectl cluster-info &> /dev/null; then
    echo "❌ Cannot connect to Kubernetes cluster. Please ensure your cluster is running."
    exit 1
fi

# Check if ArgoCD is installed
if ! kubectl get pods -n argocd &> /dev/null; then
    echo "⚠️  ArgoCD not found. Installing ArgoCD..."
    kubectl create namespace argocd --dry-run=client -o yaml | kubectl apply -f -
    kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
    
    echo "⏳ Waiting for ArgoCD to be ready..."
    kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=argocd-server -n argocd --timeout=300s
    
    echo "🔐 Getting ArgoCD admin password..."
    ARGOCD_PASSWORD=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
    echo "ArgoCD admin password: $ARGOCD_PASSWORD"
fi

# Apply the ArgoCD application
kubectl apply -f argo-app.yaml

echo "✅ ArgoCD application created successfully!"
echo ""
echo "📋 ArgoCD Information:"
echo "   Namespace: argocd"
echo "   Application: portfolio-app"
echo "   Target Namespace: portfolio"
echo ""
echo "🌐 To access ArgoCD UI:"
echo "   kubectl port-forward svc/argocd-server 8080:443 -n argocd"
echo "   Then open https://localhost:8080 in your browser"
echo "   Username: admin"
echo "   Password: $ARGOCD_PASSWORD"
echo ""
echo "🔍 To check application status:"
echo "   kubectl get application -n argocd"
echo "   kubectl get pods -n portfolio" 