# Movie Stream Monitoring Setup

## 1. Install Prometheus + Grafana via Helm
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install monitoring prometheus-community/kube-prometheus-stack \
  --namespace monitoring --create-namespace

kubectl get pods -n monitoring

# Port-forward Grafana
kubectl port-forward svc/monitoring-grafana 3000:80 -n monitoring
# Login: admin / prom-operator

## 2. Apply monitoring configs
kubectl apply -f backend-servicemonitor.yaml
kubectl apply -f grafana-datasource-configmap.yaml
kubectl apply -f grafana-dashboard-configmap.yaml
kubectl apply -f alert-rules.yaml
