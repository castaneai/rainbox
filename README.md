rainbox
============

rainbox api server

## Tips

### Authentication

Rainbox API server has **no** authentication logic and entrusts it to [Firebase Authentication](https://firebase.google.com/docs/auth)

## Testing

You can run tests with [Cloud Firestore Local Emulator](https://cloud.google.com/sdk/gcloud/reference/beta/emulators/firestore/).

```sh
gcloud beta emulators firestore start
export FIRESTORE_EMULATOR_HOST=xxxx:yyyy
go test ./...
```
