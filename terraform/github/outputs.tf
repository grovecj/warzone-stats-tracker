output "repository_url" {
  description = "URL of the GitHub repository"
  value       = github_repository.warzone_stats_tracker.html_url
}

output "repository_ssh_clone_url" {
  description = "SSH clone URL"
  value       = github_repository.warzone_stats_tracker.ssh_clone_url
}

output "repository_http_clone_url" {
  description = "HTTP clone URL"
  value       = github_repository.warzone_stats_tracker.http_clone_url
}
