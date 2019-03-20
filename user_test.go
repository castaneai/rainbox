package rainbox

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/stretchr/testify/assert"
	"testing"
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

func newTestFirebaseUser(ctx context.Context) (*auth.UserInfo, error) {
	// TODO: mock...
	return &auth.UserInfo{
		UID: "1c53d4e4-6c2a-4fb6-8ee2-be22b5f395ce",
		DisplayName: "testUser",
	}, nil
}

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	store, err := newTestFirestore(ctx)
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewUserRepository(store)
	if err != nil {
		t.Fatal(err)
	}

	fuser, err := newTestFirebaseUser(ctx)
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.SignIn(ctx, fuser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fuser.UID, user.UserID)
}
