variable "do_token" {
  description = "Digital Ocean API token"
  type        = string
  sensitive   = true
}

variable "region" {
  description = "Digital Ocean region for resources"
  type        = string
  default     = "nyc1"
}

# -----------------------------------------------------------------------------
# Existing PostgreSQL cluster
# -----------------------------------------------------------------------------

variable "postgres_cluster_id" {
  description = "ID of the existing managed PostgreSQL cluster in Digital Ocean"
  type        = string
}

variable "app_database_name" {
  description = "Name of the database to create for this application"
  type        = string
  default     = "warzone_stats_tracker"
}

variable "app_database_user" {
  description = "Name of the database user to create for this application"
  type        = string
  default     = "warzone_app"
}

# -----------------------------------------------------------------------------
# App Platform
# -----------------------------------------------------------------------------

variable "app_name" {
  description = "Name of the Digital Ocean App Platform application"
  type        = string
  default     = "warzone-stats-tracker"
}

variable "github_repo" {
  description = "GitHub repository in owner/repo format"
  type        = string
  default     = "grovecj/warzone-stats-tracker"
}

variable "deploy_branch" {
  description = "Branch to deploy from"
  type        = string
  default     = "main"
}

variable "instance_size" {
  description = "App Platform instance size slug"
  type        = string
  default     = "apps-s-1vcpu-0.5gb"
}

variable "cod_sso_token" {
  description = "Call of Duty SSO token (ACT_SSO_COOKIE) for API authentication"
  type        = string
  sensitive   = true
}

# -----------------------------------------------------------------------------
# Domain (optional)
# -----------------------------------------------------------------------------

variable "domain_name" {
  description = "Custom domain name for the app (leave empty to use DO default domain)"
  type        = string
  default     = ""
}
