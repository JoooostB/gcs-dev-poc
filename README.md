# PoC: GCS signed object URL's

Quick development proof-of-concept on generating signed object URL's for a private GCP bucket.
> **Disclaimer**: This is a proof-of-concept and not production ready code. Don't use JSON authentication in production, use a form of [Application Default Credentials](https://cloud.google.com/docs/authentication/production) instead.

## Usage

1. Create a (randomised) private bucket with associated service account and role bindings:

```hcl
resource "random_id" "eight" {
  byte_length = 8
}

resource "google_storage_bucket" "gcs_dev_poc" {
  name          = "gcs-dev-poc-${random_id.eight.hex}"
  location      = "europe-west4"
  force_destroy = true

  uniform_bucket_level_access = true

  public_access_prevention = "enforced"

  storage_class = "STANDARD"
  labels = {
    "owner"        = "joooostb"
    "purpose"      = "proof-of-concept"
    "remove-after" = "2023-12-31"
  }
}

resource "google_storage_bucket_iam_binding" "binding" {
  bucket = google_storage_bucket.gcs_dev_poc.name
  role   = "roles/storage.admin"
  members = [
    "serviceAccount:${google_service_account.gcs_dev_poc.email}",
  ]
}

resource "google_service_account" "gcs_dev_poc" {
  account_id   = "gcs-dev-poc"
  display_name = "GCS dev PoC"
  description  = "Used for upload and retrieve attachment to bucket from dev application."
}
```

2. Generate key for serviceAccount in json and store as `auth.json` in root of this project.
3. Upload a file to the bucket.
4. Edit `main.go` and change the bucket name and object name to match the file you uploaded.
5. Run `go run main.go` and copy the output URL to your browser to download the file.
6. Test again after one minute and see the URL is expired.
