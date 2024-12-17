## Kubernetes
- API Conventions (https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
- ServiceMonitor [monitoring.coreos.com/v1] (https://docs.openshift.com/container-platform/4.12/rest_api/monitoring_apis/servicemonitor-monitoring-coreos-com-v1.html)

## Helm
- https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack
- https://artifacthub.io/packages/helm/prometheus-community/prometheus


### Install Helm
``` bash
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3

$ chmod 700 get_helm.sh

$ ./get_helm.sh
```

### From Homebrew
``` bash
brew install helm
```

### Add repo
``` bash
# Add repo to local
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm repo update

# Run helm search repo prometheus-community to see the charts
helm search repo prometheus-community
```

### Install kube-prometheus-stack chart
``` bash
# Install specific chart on Kubernetes cluster
# helm install [RELEASE_NAME] prometheus-community/kube-prometheus-stack
helm install prometheus prometheus-community/kube-prometheus-stack
# helm install prometheus --namespace default --create-namespace --version 66.3.1 prometheus-community/kube-prometheus-stack

# helm install [RELEASE_NAME] prometheus-community/prometheus
# helm install prometheus prometheus-community/prometheus

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

### Remove Helm repo
``` bash
helm repo remove prometheus-community
```

### ClusterRole
``` bash
kubectl get clusterrole prometheus-grafana-clusterrole -o yaml
kubectl get clusterrole prometheus-kube-state-metrics -o yaml
kubectl get clusterrole prometheus-kube-prometheus-operator -o yaml
kubectl get clusterrole prometheus-kube-prometheus-prometheus -o yaml
```

``` yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    meta.helm.sh/release-name: prometheus
    meta.helm.sh/release-namespace: monitoring
  creationTimestamp: "2024-12-07T08:17:27Z"
  labels:
    app.kubernetes.io/instance: prometheus
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: grafana
    app.kubernetes.io/version: 11.3.1
    helm.sh/chart: grafana-8.6.4
  name: prometheus-grafana-clusterrole
  resourceVersion: "626"
  uid: 69c9c6c6-5484-4a2d-8595-7ffe2212a46d
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - secrets
  verbs:
  - get
  - watch
  - list

```

## Terraform
``` bash
# Initializes terraform working directory by installing necessary plugins (providers), setting up the backend, and initializing modules
terraform init

# Creates an execution plan, compares the desired state (defined in the configuration with the actual state of the infrastructure), and outputs the changes to be made
# Acts as a “dry run”
terraform plan

# Terraform applies the changes required to reach the new state of the infrastructure. This is where the actual creation, modification, or deletion of infrastructure happens, making it the most critical step in the Terraform workflow
# Actual changes
# terraform apply [options] [plan file]
terraform apply

terraform refresh
#or
terraform apply -refresh-only -auto-approve
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

# kubectl get [TYPE] [NAME_PREFIX] -o {yaml|json}
kubectl get statefulset prometheus-prometheus-kube-prometheus-prometheus -o yaml
# Output to file
kubectl get service prometheus-kube-prometheus-prometheus -o yaml > service.yaml
```

### Describe
- Search "Events:" to see an error if pod fail to start
``` bash
# kubectl describe [TYPE] [NAME_PREFIX]
kubectl describe statefulset prometheus-prometheus-kube-prometheus-prometheus
# Output to file
kubectl describe statefulset prometheus-prometheus-kube-prometheus-prometheus > prometheus.yaml

kubectl describe pods redis-0 -n database-monitoring
```

### Delete
``` bash
kubectl delete ([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)])
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

### Shell to pod
``` bash
# kubectl exec [POD] -- [COMMAND]
kubectl exec -it [POD_NAME] -- bash
```

### Run command
``` bash
# -c, --container string
# See the container name
# kubectl get pods [POD_NAME] -n jenkins -o yaml
# Search "containers", Property "name" will be in "containers" property
kubectl exec --namespace jenkins -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/additional/chart-admin-password && echo
```

### Log
``` bash
kubectl logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER]
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

## Deploy Go in Kubernetes cluster
``` bash
cd kubernetes-manifests/go-app

kubectl apply -f deployment.yml

kubectl apply -f service.yml

kubectl apply -f service-monitor.yml
```