provider "github" {
  token = var.github_token
  owner = var.github_owner
}

# Import the existing repository rather than creating it
import {
  to = github_repository.warzone_stats_tracker
  id = var.repository_name
}

resource "github_repository" "warzone_stats_tracker" {
  name        = var.repository_name
  description = var.repository_description
  visibility  = "public"

  has_issues   = true
  has_projects = true
  has_wiki     = false

  allow_merge_commit = false
  allow_squash_merge = true
  allow_rebase_merge = true
  allow_auto_merge   = true

  squash_merge_commit_title   = "PR_TITLE"
  squash_merge_commit_message = "PR_BODY"

  delete_branch_on_merge = true

  topics = [
    "call-of-duty",
    "warzone",
    "stats-tracker",
    "vue",
    "golang",
    "terraform",
  ]
}

resource "github_branch_protection" "main" {
  repository_id = github_repository.warzone_stats_tracker.node_id
  pattern       = var.default_branch

  required_pull_request_reviews {
    required_approving_review_count = var.required_pr_reviews
    dismiss_stale_reviews           = true
    require_code_owner_reviews      = true
  }

  required_status_checks {
    strict = true
    contexts = [
      "Backend Build & Test",
      "Frontend Build & Test",
    ]
  }

  enforce_admins = false

  allows_force_pushes = false
  allows_deletions    = false
}
