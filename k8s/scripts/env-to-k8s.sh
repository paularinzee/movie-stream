#!/bin/bash

# --- CONFIG ---
NAMESPACE="movie-stream"
BACKEND_ENV="../backend.env"
FRONTEND_ENV="../frontend.env"
BACKEND_SECRET_NAME="movie-stream-backend-secrets"
FRONTEND_CONFIGMAP_NAME="movie-stream-frontend-config"
# ---------------

set -e

echo "✅ Creating namespace if it doesn't exist..."
kubectl get namespace $NAMESPACE >/dev/null 2>&1 || kubectl create namespace $NAMESPACE

echo "🔒 Generating Backend Secret from $BACKEND_ENV..."
kubectl delete secret $BACKEND_SECRET_NAME -n $NAMESPACE --ignore-not-found
kubectl create secret generic $BACKEND_SECRET_NAME -n $NAMESPACE \
  --from-env-file=$BACKEND_ENV

echo "🧾 Generating Frontend ConfigMap from $FRONTEND_ENV..."
kubectl delete configmap $FRONTEND_CONFIGMAP_NAME -n $NAMESPACE --ignore-not-found
kubectl create configmap $FRONTEND_CONFIGMAP_NAME -n $NAMESPACE \
  --from-env-file=$FRONTEND_ENV

echo "🎉 Done! Secrets and ConfigMaps have been updated."
kubectl get secret $BACKEND_SECRET_NAME -n $NAMESPACE
kubectl get configmap $FRONTEND_CONFIGMAP_NAME -n $NAMESPACE
