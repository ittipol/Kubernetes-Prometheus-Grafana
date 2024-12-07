## Exporter
- Node Exporter
    - Exposes hardware and OS level metrics for *NIX kernels
- Blackbox Exporter
    - Metrics for “blackbox” probing of endpoints over HTTP, HTTPS, DNS, TCP, and ICMP
- Redis Exporter
    - Connects to local or remote Redis instance to provide various Redis metrics in Prometheus format
- Elasticsearch Exporter
    - Connects to local or remote Elasticsearch (ES) instance to provide various Elasticsearch metrics in Prometheus format
- Kube-state Metrics Exporter
    - Kubernetes (K8s) add-on agent which provides metrics about various Kubernetes objects, such as deployments, nodes, and pods

## k6
``` bash
# Running k6
docker-compose run --rm k6 run /scripts/script.js

docker-compose run --rm k6 run {Path in container}
```

## Prometheus
- http://localhost:9090/targets (Target list) (prometheus.yml)

## Docker
``` bash
# Running docker compose
docker-compose up -d --build
```