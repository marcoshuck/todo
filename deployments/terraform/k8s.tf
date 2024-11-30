resource "digitalocean_kubernetes_cluster" "cluster" {
  name   = "todo-k8s-cluster"
  region = "nyc1"
  version = "1.28.2-do.0"

  node_pool {
    name       = "todo-k8s-worker-pool"
    size       = "s-2vcpu-2gb"
    node_count = 3
  }
}