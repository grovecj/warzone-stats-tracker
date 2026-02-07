output "app_url" {
  description = "Default URL of the deployed application"
  value       = digitalocean_app.warzone_stats_tracker.default_ingress
}

output "app_id" {
  description = "Digital Ocean App Platform application ID"
  value       = digitalocean_app.warzone_stats_tracker.id
}

output "database_name" {
  description = "Name of the application database"
  value       = digitalocean_database_db.app_db.name
}

output "database_user" {
  description = "Name of the application database user"
  value       = digitalocean_database_user.app_user.name
}

output "database_pool_uri" {
  description = "Connection pool URI (use this in the application)"
  value       = digitalocean_database_connection_pool.app_pool.uri
  sensitive   = true
}
