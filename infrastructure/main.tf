provider "google" {
  project = "go-api-355418"
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_project_service" "project" {
  service = "run.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

  disable_dependent_services = true
}

module "cloud_run" {
  source = "./modules/cloudrun"

  name    = "api"
  project = var.project
}