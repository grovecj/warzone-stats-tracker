variable "github_token" {
  description = "GitHub personal access token with repo and admin:org permissions"
  type        = string
  sensitive   = true
}

variable "github_owner" {
  description = "GitHub username or organization that owns the repository"
  type        = string
}

variable "repository_name" {
  description = "Name of the GitHub repository"
  type        = string
  default     = "warzone-stats-tracker"
}

variable "repository_description" {
  description = "Description of the GitHub repository"
  type        = string
  default     = "Call of Duty Warzone statistics tracker. View, compare, and track squad stats. Built with Vue 3 + Go."
}

variable "default_branch" {
  description = "Default branch name"
  type        = string
  default     = "main"
}

variable "required_pr_reviews" {
  description = "Number of required PR review approvals before merging to the default branch"
  type        = number
  default     = 1
}
