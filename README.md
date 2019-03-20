rainbox
============

rainbox api server with GCP

## Testing

A Test suite requires GCP credentials which have permission to Cloud Firestore.

```sh
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/key.json
go test ./...
```