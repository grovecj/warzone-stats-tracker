provider "digitalocean" {
  token = var.do_token
}

# -----------------------------------------------------------------------------
# Reference the existing PostgreSQL cluster
# -----------------------------------------------------------------------------

data "digitalocean_database_cluster" "postgres" {
  id = var.postgres_cluster_id
}

# Create a dedicated database for this application
resource "digitalocean_database_db" "app_db" {
  cluster_id = data.digitalocean_database_cluster.postgres.id
  name       = var.app_database_name
}

# Create a dedicated user for this application
resource "digitalocean_database_user" "app_user" {
  cluster_id = data.digitalocean_database_cluster.postgres.id
  name       = var.app_database_user
}

# Restrict the app user to the app database
resource "digitalocean_database_connection_pool" "app_pool" {
  cluster_id = data.digitalocean_database_cluster.postgres.id
  name       = "${var.app_database_name}-pool"
  mode       = "transaction"
  size       = 10
  db_name    = digitalocean_database_db.app_db.name
  user       = digitalocean_database_user.app_user.name
}

# -----------------------------------------------------------------------------
# Firewall: only allow app platform to reach the database
# -----------------------------------------------------------------------------

resource "digitalocean_database_firewall" "app_fw" {
  cluster_id = data.digitalocean_database_cluster.postgres.id

  rule {
    type  = "app"
    value = digitalocean_app.warzone_stats_tracker.id
  }
}

# -----------------------------------------------------------------------------
# App Platform
# -----------------------------------------------------------------------------

resource "digitalocean_app" "warzone_stats_tracker" {
  spec {
    name   = var.app_name
    region = var.region

    service {
      name               = "api"
      instance_count     = 1
      instance_size_slug = var.instance_size

      dockerfile_path = "Dockerfile"

      github {
        repo           = var.github_repo
        branch         = var.deploy_branch
        deploy_on_push = true
      }

      http_port = 8080

      health_check {
        http_path = "/api/v1/health"
      }

      env {
        key   = "DATABASE_URL"
        value = digitalocean_database_connection_pool.app_pool.uri
        type  = "SECRET"
      }

      env {
        key   = "COD_SSO_TOKEN"
        value = var.cod_sso_token
        type  = "SECRET"
      }

      env {
        key   = "PORT"
        value = "8080"
      }

      env {
        key   = "CORS_ALLOWED_ORIGINS"
        value = var.domain_name != "" ? "https://${var.domain_name}" : "*"
      }
    }

    dynamic "domain_spec" {
      for_each = var.domain_name != "" ? [var.domain_name] : []
      content {
        domain = domain_spec.value
        type   = "PRIMARY"
      }
    }
  }
}
