package rainbox

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func newTestFirebaseApp(ctx context.Context) (*firebase.App, error) {
	return firebase.NewApp(ctx, nil)
}

func newTestFirestore(ctx context.Context) (*firestore.Client, error) {
	app, err := newTestFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}
