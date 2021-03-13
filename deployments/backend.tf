resource "google_cloud_run_service" "backend" {
  name     = "${var.name}-service"
  location = var.region

  template {
    spec {
      containers {
        image = var.image

        ports {
          container_port = var.container_port
        }

        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "256Mi"
          }
        }

        env {
          name  = "DOCSTORE_COLLECTION"
          value = "firestore://projects/${var.projectId}/databases/(default)/documents/todos?name_field=id"
        }

        env {
          name  = "GRAPHQL_ENABLE_PLAYGROUND"
          value = "true"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}


data "google_iam_policy" "no_auth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "no_auth" {
  location = google_cloud_run_service.backend.location
  project  = google_cloud_run_service.backend.project
  service  = google_cloud_run_service.backend.name

  policy_data = data.google_iam_policy.no_auth.policy_data
}
