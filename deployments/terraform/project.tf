resource "digitalocean_project" "project" {
  name        = var.project-name
  purpose     = "Web Application"
  environment = title(var.environment)

  resources = [
    digitalocean_kubernetes_cluster.cluster.urn,
    digitalocean_database_cluster.db.urn,
  ]
}
