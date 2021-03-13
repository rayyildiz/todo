resource "google_storage_bucket" "ui" {
  name          = var.ui_domain
  location      = "EU"
  force_destroy = true
  storage_class = "STANDARD"

  uniform_bucket_level_access = true

  website {
    main_page_suffix = "index.html"
    not_found_page   = "index.html"
  }
  cors {
    origin          = ["http://${var.ui_domain}"]
    method          = ["GET", "HEAD"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
}

resource "google_storage_bucket_iam_member" "viewer" {
  bucket = google_storage_bucket.ui.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}
