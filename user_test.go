package rainbox

import (
	"context"
	"firebase.google.com/go"
	"testing"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/stretchr/testify/assert"
)

const (
	testUserID = "1c53d4e4-6c2a-4fb6-8ee2-be22b5f395ce"
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
		UID:         testUserID,
		DisplayName: "testUser",
	}, nil
}

func newTestUser(ctx context.Context) (*User, error) {
	store, err := newTestFirestore(ctx)
	if err != nil {
		return nil, err
	}
	fuser, err := newTestFirebaseUser(ctx)
	if err != nil {
		return nil, err
	}
	repo := NewUserRepository(store)
	return repo.SignIn(ctx, fuser)
}

func TestSignIn(t *testing.T) {
	ctx := context.Background()
	user, err := newTestUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testUserID, user.UserID)
}
