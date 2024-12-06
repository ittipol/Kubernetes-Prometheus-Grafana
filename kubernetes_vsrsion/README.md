## Helm
- https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack
- https://artifacthub.io/packages/helm/prometheus-community/prometheus


### Installing Helm
``` bash
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3

$ chmod 700 get_helm.sh

$ ./get_helm.sh
```

### From Homebrew
``` bash
brew install helm
```

### Install Prometheus Chart
#### You can then run helm search repo prometheus-community to see the charts.
``` bash
# Add to repositories
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm repo update

# Install specific chart on Kubernetes cluster
# helm install [RELEASE_NAME] prometheus-community/kube-prometheus-stack
helm install prometheus prometheus-community/kube-prometheus-stack
# or
# helm install [RELEASE_NAME] prometheus-community/prometheus
helm install prometheus prometheus-community/prometheus

# Apply custom values.yaml
# helm install prometheus [CHART_NAME] -f values.yaml
helm install prometheus prometheus-community/kube-prometheus-stack -f values.yaml
```

### Show the chart's values
``` bash
# helm show values [CHART] [flags]
helm show values prometheus-community/kube-prometheus-stack

# Output to values.yaml
helm show values prometheus-community/kube-prometheus-stack > values.yaml
```

### Uninstall Chart
``` bash
helm uninstall [RELEASE_NAME]
```

### Upgrade Chart
``` bash
helm upgrade [RELEASE_NAME] prometheus-community/prometheus --install
```

## Kubernetes
``` bash
kubectl get all

kubectl get node

kubectl get deployment

kubectl get statefulset

kubectl get pods

kubectl get replicaset

kubectl get service

# kubectl describe [TYPE] [NAME_PREFIX]
kubectl describe statefulset prometheus-prometheus-kube-prometheus-prometheus
# Output to file
kubectl describe statefulset prometheus-prometheus-kube-prometheus-prometheus > prometheus.yaml

# kubectl get [TYPE] [NAME_PREFIX] -o {yaml|json}
kubectl get statefulset prometheus-prometheus-kube-prometheus-prometheus -o yaml
# Output to file
kubectl get service prometheus-kube-prometheus-prometheus -o yaml > service.yaml
```

### Forward a local port to a port on the Pod
``` bash
kubectl port-forward prometheus-prometheus-kube-prometheus-prometheus-0 9090
#or
kubectl port-forward svc/prometheus-operated 9090:9090

kubectl port-forward prometheus-grafana-665f5c84b9-b8jp8 3000
#or
kubectl port-forward svc/prometheus-grafana 3000:80
```

### Namespace
``` bash
kubectl create namespace monitoring
```

### CRD
``` bash
kubectl get crd
```
- **prometheuses.monitoring.coreos.com**, Used to create a prometheus instance
- **servicemonitors.monitoring.coreos.com**, Add additional targets for prometheus to scrape

### Service Monitor
- Define a set of targets for prometheus to monitor and scrape
``` yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: go-api-service-monitor
  #namespace: monitoring
  labels:
    release: prometheus # Matches kube-prometheus-stack release label
    app: prometheus
spec:
  jobLabel: job
  selector:
    matchLabels:
      app: service-api
  endpoints:
  - port: http-metrics
    interval: 10s
    # path: /api-path/metrics
```

## Grafana
### Login
- username: admin
- password: prom-operator

## Docker
### Build image in Kubernetes cluster
``` bash
minikube docker-env

# To point your shell to minikube's docker-daemon, run:
eval $(minikube -p minikube docker-env)

docker build -t go-app:v1 .

minikube ssh

docker image ls
```