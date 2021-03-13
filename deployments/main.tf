terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.59.0"
    }
  }

  backend "gcs" {
    bucket = "rayyildiz-terraform-state"
    prefix = "todo"
  }
}

provider "google" {
  project = var.projectId
  region  = var.region
}


provider "google-beta" {
  project = var.projectId
  region  = var.region
}
