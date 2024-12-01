## Exporter
- Node Exporter
- Blackbox Exporter
- Redis Exporter

## k6
``` bash
# Running k6
docker-compose run --rm k6 run /scripts/script.js

docker-compose run --rm k6 run {Path in container}
```

## Prometheus
- http://localhost:9090/targets (Target list) (prometheus.yml)

## Grafana Dashboard
### Searching dashboard ID
1. Go to https://grafana.com/grafana/dashboards/
2. Input keyword in search box (k6, redis, blackbox)
3. Click button "Copy ID to clipboard"

## Docker
``` bash
# Running docker compose
docker-compose up -d --build
```