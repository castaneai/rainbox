rainbox
============

rainbox api server with GCP

## Testing

A Test suite requires GCP credentials which have permission to Cloud Firestore.

```sh
gcloud beta emulators firestore start
export FIRESTORE_EMULATOR_HOST=localhost:8812
go test ./...
```