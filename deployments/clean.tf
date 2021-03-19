resource "google_cloud_run_service" "clean" {
  location = var.region
  name     = "${var.name}-clean-service"

  template {
    spec {
      containers {
        image = var.clean_image

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
      }
    }


    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "2"
      }
    }

  }
  autogenerate_revision_name = true

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [metadata.0.annotations]
  }
}

resource "google_cloud_scheduler_job" "clean" {
  name     = "${var.name}-clean-scheduled"
  schedule = "0 */2 * * *"

  http_target {
    http_method = "POST"
    uri         = "${google_cloud_run_service.clean.status[0].url}/clean"

    oidc_token {
      service_account_email = var.iam_service_invoker_mail
      audience              = "${google_cloud_run_service.clean.status[0].url}/clean"
    }
  }
}

