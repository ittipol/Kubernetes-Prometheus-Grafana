
resource "helm_release" "jenkins" {
  name = "jenkins"

  repository       = "https://charts.jenkins.io"
  chart            = "jenkins"
  namespace        = "jenkins"
  version          = "5.7.21"
  create_namespace = true

  # values = [file("values/jenkins-values.yaml")]
}