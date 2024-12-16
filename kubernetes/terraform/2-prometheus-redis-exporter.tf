
resource "helm_release" "redis-exporter" {
  name = "redis-exporter"

  repository       = "https://prometheus-community.github.io/helm-charts"
  chart            = "prometheus-redis-exporter"
  namespace        = "database-monitoring"
  version          = "6.8.0"
  create_namespace = true

  values = [file("values/redis-exporter-values.yaml")]
}