## Redis

### Create Namespace
``` bash
kubectl create namespace database-monitoring
```

### Apply configuration inside of the redis folder
``` bash
kubectl apply -f redis/ -n database-monitoring

kubectl get all -n database-monitoring

kubectl get secret -n database-monitoring
```

### Get a Shell to a Redis pod
``` bash
kubectl exec -n database-monitoring -it pod/redis-0 -- bash

# Inside Redis pod
redis-cli

# Authentication
# Password is stored in the secret
AUTH <password>
```

### Connect service in a different namespace
- {protocol}://{service_name}.{service_namespace}.{Kubernetes_suffix}:{service_port}

### Install Redis Exporter
- https://github.com/prometheus-community/helm-charts/tree/main/charts
- https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus-redis-exporter

``` bash
# Add repo to local
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm search repo prometheus-community

# Use chart name prometheus-community/prometheus-redis-exporter
# helm install [RELEASE_NAME] prometheus-community/prometheus-redis-exporter
helm install redis-exporter prometheus-community/prometheus-redis-exporter -f redis-exporter/redis-exporter-values.yaml -n database-monitoring --version 6.8.0

# List releases of charts
helm list -A
```

### Use service monitor to collect metrics
- https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus-redis-exporter/templates/servicemonitor.yaml

``` yaml
serviceMonitor:
  # When set true then use a ServiceMonitor to configure scraping
  enabled: true
  # Set the namespace the ServiceMonitor should be deployed
  # namespace: database-monitoring
  # Set how frequently Prometheus should scrape
  # interval: 30s
  # Set path to redis-exporter telemtery-path
  # Please set telemetryPath to /scrape if you are using multiple targets
  # telemetryPath: /metrics
  # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
  labels:
    release: prometheus
  # Set timeout for scrape
  # timeout: 10s
  # Set relabel_configs as per https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
  # relabelings: []
  # Set of labels to transfer on the Kubernetes Service onto the target.
  # targetLabels: []
  # metricRelabelings: []
  # Set tls options
  # scheme: ""
  # tlsConfig: {}
```

``` bash
# Get service monitor
kubectl get servicemonitor -n database-monitoring
```

### Upgrade value file for Redis Exporter 
``` bash
helm upgrade redis-exporter prometheus-community/prometheus-redis-exporter -f redis-exporter/redis-exporter-values.yaml -n database-monitoring
```

### Access Redis Exporter
``` bash
# End with "&", It will run this port forward in the background
kubectl port-forward svc/redis-exporter-prometheus-redis-exporter 9121:9121 -n database-monitoring &
```

### Access Prometheus
- http://localhost:9090/targets?search=redis
``` bash
# End with "&", It will run this port forward in the background
kubectl port-forward svc/prometheus-operated 9090:9090 -n monitoring &
```

### Access Grafana
``` bash
# End with "&", It will run this port forward in the background
kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring &
```

### Add Grafana dashboard
``` bash
# Add Grafana dashboard inside of the same namespace as where Grafana is in
kubectl apply -f grafana/ -n monitoring
```

### Kill port forward running in the background
``` bash
ps
#or
ps aux | grep kubectl

# To find signal numbers
kill -l

# Use PID to kill the process
kill [signal] PID

# Output:
# No output if the command is successful and the process with PID is terminated.
```

### Uninstall Redis Exporter
``` bash
helm uninstall redis-exporter -n database-monitoring
```

### Delete Redis 
``` bash
kubectl delete -f redis/ -n database-monitoring

kubectl delete -f grafana/ -n monitoring

kubectl delete namespace database-monitoring
```