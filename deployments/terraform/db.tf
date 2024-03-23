resource "digitalocean_database_cluster" "db" {
  engine     = "mysql"
  name       = "todo-db-mysql-cluster"
  version = "8"
  node_count = 1
  region     = "nyc1"
  size       = "db-s-1vcpu-1gb"
}

resource "digitalocean_database_connection_pool" "db-pool" {
  cluster_id = digitalocean_database_cluster.db.id
  db_name    = "todo-db-mysql-pool"
  mode       = "transaction"
  name       = "todo"
  size       = 20
}

resource "digitalocean_database_firewall" "db-fw" {
  cluster_id = digitalocean_database_cluster.db.id

  rule {
    type  = "k8s"
    value = digitalocean_kubernetes_cluster.cluster.id
  }
}

resource "digitalocean_database_user" "db-user" {
  cluster_id = digitalocean_database_cluster.db.id
  name       = "todo"
}